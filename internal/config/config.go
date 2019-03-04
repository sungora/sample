package config

type configTyp struct {
	SampleCustomConfig *sampleCustomConfig
}

type sampleCustomConfig struct {
	Name     string
	IsAccess bool
	Balance  float64
}
