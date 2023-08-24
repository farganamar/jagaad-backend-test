package apipkg

import (
	"encoding/json"
	"io"
	"jagaad-backend-test/config"
	"jagaad-backend-test/entities"
	"log"
	"net/http"
)

func (i *ApiPkg) FetchUsersFromAPI(conf config.Config) []entities.User {
	var result []entities.User
	for _, endpoint := range conf.URL {
		resp, err := http.Get(endpoint)
		if err != nil {
			log.Printf("Error on access this url: %v , error message : %v", endpoint, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Printf("Error on fetching this url %v", endpoint)
			continue
		}

		var users []entities.User
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error on read the data : %v", err)
			continue
		}

		err = json.Unmarshal(body, &users)
		if err != nil {
			log.Printf("error decode the response data :  %v", err)
			continue
		}

		result = append(result, users...)
	}

	return result
}
