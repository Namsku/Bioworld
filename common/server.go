package common

type Server struct {
	Port     int
	Password string
	Channels map[string]Channel
}
