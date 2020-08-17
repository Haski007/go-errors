package errors

import (
	"fmt"
	"log"
)

const (
	// RED makes text in bold red!
	RED = "\033[1;31m"
	// GREEN makes text in bold green!
	GREEN = "\033[1;32m"
	// EOC makes text in default color!
	EOC = "\033[0m"
)

// MakeError returns error with red "ERROR: " sign
func MakeError(err error) error {
	return fmt.Errorf("%sERROR:%s %v", RED, EOC, err.Error())
}

// Println makes log.Println with "ERROR: " sing
func Println(err error) {
	log.Println(MakeError(err))
}