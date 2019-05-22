package lib

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

func NewDBConn() *sql.DB {
	db, err := sql.Open("mssql", viper.GetString("db.url"))
	Must(err)
	return db
}

func GetSourceHashes(sourceHashTypeId int, client string) []string {
	db := NewDBConn()
	defer db.Close()

	fmt.Println(db)

	rows, err := db.Query("SELECT hash FROM SourceHashes WHERE SourceHashTypeId=? AND client=?", sourceHashTypeId, client)
	Must(err)

	hashes := make([]string, 0)
	for rows.Next() {
		var hash string
		Must(rows.Scan(&hash))
		hashes = append(hashes, hash)
	}

	return hashes
}

func LoadConfig(configFilename *string) {
	viper.SetDefault("server.listen", ":9101")
	viper.SetConfigFile(*configFilename)
	err := viper.ReadInConfig()
	if nil != err {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func Must(err error) {
	if nil != err {
		panic(err)
	}
}
