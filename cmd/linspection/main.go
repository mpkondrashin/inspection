package main

import (
	"context"
	"inspection/pkg/cone"
	"inspection/pkg/coneenv"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(hello)
}

func hello() (string, error) {
	cOne := coneenv.NewCloudOneNS()
	inspection := cone.ActionBypass
	status, err := cOne.GetInspectionBypassStatus(context.TODO(), GetEnv(flagAWSRegion))
	errMessage := ""
	if err != nil {
		errMessage += "(" + err.Error() + ") "
	} else {
		if status.Action == cone.ActionBypass {
			inspection = cone.ActionInspect
		}
	}
	err = cOne.SetInspectionBypass_(context.TODO(), GetEnv(flagAWSRegion), inspection)
	if err != nil {
		return "", err
	}
	return errMessage + inspection.String(), nil
}

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
		//GetEnv(flagAWSRegion),
	)
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s: missing environment variable", key)
	}
	return value
}
