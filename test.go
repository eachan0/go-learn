package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// bytes 包
	arrA := [...]byte{97, 98, 99}
	arrB := [...]byte{65, 66, 67}
	sA := arrA[:]
	sB := arrB[:]
	fmt.Println("EqualFold", bytes.EqualFold(sA, sB))
	prefixA := [1]byte{65}
	prefix := prefixA[:]
	fmt.Println("HasPrefix", bytes.HasPrefix(sA, prefix))
	fmt.Println("ContainsAny", bytes.ContainsAny(sA, "cde"))

	// crypto包
	str := "123456" // e10adc3949ba59abbe56e057f20f883e
	// e10adc3949ba59abbe56e057f20f883e
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))
}
