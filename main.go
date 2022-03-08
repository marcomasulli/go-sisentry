package sisentry

import (
	"database/sql"
	"encoding/json"
	"fmt"
	builds "go-sisentry/sisentry"
	"log"
	"time"
)

const dbName = "sqlite.db"

func main() {

	fmt.Println("Launching Sisentry. Please wait...")

	var apiBuild builds.ApiBuild

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	buildsRepository := builds.NewSQLiteRepository(db)

	if err := buildsRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	authKey := builds.GetAuthKey()

	// fmt.Println(authKey)

	for {

		fmt.Println("Getting builds...")

		buildJson := builds.GetBuilds(authKey)

		unmarshalError := json.Unmarshal(buildJson, &apiBuild)
		if unmarshalError != nil {
			log.Fatal(unmarshalError)
		}

		fmt.Println("Checking builds for failures...")

		for b := range apiBuild {
			// check if the build is failed
			if apiBuild[b].Status == "failed" {
				fmt.Printf("Failed build: %s, %s\n\r", apiBuild[b].Oid, apiBuild[b].DatamodelTitle)
				// if so, check if we already have an entry for the same oid in our sqlite db
				existingBuild, err := buildsRepository.GetByOid(apiBuild[b].Oid)
				if err != nil {
					// if we don't, create a row...
					if existingBuild == nil {
						newBuild := builds.DbBuild{
							Oid:            apiBuild[b].Oid,
							DatamodelID:    apiBuild[b].DatamodelID,
							DatamodelTitle: apiBuild[b].DatamodelTitle,
							InstanceID:     apiBuild[b].InstanceID,
							Created:        apiBuild[b].Created,
							Started:        apiBuild[b].Started,
							Completed:      apiBuild[b].Completed,
						}
						createdDbBuild, err := buildsRepository.Create(newBuild)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Printf("Created build %s in db", createdDbBuild.Oid)
						// extract the log
						logOutput := builds.GetBuildLog(authKey, apiBuild[b].DatamodelID, apiBuild[b].DatamodelTitle)
						// build the card
						teamsCard := builds.MakeTeamsCard(apiBuild[b].DatamodelTitle, logOutput.Timestamp, logOutput.ErrorMessage)
						// and send a notification to teams/whatsapp/
						r := builds.SendTeamsCard(teamsCard)

						fmt.Printf("%s", r.Body)
					}
				}
			}
		}

		// wait 60 sec
		fmt.Println("Waiting 60 seconds...")
		time.Sleep(60 * time.Second)
	}
}
