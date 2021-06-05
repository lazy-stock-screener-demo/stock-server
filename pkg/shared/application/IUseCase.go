package appcore

// IRequest interface
type IRequest struct {}

// IResponse interface
type IResponse struct {}

// IUseCase interface
type IUseCase interface{
	execute(request IRequest) IResponse
}