package config

type Wxpay struct {
	MchID                      string `mapstructure:"mch-id" json:"mchId" yaml:"mch-id"`                                                                    // 商户ID
	AppID                      string `mapstructure:"app-id" json:"appId" yaml:"app-id"`                                                                    // 绑定小程序的APPID
	NotifyUrl                  string `mapstructure:"notify-url" json:"notifyUrl" yaml:"notify-url"`                                                        // 支付回调域名
	MchCertificateSerialNumber string `mapstructure:"mch-certificate-serial-number" json:"mchCertificateSerialNumber" yaml:"mch-certificate-serial-number"` // 商户证书序列号
	MchAPIv3Key                string `mapstructure:"mch-api-v3-key" json:"mchAPIv3Key" yaml:"mch-api-v3-key"`                                              // 商户APIv3密钥
	PemPath                    string `mapstructure:"pem-path" json:"pemPath" yaml:"pem-path"`                                                              // 证书文件所在地址
}
