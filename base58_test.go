// Copyright (c) 2020 Emanuele Bellocchia
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package base58

//
// Imports
//
import (
	"bytes"
	"encoding/hex"
	"testing"
)

//
// Types
//

// Single test vector entry structure
type testVectEntry struct {
	Hex      string
	Enc      string
	CheckEnc string
}

//
// Variables
//

// Test vector for Bitcoin
var testVectBtc = []testVectEntry {
	testVectEntry {
		Hex:      "61",
		Enc:      "2g",
		CheckEnc: "C2dGTwc",
	},
	testVectEntry {
		Hex:      "626262",
		Enc:      "a3gV",
		CheckEnc: "4jF5uERJAK",
	},
	testVectEntry {
		Hex:      "636363",
		Enc:      "aPEr",
		CheckEnc: "4mT4krqUYJ",
	},
	testVectEntry {
		Hex:      "73696d706c792061206c6f6e6720737472696e67",
		Enc:      "2cFupjhnEsSn59qHXstmK2ffpLv2",
		CheckEnc: "BXF1HuEUCqeVzZdrKeJjG74rjeXxqJ7dW",
	},
	testVectEntry {
		Hex:      "00eb15231dfceb60925886b67d065299925915aeb172c06647",
		Enc:      "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L",
		CheckEnc: "13REmUhe2ckUKy1FvM7AMCdtyYq831yxM3QeyEu4",
	},
	testVectEntry {
		Hex:      "516b6fcd0f",
		Enc:      "ABnLTmg",
		CheckEnc: "237LSrY9NUUas",
	},
	testVectEntry {
		Hex:      "bf4f89001e670274dd",
		Enc:      "3SEo3LWLoPntC",
		CheckEnc: "GwDDDeduj1jpykc27e",
	},
	testVectEntry {
		Hex:      "572e4794",
		Enc:      "3EFU7m",
		CheckEnc: "FamExfqCeza",
	},
	testVectEntry {
		Hex:      "ecac89cad93923c02321",
		Enc:      "EJDM8drfXA6uyA",
		CheckEnc: "2W1Yd5Zu6WGyKVtHGMrH",
	},
	testVectEntry {
		Hex:      "10c8511e",
		Enc:      "Rt5zm",
		CheckEnc: "3op3iuGMmhs",
	},
	testVectEntry {
		Hex:      "00000000000000000000",
		Enc:      "1111111111",
		CheckEnc: "111111111146Momb",
	},
	testVectEntry {
		Hex:      "000111d38e5fc9071ffcd20b4a763cc9ae4f252bb4e48fd66a835e252ada93ff480d6dd43dc62a641155a5",
		Enc:      "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz",
		CheckEnc: "17mxz9b2TuLnDf6XyQrHjAc3UvMoEg7YzRsJkBd4VwNpFh8a1StKmCe5WtAW27Y",
	},
	testVectEntry {
		Hex:      "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f808182838485868788898a8b8c8d8e8f909192939495969798999a9b9c9d9e9fa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c9cacbcccdcecfd0d1d2d3d4d5d6d7d8d9dadbdcdddedfe0e1e2e3e4e5e6e7e8e9eaebecedeeeff0f1f2f3f4f5f6f7f8f9fafbfcfdfeff",
		Enc:      "1cWB5HCBdLjAuqGGReWE3R3CguuwSjw6RHn39s2yuDRTS5NsBgNiFpWgAnEx6VQi8csexkgYw3mdYrMHr8x9i7aEwP8kZ7vccXWqKDvGv3u1GxFKPuAkn8JCPPGDMf3vMMnbzm6Nh9zh1gcNsMvH3ZNLmP5fSG6DGbbi2tuwMWPthr4boWwCxf7ewSgNQeacyozhKDDQQ1qL5fQFUW52QKUZDZ5fw3KXNQJMcNTcaB723LchjeKun7MuGW5qyCBZYzA1KjofN1gYBV3NqyhQJ3Ns746GNuf9N2pQPmHz4xpnSrrfCvy6TVVz5d4PdrjeshsWQwpZsZGzvbdAdN8MKV5QsBDY",
		CheckEnc: "151KWPPBRzdWPr1ASeu172gVgLf1YfUp6VJyk6K9t4cLqYtFHcMa2iX8S3NJEprUcW7W5LvaPRpz7UG7puBj5STE3nKhCGt5eckYq7mMn5nT7oTTic2BAX6zDdqrmGCnkszQkzkz8e5QLGDjf7KeQgtEDm4UER6DMSdBjFQVa6cHrrJn9myVyyhUrsVnfUk2WmNFZvkWv3Tnvzo2cJ1xW62XDfUgYz1pd97eUGGPuXvDFfLsBVd1dfdUhPwxW7pMPgdWHTmg5uqKGFF6vE4xXpAqZTbTxRZjCDdTn68c2wrcxApm8hq3JX65Hix7VtcD13FF8b7BzBtwjXq1ze6NMjKgUcqpGV5XA5",
	},
}

