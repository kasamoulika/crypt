// Package config will manage all application level configurations
// config file will be taken based on the application environment
// this will be immutable as it always provides the value of the struct
package config

import (
	"fmt"
	"io"
	"os"

	"crypt.com/v2/application/constants"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

const (
	// FilePath - relative path to the config directory
	FilePath = "%s/conf/%s"

	// DefaultFilename - Filename format of default config file
	DefaultFilename = "env.default.toml"

	// EnvFilename - Filename format of env specific config file
	EnvFilename = "env.%s.toml"
)

var (
	// config : this will hold all the application configuration
	config AppConfig
)

// AppConfig is the struct definition for global configuration
type AppConfig struct {
	Application application  `toml:"application"`
	HitBtc      HitBtcConfig `toml:"hitbtc"`
	Symbol      Symbol       `toml:"symbol"`
}

// LoadConfig will load the configuration available in the conf directory available in basePath
// conf file will takes based on the env provided
func LoadConfig(basePath string, env string) {
	// reading conf based on default environment
	loadConfigFromFile(basePath, DefaultFilename, "")

	// reading env file and override conf values; if env file exists
	loadConfigFromFile(basePath, EnvFilename, env)
}

// GetConfig : will give the struct as value so that the actual conf doesn't get tampered
func GetConfig() AppConfig {
	return config
}

// loadConfigFromFile: loads config values from given basepath, filename and env
func loadConfigFromFile(basePath string, filename string, env string) {
	path := getFilePath(basePath, filename, env)
	_, err := os.Stat(path)
	if err != nil {
		log.WithError(err).Error(constants.ConfigFileNotFound)
	}

	content := readConfigFile(path)
	if _, err := toml.Decode(content, &config); err != nil {
		log.WithError(err).Panic(constants.InvalidConfigType)
	}
}

// readConfigFile: config file will be read from the given file and gives the content in string format
func readConfigFile(path string) string {
	file, err := os.Open(path)

	if err != nil {
		log.WithError(err).Error(constants.ConfigFileNotFound)
		return ""
	}
	defer file.Close()

	// Read the file content into a byte slice
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return ""
	}

	return string(data)
}

// getFilePath: gives the file path based on the environment provided
// file path will be relative to the application and determined by basePath
func getFilePath(basePath string, fileName string, env string) string {
	if env != "" {
		fileName = fmt.Sprintf(fileName, env)
	}

	path := fmt.Sprintf(FilePath, basePath, fileName)

	return path
}
