package database

func createProductTable() {
	// SQL command to create the Product table
	sqlCommand := `
CREATE TABLE IF NOT EXISTS Products (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  data TEXT NOT NULL
);`

	createTable(sqlCommand)
}
