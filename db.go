package unibookBackend

import (
	"github.com/debuconnor/dbcore"
)

func GetDbFromGcpSecret(secretVersion string) dbcore.Connection {
	resultJson := accessSecretVersion(secretVersion)
	dbInfo := parseJson(resultJson)

	db := dbcore.NewDb()
	db.SetConnection(dbInfo["host"].(string), dbInfo["port"].(string), dbInfo["username"].(string), dbInfo["password"].(string), dbInfo["dbname"].(string))

	return db
}
