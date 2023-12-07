package coneenv

import (
	"inspection/pkg/cone"
	"log"
	"os"
)

const EnvPrefix = "C1NS_"

const (
	flagApiKey    = EnvPrefix + "API_KEY"
	flagRegion    = EnvPrefix + "REGION"
	flagAccountID = EnvPrefix + "ACCOUNT_ID"
	flagAWSRegion = EnvPrefix + "AWS_REGION"
)

func NewCloudOneNS() *cone.CloudOneNS {
	return cone.NewCloudOneNS(
		GetEnv(flagApiKey),
		GetEnv(flagRegion),
		GetEnv(flagAccountID),
		GetEnv(flagAWSRegion),
	)
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s: missing environment variable", key)
	}
	return value
}
