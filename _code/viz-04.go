// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to generate box plots of the iris data variables.
package main

import (
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"go-hep.org/x/hep/csvutil/csvdriver"
)

func main() {

	// Create the plot and set its title and axis label.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// Create the box for our data.
	w := vg.Points(50)

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

		// Add the data to the plot.
		b, err := plotter.NewBoxPlot(w, float64(i), v)
		if err != nil {
			log.Fatal(err)
		}
		p.Add(b)
	}

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")

	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
