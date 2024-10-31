package service

import (
	"net/http"
)

// 远程配置类
/* type Config struct {
	apiKey string // 鉴权key
	client *http.Client
}

func NewConfig(apiKey string) *Config {
	client := &http.Client{}
	config := &Config{apiKey: apiKey, client: client}
	return config
}

// FetchConfigByKey 从远程拉取配置
func (c *Config) FetchConfigByKey(key string) (string, error) {
	url := fmt.Sprintf("http://example.com/config?apiKey=%s&key=%s", c.apiKey, key)
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
*/

// ConfigBuilder结构体用于构建Config类
/* type ConfigBuilder struct {
	apiKey  string // 鉴权key
	client  *http.Client
	timeout int
	cluster string
}

// NewConfigBuilder创建一个新的ConfigBuilder实例
func NewConfigBuilder(apiKey string) *ConfigBuilder {
	return &ConfigBuilder{
		apiKey:  apiKey,
		client:  &http.Client{},
		timeout: 60,
		cluster: "default",
	}
}

// 设置超时时间
func (cfb *ConfigBuilder) SetTimeout(timeout int) *ConfigBuilder {
	cfb.timeout = timeout
	return cfb
}

// SetCluster设置请求的下游集群
func (cfb *ConfigBuilder) SetCluster(cluster string) *ConfigBuilder {
	cfb.cluster = cluster
	return cfb
}

// Build构建Config实例
func (cfb *ConfigBuilder) Build() (*Config, error) {
	return &Config{
		apiKey:  cfb.apiKey,
		client:  cfb.client,
		timeout: cfb.timeout,
		cluster: cfb.cluster,
	}, nil

}

// Config结构体用于从远程拉取配置
type Config struct {
	apiKey  string // 鉴权key
	client  *http.Client
	timeout int
	cluster string
}

// FetchConfig执行从远程按key拉取配置的操作
func (c *Config) FetchConfigByKey(key string) (string, error) {
	return "", nil
}
*/

// Config结构体用于从远程HTTP拉取配置
type Config struct {
	apiKey  string // 鉴权key
	client  *http.Client
	timeout int
	cluster string
}

// Option是函数选项类型，用于设置Config的属性
type Option func(*Config)

// WithTimeout函数选项用于设置超时时间
func WithTimeout(timeout int) Option {
	return func(cf *Config) {
		cf.timeout = timeout
	}
}

// WithCluster函数选项用于设置调用集群
func WithCluster(cluster string) Option {
	return func(cf *Config) {
		cf.cluster = cluster
	}
}

// NewConfig创建一个新的Config实例，并根据传入的函数选项进行设置
func NewConfig(apiKey string, opts ...Option) (*Config, error) {
	cf := &Config{
		apiKey:  apiKey,
		client:  &http.Client{},
		timeout: 60,
		cluster: "default",
	}

	for _, opt := range opts {
		opt(cf)
	}
	return cf, nil
}

// FetchConfig执行从远程HTTP按指定参数拉取配置的操作
func (c *Config) FetchConfigByKey(key string) (string, error) {
	return "", nil
}
