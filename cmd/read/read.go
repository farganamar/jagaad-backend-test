package readcmd

import (
	"encoding/json"
	"fmt"
	"jagaad-backend-test/entities"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var tags string
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search users with specific tags",
		RunE: func(cmd *cobra.Command, args []string) error {
			users, err := useCase.ReadAndSearch(tags)
			if err != nil {
				log.Fatalf("Error could not get user by tags, err: %v", err)
			}

			if len(users) <= 0 {
				fmt.Println("No user found with that tags")
			}

			for _, user := range users {
				fmt.Printf("%s | %s\n", user.GUID, user.Balance)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&tags, "tags", "t", "", "Comma-separated tags (required)")
	cmd.MarkFlagRequired("tags")

	return cmd
}

func (i *ReadUseCase) ReadAndSearch(tags string) ([]entities.User, error) {
	// read all from CSV
	rows, err := i.CSVPkg.Read()
	if err != nil {
		return nil, err
	}

	arrTags := strings.Split(tags, ",")

	users := make([]entities.User, 0)
	for _, row := range rows[1:] {
		hasAllTags := hasAllTags(row, arrTags)
		if !hasAllTags {
			continue
		}

		index, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatalf("Error convert index to integer, err: %v", err)
			return nil, err
		}

		isActive := false
		if row[3] == "true" {
			isActive = true
		}

		rowTags := strings.Split(row[5], ",")

		// decode friends row data
		var friends []entities.Friend
		err = json.Unmarshal([]byte(row[6]), &friends)
		if err != nil {
			log.Fatalf("Error unmarshal friend, error: %v", err)
			return nil, err
		}

		users = append(users, entities.User{
			ID:       row[0],
			Index:    int64(index),
			GUID:     row[2],
			IsActive: isActive,
			Balance:  row[4],
			Tags:     rowTags,
			Friends:  friends,
		})
	}

	return users, nil

}

func hasAllTags(row []string, tags []string) bool {
	for _, tag := range tags {
		if !strings.Contains(row[5], tag) {
			return false
		}
	}
	return true
}
