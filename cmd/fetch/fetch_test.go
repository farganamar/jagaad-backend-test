package fetchcmd

import (
	"errors"
	"jagaad-backend-test/entities"
	apipkg "jagaad-backend-test/pkg/api"
	csvpkg "jagaad-backend-test/pkg/csv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchAndSaveUsers(t *testing.T) {
	tests := []struct {
		name            string
		mockFetchUser   []entities.User
		mockWriteCSVErr error
		isError         bool
	}{
		{
			name: "Success case",
		},
		{
			name:            "Error write to csv case",
			mockWriteCSVErr: errors.New("Error write to csv"),
			isError:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockFetch := apipkg.MockIFetch{}
			mockCSV := csvpkg.MockICSV{}

			mockFetch.On("FetchUsersFromAPI", mock.Anything).Return(test.mockFetchUser)
			mockCSV.On("Write", mock.Anything).Return(test.mockWriteCSVErr)

			useCase := FetchUseCase{
				CSVPkg:     &mockCSV,
				FetcherPkg: &mockFetch,
			}

			err := useCase.FetchAndSaveUsers()
			if test.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
