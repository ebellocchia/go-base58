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

//
// This file cointains some common constants, variables and functions for base58 package.
//

package base58

//
// Imports
//
import (
	"crypto/sha256"
	"errors"
	"math/big"
)

//
// Constants
//
const (
	// Supported alphabets
	AlphabetBitcoin = 0
	AlphabetRipple  = 1
	AlphabetFlickr  = 2
	// Checksum length
	checksumLen = 4
)

//
// Variables
//
var (
	// Base58 radix as big.Int
	bigRadix = big.NewInt(58)
	// Map from alphabet index to alphabet string
	alphabetMap = map[int]string {
		AlphabetBitcoin : "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz",
		AlphabetRipple  : "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz",
		AlphabetFlickr  : "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ",
	}
	// ErrInvalidAlphabet is returned when trying to get a not-existent alphabet
	ErrInvalidAlphabet = errors.New("The specified alphabet is not existent")
)

//
// Types
//

// Base58 structure. It basically holds the alphabet index to be used.
// The default value (0) is the Bitcoin alphabet.
type Base58Obj struct {
	AlphIdx int
}

//
// Exported functions
//

// Helper function for creating Base58Obj structuer from the alphabet index.
func New(alphIdx int) (*Base58Obj) {
	return &Base58Obj {
		AlphIdx: alphIdx,
	}
}

//
// Not-exported functions
//

// Get the alphabet string from the specified alphabet index
func getAlphabet(alphIdx int) (string, error) {
	alph, ok := alphabetMap[alphIdx]
	if !ok {
		return "", ErrInvalidAlphabet
	}
	return alph, nil
}

// Compute Base58 checksum.
// The checksum is defined as the first 4-byte of the double SHA256 of the input.
func computeCheckum(slice []byte) (chksum [checksumLen]byte) {
	// Compute the double SHA256
	hash := doubleSha256(slice)
	// The first 4-byte of the hash is the checksum
	copy(chksum[:], hash[:checksumLen])

	return chksum
}

// Compute the double SHA256 of the specified byte slice
func doubleSha256(slice []byte) []byte {
	hash1 := sha256.Sum256(slice)
	hash2 := sha256.Sum256(hash1[:])

	return hash2[:]
}
