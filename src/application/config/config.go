package config

type Config struct {
	DB     DB
	Server Server
	App    App
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type Server struct {
	Host string
	Port string
}

type App struct {
	App string
}
