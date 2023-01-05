package types

import (
	"encoding/hex"
	"strings"
)

const (
	AddressLength = 20
)

type Address [AddressLength]byte

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func BytesToAddress(b []byte) Address {
	var a Address

	size := len(b)
	min := min(size, AddressLength)

	copy(a[AddressLength-min:], b[len(b)-min:])

	return a
}

func stringToBytes(str string) []byte {
	str = strings.TrimPrefix(str, "0x")
	if len(str)%2 == 1 {
		str = "0" + str
	}

	b, _ := hex.DecodeString(str)

	return b
}

func StringToAddress(str string) Address {
	return BytesToAddress(stringToBytes(str))
}

