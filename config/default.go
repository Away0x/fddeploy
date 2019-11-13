package config

import (
	"github.com/spf13/viper"
)

// 默认配置
var defaultConfigMap = map[string]interface{}{}

// 设置配置默认值
func setupDefaultConfig() {
	for k, v := range defaultConfigMap {
		viper.SetDefault(k, v)
	}
}
