package unibookBackend

import (
	"github.com/debuconnor/dbcore"
)

func initDb() (db dbcore.Connection) {
	db = getDbFromGcpSecret(GCP_SECRET_VERSION)
	db.ConnectMysql()
	defer db.DisconnectMysql()

	if !db.IsConnected() {
		Error(ERROR_CODE_DB_CONNECTION_FAILED)
		return
	}

	return

}

func getDbFromGcpSecret(secretVersion string) dbcore.Connection {
	resultJson := accessSecretVersion(secretVersion)
	dbInfo := parseJson(resultJson)

	db := dbcore.NewDb()
	db.SetConnection(dbInfo["host"].(string), dbInfo["port"].(string), dbInfo["username"].(string), dbInfo["password"].(string), dbInfo["dbname"].(string))

	return db
}

func checkTables(db dbcore.Connection) (errCode int) {
	ddl := dbcore.NewDdl()

	if !ddl.CheckTableExists(db.GetDb(), SCHEMA_NAME_SYSTEM) {
		return ERROR_CODE_SYSTEM_TABLE_NOT_EXISTS
	}

	if !ddl.CheckTableExists(db.GetDb(), SCHEMA_NAME_CREDENTIAL) {
		return ERROR_CODE_CREDENTIAL_TABLE_NOT_EXISTS
	}

	return ERROR_NOT_FOUND
}
