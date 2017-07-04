// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to calculate both central tendency and statistical dispersion
// measures for the iris dataset.
package main

import (
	"log"

	"go-hep.org/x/hep/csvutil/csvdriver"
)

func main() {

	// Open the CSV file as a database table.
	db, err := csvdriver.Conn{
		File:    "data/iris.csv",
		Comment: '#',
		Comma:   ',',
	}.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a database transaction.
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()

	// Define a SQL query that we will execute against the CSV file.
	// Only select float columns
	query := `SELECT var1, var2, var3, var4 FROM csv`

	// Query the CSV via the SQL statement.
	rows, err := tx.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop over the float columns.

	// Get the float values from the column.

	// Calculate the Mean of the variable.

	// Calculate the Mode of the variable.

	// Calculate the Median of the variable.

	// Calculate the variance of the variable.

	// Calculate the standard deviation of the variable.

	// Output the results to standard out.
}
