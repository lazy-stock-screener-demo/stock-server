package catalogrepo

import (
	"context"
	appcore "stock-contexts/pkg/shared/application"
	mongoclient "stock-contexts/pkg/shared/infra/repo/mongodriver/config"
	catalogmapper "stock-contexts/pkg/stock-catalog/application/mapper/Catalog"
	catalogdomain "stock-contexts/pkg/stock-catalog/domain"
	catalogschema "stock-contexts/pkg/stock-catalog/infra/repo/schema"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// IRead define Read interface
type IRead interface {
	ReadStocksByTicker(VID string) (catalogdomain.StockView, appcore.ErrProps)
	ReadStockTickerList() ([]string, appcore.ErrProps)
}

// Read Define reac operation of database ODM
type Read struct {
	client *mongo.Client
}

// ReadStocksByTicker define the method to read of stock by ticker
func (r *Read) ReadStocksByTicker(VID string) (catalogdomain.StockView, appcore.ErrProps) {
	collection := r.client.Database("stock-db").Collection("catalog")
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var stock catalogschema.Catalog
	collection.FindOne(context.TODO(), bson.M{"ticker": strings.ToUpper(VID)}).Decode(&stock)
	// fmt.Printf("POST stockVID=%s <br/>", stock)

	stockView := catalogmapper.NewViewMap().ToDomain(&stock)
	return stockView, appcore.ErrProps{}
}

func (r *Read) ReadStockTickerList() ([]string, appcore.ErrProps) {
	list := []string{
		"aardvark",
		"altimeter",
		"apotactic",
		"bagonet",
		"boatlip",
		"carburant",
		"chyliferous",
		"consonance",
		"cyclospondylic",
		"dictyostele",
		"echelon",
		"estadal",
		"flaunty",
		"gesneriaceous",
		"hygienic",
		"infracentral",
		"jipijapa",
		"lipoceratous",
		"melanthaceae",
	}
	return list, appcore.ErrProps{}
}

// NewReadRepo define ReadRepo Instance
func NewReadRepo() *Read {
	return &Read{
		client: mongoclient.NewConnectedMongoDriver(),
	}
}
