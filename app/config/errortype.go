package config

// アプリケーション環境情報エラー
type ConfigError struct {
	message string
}

func (c *ConfigError) Error() string {
	return c.message
}

func NewConfigError(message string) *ConfigError {
	return &ConfigError{message: message}
}