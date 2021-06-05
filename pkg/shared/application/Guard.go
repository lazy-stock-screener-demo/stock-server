package appcore

import "fmt"

type GuardResult struct {
	Passed  bool
	Message string
}
type Guard struct{}

// AgainstAtLeast check char limit
func (g *Guard) AgainstAtLeast(limitNum int, text string) GuardResult {
	if len(text) > limitNum {
		return GuardResult{
			Passed: true,
		}
	}
	return GuardResult{
		Passed:  false,
		Message: fmt.Sprintf("Text length is not at least %v chars", limitNum),
	}
}

// AgainstAtMost check char limit
func (g *Guard) AgainstAtMost(limitNum int, text string) GuardResult {
	if len(text) < limitNum {
		return GuardResult{
			Passed: true,
		}
	}
	return GuardResult{
		Passed:  false,
		Message: fmt.Sprintf("Text length is not at most %v chars", limitNum),
	}
}
func (g *Guard) AgainstNil() {}

func NewGurad() Guard {
	return Guard{}
}
