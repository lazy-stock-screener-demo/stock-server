package appcore

// IEitherError define any struct that implementing these method is belong to this type
type IEitherError interface {
	IsError() bool
	IsSuccess() bool
}

// // EitherError struct
// type EitherError struct {
// 	IEitherError
// 	Result IResult
// }

// EitherError struct
type EitherError struct {
	IEitherError
	Result IResult
}

// // NewEitherErr constructor
// func NewEitherErr(SuccessOrErr IEitherError, Result IResult) EitherError {
// 	return EitherError{
// 		SuccessOrErr,
// 		Result,
// 	}
// }

// NewEitherErr constructor
func NewEitherErr(SuccessOrErr IEitherError, Result IResult) EitherError {
	return EitherError{
		SuccessOrErr,
		Result,
	}
}

// ErrorStruct encapsulate result with error class
type ErrorStruct struct{}

// IsError define method that return bool value when success
func (e *ErrorStruct) IsError() bool {
	return true
}

// IsSuccess define method that return bool value when Error
func (e *ErrorStruct) IsSuccess() bool {
	return false
}

// NewErr Function
func NewErr() *ErrorStruct {
	return &ErrorStruct{}
}

// SuccessStruct encapsulate result with success class
type SuccessStruct struct{}

// IsError define method that return bool value when success
func (s *SuccessStruct) IsError() bool {
	return false
}

// IsSuccess define method that return bool value when success
func (s *SuccessStruct) IsSuccess() bool {
	return true
}

// NewSuccess Function
func NewSuccess() *SuccessStruct {
	return &SuccessStruct{}
}
