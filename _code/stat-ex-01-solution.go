// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to calculate both central tendency and statistical dispersion
// measures for the iris dataset.
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
	// Only select float columns
	query := `SELECT var1, var2, var3, var4 FROM csv`

	// Query the CSV via the SQL statement.
	rows, err := tx.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var vars [4][]float64

	for rows.Next() {
		var v1, v2, v3, v4 float64
		if err = rows.Scan(&v1, &v2, &v3, &v4); err != nil {
			log.Fatal(err)
		}
		vars[0] = append(vars[0], v1)
		vars[1] = append(vars[1], v2)
		vars[2] = append(vars[2], v3)
		vars[3] = append(vars[3], v4)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Loop over the float columns.
	for i, v := range vars {
		// Calculate the Mean of the variable.
		meanVal := stat.Mean(v, nil)

		// Calculate the Mode of the variable.
		modeVal, modeCount := stat.Mode(v, nil)

		// Calculate the variance of the variable.
		varianceVal := stat.Variance(v, nil)

		// Calculate the standard deviation of the variable.
		stdDevVal := stat.StdDev(v, nil)

		// Output the results to standard out.
		fmt.Printf("\nvar-%d Summary Statistics:\n", i+1)
		fmt.Printf("Mean value: %0.2f\n", meanVal)
		fmt.Printf("Mode value: %0.2f\n", modeVal)
		fmt.Printf("Mode count: %d\n", int(modeCount))
		fmt.Printf("Variance value: %0.2f\n", varianceVal)
		fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	}
}
