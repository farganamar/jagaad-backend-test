package fetchcmd

import (
	"jagaad-backend-test/config"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "fetch",
		Short: "Fetch users from APIs and save to CSV",
		RunE: func(cmd *cobra.Command, args []string) error {
			return useCase.FetchAndSaveUsers()
		},
	}
}

func (i *FetchUseCase) FetchAndSaveUsers() error {
	users := i.FetcherPkg.FetchUsersFromAPI(config.ConfigObj)
	err := i.CSVPkg.Write(users)

	if err != nil {
		return err
	}

	return nil
}
