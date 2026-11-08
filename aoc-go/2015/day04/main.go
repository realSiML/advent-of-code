package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const INPUT string = "ckczppom"

func ValidateHash(hash []byte) bool {
	return strings.HasPrefix(hex.EncodeToString(hash), "000000")
}

func main() {
	i := 1
	for {
		data := []byte(INPUT + strconv.Itoa(i))
		hash := md5.Sum(data)
		if ValidateHash(hash[:]) {
			fmt.Println(string(data), hex.EncodeToString(hash[:]))
			break
		}

		i += 1
	}
}
