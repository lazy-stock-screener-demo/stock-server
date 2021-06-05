package catalogmapper

import (
	catalogdomain "stock-contexts/pkg/stock-catalog/domain"
	catalogschema "stock-contexts/pkg/stock-catalog/infra/repo/schema"
)

// NewProfitToDomain define a interface to ProfitDTO
func NewProfitToDomain(raw *catalogschema.Catalog) catalogdomain.ProfitDetails {
	operatingCashFlow := catalogdomain.NewOperatingCashFlow(
		catalogdomain.OperatingCashFlowProps{
			Value: raw.Operating_cash_flow,
		})
	netIncome := catalogdomain.NewNetIncome(
		catalogdomain.NetIncomeProps{
			Value: raw.Net_income,
		})
	// fmt.Println("raw.NetIncome", raw.Net_income)
	profitDetails := catalogdomain.NewProfitDetails(
		catalogdomain.ProfitDetailsProps{
			OperatingCashFlow: operatingCashFlow,
			NetIncome:         netIncome,
		})
	return profitDetails
}
