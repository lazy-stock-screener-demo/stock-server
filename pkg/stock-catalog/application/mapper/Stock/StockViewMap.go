package catalogmapper

type ViewMap struct{}

func (v ViewMap) ToDTO(stockView catalogdomain.StockView) catalogdto.CatalogDTO {
	return
}

func NewViewMap() ViewMap {
	return ViewMap{}
}
