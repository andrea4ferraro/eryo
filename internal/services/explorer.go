package services

import (
	"fmt"

	"blockchain-explorer-cli/internal/commands"
)

func Execute(command string, value string) {

	switch command {

	case "address":

		commands.Address(value)

	case "tx":

		commands.Tx(value)

	case "block":

		commands.Block(value)

	case "token":

		commands.Token(value)

	default:

		fmt.Println("Unknown command")

	}

}
