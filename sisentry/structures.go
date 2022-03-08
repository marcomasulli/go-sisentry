package sisentry

type ApiBuild []struct {
	BuildType        string `json:"buildType"`
	Completed        string `json:"completed"`
	Created          string `json:"created"`
	DatamodelID      string `json:"datamodelId"`
	DatamodelTitle   string `json:"datamodelTitle"`
	DatamodelType    string `json:"datamodelType"`
	IndexSize        string `json:"indexSize"`
	InstanceID       string `json:"instanceId"`
	Oid              string `json:"oid"`
	RowLimit         int64  `json:"rowLimit"`
	SchemaLastUpdate string `json:"schemaLastUpdate"`
	Source           string `json:"source"`
	Started          string `json:"started"`
	Status           string `json:"status"`
}

type DbBuild struct {
	Id             int64
	Oid            string
	DatamodelID    string
	DatamodelTitle string
	InstanceID     string
	Created        string
	Started        string
	Completed      string
}

type ApiBuildLog []struct {
	BuildSeq     int64  `json:"buildSeq"`
	CubeID       string `json:"cubeId"`
	DataModelOid string `json:"dataModelOid"`
	Message      string `json:"message"`
	ServerID     string `json:"serverId"`
	ServerName   string `json:"serverName"`
	Timestamp    string `json:"timestamp"`
	Title        string `json:"title"`
	TraceLevel   int64  `json:"traceLevel"`
	Type         string `json:"type"`
	TypeValue    struct {
		AdditionalInfo interface{} `json:"additionalInfo"`
		ColumnName     string      `json:"columnName"`
		CountRecords   int64       `json:"countRecords"`
		Dependencies   struct {
			SortedDependencies []struct {
				Dependencies []interface{} `json:"dependencies"`
				Name         string        `json:"name"`
				Table        string        `json:"table"`
				Type         string        `json:"type"`
			} `json:"sortedDependencies"`
		} `json:"dependencies"`
		Description  string `json:"description"`
		Message      string `json:"message"`
		TableName    string `json:"tableName"`
		Title        string `json:"title"`
		TotalRecords int64  `json:"totalRecords"`
	} `json:"typeValue"`
	Verbosity string `json:"verbosity"`
}

type LogOutput struct {
	DatamodelTitle string `json:"datamodelTitle"`
	Timestamp      string `json:"timestamp"`
	ErrorMessage   string `json:"errorMessage"`
}

type TeamsCard struct {
	Context    string     `json:"@context"`
	Type       string     `json:"@type"`
	Sections   []Sections `json:"sections"`
	Summary    string     `json:"summary"`
	ThemeColor string     `json:"themeColor"`
}

type Sections struct {
	ActivityTitle string  `json:"activityTitle"`
	Facts         []Facts `json:"facts"`
	Markdown      bool    `json:"markdown"`
}

type Facts struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
