package config

type Wx struct {
	AppID     string `mapstructure:"appId" json:"appId" yaml:"appId"`             // 是否启用
	AppSecret string `mapstructure:"appSecret" json:"appSecret" yaml:"appSecret"` // CRON表达式
}
