package hashimoto

import (
	"github.com/consensusdb/value"
	"github.com/consensusdb/value-rpc/client"
	"github.com/pkg/errors"
)

const defaultAddress = "localhost:9777"

type hashimotoClient struct {
	address string
	cli     client.Client
}

func NewClient(address, socks5 string) (Client, error) {

	return &hashimotoClient{
		address: address,
		cli:     client.NewClient(address, socks5),
	}, nil
}

func NewDefaultClient() (Client, error) {
	return NewClient(defaultAddress, "")
}

func (t *hashimotoClient) Close() {
	t.cli.Close()
}

func (t *hashimotoClient) VRpcClient() client.Client {
	return t.cli
}

func (t *hashimotoClient) Status() (string, error) {
	result, err := t.cli.CallFunction("status", nil)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", errors.New("empty result")
	}
	return result.String(), err
}

func (t *hashimotoClient) Generate(blockNum uint64) error {

	args := value.Long(int64(blockNum))

	_, err := t.cli.CallFunction("generate", args)
	return err

}

func (t *hashimotoClient) FullHash(blockNum uint64, hashNoNonce []byte, nonce uint64) ([]byte, []byte, error) {

	args := value.EmptyList().
		Append(value.Long(int64(blockNum))).
		Append(value.Raw(hashNoNonce, true)).
		Append(value.Long(int64(nonce)))

	res, err := t.cli.CallFunction("fullHash", args)
	if err != nil {
		return nil, nil, err
	}

	if res == nil || res.Kind() != value.LIST {
		return nil,nil, errors.Errorf("invalid result %v", res)
	}

	resList := res.(value.List)
	digest := resList.GetStringAt(0)
	result := resList.GetStringAt(1)

	if digest == nil || result == nil {
		return nil, nil, errors.Errorf("empty field in result %v", resList.String())
	}

	return digest.Raw(), result.Raw(), nil
}

func (t *hashimotoClient) Stop(token string) error {

	args := value.Utf8(token)

	_, err := t.cli.CallFunction("stop", args)
	return err

}
