package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	r, err := os.Open("test/test1.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	abi, err := abi.JSON(r)
	if err != nil {
		log.Fatal(err)
		return
	}

	spew.Dump(abi)

	for key, val := range abi.Methods {
		fmt.Println("====================================")
		spew.Dump(key)
		fmt.Println("input ====================================")
		spew.Dump(val.Inputs)
		fmt.Println("outputs ====================================")
		spew.Dump(val.Outputs)
	}
}
