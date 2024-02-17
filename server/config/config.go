package config

type Config struct {
	App    App
	Server Server
	DB     DB
	Jwt    Jwt
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type Jwt struct {
	Secret          string
	JwtValidityTime int
}
