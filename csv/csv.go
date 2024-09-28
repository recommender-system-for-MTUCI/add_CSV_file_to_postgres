package csv

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func LoadData() dataframe.DataFrame {
	file, err := os.Open("dataset_1.3.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	dataFrame := dataframe.ReadCSV(file)
	return dataFrame
}
