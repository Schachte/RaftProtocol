package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
)

const RaftRPC = 1

type RaftConfig struct {
	NodeIdentifier     string
	StorageLocation    string
	BindAddr           string
	BindPort           int
	FiniteStateMachine raft.FSM
}

func SetupRaft(cfg *RaftConfig, bootstrapEnabled bool) (*raft.Raft, error) {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(cfg.NodeIdentifier)
	baseDir := filepath.Join(cfg.StorageLocation, cfg.NodeIdentifier)

	logStore, err := boltdb.NewBoltStore(filepath.Join(baseDir, "logs.dat"))
	if err != nil {
		return nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, "logs.dat"), err)
	}

	stableStore, err := boltdb.NewBoltStore(filepath.Join(baseDir, "stable.dat"))
	if err != nil {
		return nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, "stable.dat"), err)
	}

	fileSnapshotStore, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf(`raft.NewFileSnapshotStore(%q, ...): %v`, baseDir, err)
	}

	transport, err := raft.NewTCPTransport(fmt.Sprintf("%s:%d", cfg.BindAddr, cfg.BindPort), nil, 5, 10*time.Second, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	rs := &ReminderService{}

	r, err := raft.NewRaft(
		raftConfig,
		rs,
		logStore,
		stableStore,
		fileSnapshotStore,
		transport,
	)

	if err != nil {
		return nil, fmt.Errorf("raft.NewRaft: %v", err)
	}

	if bootstrapEnabled {
		cfg := raft.Configuration{
			Servers: []raft.Server{
				{
					Suffrage: raft.Voter,
					ID:       raft.ServerID(cfg.NodeIdentifier),
					Address:  raft.ServerAddress(fmt.Sprintf("%s:%d", cfg.BindAddr, cfg.BindPort)),
				},
			},
		}

		f := r.BootstrapCluster(cfg)
		if err := f.Error(); err != nil {
			return nil, fmt.Errorf("raft.Raft.BootstrapCluster: %v", err)
		}
	}

	return r, nil
}
