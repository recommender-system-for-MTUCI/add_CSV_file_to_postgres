package csv

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func LoadData() dataframe.DataFrame {
	file, err := os.Open("dataset_23k_v2.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	dataFrame := dataframe.ReadCSV(file)
	return dataFrame
}
