package catalogmapper

import (
	catalogdomain "stock-contexts/pkg/stock-catalog/domain"
	catalogschema "stock-contexts/pkg/stock-catalog/infra/repo/schema"
)

// NewSafetyToDomain define a interface to SafetyDTO
func NewSafetyToDomain(raw *catalogschema.Catalog) catalogdomain.SafetyDetails {
	longTermDebt := catalogdomain.NewLongTermDebt(
		catalogdomain.LongTermDebtProps{
			Value: raw.Long_term_debt,
		})
	currentRatio := catalogdomain.NewCurrentRatio(
		catalogdomain.CurrentRatioProps{
			Value: raw.Current_ratio,
		})
	safetyDetails := catalogdomain.NewSafetyDetails(
		catalogdomain.SafetyDetailsProps{
			LongTermDebt: longTermDebt,
			CurrentRatio: currentRatio,
		})
	return safetyDetails
}
