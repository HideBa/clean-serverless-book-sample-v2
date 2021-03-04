package adapter

import (
	"clean-serverless-book-sample-v2/domain"

	"github.com/memememomo/nomof"
	"github.com/pkg/errors"
)

type ProductOperator struct {
	Client *ResourceTableOperator
	Mapper *DynamoModelMapper
}

func (p *ProductOperator) CreateProduct(productModel *domain.ProductModel) (*domain.ProductModel, error) {
	productResource := NewProductResource(productModel, p.Mapper)

	err := p.Mapper.PutResource(productResource)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &productResource.ProductModel, nil
}

func (p *ProductOperator) GetProducts() ([]*domain.ProductModel, error) {
	table, err := p.Client.ConnectTable()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fb := nomof.NewBuilder()
	fb.BeginsWith("PK", p.Mapper.GetEntityNameFromStruct(ProductResource{}))

	var productResource []ProductResource
	err = table.Scan().Filter(fb.JoinAnd(), fb.Arg...).All(&productResource)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var products = make([]*domain.ProductModel, len(productResource))
	for i := range productResource {
		products[i] = &productResource[i].ProductModel
	}
	return products, nil
}
