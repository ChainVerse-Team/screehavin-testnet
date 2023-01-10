package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

type Config struct {
	StakingContract   string
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