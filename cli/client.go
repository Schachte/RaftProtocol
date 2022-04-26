package cli

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v2"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

func InvokeClient() {
	conf := config{
		Server: configServer{
			Port: v.GetInt(serverPort),
		},
		Raft: configRaft{
			NodeId:    v.GetString(raftNodeId),
			Port:      v.GetInt(raftPort),
			VolumeDir: v.GetString(raftVolDir),
		},
	}

	log.Printf("%+v\n", conf)

	// Preparing badgerDB
	badgerOpt := badger.DefaultOptions(conf.Raft.VolumeDir)
	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		if err := badgerDB.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error close badgerDB: %s\n", err.Error())
		}
	}()

	var raftBinAddr = fmt.Sprintf(":%d", conf.Raft.Port)

	raftConf := raft.DefaultConfig()
	raftConf.LocalID = raft.ServerID(conf.Raft.NodeId)
	raftConf.SnapshotThreshold = 1024

	fsmStore := fsm.NewBadger(badgerDB)

	store, err := raftboltdb.NewBoltStore(filepath.Join(conf.Raft.VolumeDir, "raft.dataRepo"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Wrap the store in a LogCache to improve performance.
	cacheStore, err := raft.NewLogCache(raftLogCacheSize, store)
	if err != nil {
		log.Fatal(err)
		return
	}

	snapshotStore, err := raft.NewFileSnapshotStore(conf.Raft.VolumeDir, raftSnapShotRetain, os.Stdout)
	if err != nil {
		log.Fatal(err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", raftBinAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	transport, err := raft.NewTCPTransport(raftBinAddr, tcpAddr, maxPool, tcpTimeout, os.Stdout)
	if err != nil {
		log.Fatal(err)
		return
	}

	raftServer, err := raft.NewRaft(raftConf, fsmStore, cacheStore, store, snapshotStore, transport)
	if err != nil {
		log.Fatal(err)
		return
	}

	// always start single server as a leader
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(conf.Raft.NodeId),
				Address: transport.LocalAddr(),
			},
		},
	}

	raftServer.BootstrapCluster(configuration)

	srv := server.New(fmt.Sprintf(":%d", conf.Server.Port), badgerDB, raftServer)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}

	return
}
