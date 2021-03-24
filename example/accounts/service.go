package accounts

import (
	"github.com/apoprotsky/go-services"
	"github.com/apoprotsky/go-services/example/database"
)

// Account structure
type Account struct {
	Name  string
	Email string
}

// Service structure
// Public fields of service structure will be initialized without GoService method
type Service struct {
	DatabaseService *database.Service
}

// GetByName returns Account structure by name
func (svc *Service) GetByName(name string) *Account {
	//var email string
	//row := svc.DatabaseService.QueryRow("SELECT email FROM accounts WHERE name = $1", name)
	//_ = row.Scan(&email)
	return &Account{
		Name:  name,
		Email: "admin@example.com",
	}
}

// GetService returns Service instance
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
