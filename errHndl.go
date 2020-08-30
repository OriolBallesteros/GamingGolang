package main

import "fmt"

func errHndl(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}
