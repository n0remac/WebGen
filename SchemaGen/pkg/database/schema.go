package database

func createSchemaTable() {
	// SQL command to create the Schema table
	sqlCommand := `
CREATE TABLE IF NOT EXISTS Schemas (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  data TEXT NOT NULL
);`

	createTable(sqlCommand)
}
