package model

import (
	"encoding/json"
	"SchemaGen/gen/proto/schema"
	"SchemaGen/pkg/database"

	"github.com/upper/db/v4"
)

type Schema struct {
	ID   int    `db:"id,omitempty"`
	Data string `db:"data"`
}

func CreateSchema(m *schema.Schema) (*schema.Schema, error) {
	sess := database.GetSession()

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	newSchema := &Schema{
		Data: string(data),
	}

	err = sess.Collection("schemas").InsertReturning(newSchema)
	if err != nil {
		return nil, err
	}

	m.Id = int32(newSchema.ID)
	return m, nil
}

func GetSchemaFromDB(id int32) (*schema.Schema, error) {
	sess := database.GetSession()
	var dbSchema Schema

	res := sess.Collection("schemas").Find(db.Cond{"id": id})
	err := res.One(&dbSchema)
	if err != nil {
		return nil, err
	}

	var m schema.Schema
	err = json.Unmarshal([]byte(dbSchema.Data), &m)
	if err != nil {
		return nil, err
	}

	m.Id = int32(dbSchema.ID)
	return &m, nil
}

func UpdateSchemaInDB(m *schema.Schema) (*schema.Schema, error) {
	sess := database.GetSession()
	var dbSchema Schema

	res := sess.Collection("schemas").Find(db.Cond{"id": m.Id})
	err := res.One(&dbSchema)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	updatedSchema := &Schema{
		ID:   dbSchema.ID,
		Data: string(data),
	}
	err = res.Update(updatedSchema)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func DeleteSchemaFromDB(id int32) error {
	sess := database.GetSession()

	res := sess.Collection("schemas").Find(db.Cond{"id": id})
	err := res.Delete()
	return err
}
