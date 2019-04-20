package users

type Config struct {
	Name     string  `yaml:"Name"`
	IsAccess bool    `yaml:"IsAccess"`
	Balance  float64 `yaml:"Balance"`
}
