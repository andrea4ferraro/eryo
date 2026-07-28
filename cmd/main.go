package main

import (
	"fmt"
	"os"

	"blockchain-explorer-cli/internal/services"
)

func main() {

	if len(os.Args) < 3 {

		fmt.Println("Usage:")
		fmt.Println("address <id>")
		fmt.Println("tx <hash>")
		fmt.Println("block <number>")
		fmt.Println("token <symbol>")

		return
	}

	services.Execute(os.Args[1], os.Args[2])

}
