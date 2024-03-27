package unibookBackend

const (
	SCHEMA_NAME_SYSTEM     = "system_config"
	SCHEMA_NAME_CREDENTIAL = "platform_credential"

	PLATFORM_NVR = "2"
	PLATFORM_SC  = "1"

	ERROR_NOT_FOUND                        = -1
	ERROR_CODE_DB_CONNECTION_FAILED        = 100
	ERROR_CODE_SYSTEM_TABLE_NOT_EXISTS     = 101
	ERROR_CODE_CREDENTIAL_TABLE_NOT_EXISTS = 102

	ERROR_CODE_PARSE_JSON = 201

	ERROR_CODE_CREATE_SECRETMANAGER_CLIENT = 301
	ERROR_CODE_ACCESS_SECRET_VERSION       = 302
	ERROR_CODE_DATA_CORRUPTION             = 303
)

var (
	IS_DEBUG           = false
	GCP_SECRET_VERSION = ""

	SC_CONFIGS  = []string{"sc_login_url", "sc_scrape_url", "sc_scrape_page_count", "sc_scrape_selector", "sc_session_key"}
	NVR_CONFIGS = []string{"nvr_login_url", "nvr_scrape_url", "nvr_scrape_selector"}
)

func SetDebug(mode bool) {
	IS_DEBUG = mode
}

func SetGcpSecretVersion(version string) {
	GCP_SECRET_VERSION = version
}
