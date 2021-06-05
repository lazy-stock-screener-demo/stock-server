package identityschema

import (
	"time"

	"github.com/google/uuid"
)

// User define a struct
type User struct {
	Sequence  int       `gorm:"column:sequence;autoIncrement:true"`
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	LicenseID uuid.UUID
	AccountID string `gorm:"column:account_id"`
	UserName  string `gorm:"column:customer_name"`
	UserPWD   []byte `gorm:"column:customer_pwd"`
	Email     string `gorm:"column:email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName define own table name
func (c *User) TableName() string {
	return "customer"
}
