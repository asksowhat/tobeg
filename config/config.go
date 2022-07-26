package config

type Config struct {
	BaseConfig   BaseConfig   `mapstructure:"base" json:"base"`
	ServerConfig ServerConfig `mapstructure:"server" json:"server"`
	AlipayConfig AlipayConfig `mapstructure:"alipay" json:"alipay"`
}

type BaseConfig struct {
	Title   string   `mapstructure:"title" json:"title"`
	Url     string   `mapstructure:"url" json:"url"`
	Favicon string   `mapstructure:"favicon" json:"favicon"`
	Thank   string   `mapstructure:"thank" json:"thank"`
	ToSells []string `mapstructure:"toSells" json:"toSells"`
}

type ServerConfig struct {
	Port int `mapstructure:"port" json:"port"`
}

type AlipayConfig struct {
	AppId             string `mapstructure:"appid" json:"appid"`
	PrivateKey        string `mapstructure:"private_key" json:"private_key"`
	PublicContentRSA2 string `mapstructure:"public_content_rsa2" json:"public_content_rsa2"`
	AppPublicContent  string `mapstructure:"app_public_content" json:"app_public_content"`
}
