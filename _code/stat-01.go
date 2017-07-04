// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to calculate means, modes, and medians.
package main

import (
	"fmt"
	"log"

	"go-hep.org/x/hep/csvutil/csvdriver"
	"gonum.org/v1/gonum/stat"
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
	query := `SELECT var1 FROM csv`

	// Query the CSV via the SQL statement.  Here we will only get
	// the petal length, petal width, and species for all rows
	// where the species is "Iris-versicolor".
	rows, err := tx.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var sepalLength []float64

	// Output the results of the query to standard out.
	for rows.Next() {
		var sepal float64
		if err = rows.Scan(&sepal); err != nil {
			log.Fatal(err)
		}
		sepalLength = append(sepalLength, sepal)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(sepalLength, nil)

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
}
