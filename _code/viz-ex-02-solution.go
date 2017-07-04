// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise2

// Sample program to generate a histogram of diabetes bmi values.
package main

import (
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
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
	v := make(plotter.Values, len(bmi))
	for i, val := range bmi {
		v[i] = val
	}

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Histogram of a BMI"

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		log.Fatal(err)
	}

	// Normalize the histogram.
	h.Normalize(1)

	// Add the histogram to the plot.
	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "bmi_hist.png"); err != nil {
		log.Fatal(err)
	}

}
