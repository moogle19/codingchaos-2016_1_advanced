package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const key = "chaostreffosnabrueck"
const a = 97

func main() {
	cleartext, _ := ioutil.ReadFile(os.Args[1])

	outcipher := ""
	for i := 0; i < len(cleartext); i++ {
		val := cleartext[i] - 97
		if val >= 0 && val < 26 {
			sub := (key[i%len(key)] - 97)
			for sub > val {
				val += 26
			}
			val -= sub
			val %= 26
		}
		val += 97
		outcipher += string(val)
	}
	fmt.Println(outcipher)
}
