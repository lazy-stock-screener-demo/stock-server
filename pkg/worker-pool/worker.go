package worker

import (
	"fmt"
	ternarySearchTree "stock-contexts/pkg/shared/services/ternarySearchTree"
	createTree "stock-contexts/pkg/stock-catalog/application/useCase/createAutoCompleteTree"
	"sync"
)

var treeInstanceChanSington chan *ternarySearchTree.Tree

func ExecuteWorkerJob(workerControllerChan chan func() *ternarySearchTree.Tree, wg *sync.WaitGroup) {
	defer wg.Done()
	for workerController := range workerControllerChan {
		fmt.Println("[Working Run]")
		// Trigger worker controller just like trigger by API
		// But here we trigger in another thread.
		NewTreeInstanceChanSington() <- workerController()
	}
}

func NewTreeInstanceChanSington() chan *ternarySearchTree.Tree {
	if treeInstanceChanSington == nil {
		treeInstanceChanSington = make(chan *ternarySearchTree.Tree, 10)
		return treeInstanceChanSington
	}
	return treeInstanceChanSington
}

func NewWorker() {
	workerControllerChan := make(chan func() *ternarySearchTree.Tree, 10)
	wg := new(sync.WaitGroup)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go ExecuteWorkerJob(workerControllerChan, wg)
	}
	workerControllerChan <- createTree.WorkerExecute()

	close(workerControllerChan)
}
