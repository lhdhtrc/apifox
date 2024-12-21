package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	config, body := ReadConfig()

	method := "POST"
	importApiUrl := fmt.Sprintf("https://api.apifox.com/v1/projects/%s/import-openapi", config.ProjectId)

	//_ = os.WriteFile("request.json", body, 0755)

	request, err := http.NewRequest(method, importApiUrl, strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}

	request.Header.Add("X-Apifox-Api-Version", "2024-03-28")
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
	if res.StatusCode == 200 {
		fmt.Println("sync success")
	} else {
		fmt.Println(string(b))
	}
}
