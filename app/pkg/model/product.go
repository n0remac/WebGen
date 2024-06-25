package model

import (
	"encoding/json"
	"CodeGen/gen/proto/product"
	"CodeGen/pkg/database"

	"github.com/upper/db/v4"
)

type Product struct {
	ID   int    `db:"id,omitempty"`
	Data string `db:"data"`
}

func CreateProduct(m *product.Product) (*product.Product, error) {
	sess := database.GetSession()

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	newProduct := &Product{
		Data: string(data),
	}

	err = sess.Collection("products").InsertReturning(newProduct)
	if err != nil {
		return nil, err
	}

	m.Id = int32(newProduct.ID)
	return m, nil
}

func GetProductFromDB(id int32) (*product.Product, error) {
	sess := database.GetSession()
	var dbProduct Product

	res := sess.Collection("products").Find(db.Cond{"id": id})
	err := res.One(&dbProduct)
	if err != nil {
		return nil, err
	}

	var m product.Product
	err = json.Unmarshal([]byte(dbProduct.Data), &m)
	if err != nil {
		return nil, err
	}

	m.Id = int32(dbProduct.ID)
	return &m, nil
}

func UpdateProductInDB(m *product.Product) (*product.Product, error) {
	sess := database.GetSession()
	var dbProduct Product

	res := sess.Collection("products").Find(db.Cond{"id": m.Id})
	err := res.One(&dbProduct)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	updatedProduct := &Product{
		ID:   dbProduct.ID,
		Data: string(data),
	}
	err = res.Update(updatedProduct)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func DeleteProductFromDB(id int32) error {
	sess := database.GetSession()

	res := sess.Collection("products").Find(db.Cond{"id": id})
	err := res.Delete()
	return err
}
