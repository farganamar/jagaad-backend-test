package apipkg

import (
	"jagaad-backend-test/config"
	"jagaad-backend-test/entities"
)

type IFetch interface {
	FetchUsersFromAPI(conf config.Config) []entities.User
}

type ApiPkg struct{}
