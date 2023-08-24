package apipkg

import (
	"jagaad-backend-test/config"
	"jagaad-backend-test/entities"
)

//go:generate mockery --name IFetch --inpackage --case=underscore
type IFetch interface {
	FetchUsersFromAPI(conf config.Config) []entities.User
}

type ApiPkg struct{}
