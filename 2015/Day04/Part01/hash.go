package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	secretKey := "yzbqklnj" //  input
	number := 1

	for {
		data := secretKey + strconv.Itoa(number) // Combine  secret key + current number

		hash := md5.Sum([]byte(data)) // Compute the MD5 hash

		hexString := hex.EncodeToString(hash[:]) // Convert the hash to a hexadecimal string

		// Check if the first five characters are zeroes

		if hexString[:5] == "00000" {
			fmt.Printf("The lowest number is %d and its hash is %s\n", number, hexString)
			break
		}

		number++ // ++ up to 282749

	}
}

/*


number++  up to :   282749
last data 'll be:  yzbqklnj282749
hash :             [0 0 2 198 85 223 119 56 36 110 136 246 193 196 62 183]
hexString :         000002c655df7738246e88f6c1c43eb7



for loop iterate over numbers starting from 1.
For each number, it appends the number to the secret key,
 computes the MD5 hash of the resulting string, and then converts the hash to a hexadecimal string.
 It checks if the first five characters of the hexadecimal string are zeroes.
If they are, it prints the number and its hash; otherwise, it increments the number and continues the loop.
*/
