package adapter

import (
	"clean-serverless-book-sample-v2/adapter"
	"clean-serverless-book-sample-v2/domain"
	"clean-serverless-book-sample-v2/mocks"
	"clean-serverless-book-sample-v2/registry"
	"fmt"
	"testing"
	"time"

	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestProductOperator_CreateProduct(t *testing.T) {
	tables := mocks.SetupDB(t)
	defer tables.Cleanup()

	product := domain.NewProductModel("テスト", 100, time.Now())

	operator := registry.GetFactory().BuildProductOperator()

	_, err := operator.CreateProduct(product)
	assert.NoError(t, err)

	result, err := getProductResource(1)
	assert.NoError(t, err)

	assert.Equal(t, product.Name, result.Name)
	assert.Equal(t, product.Price, result.Price)
	assert.Equal(t, product.ReleaseDate.Format(DatetimeFormat), result.ReleaseDate.Format(DatetimeFormat))
}

func getProductResource(id uint64) (*adapter.ProductResource, error) {
	client := registry.GetFactory().BuildResourceTableOperator()

	table, err := client.ConnectTable()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result adapter.ProductResource
	err = table.Get("PK", fmt.Sprintf("ProductResource-%011d", id)).Range("SK", dynamo.Equal, fmt.Sprintf("%011d", id)).One(&result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &result, nil
}
