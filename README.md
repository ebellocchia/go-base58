# Base58 package
[![Build Status](https://travis-ci.com/ebellocchia/go-base58.svg?branch=master)](https://travis-ci.com/ebellocchia/go-base58)
[![codecov](https://codecov.io/gh/ebellocchia/go-base58/branch/master/graph/badge.svg)](https://codecov.io/gh/ebellocchia/go-base58)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/ebellocchia/go-base58/master/LICENSE)

## Introduction

This simple package implements the Base58 encoding/decoding, with and without checksum.

## Installation

The package can be installed by simply running:

    go get -u github.com/ebellocchia/go-base58

## Usage

First of all, a base58 object shall be created using the *New* function, by specifying the alphabet to be used.\
The possible alphabets are:
- Bitcoin: *base58.AlphabetBitcoin*
- Ripple: *base58.AlphabetRipple*
- Flickr: *base58.AlphabetFlickr*

If the object is created without using the *New* function, the Bitcoin alphabet will be used by default.\
There are 4 APIs that can be used:
- *Encode([]byte) string*: encode bytes into string
- *CheckEncode([]byte) string*: encode bytes into string with checksum
- *Decode(string) ([]byte, error)*: decode string back into bytes, return error if format is not valid
- *CheckDecode(string) ([]byte, error)*: decode string with checksum back into bytes, return error if format or checksum is not valid

**Example**

    package main

    import (
      "github.com/ebellocchia/go-base58"
      "fmt"
      "encoding/hex"
    )

    func main() {
        // Create base58 object using Bitcoin alphabet
        base58Btc := base58.New(base58.AlphabetBitcoin)

        // Some bytes to encode
        data_bytes, _ := hex.DecodeString("bf4f89001e670274dd")

        // Encode
        // It returns an empty string if the alphabet is not valid
        enc := base58Btc.Encode(data_bytes)
        fmt.Printf("Encode: %s\n", enc)
        // CheckEncode
        // It returns an empty string if the alphabet is not valid
        check_enc := base58Btc.CheckEncode(data_bytes)
        fmt.Printf("Checksum encode: %s\n", check_enc)

        // Decode
        // It returns error if the alphabet or the string format is not valid
        dec, err := base58Btc.Decode(enc)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Decode: %s\n", hex.EncodeToString(dec))
        // CheckDecode
        // It returns error if the alphabet, the string format or the checksum is not valid
        check_dec, err := base58Btc.CheckDecode(check_enc)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Checksum decode: %s\n", hex.EncodeToString(check_dec))
    }


## License

This software is available under the MIT license.
