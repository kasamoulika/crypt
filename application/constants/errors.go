package constants

// Application Internal Specific
const (
	//ConfigFileNotFound .. config file unavailable
	ConfigFileNotFound = "Config file not found"

	//InvalidConfigType : Invalid types in config
	InvalidConfigType = "Invalid types in config"
)

const (
	HitBtcError = "HITBTC_INTERNAL_ERROR"
)

// Application Status /Request Specific
const (
	//Error400 : Error string 400 error code
	Error400 string = "Bad Request"

	//Error401 : Error string 401 error code
	Error401 string = "Unauthorized access"

	//Error403 : Error string 403 error code
	Error403 string = "Not Authorized"

	//Error404 : Error string 404 error code
	Error404 string = "Page Not Found"

	//Error500 : Error string 500 error code
	Error500 string = "Internal Server Error"
)
