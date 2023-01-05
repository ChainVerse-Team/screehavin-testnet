package config

import (
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func GetEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func GetEnvInt(key string, fallback int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		valInt, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			color.Red(err.Error())
			return 0
		}
		return valInt
	}
	return fallback
}

type Config struct {
	AccountPrivateKey string
	Recipient1        string
	Recipient2        string
	Recipient3        string
	Recipient4        string
	NftHolderPriKey1  string
	NftHolderPriKey2  string
	NftHolderPriKey3  string
	NftHolderPriKey4  string
	NftContract       string
	StakingContract   string
	V1Address         string
	V2Address         string
	V3Address         string
	V4Address         string
	V1AddrPriKey      string
	V2AddrPriKey      string
	V3AddrPriKey      string
	V4AddrPriKey      string
	JsonRPC_url    string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	c := &Config{}
	c.StakingContract = GetEnv("STAKING_CONTRACT", "0x0000000000000000000000000000000000001001")
	c.JsonRPC_url = GetEnv("JSONRPC_URL", "http://103.138.113.121:8545/")
	return c
}