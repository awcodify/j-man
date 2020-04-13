package config

import (
	"database/sql"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/awcodify/j-man/utils"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Join(filepath.Dir(b), "../")
)

const projectDirName = "j-man"

// Config will be seperated between production and development
type Config struct {
	JMeter struct {
		Path string
	}
	Server struct {
		Host string
		Port string
	}
	HTML  HTML
	DB    Database
	OAuth OAuth
	Redis Redis
}

// HTML used for templating
type HTML struct {
	Root   string
	Layout layout
}

type layout struct {
	Root     string
	BaseHTML string
}

// Database will used for managing our state and scripts
type Database struct {
	DSN string `yaml:"dsn"`
}

//OAuth for google sign in
type OAuth struct {
	GoogleClientID     string
	GoogleClientSecret string
	CallbackURL        string
	Scopes             string
	Expiration         int
	Endpoint           string
}

// Redis will store our cache
type Redis struct {
	Host     string
	Password string
	DB       int
}

// ConnectRedis will initiate connection to redis
func (cfg Config) ConnectRedis() *redis.Client {
	options := &redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	return redis.NewClient(options)
}

// GetGoogleOAuthConfig will parse config for google oauth
func (config *Config) GetGoogleOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  config.OAuth.CallbackURL,
		ClientID:     config.OAuth.GoogleClientID,
		ClientSecret: config.OAuth.GoogleClientSecret,
		Scopes:       strings.Split(config.OAuth.Scopes, ","),
		Endpoint:     google.Endpoint,
	}
}

// ConnectDB based on driver being used
func (cfg Config) ConnectDB() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", cfg.DB.DSN)
	utils.DieIf(err)

	return db, err
}

// New will instantiate config from production or development
func New() (*Config, error) {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}
	envPath := "/.env." + env + ".local"

	err := godotenv.Load(string(rootPath) + envPath)
	if err != nil {
		return nil, err
	}

	godotenv.Load() // for production ( .env only )

	var c Config
	c.JMeter.Path = os.Getenv("JMETER_PATH")
	c.Server.Host = os.Getenv("SERVER_HOST")
	c.Server.Port = os.Getenv("SERVER_PORT")
	c.HTML.Root = os.Getenv("HTML_ROOT")
	c.HTML.Layout.Root = os.Getenv("HTML_LAYOUT_ROOT")
	c.HTML.Layout.BaseHTML = os.Getenv("HTML_LAYOUT_BASE_HTML")
	c.DB.DSN = os.Getenv("DB_DSN")
	c.OAuth.GoogleClientID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	c.OAuth.GoogleClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	c.OAuth.CallbackURL = os.Getenv("GOOGLE_OAUTH_CALLBACK_URL")
	c.OAuth.Scopes = os.Getenv("GOOGLE_OAUTH_SCOPES")

	expiration, err := strconv.Atoi(os.Getenv("GOOGLE_OAUTH_EXPIRATION"))
	if err != nil {
		return nil, err
	}
	c.OAuth.Expiration = expiration

	return &c, nil
}
