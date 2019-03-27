package eth

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func NewAccount(keydir, passphrase string) (*keystore.KeyStore,
	accounts.Account, error) {
	store := keystore.NewKeyStore(keydir, keystore.StandardScryptN,
		keystore.StandardScryptP)

	account, err := store.NewAccount(passphrase)
	if nil != err {
		return nil, accounts.Account{}, err
	}

	return store, account, nil
}
