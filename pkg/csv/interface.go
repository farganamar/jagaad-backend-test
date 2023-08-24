package csvpkg

import "jagaad-backend-test/entities"

//go:generate mockery --name ICSV --inpackage --case=underscore
type ICSV interface {
	Write([]entities.User) error
	Read() ([][]string, error)
}

type CSVPkg struct{}
