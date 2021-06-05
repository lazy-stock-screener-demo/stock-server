package customerselfschema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	Customers []Customer `gorm:"foreignKey:LicenseID;references:id"`
}

// TableName define own table name
func (c *License) TableName() string {
	return "license"
}

// BeforeCreate define a hook
func (c *License) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
