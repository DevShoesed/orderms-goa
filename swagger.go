package msorderms

import (
	"log"
	"orderms/internal/gateways/http/gen/swagger"
)

type swaggerSvc struct {
	logger *log.Logger
}

// NewSwagger return the swagger service implementation
func NewSwagger(logger *log.Logger) swagger.Service {
	return &swaggerSvc{logger}
}
