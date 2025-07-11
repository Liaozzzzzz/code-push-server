package config

type Config struct {
	General General
}

type General struct {
	AppName            string `default:"code-push-server"`
	Version            string `default:"v0.0.1"`
	Debug              bool
	DisableSwagger     bool
	DisablePrintConfig bool
	HTTP               struct {
		Addr            string `default:":8040"`
		ShutdownTimeout int    `default:"10"` // seconds
		ReadTimeout     int    `default:"60"` // seconds
		WriteTimeout    int    `default:"60"` // seconds
		IdleTimeout     int    `default:"10"` // seconds
		CertFile        string
		KeyFile         string
	}
}
