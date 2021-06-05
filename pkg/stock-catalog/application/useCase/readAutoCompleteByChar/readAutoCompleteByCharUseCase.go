package readautocompletebychar

import (
	"fmt"
	appcore "stock-contexts/pkg/shared/application"
	ternarySearchTreeService "stock-contexts/pkg/shared/services/ternarySearchTree"
	autocomplete "stock-contexts/pkg/worker-pool"
	"time"
)

var TSTTreeSington *ternarySearchTreeService.Tree

type iuseCase interface {
	execute(req ReqDTO) ([]string, appcore.EitherError)
}

type useCase struct {
	TSTTreeChan chan *ternarySearchTreeService.Tree
}

func (u *useCase) execute(req ReqDTO) ([]string, appcore.EitherError) {
	if TSTTreeSington == nil {
		for {
			select {
			case TSTTree := <-u.TSTTreeChan:
				TSTTreeSington = TSTTree
				close(u.TSTTreeChan)
				return TSTTreeSington.Search(req.StockChar), appcore.NewEitherErr(
					appcore.NewSuccess(),
					appcore.NewResultOk(),
				)
			default:
				fmt.Println("WAITING...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
	return TSTTreeSington.Search(req.StockChar), appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

func newUseCase() *useCase {
	return &useCase{
		TSTTreeChan: autocomplete.NewTreeInstanceChanSington(), // => wait here
	}
}
