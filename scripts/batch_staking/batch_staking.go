package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"screehavin-testnet/scripts/build/staking"
	"screehavin-testnet/scripts/config"
	"screehavin-testnet/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
)

var (
	privateKeysTransferTo  = "./scripts/keys.txt"
	JsonrpcUrl             string
	ErrCastPubKeyToECDSA   = errors.New("error casting public key to ECDSA")
	StakingContractAddress common.Address
)

func initRPC(c *config.Config) {
	JsonrpcUrl = c.JsonRPC_url
}

func initConfig() {
	//load dot env
	c := config.NewConfig()
	initRPC(c)
	StakingContractAddress = common.Address(types.StringToAddress(c.StakingContract))
}

func main() {
	initConfig()

	runStaking()
}

func runStaking() {
	client, err := ethclient.Dial(JsonrpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	keys := loadPriKeys()

	err = batchStaking(keys, client)
	if err != nil {
		log.Fatal(err)
	}
}

func batchStaking(priKeys []string, client *ethclient.Client) error {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	instance, err := staking.NewStaking(StakingContractAddress, client)
	if err != nil {
		return err
	}

	for _, key := range priKeys {
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			return err
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return ErrCastPubKeyToECDSA
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return err
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			return err
		}

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(1e18) // stake 1 ether
		auth.Value = auth.Value.Mul(auth.Value, big.NewInt(10))
		auth.GasLimit = uint64(1e12)
		auth.GasPrice = gasPrice

		// TODO
		tx, err := instance.Stake(auth, "ParthLabs")
		if err != nil {
			return err
		}

		color.Yellow("tx sent: %s\n", tx.Hash().Hex())
		color.Blue("Staked Successfully !!")
	}

	return nil
}

// loadPriKeys reads lines from a file text
func loadPriKeys() []string {
	var arr []string
	body, err := ioutil.ReadFile(privateKeysTransferTo)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))
	arr = append(arr, string(body))
	return arr
}
