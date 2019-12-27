package config

//ref: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64

import (
	"fmt"
	"os"
	"utilities"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

//Config ...
type Config struct {
	ConfigFileDir string
	SSFC          SFC `yaml:"sfc"`
}

// type Snort struct {
// 	LogDir         string `yaml:"log_dir", envconfig:"SNORT_LOG_DIR"`             //e.g. "/var/log/snort/alert"
// 	IPBlacklistDir string `yaml:"ip_blacklist_dir", envconfig:"IP_BLACKLIST_DIR"` // e.g.  "/etc/snort/rules/iplists/black_list.rules"
// 	Interfaces     string `yaml:"interfaces", envconfig:"SNORT_INTERFACES"`       //e.g. "enp0s3:enp0s8"
// 	ConfDir        string `yaml:"conf_dir", envconfig:"SNORT_CONF_DIR"`           //e.g. "/etc/snort/snort.conf"
// }

type SFC struct {
	Info  []NetworkFunctionsInfo `yaml:"network_functions_info,flow", envconfig:"NETWORK_FUNCTIONS_INFO"`
	Chain []string               `yaml:"chain,flow", envconfig:"CHAIN_LIST"`
}

type NetworkFunctionsInfo struct {
	NameOfFunction string `yaml:"name_of_function", envconfig:"FUNCTION_NAME"`
	TypeOfFunction string `yaml:"type_of_function", envconfig:"TYPE_OF_FUNCTION"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(configDir string, cfg *Config) {
	f, err := os.Open(configDir)
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

//Set sets configuration and also sets env variables
func Set(configDir string) *Config {
	//todo: check whether env variable is being set or not
	var cfg Config
	cfg.ConfigFileDir = configDir
	readFile(configDir, &cfg)
	readEnv(&cfg)
	//fmt.Printf("%+v", cfg)
	return &cfg
}

//CheckIfConfigFileExists ...
func (cfg *Config) CheckIfConfigFileExists() bool {
	return utilities.IsFileExist(cfg.ConfigFileDir)
}
