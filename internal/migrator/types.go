package migrator

var WorkdirPath string
var SchemasPath string

type Migration struct {
	Id        int    `json:"id"`
	Name      string `json:"Name"`
	Version   int    `json:"Version"`
	CreatedAt int64  `json:"CreatedAt"`
}

type MigrationStash struct {
	Name         string `json:"Name"`
	Version      int    `json:"Version"`
	UpFileBody   string `json:"UpFileBody"`
	DownFileBody string `json:"DownFileBody"`
	ChangeTime   int64  `json:"ChangeTime"`
}
