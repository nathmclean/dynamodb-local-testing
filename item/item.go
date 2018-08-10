package item

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rs/xid"
	"time"
)

// Item represents the schema for the Dynamo Table
type Item struct {
	Id          string    `dynamo:"item_id,hash"`
	Name        string    `dynamo:"name"`
	Description string    `dynamo:"description"`
	CreatedAt   time.Time `dynamo:"created_at"`
	UpdatedAt   time.Time `dynamo:"updated_at"`
}

// ItemService holds out dynamo client
type ItemService struct {
	itemTable dynamo.Table
}

// NewItemService creates a new item service with a dynamo client setup to talk to the provided table name
func NewItemService(itemTableName string) (*ItemService, error) {
	dynamoTable, err := newDynamoTable(itemTableName, "")
	if err != nil {
		return nil, err
	}
	return &ItemService{
		itemTable: dynamoTable,
	}, nil
}

// newDynamoTable creates a client that can interact with a specific Dynamo Table. Can optionally be setup to talk to dynamo
// local, useful for testing.
func newDynamoTable(tableName, endpoint string) (dynamo.Table, error) {
	if tableName == "" {
		return dynamo.Table{}, fmt.Errorf("you must supply a table name")
	}
	cfg := aws.Config{}
	cfg.Region = aws.String("eu-west-2")
	if endpoint != "" {
		cfg.Endpoint = aws.String(endpoint)
	}
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &cfg)
	table := db.Table(tableName)
	return table, nil
}

// CreateItem puts a new item into Dynamo. Generates an ID and sets the CreatedAt and UpdatedAt values.
// Performs no validation
func (i *ItemService) CreateItem(item *Item) error {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	item.Id = xid.New().String()
	return i.itemTable.Put(item).Run()
}

// GetItem Gets an item from DynamoDb. Performs no validation, such as if the ID is set.
func (i *ItemService) GetItem(item *Item) error {
	return i.itemTable.Get("item_id", item.Id).One(item)
}