// Test vector for Ripple
var testVectXrp = []testVectEntry {
	testVectEntry {
		Hex:      "61",
		Enc:      "pg",
		CheckEnc: "UpdGTAc",
	},
	testVectEntry {
		Hex:      "626262",
		Enc:      "2sgV",
		CheckEnc: "hjEnuNRJwK",
	},
	testVectEntry {
		Hex:      "636363",
		Enc:      "2PNi",
		CheckEnc: "hmThkiq7YJ",
	},
	testVectEntry {
		Hex:      "73696d706c792061206c6f6e6720737472696e67",
		Enc:      "pcEuFj68N1S8n9qHX1tmKpCCFLvp",
		CheckEnc: "BXErHuN7UqeVzZdiKeJjGfhijeXxqJfdW",
	},
	testVectEntry {
		Hex:      "00eb15231dfceb60925886b67d065299925915aeb172c06647",
		Enc:      "r4Srf52g9jJgTHDrVXjvLUN8ZuQsiJDN9L",
		CheckEnc: "rsRNm76epck7KyrEvMfwMUdtyYq3sryxMsQeyNuh",
	},
	testVectEntry {
		Hex:      "516b6fcd0f",
		Enc:      "wB8LTmg",
		CheckEnc: "psfLSiY947721",
	},
	testVectEntry {
		Hex:      "bf4f89001e670274dd",
		Enc:      "sSNosLWLoP8tU",
		CheckEnc: "GADDDedujrjFykcpfe",
	},
	testVectEntry {
		Hex:      "572e4794",
		Enc:      "sNE7fm",
		CheckEnc: "E2mNxCqUez2",
	},
	testVectEntry {
		Hex:      "ecac89cad93923c02321",
		Enc:      "NJDM3diCXwauyw",
		CheckEnc: "pWrYdnZuaWGyKVtHGMiH",
	},
	testVectEntry {
		Hex:      "10c8511e",
		Enc:      "Rtnzm",
		CheckEnc: "soFs5uGMm61",
	},
	testVectEntry {
		Hex:      "00000000000000000000",
		Enc:      "rrrrrrrrrr",
		CheckEnc: "rrrrrrrrrrhaMomb",
	},
	testVectEntry {
		Hex:      "000111d38e5fc9071ffcd20b4a763cc9ae4f252bb4e48fd66a835e252ada93ff480d6dd43dc62a641155a5",
		Enc:      "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz",
		CheckEnc: "rfmxz9bpTuL8DCaXyQiHjwcs7vMoNgfYzR1JkBdhVA4FE632rStKmUenWtwWpfY",
	},
	testVectEntry {
		Hex:      "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f808182838485868788898a8b8c8d8e8f909192939495969798999a9b9c9d9e9fa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c9cacbcccdcecfd0d1d2d3d4d5d6d7d8d9dadbdcdddedfe0e1e2e3e4e5e6e7e8e9eaebecedeeeff0f1f2f3f4f5f6f7f8f9fafbfcfdfeff",
		Enc:      "rcWBnHUBdLjwuqGGReWNsRsUguuASjAaRH8s91pyuDRTSn41Bg45EFWgw8NxaVQ53c1exkgYAsmdYiMHi3x95f2NAP3kZfvccXWqKDvGvsurGxEKPuwk83JUPPGDMCsvMM8bzma469z6rgc41MvHsZ4LmPnCSGaDGbb5ptuAMWPt6ihboWAUxCfeASg4Qe2cyoz6KDDQQrqLnCQE7WnpQK7ZDZnCAsKX4QJMc4Tc2BfpsLc6jeKu8fMuGWnqyUBZYzwrKjoC4rgYBVs4qy6QJs41fhaG4uC94pFQPmHzhxF8SiiCUvyaTVVzndhPdije161WQAFZ1ZGzvbdwd43MKVnQ1BDY",
		CheckEnc: "rnrKWPPBRzdWPirwSeurfpgVgLCrYC7FaVJykaK9thcLqYtEHcM2p5X3Ss4JNFi7cWfWnLv2PRFzf7GfFuBjnSTNs8K6UGtneckYqfmM8n8TfoTT5cpBwXazDdqimGU8k1zQkzkz3enQLGDjCfKeQgtNDmh7NRaDMSdBjEQV2acHiiJ89myVyy67i1V8C7kpWm4EZvkWvsT8vzopcJrxWapXDC7gYzrFd9fe7GGPuXvDECL1BVdrdCd76PAxWfFMPgdWHTmgnuqKGEEavNhxXFwqZTbTxRZjUDdT8a3cpAicxwFm36qsJXanH5xfVtcDrsEE3bfBzBtAjXqrzea4MjKg7cqFGVnXwn",
	},
}

