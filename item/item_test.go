package item

import (
	"fmt"
	"github.com/nathmclean/dynamodb-local-testing/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestItemService_CreateItem(t *testing.T) {
	cases := []struct {
		name string
		item *Item
		err  bool // Whether we expect an error back or not
	}{
		{
			name: "created successfully",
			item: &Item{
				Name:        "spoon",
				Description: "shiny",
			},
		},
	}

	service, err := newItemService()
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := service.CreateItem(c.item)
			if c.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, time.Time{}, c.item.CreatedAt)
				assert.NotEqual(t, time.Time{}, c.item.UpdatedAt)
			}
		})
	}
}

func TestItemService_GetItem(t *testing.T) {
	cases := []struct {
		name   string
		item   *Item
		create bool
		err    bool
	}{
		{
			name: "success",
			item: &Item{
				Name:        "spoon",
				Description: "shiny",
			},
			create: true,
			err:    false,
		},
		{
			name: "item does not exist",
			item: &Item{
				Id: "idontexist",
			},
			err: true,
		},
	}

	service, err := newItemService()
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.create {
				err := service.CreateItem(c.item)
				if err != nil {
					t.Fatal(err)
				}
			}
			retrievedItem := &Item{
				Id: c.item.Id,
			}
			err := service.GetItem(retrievedItem)
			if c.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.item.Name, retrievedItem.Name)
			}
		})
	}
}

// newItemService sets up a new ItemService with new, randomly named Dynamo table, so that our tests don't collide.
func newItemService() (*ItemService, error) {
	tableName, err := test_utils.CreateTable(Item{})
	if err != nil {
		return nil, fmt.Errorf("failed to set up table. %s", err)
	}

	db, err := newDynamoTable(tableName, "http://localhost:9000")
	if err != nil {
		return nil, err
	}
	service := ItemService{
		itemTable: db,
	}
	return &service, nil
}
