package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	httpMainnet := os.Getenv("MAINNET_RPC_URL")

	client, err := ethclient.Dial(httpMainnet)
	if err != nil {
		log.Fatal(err)
	}

	// LendingPool address
	address := common.HexToAddress("0x70b97A0da65C15dfb0FFA02aEE6FA36e507C2762")

	lendingVault, err := NewMain(address, client)
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
		FromBlock: big.NewInt(19625941),
		ToBlock:   big.NewInt(19625943),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		if vLog.Topics[0].Hex() == "0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba" {
			auction, err := lendingVault.ParseAuction(vLog)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(auction)
		}
	}
}
