package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ReadConfig() (*ApiFoxConfig, string) {
	currentDir, ce := os.Getwd()
	if ce != nil {
		panic(ce)
	}

	file, fe := os.ReadFile(filepath.Join(currentDir, "apifox.config.json"))
	if fe != nil {
		panic(fe)
	}

	var config ApiFoxConfig
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	swagger, se := os.ReadFile(filepath.Join(currentDir, config.Body.Input))
	if se != nil {
		panic(se)
	}
	config.Body.Input = string(swagger)

	body, be := json.Marshal(config.Body)
	if be != nil {
		panic(be)
	}

	return &config, string(body)
}
