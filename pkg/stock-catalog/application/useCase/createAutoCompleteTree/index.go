package autocompletesearch

import (
	ternarySearchTree "stock-contexts/pkg/shared/services/ternarySearchTree"
)

var w = newWorkerController()

func WorkerExecute() func() *ternarySearchTree.Tree {
	return func() *ternarySearchTree.Tree {
		return w.executeImpl()
	}
}
