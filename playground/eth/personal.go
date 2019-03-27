package eth

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func ListAndUnlockAccounts(keydir, passphrase string) (*keystore.KeyStore,
	[]accounts.Account, error) {
	store := keystore.NewKeyStore(keydir, keystore.StandardScryptN,
		keystore.StandardScryptP)

	accounts := store.Accounts()
	for _, a := range accounts {
		if err := store.Unlock(a, passphrase); nil != err {
			return nil, nil, err
		}
	}

	return store, accounts, nil
}

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
