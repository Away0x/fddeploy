package cmd

import (
	"fmt"
	"os"

	"github.com/Away0x/fddeploy/config"

	"github.com/spf13/cobra"
)

const (
	// 默认配置文件路径
	defaultConfigFilePath = "config.yaml"
	// 配置文件格式
	configFileType = "yaml"
)

var (
	configFilePath string

	rootCmd = &cobra.Command{
		Use: "ffdeploy",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	// 配置文件路径: --cofig config_path
	rootCmd.PersistentFlags().
		StringVar(&configFilePath, "config", defaultConfigFilePath, "config file (default is $APP/"+defaultConfigFilePath+")")
}

func initConfig() {
	// setup config
	if configFilePath == "" {
		configFilePath = defaultConfigFilePath
	}

	config.InitConfig(configFilePath, configFileType)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
