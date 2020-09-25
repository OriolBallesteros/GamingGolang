package errhndl

import "fmt"

func ErrHndl(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}
