package main

import (
	"fmt"
	"os"
)

// TODO lifecycle:
//  https://github.com/google/wire/issues/107
//  `I don't think the Wire tool needs to have explicit support for this.`
//  https://github.com/google/wire/issues/107#issuecomment-455257936

func main() {
	sigChan := make(chan os.Signal, 1)

	service, err := WireService()
	fmt.Println(service)
	fmt.Println(err)

	<-sigChan
}
