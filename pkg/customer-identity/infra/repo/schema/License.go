package identityschema

import (
	"time"

	"github.com/google/uuid"
)

// License define a struct
type License struct {
	// gorm.Model
	Sequence    int       `gorm:"column:sequence;autoIncrement:true"`
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	LicenseName string    `gorm:"column:license_name;notNull"`
	IsEnable    bool      `gorm:"column:is_enable"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	//Customers defines the relationship between license and customer
	User []User `gorm:"foreignKey:LicenseID;references:id"`
}

// TableName define own table name
func (c *License) TableName() string {
	return "license"
}
