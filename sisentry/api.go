package sisentry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAuthKey() string {

	url := fmt.Sprintf("%s%s", SisenseRootUrl, SisenseAuthPath)

	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(SisenseAuthData))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return fmt.Sprintf("Bearer %s", res["access_token"])
}

func GetBuilds(authKey string) []byte {

	url := fmt.Sprintf("%s%s", SisenseRootUrl, SisenseBuildsPath)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("authorization", authKey)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func GetBuildLog(authKey string, datamodelId string, datamodelTitle string) *LogOutput {

	var apiBuildLog ApiBuildLog

	url := fmt.Sprintf("%sv1/elasticubes/%s/buildLogs", SisenseRootUrl, datamodelId)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("authorization", authKey)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	unmarshalError := json.Unmarshal(body, &apiBuildLog)
	if unmarshalError != nil {
		log.Fatal(unmarshalError)
	}

	logOutput := LogOutput{}

	for l := range apiBuildLog {
		if apiBuildLog[l].Verbosity == "Error" {
			logOutput.DatamodelTitle = datamodelTitle
			logOutput.Timestamp = apiBuildLog[l].Timestamp
			logOutput.ErrorMessage = apiBuildLog[l].Message
		}
		fmt.Println(datamodelTitle)
		fmt.Println(apiBuildLog[l].Timestamp)
		fmt.Println(apiBuildLog[l].Message)
	}

	return &logOutput

}
