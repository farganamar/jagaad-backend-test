package fetchcmd

import (
	apipkg "jagaad-backend-test/pkg/api"
	csvpkg "jagaad-backend-test/pkg/csv"
)

type IFetchUseCase interface {
	FetchAndSaveUsers() error
}

type FetchUseCase struct {
	CSVPkg     csvpkg.ICSV
	FetcherPkg apipkg.IFetch
}

var (
	useCase FetchUseCase
)

func init() {
	useCase = FetchUseCase{
		CSVPkg:     &csvpkg.CSVPkg{},
		FetcherPkg: &apipkg.ApiPkg{},
	}
}
