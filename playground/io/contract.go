package io

import (
	"encoding/hex"
	"io/ioutil"
)

// LoadContractCode reads in the code of a contract encoded as hex in
// the given path
func LoadContractCode(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}

	return hex.DecodeString(string(data))
}
