package config

import "errors"

// Certificate configure the settings for a certificate
type Certificate struct {
	Cert string `env:"OUTIS_CERTIFICATE_CERT"`
	Key  string `env:"OUTIS_CERTIFICATE_KEY"`
	CA   string `env:"OUTIS_CERTIFICATE_CA"`
}

// Validate validate data structure settings
func (c *Certificate) Validate() error {
	if c.Cert == "" {
		return errors.New("defines the value of the certificate")
	}

	if c.Key == "" {
		return errors.New("defines the key path")
	}

	if c.CA == "" {
		return errors.New("defines the certificate authority path")
	}

	return nil
}