// Test vector for Flickr
var testVectFlickr = []testVectEntry {
	testVectEntry {
		Hex:      "61",
		Enc:      "2F",
		CheckEnc: "c2CgsWB",
	},
	testVectEntry {
		Hex:      "626262",
		Enc:      "z3Fu",
		CheckEnc: "4Jf5Ueqiaj",
	},
	testVectEntry {
		Hex:      "636363",
		Enc:      "zoeR",
		CheckEnc: "4Ls4KRQtxi",
	},
	testVectEntry {
		Hex:      "73696d706c792061206c6f6e6720737472696e67",
		Enc:      "2BfUPJGMeSrM59QhwSTLj2EEPkV2",
		CheckEnc: "bwf1hUetcQDuZyCRjDiJg74RJDwXQi7Cv",
	},
	testVectEntry {
		Hex:      "00eb15231dfceb60925886b67d065299925915aeb172c06647",
		Enc:      "1nr17HzF9JiFshd1uwJVkceMyUp3Ride9k",
		CheckEnc: "13qeLtGD2BKtjY1fVm7amcCTYxQ831YXm3pDYeU4",
	},
	testVectEntry {
		Hex:      "516b6fcd0f",
		Enc:      "abMksLF",
		CheckEnc: "237krRx9nttzS",
	},
	testVectEntry {
		Hex:      "bf4f89001e670274dd",
		Enc:      "3reN3kvkNoMTc",
		CheckEnc: "gWdddDCUJ1JPYKB27D",
	},
	testVectEntry {
		Hex:      "572e4794",
		Enc:      "3eft7L",
		CheckEnc: "fzLeXEQcDZz",
	},
	testVectEntry {
		Hex:      "ecac89cad93923c02321",
		Enc:      "eidm8CREwa6UYa",
		CheckEnc: "2v1xC5yU6vgYjuThgmRh",
	},
	testVectEntry {
		Hex:      "10c8511e",
		Enc:      "qT5ZL",
		CheckEnc: "3NP3HUgmLGS",
	},
	testVectEntry {
		Hex:      "00000000000000000000",
		Enc:      "1111111111",
		CheckEnc: "111111111146mNLA",
	},
	testVectEntry {
		Hex:      "000111d38e5fc9071ffcd20b4a763cc9ae4f252bb4e48fd66a835e252ada93ff480d6dd43dc62a641155a5",
		Enc:      "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ",
		CheckEnc: "17LXZ9A2sUkMdE6wYpRhJaB3tVmNeF7xZqSiKbC4uWnPfG8z1rTjLcD5vTav27x",
	},
	testVectEntry {
		Hex:      "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f808182838485868788898a8b8c8d8e8f909192939495969798999a9b9c9d9e9fa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c9cacbcccdcecfd0d1d2d3d4d5d6d7d8d9dadbdcdddedfe0e1e2e3e4e5e6e7e8e9eaebecedeeeff0f1f2f3f4f5f6f7f8f9fafbfcfdfeff",
		Enc:      "1Bvb5hcbCkJaUQggqDve3q3cFUUWrJW6qhM39S2YUdqsr5nSbFnHfPvFaMeX6upH8BSDXKFxW3LCxRmhR8X9H7zeWo8Ky7VBBwvQjdVgV3U1gXfjoUaKM8icoogdmE3VmmMAZL6nG9ZG1FBnSmVh3ynkLo5Erg6dgAAH2TUWmvoTGR4ANvWcXE7DWrFnpDzBYNZGjddpp1Qk5Epftv52pjtydy5EW3jwnpimBnsBzb723kBGJDjUM7mUgv5QYcbyxZa1jJNEn1Fxbu3nQYGpi3nS746gnUE9n2PpoLhZ4XPMrRREcVY6suuZ5C4oCRJDSGSvpWPySygZVACaCn8mju5pSbdx",
		CheckEnc: "151jvoobqZCvoR1arDU172FuFkE1xEtP6uiYK6j9T4BkQxTfhBmz2Hw8r3niePRtBv7v5kVzoqPZ7tg7PUbJ5rse3MjGcgT5DBKxQ7LmM5Ms7NssHB2baw6ZdCQRLgcMKSZpKZKZ8D5pkgdJE7jDpFTedL4teq6dmrCbJfpuz6BhRRiM9LYuYYGtRSuMEtK2vLnfyVKvV3sMVZN2Bi1Xv62wdEtFxZ1PC97DtggoUwVdfEkSbuC1CECtGoWXv7PmoFCvhsLF5UQjgff6Ve4XwPaQysAsXqyJcdCsM68B2WRBXaPL8GQ3iw65hHX7uTBd13ff8A7bZbTWJwQ1ZD6nmJjFtBQPgu5wa5",
	},
}

