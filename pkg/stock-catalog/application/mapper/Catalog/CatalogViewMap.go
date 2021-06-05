package catalogmapper

import (
	catalogdto "stock-contexts/pkg/stock-catalog/application/dto"
	catalogdomain "stock-contexts/pkg/stock-catalog/domain"
	catalogschema "stock-contexts/pkg/stock-catalog/infra/repo/schema"
)

type ViewMap struct{}

// NewToDomain Implemented
func (v ViewMap) ToDomain(raw *catalogschema.Catalog) catalogdomain.StockView {
	stockTicker := catalogdomain.NewStockTicker(
		catalogdomain.StockTickerProps{
			Value: raw.Ticker,
		})

	stockView := catalogdomain.NewStockView(
		catalogdomain.StockViewProps{
			StockTicker:   stockTicker,
			ProfitDetails: NewProfitToDomain(raw),
			SafetyDetails: NewSafetyToDomain(raw),
		},
	)
	return stockView
}

// ToDTO define a interface to CatalogDTO
func (v ViewMap) ToDTO(stockView catalogdomain.StockView) catalogdto.CatalogDTO {
	return catalogdto.CatalogDTO{
		StockTicker: stockView.GetStockTicker().GetValue(),
		ProfitDetails: catalogdto.ProfitDTO{
			NetIncome:         stockView.GetProfitDetails().GetNetIncome().GetValue(),
			OperatingCashFlow: stockView.GetProfitDetails().GetOperatingCashFlow().GetValue(),
		},
		SafetyDetails: catalogdto.SafetyDTO{
			LongTermDebt: stockView.GetSafetyDetails().GetLongTermDebt().GetValue(),
			CurrentRatio: stockView.GetSafetyDetails().GetCurrentRatio().GetValue(),
		},
	}
}

func NewViewMap() ViewMap {
	return ViewMap{}
}

// // ToDomain Static Method
// func ToDomain(raw *catalogschema.Catalog) (stockView appcore.DataProps, err appcore.ErrProps) {
// 	stockTickerOrErr := catalogdomain.NewStockTicker(catalogdomain.StockTickerProps{
// 		Value: raw.Ticker,
// 	})
// 	// fmt.Println(fmt.Sprintf("%p", stockTickerOrErr.GetData()))
// 	stockViewOrErr := catalogdomain.NewStockView(
// 		catalogdomain.StockViewProps{
// 			StockTicker: stockTickerOrErr.GetData().Data.(catalogdomain.StockTicker),
// 		},
// 	)
// 	if stockViewOrErr.IsFailure {
// 		return appcore.DataProps{}, stockViewOrErr.GetErr()
// 	}
// 	fmt.Println(fmt.Sprintf("%p", stockViewOrErr.GetData().Data))
// 	return stockViewOrErr.GetData(), appcore.ErrProps{}
// }

// // ToDTO Static Method
// func ToDTO(stockView catalogdomain.StockView) catalogdto.StockDTO {
// 	return catalogdto.StockDTO{
// 		StockVID: stockView.GetStockVID().GetValue(),
// 	}
// }

// a, _ := json.Marshal(result)
// b := string(a)
// c := json.Unmarshal(b)
