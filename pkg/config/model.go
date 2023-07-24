package config

type Config struct {
	BindAddress   string `yaml:"bind-address"`
	Authorization string `yaml:"authorization"`

	DSN      string `yaml:"dsn"`
	DebugDSN string `yaml:"debug-dsn"`

	CutOffTime string `yaml:"cutoff-time"`
	Production bool   `yaml:"production"`
	Cronjob    string `yaml:"cron-job"`

	LogLevel       string `yaml:"log-level"`
	LogDevelopment bool   `yaml:"log-development"`
}
