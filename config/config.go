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
	// UnixSocketPermission defaults to 0o770 so only the owner and group can access the socket.
	// Changed from upstream default (0o640) to allow group members to connect without sudo.
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
}
