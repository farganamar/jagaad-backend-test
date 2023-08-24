package csvpkg

import "jagaad-backend-test/entities"

type ICSV interface {
	Write([]entities.User) error
	Read() ([][]string, error)
}

type CSVPkg struct{}
