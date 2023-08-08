package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

var solidityInterface []byte

var parsedSolidityInterface map[string]bind.CallOpts

type HexString = string
type InvokeSolidityContractInput struct {
	NodeURL       string                 `json:"NodeURL"`       // url to the ethereum node
	ABIFilePath   string                 `json:"ABIFilePath"`      // filePath to the abi file
	ContractName  string                 `json:"ContractName"`  // name of the contract you want to invoke
	Input         map[string]interface{} `json:"Input"`         // this input is going to get encoded into abi
	From          HexString              `json:"From"`          // the sender of the 'transaction'
	To            HexString              `json:"To"`            // the destination contract (nil for contract creation)
	Gas           uint64                 `json:"Gas"`           // if 0, the call executes with near-infinite gas
	GasPrice      int64                  `json:"GasPrice"`      // wei <-> gas exchange ratio
	GasFeeCap     int64                  `json:"GasFeeCap"`     // EIP-1559 fee cap per gas.
	GasTipCap     int64                  `json:"GasTipCap"`     // EIP-1559 tip per gas.
	Value         int64                  `json:"Value"`         // amount of wei sent along with the call
}

func (a *App) InvokeSolidityContract(name string, input InvokeSolidityContractInput) string {
	from, to := common.HexToAddress(input.From), common.HexToAddress(input.To)

	client, err := ethclient.Dial(input.NodeURL)
	if err != nil {
		return err.Error()
	}

	blockN, err := client.BlockNumber(NewApp().ctx)
	if err != nil {
		return err.Error()
	}

	data, check := abi.ABI.Methods[input.ContractName]
	data ,err := abi.ABI.Pack(name, input.Input)
	if err != nil {
		return err.Error()
	}

	types.NewTransaction(blockN)
	types.SignNewTx()
	client.TransactionReceipt()
	types.SignNewTx()
	err := client.SendTransaction(NewApp().ctx, &)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) DialToEthereumNode(name string) {

}

func (a *App) stuff() {
	r, err := os.Open("name")
	if err != nil {
		log.Fatal(err)
		return
	}

	abi, err := abi.JSON(r)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(abi.Constructor.RawName)
}

func (a *App) returnAllMethods(filepath string) string {
	errHandle := func(err error) string {
		log.Println(log.Flags(), err)
		return ""
	}
	r, err := os.Open(filepath)
	defer r.Close()
	if err != nil {
		return errHandle(err)
	}

	abi, err := abi.JSON(r)
	if err != nil {
		return errHandle(err)
	}

	data, err := json.Marshal(abi.Methods)
	if err != nil {
		return errHandle(err)
	}

	return string(data)
}
