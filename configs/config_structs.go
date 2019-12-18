package configs

import (
	"time"
)

type DefaultConfig struct {
	APIPort string `yaml:"api_port"`

	CookieAuthName          string        `yaml:"cookie_auth_name"`
	CookieAuthDurationHours time.Duration `yaml:"cookie_auth_duration_hours"`

	AccessLog string `yaml:"access_log"`
	ErrorLog  string `yaml:"error_log"`

	UsersImageUploadPath   string `yaml:"users_image_upload_path"`
	FilmsImageUploadPath   string `yaml:"films_image_upload_path"`
	PersonsImageUploadPath string `yaml:"persons_image_upload_path"`

	MongoUser   string `yaml:"mongo_user"`
	MongoPwd    string `yaml:"mongo_pwd"`
	MongoHost   string `yaml:"mongo_host"`
	MongoDbName string `yaml:"mongo_db_name"`

	CountersCollectionName     string `yaml:"counters_collection_name"`
	UsersCollectionName        string `yaml:"users_collection_name"`
	CookiesCollectionName      string `yaml:"cookies_collection_name"`
	FilmsCollectionName        string `yaml:"films_collection_name"`
	PersonsCollectionName      string `yaml:"persons_collection_name"`
	LikesCollectionName        string `yaml:"likes_collection_name"`
	ReviewsCollectionName      string `yaml:"reviews_collection_name"`
	ListsCollectionName        string `yaml:"lists_collection_name"`
	SubscriptionCollectionName string `yaml:"subscriptions_collection_name"`

	UserTargetName         string `yaml:"user_target_name"`
	FilmTargetName         string `yaml:"film_target_name"`
	CookieTargetName       string `yaml:"cookie_target_name"`
	PersonTargetName       string `yaml:"person_target_name"`
	ReviewTargetName       string `yaml:"review_target_name"`
	LikeTargetName         string `yaml:"like_target_name"`
	StarTargetName         string `yaml:"star_target_name"`
	ListTargetName         string `yaml:"list_target_name"`
	SubscriptionTargetName string `yaml:"subscription_target_name"`
	FilmImageTargetName    string `yaml:"film_image_target_name"`
	PersonImageTargetName  string `yaml:"person_image_target_name"`
	DefaultImageName       string `yaml:"default_image_name"`

	CSRFHeader     string `yaml:"csrf_header"`
	CSRFCookieName string `yaml:"csrf_cookie_name"`

	DefaultEntriesLimit int `yaml:"default_entries_limit"`

	DefaultOrderBy string `yaml:"default_order_by"`

	DefaultYearMin int `yaml:"default_year_min"`
	DefaultYearMax int `yaml:"default_year_max"`

	MaxFileUploadSize int64 `yaml:"max_file_upload_size"`


	SessionsMicroservicePort string `yaml:"sessions_microservice_port"`
	UsersMicroservicePort    string `yaml:"users_microservice_port"`
}

var Default DefaultConfig

type HeadersConfig struct {
	HeadersMap map[string]string `yaml:"headers"`
}

var Headers HeadersConfig
