package database

func create{{.ModelName}}Table() {
	// SQL command to create the {{.ModelName}} table
	sqlCommand := `
CREATE TABLE IF NOT EXISTS {{.ModelName | title}}s (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  data TEXT NOT NULL
);`

	createTable(sqlCommand)
}
