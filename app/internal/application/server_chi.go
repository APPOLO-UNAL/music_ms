// This script contains the configuration of the chi router
package application

// ConfigServerChi configures the chi router
type ConfigServerChi struct {
	// Addr is the address where the server will listen to
	Addr string
	// ReadTimeout is the maximum duration for reading the entire request, including the body
	ReadTimeout int
	// WriteTimeout is the maximum duration before timing out writes of the response
	WriteTimeout int
	// IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled
	IdleTimeout int
	// Port is the port where the server will listen to
	Port int
}

// NewConfigServerChi creates a new ConfigServerChi
func NewServerChi(cfg *ConfigServerChi) *ConfigServerChi {
	// Default Config
	defaultConf := &ConfigServerChi{
		Addr:         "8080",
		ReadTimeout:  5,
		WriteTimeout: 10,
		IdleTimeout:  15,
		Port:         8080,
	}
	if cfg != nil {
		defaultConf = cfg
	}
	return defaultConf
}

// Run runs the server

func (s *ConfigServerChi) Run() (err error) {

	// Depedencies
	// - Database Connection

	// Repository

	// Service

	// Handler

	// Router

	// Middlewares

	// Endpoints

	// - Endpoints Music

	// - Endpoints Artist

	// - Endpoints Album
	return
}
