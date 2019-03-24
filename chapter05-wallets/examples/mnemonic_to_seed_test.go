package main_test

import (
	"encoding/hex"
	"testing"

	"github.com/sammyne/bip39"
)

func MnemonicToSeed() {}

func Test_MnemonicToSeed(t *testing.T) {
	type expect struct {
		mnemonic string
		seed     string
	}

	testCases := []struct {
		description string
		entropy     string
		passphrase  string
		expect      expect
	}{
		{
			"128-bit entropy mnemonic code, no passphrase, resulting seed",
			"0c1e24e5917779d297e14d45f14e1a1a",
			"",
			expect{
				"army van defense carry jealous true garbage claim echo media make crunch",
				"5b56c417303faa3fcba7e57400e120a0ca83ec5a4fc9ffba757fbe63fbd77a89a1a3be4c67196f57c39a88b76373733891bfaba16ed27a813ceed498804c0570",
			},
		},
		{
			"128-bit entropy mnemonic code, with passphrase, resulting seed",
			"0c1e24e5917779d297e14d45f14e1a1a",
			"SuperDuperSecret",
			expect{
				"army van defense carry jealous true garbage claim echo media make crunch",
				"3b5df16df2157104cfdd22830162a5e170c0161653e3afe6c88defeefb0818c793dbb28ab3ab091897d0715861dc8a18358f80b79d49acf64142ae57037d1d54",
			},
		},
		{
			"256-bit entropy mnemonic code, no passphrase, resulting seed",
			"2041546864449caff939d32d574753fe684d3c947c3346713dd8423e74abcf8c",
			"",
			expect{
				"cake apple borrow silk endorse fitness top denial coil riot stay wolf luggage oxygen faint major edit measure invite love trap field dilemma oblige",
				"3269bce2674acbd188d4f120072b13b088a0ecf87c6e4cae41657a0bb78f5315b33b3a04356e53d062e55f1e0deaa082df8d487381379df848a6ad7e98798404",
			},
		},
	}

	for i, c := range testCases {
		entropy, _ := hex.DecodeString(c.entropy)
		if mnemonic, err := bip39.NewMnemonic(entropy); nil != err {
			t.Fatalf("#%d [%s] failed to generate mnemonic: %v", i, c.description, err)
		} else if mnemonic != c.expect.mnemonic {
			t.Fatalf("#%d [%s] wrong mnemonic: got %s, expect %s", i, c.description, mnemonic, c.expect.mnemonic)
		}

		if seed, err := bip39.GenerateSeed(c.expect.mnemonic, c.passphrase); nil != err {
			t.Fatalf("#%d [%s] unexpected error during generating seed: %v", i, c.description, err)
		} else if got := hex.EncodeToString(seed); got != c.expect.seed {
			t.Fatalf("#%d [%s] wrong seed: got %s, expect %s", i, c.description, got, c.expect.seed)
		}
	}
}
