package sisentry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MakeTeamsCard(datamodelTitle string, timestamp string, errorMessage string) *TeamsCard {

	facts := []Facts{
		{Name: "TimeStamp", Value: fmt.Sprintf("Failed Build: %s", timestamp)},
		{Name: "Error Log", Value: errorMessage},
	}

	sections := []Sections{
		{
			ActivityTitle: fmt.Sprintf("Failed Build: %s", datamodelTitle),
			Markdown:      false,
			Facts:         facts,
		},
	}

	teamsCard := TeamsCard{
		Context:    "http://schema.org/extensions",
		Type:       "MessageCard",
		Summary:    fmt.Sprintf("Failed Build: %s", datamodelTitle),
		ThemeColor: "0076D7",
		Sections:   sections,
	}

	return &teamsCard
}

func SendTeamsCard(teamsCard *TeamsCard) *http.Response {

	url := TeamsConnectorUrl
	teamsCardJson, err := json.Marshal(teamsCard)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(teamsCardJson))
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
