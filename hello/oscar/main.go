package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type OscarData struct {
	Index string
	Year  string
	Age   string
	Name  string
	Movie string
}

func ReadOscarCSV() []OscarData {

	csvFile, err := os.Open("oscar.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	var oscars []OscarData
	for _, line := range csvLines {
		d := OscarData{
			Index: line[0],
			Year:  line[1],
			Age:   line[2],
			Name:  line[3],
			Movie: line[4],
		}

		oscars = append(oscars, d)
	}

	return oscars
}

func main() {
	x := ReadOscarCSV()
	records := x[1:]

	m := map[string]int{}
	for _, record := range records {
		m[record.Name]++
	}

	for name, count := range m {
		if count > 1 {
			fmt.Println(name)
		}
	}
}
