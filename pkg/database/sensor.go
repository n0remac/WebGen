package database

func createSensorTable() {
	// SQL command to create the Sensor table
	sqlCommand := `
CREATE TABLE IF NOT EXISTS Sensors (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  data TEXT NOT NULL
);`

	createTable(sqlCommand)
}
