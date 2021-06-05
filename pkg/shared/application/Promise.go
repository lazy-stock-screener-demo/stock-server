package appcore

import "sync"

// https://stackoverflow.com/questions/35926173/implementing-promise-with-channels-in-go

// Promise in Golang
type Promise struct {
	wg  sync.WaitGroup
	res string
	err error
}

// Then method in Promise
func (p *Promise) Then(r func(string), e func(error)) {
	go func() {
			p.wg.Wait()
			if p.err != nil {
					e(p.err)
					return
			}
			r(p.res)
	}()
}

// NewPromise Func
func NewPromise(f func() (string, error)) *Promise {
	p := &Promise{}
	p.wg.Add(1)
	go func() {
			p.res, p.err = f()
			p.wg.Done()
	}()
	return p
}