package readcmd

import (
	"errors"
	csvpkg "jagaad-backend-test/pkg/csv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReadAndSearch(t *testing.T) {
	var mockReadCsv = [][]string{
		{"ID", "Index", "GUID", "IsActive", "Balance", "Tags", "Friends"},
		{"64d39b0582ec3cff5fc7f24e", "0", "03ee84da-5a54-493f-8438-60bad7ab6e2a", "true", "$2,633.92", "pariatur,qui,ea,culpa,laboris,laboris,minim", "[{\"id\":0,\"name\":\"Koch Valdez\"},{\"id\":1,\"name\":\"Kramer Bush\"},{\"id\":2,\"name\":\"Townsend Church\"}]"},
		{"64d39b0506da562378c0f321", "1", "bfb84280-21a2-4199-833f-4f54d98fb15b", "true", "$3,626.10", "irure,nostrud,ipsum,consectetur,consectetur,occaecat,consectetur", "[{\"id\":0,\"name\":\"Therese Dorsey\"},{\"id\":1,\"name\":\"Gilliam Stephens\"},{\"id\":2,\"name\":\"Leblanc Odonnell\"}]"},
	}
	tests := []struct {
		name           string
		tagSearch      string
		mockReadCsvRes [][]string
		mockReadCsvErr error
		isError        bool
	}{
		{
			name:           "Success case",
			tagSearch:      "irure",
			mockReadCsvRes: mockReadCsv,
		},
		{
			name:           "Error read csv",
			mockReadCsvErr: errors.New("Error read csv"),
			isError:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCSV := csvpkg.MockICSV{}

			mockCSV.On("Read", mock.Anything).Return(test.mockReadCsvRes, test.mockReadCsvErr)

			useCase := ReadUseCase{
				CSVPkg: &mockCSV,
			}

			_, err := useCase.ReadAndSearch(test.tagSearch)
			if test.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
