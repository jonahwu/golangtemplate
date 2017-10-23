package main

import (
	"fmt"
)

func retmore() (string, error) {
	return "aa", nil
}

func main() {
	if ret, err := retmore(); err != nil {
		fmt.Println(err)
	} else {
		//process ret
		fmt.Println(ret)
	}
}
