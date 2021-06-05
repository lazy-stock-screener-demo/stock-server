package catalogdomain

import (
	appcore "stock-contexts/pkg/shared/application"
	domaincore "stock-contexts/pkg/shared/domain/core"
)

// ProfitDetailsProps struct
type ProfitDetailsProps struct {
	OperatingCashFlow
	NetIncome
}

// ProfitDetails define
type ProfitDetails struct {
	valueObject domaincore.IValueObject
	props       ProfitDetailsProps
	Result      appcore.IResult
}

func (p ProfitDetails) GetOperatingCashFlow() OperatingCashFlow {
	return p.props.OperatingCashFlow
}

func (p ProfitDetails) GetNetIncome() NetIncome {
	return p.props.NetIncome
}

func NewProfitDetails(props ProfitDetailsProps) ProfitDetails {
	return ProfitDetails{
		valueObject: domaincore.NewValueObject(props),
		props:       props,
		Result:      appcore.NewResultOk(),
	}
}
