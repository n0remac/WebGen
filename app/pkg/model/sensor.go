package model

import (
	"encoding/json"
	"CodeGen/gen/proto/sensor"
	"CodeGen/pkg/database"

	"github.com/upper/db/v4"
)

type Sensor struct {
	ID   int    `db:"id,omitempty"`
	Data string `db:"data"`
}

func CreateSensor(m *sensor.Sensor) (*sensor.Sensor, error) {
	sess := database.GetSession()

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	newSensor := &Sensor{
		Data: string(data),
	}

	err = sess.Collection("sensors").InsertReturning(newSensor)
	if err != nil {
		return nil, err
	}

	m.Id = int32(newSensor.ID)
	return m, nil
}

func GetSensorFromDB(id int32) (*sensor.Sensor, error) {
	sess := database.GetSession()
	var dbSensor Sensor

	res := sess.Collection("sensors").Find(db.Cond{"id": id})
	err := res.One(&dbSensor)
	if err != nil {
		return nil, err
	}

	var m sensor.Sensor
	err = json.Unmarshal([]byte(dbSensor.Data), &m)
	if err != nil {
		return nil, err
	}

	m.Id = int32(dbSensor.ID)
	return &m, nil
}

func UpdateSensorInDB(m *sensor.Sensor) (*sensor.Sensor, error) {
	sess := database.GetSession()
	var dbSensor Sensor

	res := sess.Collection("sensors").Find(db.Cond{"id": m.Id})
	err := res.One(&dbSensor)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	updatedSensor := &Sensor{
		ID:   dbSensor.ID,
		Data: string(data),
	}
	err = res.Update(updatedSensor)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func DeleteSensorFromDB(id int32) error {
	sess := database.GetSession()

	res := sess.Collection("sensors").Find(db.Cond{"id": id})
	err := res.Delete()
	return err
}
