package config

import (
	"github.com/apoprotsky/go-services"
)

// Service structure
type Service struct{}

// GetDBPath returns path to database file
func (svc *Service) GetDBPath() string {
	return "app.db"
}

// GetService returns Service instance
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
