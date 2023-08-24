package readcmd

import (
	csvpkg "jagaad-backend-test/pkg/csv"
)

type IReadUseCase interface {
	FetchAndSaveUsers() error
}

type ReadUseCase struct {
	CSVPkg csvpkg.ICSV
}

var (
	useCase ReadUseCase
)

func init() {
	useCase = ReadUseCase{
		CSVPkg: &csvpkg.CSVPkg{},
	}
}
