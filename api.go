package hashimoto

import (
	client "github.com/consensusdb/value-rpc/valueclient"
)

type Client interface {

	RPC() client.Client

	Status() (string, error)

	Generate(blockNum uint64) error

	FullHash(blockNum uint64, hashNoNonce []byte, nonce uint64) (digest []byte, result []byte, err error)

	Stop(token string) error

	Close() error
}
