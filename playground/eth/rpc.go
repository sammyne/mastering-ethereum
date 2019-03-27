package eth

import "github.com/ethereum/go-ethereum/ethclient"

func Dial(url ...string) (*ethclient.Client, error) {
	if len(url) > 0 {
		return ethclient.Dial(url[0])
	}

	return ethclient.Dial(ProviderURL())
}
