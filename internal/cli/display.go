package cli

import "fmt"

func Display(response string) {
	//TODO: parse RESP
	fmt.Println(response)
}

func ShowError(err error) {
	fmt.Println(err.Error())
}