// Tests for base58 encoded strings with invalid checksum
var testVectChksumInvalid = []string {
	"237LSrY9NUUar",
	"GwDDDeduj1jpykc27a",
	"2W1Yd5Zu6WGyKVtHGMrJ",
}

// Tests for base58 encoded strings with invalid encoding
var testVectEncodingInvalid = []string {
	"237LSrYONUUar",
	"GwDDDeduj1jpykc27I",
	"2WlYd5Zu6WGyKVtHGMrJ",
}

//
// Functions
//

// Test encoding for a generic base58 object
func GenericTestEncoder(t *testing.T, testEntries []testVectEntry, obj *Base58Obj) {
	for _, currTest := range testEntries {
		// Convert hex to bytes
		raw, _ := hex.DecodeString(currTest.Hex)

		// Encode
		enc := obj.Encode(raw)
		if enc != currTest.Enc {
			t.Errorf("Encoding was incorrect: expected %s, got: %s", currTest.Enc, enc)
		}

		// CheckEncode
		checkEnc := obj.CheckEncode(raw)
		if checkEnc != currTest.CheckEnc {
			t.Errorf("Checksum encoding was incorrect: expected %s, got: %s", currTest.CheckEnc, checkEnc)
		}
	}
}

// Test decoding for a generic base58 object
func GenericTestDecoder(t *testing.T, testEntries []testVectEntry, obj *Base58Obj) {
	for _, currTest := range testEntries {
		// Convert hex to bytes
		raw, _ := hex.DecodeString(currTest.Hex)

		// Decode
		dec, err := obj.Decode(currTest.Enc)
		if err != nil {
			t.Errorf("Decoding (%s) returned error: %s", currTest.Hex, err.Error())
		}
		if bytes.Compare(dec, raw) != 0 {
			t.Errorf("Decoding was incorrect: expected %v, got: %v", raw, dec)
		}

		// CheckDecode
		checkDec, err := obj.CheckDecode(currTest.CheckEnc)
		if err != nil {
			t.Errorf("Checksum decoding (%s) returned error: %s", currTest.Hex, err.Error())
		}
		if bytes.Compare(checkDec, raw) != 0 {
			t.Errorf("Checksum decoding was incorrect: expected %v, got: %v", raw, checkDec)
		}
	}
}

