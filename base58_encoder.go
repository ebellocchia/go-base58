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
// This file cointains encoding functions for base58 package.
//

package base58

//
// Imports
//
import (
	"math/big"
)

//
// Variables
//
var (
	// Zero as big.Int
	bigZero = big.NewInt(0)
)

//
// Exported functions
//

// Encode the specified bytes to Base58 format
func (obj *Base58Obj) Encode(input []byte) string {
	// Get alphabet
	alphabet, err := getAlphabet(obj.AlphIdx)
	if err != nil {
		return ""
	}

	// Create slice for encoded output
	enc := make([]byte, 0, getOutputLength(input))

	// Convert bytes to big integer
	encVal := byteSliceToBigInt(input)

	// Get encoding bytes from integer
	mod := new(big.Int)
	for encVal.Cmp(bigZero) > 0 {
		encVal.DivMod(encVal, bigRadix, mod)
		enc = append(enc, alphabet[mod.Int64()])
	}

	// Pad encoding depending on the number of initial zeros
	enc = padEncoding(enc, input, alphabet)

	// Reverse bytes slice
	reverseByteSlice(enc)

	// Convert to string
	return string(enc)
}

// Encode the specified bytes to Base58 format, by adding the checksum.
func (obj *Base58Obj) CheckEncode(input []byte) string {
	// Create slice for data with checksum
	dataWithChksum := make([]byte, 0, len(input) + checksumLen)
	dataWithChksum = append(dataWithChksum, input[:]...)

	// Compute checksum and append it
	chksum := computeCheckum(input)
	dataWithChksum = append(dataWithChksum, chksum[:]...)

	// Encode the final slice
	return obj.Encode(dataWithChksum)
}

//
// Not-exported functions
//

// Compute the Base58 output length from the input bytes.
// By definition, the output length is ~138% of input length.
func getOutputLength(slice []byte) int {
	return (len(slice) * 138 / 100) + 1
}

// Reverse the specified byte slice
func reverseByteSlice(slice []byte) {
	for i, j := 0, len(slice) - 1; i < j; i, j = i + 1, j - 1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Convert the specified byte slice to big.Int
func byteSliceToBigInt(slice []byte) (*big.Int) {
	bigVal := new(big.Int)
	bigVal.SetBytes(slice)

	return bigVal
}

// Pad encoding by adding the first alphabet letter as many times as the number of leading zeros in the original bytes.
func padEncoding(enc []byte, input []byte, alphabet string) []byte {
	for _, b := range(input) {
		if b != 0 {
			break
		}
		enc = append(enc, alphabet[0])
	}

	return enc
}
