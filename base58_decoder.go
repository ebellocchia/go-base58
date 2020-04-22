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
// This file cointains decoding functions for base58 package.
//

package base58

//
// Imports
//
import (
	"bytes"
	"errors"
	"math/big"
	"strings"
)

//
// Variables
//
var (
	// ErrInvalidFormat is returned when trying to decode a string in invalid format
	ErrInvalidFormat = errors.New("The specified string is not a valid Base58 format")
	// ErrInvalidChecksum is returned when trying to decode a string with invalid checksum
	ErrInvalidChecksum = errors.New("The checksum of the specified string is not valid")
)

//
// Exported functions
//

// Decode the specified string in Base58 format to bytes.
func (obj *Base58Obj) Decode(input string) ([]byte, error) {
	// Get alphabet
	alphabet, err := getAlphabet(obj.AlphIdx)
	if err != nil {
		return nil, err
	}

	decVal := big.NewInt(0)

	// Convert the string back to big integer
	mult := big.NewInt(1)
	tmp  := new(big.Int)

	for i := len(input) - 1; i >= 0; i-- {
		// Find character in the alphabet
		chrIdx := strings.IndexByte(alphabet, input[i])
		// Format error if not found
		if chrIdx == -1 {
			return nil, ErrInvalidFormat
		}
		// Update value: val += mult * chrIdx
		tmp.SetInt64(int64(chrIdx))
		tmp.Mul(mult, tmp)
		decVal.Add(decVal, tmp)
		// Increase multiplier: mult = mult * 58
		mult.Mul(mult, bigRadix)
	}

	// Pad decoding depending on the number of the first alphabet letter
	dec := padDecoding(decVal.Bytes(), input, alphabet)

	return dec, nil
}

// Decode the specified string in Base58 format to bytes, by removing and verifying the checksum.
func (obj *Base58Obj) CheckDecode(input string) ([]byte, error) {
	// Decode string
	dec, err := obj.Decode(input)
	if err != nil {
		return nil, err
	}

	// Get data and checksum parts
	chksumIdx := len(dec) - checksumLen
	chksumPart, dataPart := dec[chksumIdx:], dec[:chksumIdx]

	// Compute again checksum on data
	compChksum := computeCheckum(dataPart)

	// Verify checksum
	if bytes.Compare(chksumPart, compChksum[:]) != 0 {
		return nil, ErrInvalidChecksum
	}

	return dataPart, nil
}

//
// Not-exported functions
//

// Pad decoding by adding zeros as many times as the number of leading first alphabet characters in the original string.
func padDecoding(dec []byte, input string, alphabet string) []byte {
	// Compute the number of zeros to be added
	zerosCnt := countLeadingFirstAlphChar(input, alphabet)

	// Pad bytes with zeros
	decPadded := make([]byte, len(dec) + zerosCnt)
	copy(decPadded[zerosCnt:], dec)

	return decPadded
}

// Count the number of leading first alphabet characters.
func countLeadingFirstAlphChar(input string, alphabet string) int {
	var charCnt int
	for charCnt = 0; charCnt < len(input); charCnt++ {
		if input[charCnt] != alphabet[0] {
			break;
		}
	}

	return charCnt
}
