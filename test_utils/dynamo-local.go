package test_utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rs/xid"
)

// Creates a new, randomly named, Dynamo table using the interface provided
func CreateTable(table interface{}) (string, error) {
	cfg := aws.Config{
		Endpoint: aws.String("http://localhost:9000"),
		Region:   aws.String("eu-west-2"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	}
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &cfg)
	tableName := xid.New().String()
	err := db.CreateTable(tableName, table).Run()
	if err != nil {
		return "", err
	}
	return tableName, nil
}
