package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ReadConfig() (*ApiFoxConfig, []byte) {
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

	b, be := os.ReadFile(filepath.Join(currentDir, config.Body.Input))
	if be != nil {
		panic(be)
	}

	var swagger interface{}
	if err := json.Unmarshal(b, &swagger); err != nil {
		panic(err)
	}

	swaggerStr, sse := json.Marshal(swagger)
	if sse != nil {
		panic(sse)
	}

	config.Body.Input = string(swaggerStr)

	body, be := json.Marshal(config.Body)
	if be != nil {
		panic(be)
	}

	return &config, body
}
