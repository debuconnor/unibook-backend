package unibookBackend

import (
	"github.com/debuconnor/dbcore"
	sc "github.com/debuconnor/sc-scraper"
)

func getUserCredential(db dbcore.Connection, userid, platform string) (credential map[string]string) {
	dml := dbcore.NewDml()

	dml.SelectAll()
	dml.From(SCHEMA_NAME_CREDENTIAL)
	dml.Where("", "userid", "=", userid)
	dml.Where("AND", "platform", "=", platform)
	credential = dml.Execute(db.GetDb())[0]

	return
}

func getConfigs(db dbcore.Connection, keys []string) (configs []map[string]string) {
	dml := dbcore.NewDml()

	dml.SelectAll()
	dml.From(SCHEMA_NAME_SYSTEM)
	for _, key := range keys {
		dml.Where("OR", "config_key", "=", key)
	}

	configs = dml.Execute(db.GetDb())

	return

}

func initSc(db dbcore.Connection, userid, platform string) sc.SessionController {
	credential := getUserCredential(db, userid, platform)
	configs := getConfigs(db, SC_CONFIGS)
	loginUrl := ""
	sessionKey := ""

	for _, row := range configs {
		key := row["config_key"]
		value := row["config_value"]

		switch key {
		case "sc_login_url":
			loginUrl = value
		case "sc_session_key":
			sessionKey = value
		}
	}

	sc.Init(db.GetDb())
	session := sc.NewSession(loginUrl, credential["userid"], credential["userpw"], sessionKey)

	if !session.HasLoginSession() {
		session.GetLoginSession(db.GetDb())
	}

	return session
}
