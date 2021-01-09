package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	secret := "bgvyzdsv"
	num := 0
	prefix := "111111"
	for prefix != "000000" {
		num++
		key := fmt.Sprintf("%s%d", secret, num)
		data := []byte(key)
		sum := md5.Sum(data)
		prefix = fmt.Sprintf("%x", sum)[0:6]
	}
	fmt.Printf("Answer: %d\n", num)
}
