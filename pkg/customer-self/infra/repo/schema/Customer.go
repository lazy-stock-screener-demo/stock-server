package customerselfschema

import (
	domaincore "stock-contexts/pkg/shared/domain/core"
	idcore "stock-contexts/pkg/shared/domain/id"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customer define a struct
type Customer struct {
	// gorm.Model
	Sequence     int       `gorm:"column:sequence;autoIncrement:true"`
	ID           uuid.UUID `gorm:"column:id;primaryKey"`
	LicenseID    uuid.UUID
	CustomerName string `gorm:"column:customer_name;not null;index:idx_customer_name,unique"`
	CustomerPWD  []byte `gorm:"column:customer_pwd;not null"`
	Email        string `gorm:"column:email"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TableName define own table name
func (c *Customer) TableName() string {
	return "customer"
}

// // BeforeCreate define a hook
// func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
// 	c.ID = uuid.New()
// 	return
// }

func (u *Customer) AfterCreate(tx *gorm.DB) (err error) {
	domaincore.DispatchEvents(idcore.NewIDEntity(u.ID.String()))
	return
}

type CustomerView struct {
	CustomerName string
	Email        string
	LicenseName  string
}
