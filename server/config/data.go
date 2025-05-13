package config

type Config struct {
	Web      WebConfig      `json:"web"`
	Database DatabaseConfig `json:"database"`
	Storage  StorageConfig  `json:"storage"`
}

type WebConfig struct {
	Address string `json:"address"`
	BaseUrl string `json:"baseUrl"`
}

type DatabaseConfig struct {
	Type string `json:"type"`
	Dsn  string `json:"dsn"`
}

type StorageConfig struct {
	Type    string            `json:"type"`
	Options map[string]string `json:"options"`
}
