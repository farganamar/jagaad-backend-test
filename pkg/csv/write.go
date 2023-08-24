package csvpkg

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"jagaad-backend-test/entities"
	"log"
	"os"
	"strings"
)

func (i *CSVPkg) Write(users []entities.User) error {
	file, err := os.Create("user.csv")
	if err != nil {
		log.Fatalf("Error creating csv file: %v", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	err = writer.Write([]string{
		"ID",
		"Index",
		"GUID",
		"IsActive",
		"Balance",
		"Tags",
		"Friends",
	})
	if err != nil {
		return err
	}

	// Write user data
	for _, user := range users {
		isActive := "false"
		if user.IsActive {
			isActive = "true"
		}
		friends, err := json.Marshal(user.Friends)
		if err != nil {
			log.Fatalf("Error marshal friends from user : %v", err)
		}
		err = writer.Write([]string{
			fmt.Sprint(user.ID),
			fmt.Sprint(user.Index),
			user.GUID,
			isActive,
			user.Balance,
			strings.Join(user.Tags, ","),
			string(friends),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
