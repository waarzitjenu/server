package types

// Config contains the possible configuration parameters that are available in the settings.json file.
type Config struct {
	Debug               bool                `json:"debug"`
	ServerConfiguration ServerConfiguration `json:"server"`
}

// ServerConfiguration contains configuration parameters specific to the server.
type ServerConfiguration struct {
	Port uint `json:"port"`
	TLS  TLS  `json:"tls"`
}

// TLS contains TLS configuration parameters.
type TLS struct {
	Enabled     bool                  `json:"enabled"` // Enabled tells the server to use TLS. One would set this to false when running it locally, or using a reverse proxy setup instead of relying on the built-in TLS.
	Certificate CertificateProperties `json:"certificate"`
}

// CertificateProperties contains the public key, private key and CA bundle commonly used in TLS communication.
type CertificateProperties struct {
	PublicKey  string `json:"crt"`             // The public key, usually a .crt file.
	PrivateKey string `json:"key"`             // The private key, usually a .key file.
	CABundle   string `json:"cacrt,omitempty"` // The CA Bundle (e.g. public keys of Let's Encrypt root CAs), usually a .crt or .pem file.
}
