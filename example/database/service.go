package database

import (
	"database/sql"
	"github.com/apoprotsky/go-services"
	"github.com/apoprotsky/go-services/example/config"
	_ "github.com/mattn/go-sqlite3"
)

// Service structure
type Service struct {
	*sql.DB
}

// GoService method can be used to initialize service
func (svc *Service) GoService(configService *config.Service) {
	svc.DB, _ = sql.Open("sqlite3", configService.GetDBPath())
}

// GetService returns Service instance
// using service type for initialization
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
