package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	length := getLength()
	str := bruteForce(length)
	fmt.Println(str)
}

func bruteForce(length int) string {
	charlist := "QWERTYUIOPASDFGHJKLZXCVBNM"
	password := ""
	for i := 0; i < length; i++ {
		var delta int64 = 0
		char := ""
		for _, c := range charlist {
			cmd := exec.Command("./crypto", password+string(c))
			before := time.Now()
			cmd.Run()
			del := time.Now().Sub(before)
			if int64(del) > delta {
				char = string(c)
				delta = int64(del)
			}
		}
		password += char
	}
	return password
}

func getLength() int {
	length := 0
	str := ""
	for {
		cmd := exec.Command("./crypto", str)
		//err := syscall.Exec("crypto", []string{"crypto", str}, nil)
		err := cmd.Run()
		if strings.Contains(err.Error(), "2") {
			return length - 1
		}
		str += "a"
		length++
	}
}
