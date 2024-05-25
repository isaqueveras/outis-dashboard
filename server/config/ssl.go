package config

import "errors"

// SSL configure the settings for a SSL
type SSL struct {
	Mode        string `env:"OUTIS_SSL_MODE"`
	Certificate Certificate
}

// Validate validate data structure settings
func (c *SSL) Validar() error {
	switch c.Mode {
	case "disable", "require", "verify-ca", "verify-full":
	default:
		return errors.New("set a valid SSL operating mode: disable, require, verify-ca, verify-full")
	}

	if c.Mode == "disable" {
		return nil
	}

	if err := c.Certificate.Validate(); err != nil {
		return err
	}

	return nil
}
