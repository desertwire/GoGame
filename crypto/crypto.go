package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

//Alias for shorter print function name
var print   = 	fmt.Print
var println = 	fmt.Println
var printf 	= 	fmt.Printf
var scanf	= 	fmt.Scanf

func MD5Checksum(message string) {
	data := []byte(message)
	fmt.Printf("%x", md5.Sum(data))
}

func Sha1Checksum(message string) {
	//s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(message))
	bs := h.Sum(nil)
	println(message)
	printf("%x\n", bs)
}

func Sha256Checksum(message string) {
	sum := sha256.Sum256([]byte(message))
	fmt.Printf("%x", sum)
}