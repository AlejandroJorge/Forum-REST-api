package config

type Parameters struct {
	DbFileName   string
	DbFolderName string
	Port         uint
	AuthSecret   []byte
}

var params Parameters
var isParamsInitialized = false

var defaultParams Parameters = Parameters{
	DbFolderName: "data",
	DbFileName:   "database.sqlite",
	Port:         3000,
	AuthSecret:   []byte("weaksecret"),
}

func GetParams() Parameters {
	if !isParamsInitialized {
		initializeConfigParameters()
	}

	return params
}

func initializeConfigParameters() {
	var ok bool
	if params.DbFolderName, ok = getEnvString("DB_FOLDER_NAME"); !ok {
		params.DbFolderName = defaultParams.DbFolderName
	}
	if params.DbFileName, ok = getEnvString("DB_FILE_NAME"); !ok {
		params.DbFileName = defaultParams.DbFileName
	}
	if params.Port, ok = getEnvUint("PORT"); !ok {
		params.Port = defaultParams.Port
	}
	if params.AuthSecret, ok = getEnvBytes("AUTH_SECRET"); !ok {
		params.AuthSecret = defaultParams.AuthSecret
	}

	isParamsInitialized = true
}