// Test Base58 encoding
func TestEncoder(t *testing.T) {
	// Test all alphabets
	GenericTestEncoder(t, testVectBtc, New(AlphabetBitcoin))
	GenericTestEncoder(t, testVectXrp, New(AlphabetRipple))
	GenericTestEncoder(t, testVectFlickr, New(AlphabetFlickr))
}

// Test Base58 decoding
func TestDecoder(t *testing.T) {
	// Test all alphabets
	GenericTestDecoder(t, testVectBtc, New(AlphabetBitcoin))
	GenericTestDecoder(t, testVectXrp, New(AlphabetRipple))
	GenericTestDecoder(t, testVectFlickr, New(AlphabetFlickr))
}

// Test invalid checksum
func TestInvalidChecksum(t *testing.T) {
	base58Btc := New(AlphabetBitcoin)

	for _, currTest := range testVectChksumInvalid {
		// CheckDecode
		_, err := base58Btc.CheckDecode(currTest)
		if err != ErrInvalidChecksum {
			t.Errorf("Checksum decoding (%s) with invalid checksum returned wrong error", currTest)
		}
	}
}

// Test invalid encoding
func TestInvalidEncoding(t *testing.T) {
	base58Btc := New(AlphabetBitcoin)

	for _, currTest := range testVectEncodingInvalid {
		// Decode
		_, err := base58Btc.Decode(currTest)
		if err != ErrInvalidFormat {
			t.Errorf("Decoding (%s) with invalid encoding returned wrong error", currTest)
		}
		// CheckDecode
		_, err = base58Btc.CheckDecode(currTest)
		if err != ErrInvalidFormat {
			t.Errorf("Checksum decoding (%s) with invalid encoding returned wrong error", currTest)
		}
	}
}

// Test invalid alphabet
func TestInvalidAlphabet(t *testing.T) {
	// Create with invalid alphabet
	base58Obj := New(3)

	// Encode
	enc := base58Obj.Encode([]byte("test"))
	if enc != "" {
		t.Errorf("Encoding with invalid alphabet returned not-empty result")
	}

	// CheckEncode
	enc = base58Obj.CheckEncode([]byte("test"))
	if enc != "" {
		t.Errorf("Checksum encoding with invalid alphabet returned not-empty result")
	}

	// Decode
	dec, err := base58Obj.Decode("test")
	if dec != nil {
		t.Errorf("Decoding with invalid alphabet returned not-nil result")
	}
	if err != ErrInvalidAlphabet {
		t.Errorf("Decoding with invalid alphabet returned wrong error")
	}

	// CheckDecode
	dec, err = base58Obj.CheckDecode("test")
	if dec != nil {
		t.Errorf("Checksum decoding with invalid alphabet returned not-nil result")
	}
	if err != ErrInvalidAlphabet {
		t.Errorf("Checksum decoding with invalid alphabet returned wrong error")
	}
}
