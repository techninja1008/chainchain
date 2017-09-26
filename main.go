package main

import (
	"log"
	"os"

	log15 "github.com/ethereum/go-ethereum/log"
	"github.com/techninja1008/chainchain/cmd"
)

func main() {
	log15.Root().SetHandler(log15.StdoutHandler)
	if err := cmd.CCHCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
