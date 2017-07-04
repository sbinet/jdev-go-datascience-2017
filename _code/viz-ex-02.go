// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template2

// Sample program to generate a histogram of diabetes bmi values.
package main

import (
	"log"

	"go-hep.org/x/hep/csvutil"
)

func main() {

	fileName := "data/diabetes.csv"
	tbl, err := csvutil.Open(fileName)
	if err != nil {
		log.Fatalf("could not open %s: %v\n", fileName, err)
	}
	defer tbl.Close()

	// Specify the delimiter and comment character.
	tbl.Reader.Comma = ','
	tbl.Reader.Comment = '#'

	// Read all non-header rows.
	rows, err := tbl.ReadRows(1, -1)
	if err != nil {
		log.Fatalf("could read rows [1, -1): %v\n", err)
	}
	defer rows.Close()

	// Loop over the rows and fill the bmi data slice.
	var bmi []float64
	for rows.Next() {
		var v float64
		if err = rows.Scan(&v); err != nil {
			log.Fatal(err)
		}
		bmi = append(bmi, v)
	}

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.

	// Make a plot and set its title.

	// Create a histogram of our values drawn
	// from the standard normal.

	// Normalize the histogram.

	// Add the histogram to the plot.

	// Save the plot to a PNG file.
}
