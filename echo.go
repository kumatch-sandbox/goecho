package goecho

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	fmt.Println(getArgsString())
}

func getArgsString() string {
	return strings.Join(os.Args[1:], " ")
}
