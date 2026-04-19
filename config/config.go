package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config holds the configuration for the headscale server.
type Config struct {
	ServerURL          string        `mapstructure:"server_url"`
	Addr               string        `mapstructure:"listen_addr"`
	MetricsAddr        string        `mapstructure:"metrics_listen_addr"`
	GRPCAddr           string        `mapstructure:"grpc_listen_addr"`
	GRPCAllowInsecure  bool          `mapstructure:"grpc_allow_insecure"`
	EphemeralNodeInactivityTimeout time.Duration `mapstructure:"ephemeral_node_inactivity_timeout"`
	NodeUpdateCheckInterval        time.Duration `mapstructure:"node_update_check_interval"`
	IPPrefixes         []string      `mapstructure:"ip_prefixes"`
	PrivateKeyPath     string        `mapstructure:"private_key_path"`
	NoisePrivateKeyPath string       `mapstructure:"noise_private_key_path"`
	BaseDomain         string        `mapstructure:"base_domain"`
	DNSConfig          DNSConfig     `mapstructure:"dns_config"`
	DBtype             string        `mapstructure:"db_type"`
	DBpath             string        `mapstructure:"db_path"`
	DBhost             string        `mapstructure:"db_host"`
	DBport             int           `mapstructure:"db_port"`
	DBname             string        `mapstructure:"db_name"`
	DBuser             string        `mapstructure:"db_user"`
	DBpass             string        `mapstructure:"db_pass"`
	TLSLetsEncryptHostname    string `mapstructure:"tls_letsencrypt_hostname"`
	TLSLetsEncryptCacheDir    string `mapstructure:"tls_letsencrypt_cache_dir"`
	TLSLetsEncryptChallenge   string `mapstructure:"tls_letsencrypt_challenge_type"`
	TLSCertPath               string `mapstructure:"tls_cert_path"`
	TLSKeyPath                string `mapstructure:"tls_key_path"`
	ACMEEmail                 string `mapstructure:"acme_email"`
	ACMEUrl                   string `mapstructure:"acme_url"`
	OIDCConfig                OIDCConfig `mapstructure:"oidc"`
	Logtail                   LogtailConfig `mapstructure:"logtail"`
	RandomizeClientPort       bool   `mapstructure:"randomize_client_port"`
	ACLPolicyPath             string `mapstructure:"acl_policy_path"`
	UnixSocket                string `mapstructure:"unix_socket"`
	UnixSocketPermission      uint32 `mapstructure:"unix_socket_permission"`
}

// DNSConfig holds DNS-related configuration.
type DNSConfig struct {
	OverrideLocalDNS bool     `mapstructure:"override_local_dns"`
	Nameservers      []string `mapstructure:"nameservers"`
	RestrictedNameservers map[string][]string `mapstructure:"restricted_nameservers"`
	Domains          []string `mapstructure:"domains"`
	MagicDNS         bool     `mapstructure:"magic_dns"`
	BaseDomain       string   `mapstructure:"base_domain"`
}

// OIDCConfig holds OpenID Connect configuration.
type OIDCConfig struct {
	Issuer           string            `mapstructure:"issuer"`
	ClientID         string            `mapstructure:"client_id"`
	ClientSecret     string            `mapstructure:"client_secret"`
	Scope            []string          `mapstructure:"scope"`
	ExtraParams      map[string]string `mapstructure:"extra_params"`
	AllowedDomains   []string          `mapstructure:"allowed_domains"`
	AllowedUsers     []string          `mapstructure:"allowed_users"`
	StripEmaildomain bool              `mapstructure:"strip_email_domain"`
}

// LogtailConfig holds Logtail/Tailscale logging configuration.
type LogtailConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// LoadConfig reads and parses the configuration file using viper.
func LoadConfig(path string) (*Config, error) {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/headscale/")
		viper.AddConfigPath("$HOME/.headscale")
		viper.AddConfigPath(".")
	}

	viper.SetEnvPrefix("headscale")
	viper.AutomaticEnv()

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// setDefaults configures sensible default values for optional settings.
func setDefaults() {
	viper.SetDefault("listen_addr", "0.0.0.0:8080")
	viper.SetDefault("metrics_listen_addr", "127.0.0.1:9090")
	viper.SetDefault("grpc_listen_addr", "0.0.0.0:50443")
	viper.SetDefault("grpc_allow_insecure", false)
	viper.SetDefault("ephemeral_node_inactivity_timeout", "120s")
	viper.SetDefault("node_update_check_interval", "10s")
	viper.SetDefault("db_type", "sqlite3")
	viper.SetDefault("db_path", "/var/lib/headscale/db.sqlite")
	viper.SetDefault("unix_socket", "/var/run/headscale/headscale.sock")
	viper.SetDefault("unix_socket_permission", "0o770")
	viper.SetDefault("randomize_client_port", false)
	viper.SetDefault("logtail.enabled", false)
	viper.SetDefault("dns_config.magic_dns", true)
}
