package models

type Config struct {
	DbConnStr  string `mapstructure:"CONNECTION_STRING"`
	StartPath  string `mapstructure:"START_PATH"`
	MachineStr string `mapstructure:"MACHINE"`
	BatchSize  int    `mapstructure:"BATCH_SIZE"`
	Debug      bool   `mapstructure:"DEBUG"`
}
