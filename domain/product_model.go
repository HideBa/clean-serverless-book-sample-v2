package domain

import "time"

type ProductModel struct {
	ID          int64
	Name        string
	Price       int
	ReleaseDate time.Time
}

func NewProductModel(name string, price int, releaseDate time.Time) *ProductModel {
	return &ProductModel{Name: name, Price: price, ReleaseDate: releaseDate}
}
