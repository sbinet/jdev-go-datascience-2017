// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to generate a histogram of the iris data variables.
package main

import (
	"fmt"
	"log"

	"go-hep.org/x/hep/csvutil/csvdriver"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
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
	for i, col := range vars {
		// Create a plotter.Values value and fill it with the
		// values from the respective column of the CSV file.
		v := make(plotter.Values, len(col))
		for j, vv := range col {
			v[j] = vv
		}

		// Make a plot and set its title.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a var-%d", i+1)

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
		if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("var-%d_hist.png", i+1)); err != nil {
			log.Fatal(err)
		}
	}
}
