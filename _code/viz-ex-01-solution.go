// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to generate a box plot of diabetes bmi values.
package main

import (
	"log"

	"go-hep.org/x/hep/csvutil"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	// Create the plot and set its title and axis label.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Y.Label.Text = "Values"

	// Create the box for our data.
	w := vg.Points(50)

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the CSV file.
	v := make(plotter.Values, len(bmi))
	for i, val := range bmi {
		v[i] = val
	}

	// Add the data to the plot.
	box, err := plotter.NewBoxPlot(w, 0, v)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(box)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.
	p.NominalX("bmi")

	if err := p.Save(4*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
