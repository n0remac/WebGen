package model

import (
	"encoding/json"
	"{{.ProjectName}}/gen/proto/{{.ModelName | lower}}"
	"{{.ProjectName}}/pkg/database"

	"github.com/upper/db/v4"
)

type {{.ModelName}} struct {
	ID   int    `db:"id,omitempty"`
	Data string `db:"data"`
}

func Create{{.ModelName}}(m *{{.ModelName | lower}}.{{.ModelName}}) (*{{.ModelName | lower}}.{{.ModelName}}, error) {
	sess := database.GetSession()

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	new{{.ModelName}} := &{{.ModelName}}{
		Data: string(data),
	}

	err = sess.Collection("{{.ModelName | lower}}s").InsertReturning(new{{.ModelName}})
	if err != nil {
		return nil, err
	}

	m.Id = int32(new{{.ModelName}}.ID)
	return m, nil
}

func Get{{.ModelName}}FromDB(id int32) (*{{.ModelName | lower}}.{{.ModelName}}, error) {
	sess := database.GetSession()
	var db{{.ModelName}} {{.ModelName}}

	res := sess.Collection("{{.ModelName | lower}}s").Find(db.Cond{"id": id})
	err := res.One(&db{{.ModelName}})
	if err != nil {
		return nil, err
	}

	var m {{.ModelName | lower}}.{{.ModelName}}
	err = json.Unmarshal([]byte(db{{.ModelName}}.Data), &m)
	if err != nil {
		return nil, err
	}

	m.Id = int32(db{{.ModelName}}.ID)
	return &m, nil
}

func Update{{.ModelName}}InDB(m *{{.ModelName | lower}}.{{.ModelName}}) (*{{.ModelName | lower}}.{{.ModelName}}, error) {
	sess := database.GetSession()
	var db{{.ModelName}} {{.ModelName}}

	res := sess.Collection("{{.ModelName | lower}}s").Find(db.Cond{"id": m.Id})
	err := res.One(&db{{.ModelName}})
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	updated{{.ModelName}} := &{{.ModelName}}{
		ID:   db{{.ModelName}}.ID,
		Data: string(data),
	}
	err = res.Update(updated{{.ModelName}})
	if err != nil {
		return nil, err
	}

	return m, nil
}

func Delete{{.ModelName}}FromDB(id int32) error {
	sess := database.GetSession()

	res := sess.Collection("{{.ModelName | lower}}s").Find(db.Cond{"id": id})
	err := res.Delete()
	return err
}
