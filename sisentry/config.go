package sisentry

import (
	"fmt"
	"os"
)

var (
	SisenseRootUrl    = "https://domain.sisense.com/api/"
	SisenseAuthPath   = "v1/authentication/login"
	SisenseBuildsPath = "v2/builds"
	SisenseUserName   = os.Getenv("SISENSE_UN")
	SisensePassword   = os.Getenv("SISENSE_PW")
	TeamsConnectorUrl = os.Getenv("TEAMS_CONNECTOR_URL")
	SisenseAuthData   = []byte(fmt.Sprintf(`{
		"username": "%s",
		"password": "%s"
	}`, SisenseUserName, SisensePassword))
)
