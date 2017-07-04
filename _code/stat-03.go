// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to calculate standard deviation and variance.
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

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n\n", stdDevVal)
}
