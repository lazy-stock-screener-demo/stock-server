package migration

import (
	customerselfschema "stock-contexts/pkg/customer-self/infra/repo/schema"
	gormclient "stock-contexts/pkg/shared/infra/repo/gorm/config"

	"gorm.io/gorm"
)

type Migrate struct {
	client *gorm.DB
}

func (m *Migrate) Execute() {
	m.client.AutoMigrate(customerselfschema.License{}, customerselfschema.Customer{})
}

func NewMigration() *Migrate {
	return &Migrate{
		client: gormclient.NewConnectedGorm(),
	}
}
