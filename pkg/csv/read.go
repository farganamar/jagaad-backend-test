package csvpkg

import (
	"encoding/csv"
	"os"
)

func (i *CSVPkg) Read() ([][]string, error) {
	file, err := os.Open("user.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
