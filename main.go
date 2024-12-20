package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	client := &http.Client{}
	config, body := ReadConfig()

	method := "POST"
	importApiUrl := fmt.Sprintf("https://api.apifox.com/v1/projects/%s/import-openapi", config.ProjectId)

	request, err := http.NewRequest(method, importApiUrl, strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	request.Header.Add("X-Apifox-Api-Version", time.Now().Format(time.DateTime))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(b))
}
