package example

import (
	"github.com/go-playground/validator/v10"
	example "github.com/joaosoft/example/domain/config"
	"github.com/joaosoft/logger"
)

// ExampleOption ...
type ExampleOption func(client *Example)

// Reconfigure ...
func (example *Example) Reconfigure(options ...ExampleOption) {
	for _, option := range options {
		option(example)
	}
}

// WithConfiguration ...
func WithConfiguration(config *example.ExampleConfig) ExampleOption {
	return func(client *Example) {
		client.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) ExampleOption {
	return func(example *Example) {
		example.logger = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) ExampleOption {
	return func(example *Example) {
		example.logger.SetLevel(level)
	}
}

// WithValidator ...
func WithValidator(validator *validator.Validate) ExampleOption {
	return func(example *Example) {
		example.validator = validator
	}
}
