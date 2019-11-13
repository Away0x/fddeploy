package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// InitConfig 初始化 config
func InitConfig(configFilePath, configFileType string) {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查 %s 配置文件是否存在: %v", configFilePath, err))
	}

	// 设置配置默认值
	setupDefaultConfig()

	// 环境变量 (设置环境变量: export FDDEPLOY_XX=xxx)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FDDEPLOY") // 环境变量前缀
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// String viper GetString
func String(key string) string {
	return viper.GetString(key)
}

// DefaultString viper GetString
func DefaultString(key string, defaultVal string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultVal
	}

	return v
}

// Int viper GetInt
func Int(key string) int {
	return viper.GetInt(key)
}

// DefaultInt viper GetInt
func DefaultInt(key string, defaultVal int) int {
	v := viper.GetInt(key)
	if v == 0 {
		return defaultVal
	}

	return v
}

// Bool viper GetBool
func Bool(key string) bool {
	return viper.GetBool(key)
}

// StringMap viper GetStringMap
func StringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}
