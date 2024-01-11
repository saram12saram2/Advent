package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	secretKey := "yzbqklnj"
	number := 1

	for {
		data := secretKey + strconv.Itoa(number)

		hash := md5.Sum([]byte(data))

		hexString := hex.EncodeToString(hash[:])

		// if the first 6 characters are zeroes
		if hexString[:6] == "000000" {
			fmt.Printf("The lowest number is %d and its hash is %s\n", number, hexString)
			break
		}

		number++
	}
}
