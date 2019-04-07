package tool

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func CompileSolidity(source string, contracts ...string) (
	[][]byte, error) {

	out, err := compiler.CompileSolidity("solc", source)
	if nil != err {
		return nil, err
	}

	bytecodes := make([][]byte, len(contracts))
	for i, c := range contracts {
		for name, C := range out {
			if strings.HasSuffix(name, c) {
				bytecodes[i] = hexutil.MustDecode(C.Code)
				break
			}
		}

		if nil == bytecodes[i] {
			return nil, errors.New("missing bytecodes for " + c)
		}
	}

	return bytecodes, nil
}
