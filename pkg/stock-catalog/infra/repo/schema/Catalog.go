package catalogschema

import (
	"github.com/google/uuid"
)

type Catalog struct {
	Stock_id            uuid.UUID
	Ticker              string
	Operating_cash_flow []map[string]float32
	Net_income          []map[string]float32
	Current_ratio       []map[string]float64
	Long_term_debt      []map[string]float32
	Free_cash_flow      []map[string]float32
}

// TableName define own table name
func (c *Catalog) TableName() string {
	return "catalog"
}
