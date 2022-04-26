package config

type HttpConfig struct {
	BindAddr       string
	BindPort       int
	NodeIdentifier string
}

const (
	ADD_NODE          = "add-node"
	REMOVE_NODE       = "remove-node"
	HTTP_ADDR         = "http-addr"
	RAFT_ADDR         = "raft-addr"
	HTTP_PORT         = "http-port"
	BOOTSTRAP         = "bootstrap"
	RAFT_PORT         = "raft-port"
	NODE_IDENTIFIER   = "identifier"
	SNAPSHOT_LOCATION = "snapshot-location"
	COMMIT_LOCATION   = "commit-location"
	STORE_LOCATION    = "store-location"
	BASE_LOCATION     = "base-location"
	ACTION            = "action"
	KEY               = "key"
	VALUE             = "value"
)

type RpcCommand string

const (
	ADD_KEY        RpcCommand = RpcCommand("ADD")
	UPDATE_KEY     RpcCommand = RpcCommand("UPDATE")
	REMOVE_KEY     RpcCommand = RpcCommand("REMOVE")
	RETRIEVE_VALUE RpcCommand = RpcCommand("RETRIEVE")
)
