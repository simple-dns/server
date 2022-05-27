package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

type AuthConfig struct {
	Username string `envconfig:"WEB_USERNAME" default:""`
	Password string `envconfig:"WEB_PASSWORD" default:""`
}

type HttpConfig struct {
	Auth      AuthConfig
	HttpPort  string `envconfig:"WEB_PORT" default:"80"`
	RpcSecret string `envconfig:"RPC_SECRET" default:""`
}

type DnsConfig struct {
	DnsPort         string `envconfig:"DNS_PORT" default:"53"`
	ForwarderServer string `envconfig:"FORWARDER_SERVER" default:""`
	DomainSuffix    string `envconfig:"DOMAIN_SUFFIX" default:""`
}
type Config struct {
	Http  HttpConfig
	Dns   DnsConfig
	Debug bool `envconfig:"DEBUG" default:"false"`
}

type configCache struct {
	ConfigCache *Config
	Lock        sync.Mutex
}

var cache = &configCache{}

func GetConfig() Config {
	if cache.ConfigCache != nil {
		return *cache.ConfigCache
	}
	cache.Lock.Lock()
	defer cache.Lock.Unlock()
	cache.ConfigCache = &Config{}
	err := envconfig.Process("", cache.ConfigCache)
	if err != nil {
		log.Fatal(err.Error())
	}
	return *cache.ConfigCache
}

func getAuthConfig() AuthConfig {
	return GetConfig().Http.Auth
}
func GetDnsConfig() DnsConfig {
	return GetConfig().Dns
}
func GetHttpConfig() HttpConfig {
	return GetConfig().Http
}

func GetCredentials() map[string]string {
	authConfig := getAuthConfig()
	if authConfig.Username != "" {
		return map[string]string{
			authConfig.Username: authConfig.Password,
		}
	} else {
		return map[string]string{}
	}
}

func GetDebug() bool {
	return GetConfig().Debug
}
func GetRpcSecret() string {
	return GetConfig().Http.RpcSecret
}

func GetDomainSuffix() string {
	return GetConfig().Dns.DomainSuffix
}
