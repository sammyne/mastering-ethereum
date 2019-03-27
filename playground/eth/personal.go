package eth

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func ExportAccounts(keydir, passphrase string) ([][]byte, error) {
	store := keystore.NewKeyStore(keydir, keystore.StandardScryptN,
		keystore.StandardScryptP)

	accounts := store.Accounts()
	if 0 == len(accounts) {
		return nil, errors.New("no account to export")
	}

	keyJSON := make([][]byte, len(accounts))
	for i, a := range accounts {
		var err error
		if keyJSON[i], err = store.Export(a, passphrase, passphrase); nil != err {
			return nil, err
		}
	}

	return keyJSON, nil
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

func UnlockAccounts(keydir, passphrase string) (*keystore.KeyStore,
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
