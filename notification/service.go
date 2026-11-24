package notification

import (
	"context"
	"keka/errors"
	"log"
)

type Service interface {
	Call(ctx context.Context, req *CallRequest) (*CallResponse, error)
}

type MockCallService struct{}

func NewMockCallService() *MockCallService {
	return &MockCallService{}
}

func (*MockCallService) Call(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	if req.Message == "+919900445566" {
		return nil, errors.ErrPhoneNumUnknown
	}

	log.Printf("call performed\n")

	return &CallResponse{
		CallSid: "123",
		Status:  "sent",
	}, nil
}

type CallRequest struct {
	To      string
	Message string
}

type CallResponse struct {
	CallSid string
	Status  string
}
