package s3

import (
	"context"
	"docto/constants"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Connect() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion(constants.S3_REGION))
	if err != nil {
		log.Fatal(err)
	}

	return s3.NewFromConfig(cfg)
}
