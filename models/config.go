package models

type Config struct {
	DbConnStr string `mapstructure:"db_conn"`
	StartPath string `mapstructure:"startpath"`
}
