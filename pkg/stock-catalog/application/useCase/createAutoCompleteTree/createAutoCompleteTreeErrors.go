package autocompletesearch

import (
	appcore "stock-contexts/pkg/shared/application"
)

func NewTickerNotFoundErr(err interface{}) appcore.Result {
	// t = template.Must(template.New("t2").
	// 	Parse("Couldn't find a stock by ticker view id: {{.a}}\n"))
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "Couldn't find any stock ticker",
			Type:    "TickerNotFound",
			Error:   err,
		},
	)
}
