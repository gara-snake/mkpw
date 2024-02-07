package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const ALPHA_B = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHA_L = "abcdefghijklmnopqrstuvwxyz"
const MARK = "@_-+&%$#"

func main() {

	length := 16

	ret := ""

	for i := 0; i < length; i++ {

		t := random(3)

		switch t {
		case 0:
			ret = ret + getChar(ALPHA_B)
		case 1:
			ret = ret + getChar(ALPHA_L)
		case 2:
			ret = ret + getChar(MARK)
		}

	}

	print(ret)
}

func random(len int64) int64 {

	n, err := rand.Int(rand.Reader, big.NewInt(len))
	if err != nil {
		panic(err)
	}

	return n.Int64()
}

func getChar(str string) string {
	i := random(int64(len(str)))
	arr := strings.Split(str, "")
	c := arr[i]

	return c
}
