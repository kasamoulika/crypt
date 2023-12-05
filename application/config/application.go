package config

// application : struct to hold application level configs
type application struct {
	Name       string `toml:"app_name"`
	ListenPort int    `toml:"listen_port"`
}

// HitBtcConfig holds HitBtc service level configs
type HitBtcConfig struct {
	BaseUrl string `toml:"base_url"`
	Mock    bool   `toml:"mock"`
}
