package main

import (
	"fmt"
	"os"
)

func main() {
	sigChan := make(chan os.Signal, 1)

	service, err := WireService()
	fmt.Println(service)
	fmt.Println(err)

	service.Start()
	<-sigChan
	service.Stop()

}
