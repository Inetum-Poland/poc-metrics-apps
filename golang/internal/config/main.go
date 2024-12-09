package config

type config struct {
	WebApp struct {
		Port int
		Name string
	}
	Otel struct {
		Host string
		Port int
	}
	Mongo struct {
		Host string
		Port int
		User string
		Pass string
	}
}

var C config
