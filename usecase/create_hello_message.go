package usecase

type ICreateHelloMessage interface {
	Execute(req *CreateHelloMessageRequest) (*CreateHelloMessageResponse, error)
}

type CreateHelloMessageRequest struct {
	Name string
}

type CreateHelloMessageResponse struct {
	Message string
}
