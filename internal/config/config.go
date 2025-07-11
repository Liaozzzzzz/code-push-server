package config

type Config struct {
	General    General
	Database   Database
	Log        Log
	Cache      Cache
	Security   Security
	Services   Services
	Monitoring Monitoring
	Business   Business
	Developer  Developer
}

type General struct {
	AppName        string `default:"code-push-server"`
	Version        string `default:"v0.0.1"`
	Debug          bool
	DisableSwagger bool
	HTTP           struct {
		Addr            string `default:":8040"`
		ShutdownTimeout int    `default:"10"` // seconds
		ReadTimeout     int    `default:"60"` // seconds
		WriteTimeout    int    `default:"60"` // seconds
		IdleTimeout     int    `default:"10"` // seconds
		CertFile        string
		KeyFile         string
	}
}

type Database struct {
	Driver          string `default:"sqlite"`
	DSN             string `default:"data/app.db"`
	MaxOpenConns    int    `default:"25"`
	MaxIdleConns    int    `default:"10"`
	ConnMaxLifetime int    `default:"300"` // seconds

	Migration struct {
		AutoMigrate   bool   `default:"true"`
		MigrationPath string `default:"migrations"`
		SeedData      bool
	}

	Connection struct {
		MaxRetries    int `default:"3"`
		RetryInterval int `default:"5"`  // seconds
		PingTimeout   int `default:"10"` // seconds
	}

	Logging struct {
		LogLevel           string `default:"warn"`
		SlowQueryThreshold int    `default:"1000"` // milliseconds
		LogQueries         bool
	}
}

type Log struct {
	Level      string `default:"info"`
	Format     string `default:"json"`
	Output     string `default:"stdout"`
	FilePath   string
	MaxSize    int  `default:"100"` // MB
	MaxBackups int  `default:"7"`
	MaxAge     int  `default:"30"` // days
	Compress   bool `default:"true"`
}

type Cache struct {
	Type  string `default:"memory"`
	TTL   int    `default:"3600"` // seconds
	Redis struct {
		Addr     string `default:"localhost:6379"`
		Password string
		DB       int `default:"0"`
		PoolSize int `default:"10"`
	}
}

type Security struct {
	JWTSecret     string `default:"default-jwt-secret"`
	JWTExpiration int    `default:"86400"` // seconds
	EncryptionKey string `default:"default-encryption-key"`

	CORS struct {
		AllowedOrigins   []string `default:"[\"*\"]"`
		AllowedMethods   []string `default:"[\"GET\", \"POST\", \"PUT\", \"DELETE\", \"OPTIONS\"]"`
		AllowedHeaders   []string `default:"[\"*\"]"`
		ExposedHeaders   []string
		AllowCredentials bool `default:"true"`
		MaxAge           int  `default:"86400"` // seconds
	}

	RateLimit struct {
		Enable            bool `default:"true"`
		RequestsPerMinute int  `default:"100"`
		BurstSize         int  `default:"20"`
		CleanupInterval   int  `default:"300"` // seconds
	}

	Auth struct {
		TokenType              string `default:"Bearer"`
		TokenHeader            string `default:"Authorization"`
		RefreshTokenExpiration int    `default:"604800"` // seconds
	}

	Password struct {
		MinLength        int  `default:"8"`
		RequireUppercase bool `default:"true"`
		RequireLowercase bool `default:"true"`
		RequireNumbers   bool `default:"true"`
		RequireSymbols   bool `default:"true"`
		MaxAttempts      int  `default:"5"`
		LockoutDuration  int  `default:"900"` // seconds
	}
}

type Services struct {
	Email   EmailService
	Storage StorageService
}

type EmailService struct {
	Provider string `default:"smtp"`
	Host     string
	Port     int `default:"587"`
	Username string
	Password string
	From     string
}

type StorageService struct {
	Provider       string `default:"local"`
	LocalPath      string `default:"storage"`
	Region         string
	Bucket         string
	AccessKey      string
	SecretKey      string
	Endpoint       string
	ForcePathStyle bool
}

type Monitoring struct {
	EnableMetrics   bool   `default:"false"`
	MetricsAddr     string `default:":9090"`
	EnableTracing   bool   `default:"false"`
	TracingEndpoint string
}

type Business struct {
	Package PackageConfig
	Version VersionConfig
	User    UserConfig
}

type PackageConfig struct {
	MaxSize      int64    `default:"104857600"` // 100MB
	AllowedTypes []string `default:"[\"ipa\", \"apk\", \"zip\"]"`
	StoragePath  string   `default:"packages"`
}

type VersionConfig struct {
	MaxVersions     int  `default:"100"`
	AutoCleanup     bool `default:"true"`
	CleanupInterval int  `default:"86400"` // seconds
}

type UserConfig struct {
	MaxUsers       int    `default:"1000"`
	DefaultRole    string `default:"developer"`
	SessionTimeout int    `default:"3600"` // seconds
}

type Developer struct {
	Name  string
	Email string
}
