package config

type configMain struct {
	SampleConfig *sampleCustomConfig
}

type sampleCustomConfig struct {
	Name     string
	IsAccess bool
	Balance  float64
}

var SampleCustomConfig *sampleCustomConfig
