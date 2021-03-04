package adapter

import (
	"clean-serverless-book-sample-v2/domain"
	"time"
)

type ProductResource struct {
	ResourceSchema
	DynamoResourceBase
	domain.ProductModel
	Mapper *DynamoModelMapper `dynamo:"-"`
}

func NewProductResource(
	productModel *domain.ProductModel, mapper *DynamoModelMapper) *ProductResource {
	return &ProductResource{
		ProductModel: *productModel, Mapper: mapper,
	}
}

func (p *ProductResource) EntityName() string {
	return p.Mapper.GetEntityNameFromStruct(*p)
}

func (p *ProductResource) PK() string {
	return p.Mapper.GetPK(p)
}

func (p *ProductResource) SetPK() {
	p.ResourceSchema.PK = p.PK()
}

func (p *ProductResource) SK() string {
	return p.Mapper.GetSK(p)
}

func (p *ProductResource) SetSK() {
	p.ResourceSchema.SK = p.SK()
}

func (p *ProductResource) ID() uint64 {
	return uint64(p.ProductModel.ID)
}

func (p *ProductResource) SetID(id uint64) {
	p.ProductModel.ID = id
}

func (p *ProductResource) Version() int {
	return p.DynamoResourceBase.Version
}

func (p *ProductResource) SetVersion(v int) {
	p.DynamoResourceBase.Version = v
}

func (p *ProductResource) CreatedAt() time.Time {
	return p.DynamoResourceBase.CreatedAt
}

func (p *ProductResource) SetCreatedAt(t time.Time) {
	p.DynamoResourceBase.CreatedAt = t
}

func (p *ProductResource) UpdatedAt() time.Time {
	return p.DynamoResourceBase.UpdatedAt
}

func (p *ProductResource) SetUpdatedAt(t time.Time) {
	p.DynamoResourceBase.UpdatedAt = t
}
