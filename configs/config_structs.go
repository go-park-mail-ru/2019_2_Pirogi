package configs

import "time"

type DefaultConfig struct {
	CookieAuthName          string        `yaml:"cookie_auth_name"`
	CookieAuthDurationHours time.Duration `yaml:"cookie_auth_duration_hours"`
	AccessLog               string        `yaml:"access_log"`
	ErrorLog                string        `yaml:"error_log"`
	UsersImageUploadPath    string        `yaml:"users_image_upload_path"`
	FilmsImageUploadPath    string        `yaml:"films_image_upload_path"`
	MongoUser               string        `yaml:"mongo_user"`
	MongoPwd                string        `yaml:"mongo_pwd"`
	MongoHost               string        `yaml:"mongo_host"`
	MongoDbName             string        `yaml:"mongo_db_name"`
	UsersCollectionName     string        `yaml:"users_collection_name"`
	FilmsCollectionName     string        `yaml:"films_collection_name"`
	CookiesCollectionName   string        `yaml:"cookies_collection_name"`
	CountersCollectionName  string        `yaml:"counters_collection_name"`
	UserTargetName          string        `yaml:"user_target_name"`
	FilmTargetName          string        `yaml:"film_target_name"`
	CookieTargetName        string        `yaml:"cookie_target_name"`
	APIPort                 string        `yaml:"api_port"`
	CertFile                string        `yaml:"cert_file"`
	KeyFile                 string        `yaml:"key_file"`
}

var Default DefaultConfig

type HeadersConfig struct {
	HeadersMap map[string]string `yaml:"headers"`
}

var Headers HeadersConfig