package config

type Config struct {
	DB     DB
	Server Server
	App    App
}

type DB struct {
	Username string
	Password string
	DBName   string
	Host     string
	Port     string
}

type Server struct {
	Host string
	Port string
}

type App struct {
	App string
}
