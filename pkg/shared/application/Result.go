package appcore

// // DataProps define app data
// type DataProps struct {
// 	Data interface{}
// }

// ErrProps define app error
type ErrProps struct {
	Message string
	Type    string
	Error   interface{}
}

// // ResultProps Struct
// type ResultProps struct {
// 	IsSuccess bool
// 	IsFailure bool
// 	errData   ErrProps
// 	data      DataProps
// }

// // IResult define interface
// type IResult interface {
// 	GetData() DataProps
// 	GetErr() ErrProps
// }

// // NewResultProps as constructor
// func NewResultProps(isSuccess bool, errData ErrProps, data DataProps) ResultProps {
// 	return ResultProps{
// 		IsSuccess: isSuccess,
// 		IsFailure: !isSuccess,
// 		errData:   errData,
// 		data:      data,
// 	}
// }

// // Result Struct
// type Result struct {
// 	ResultProps
// }

// // GetData Method for access value
// func (r Result) GetData() DataProps {
// 	// err := errors.New("can't get value; Use GetErr instead")
// 	// if !r.IsSuccess {
// 	// 	return nil, err
// 	// }
// 	return r.data
// }

// // GetErr Method for access error value
// func (r Result) GetErr() ErrProps {
// 	return r.errData
// }

// // NewResultOk encapsulate domain object
// func NewResultOk(value interface{}) Result {
// 	return Result{NewResultProps(true, ErrProps{}, DataProps{Data: value})}
// }

// // NewResultFailed encapsulate domain object
// func NewResultFailed(e ErrProps) Result {
// 	return Result{NewResultProps(false, e, DataProps{})}
// }

// // NewResultUseCaseErr define base errpr struct
// func NewResultUseCaseErr(e ErrProps) Result {
// 	return Result{
// 		NewResultProps(
// 			true,
// 			e,
// 			DataProps{},
// 		),
// 	}
// }

// // NewResultCombine all the results and return a fail result
// func NewResultCombine(results []Result) Result {
// 	for _, result := range results {
// 		if result.IsFailure {
// 			return result
// 		}
// 	}
// 	return NewResultOk(nil)
// }

type ResultProps struct {
	isSuccess bool
	errData   ErrProps
}

// NewResultProps2 as constructor
func NewResultProps(isSuccess bool, errData ErrProps) ResultProps {
	return ResultProps{
		isSuccess: isSuccess,
		errData:   errData,
	}
}

// IResult define interface
type IResult interface {
	GetErr() ErrProps
	IsSuccess() bool
}

// Result2 Struct
type Result struct {
	ResultProps
}

// GetErr Method for access error value
func (r Result) GetErr() ErrProps {
	return r.errData
}

// IsSuccess Method for access error value
func (r Result) IsSuccess() bool {
	return r.ResultProps.isSuccess
}

// NewResultOk encapsulate domain object
func NewResultOk() Result {
	return Result{NewResultProps(true, ErrProps{})}
}

// NewResultFailed encapsulate domain object
func NewResultFailed(e ErrProps) Result {
	return Result{NewResultProps(false, e)}
}

// NewResultUseCaseErr define base errpr struct
func NewResultUseCaseErr(e ErrProps) Result {
	return Result{
		NewResultProps(
			true,
			e,
		),
	}
}

// NewResultCombine all the results and return a fail result
func NewResultCombine(results []Result) Result {
	for _, result := range results {
		if !result.IsSuccess() {
			return result
		}
	}
	return NewResultOk()
}
