// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to generate a box plot of diabetes bmi values.
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

	// Create the plot and set its title and axis label.

	// Create the box for our data.

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the CSV file.

	// Add the data to the plot.

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.

	// Save the plot.
}
