package tcgaming

// Config represents the TC Gaming API configuration
// Based on Java SDK's TCGamingConfigObj
type Config struct {
	APIUrl       string
	MerchantCode string
	DESKey       string
	SHA256Key    string
}

// NewConfig creates a new configuration
func NewConfig(apiUrl, merchantCode, desKey, sha256Key string) *Config {
	return &Config{
		APIUrl:       apiUrl,
		MerchantCode: merchantCode,
		DESKey:       desKey,
		SHA256Key:    sha256Key,
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.APIUrl == "" {
		return NewAPIError(-1, "API URL is required")
	}
	if c.MerchantCode == "" {
		return NewAPIError(-1, "Merchant code is required")
	}
	if c.DESKey == "" {
		return NewAPIError(-1, "DES key is required")
	}
	if c.SHA256Key == "" {
		return NewAPIError(-1, "SHA256 key is required")
	}
	return nil
}