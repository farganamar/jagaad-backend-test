package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	URL []string
}

var (
	ConfigObj = Config{}
)

func Init() error {
	currDir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
		return err
	}

	jsonConf, err := os.Open(fmt.Sprintf("%s%s", currDir, "/config/config.json"))

	if err != nil {
		log.Fatalf("Failed to find config.json file, ypu must created it. error message: %v", err)
		return err
	}

	configByte, err := io.ReadAll(jsonConf)
	if err != nil {
		log.Fatalf("Failed to read the io data : %v", err)
		return err
	}

	err = json.Unmarshal(configByte, &ConfigObj)
	if err != nil {
		log.Fatalf("Error reading config.json, %v", err)
		return err
	}

	return nil

}
