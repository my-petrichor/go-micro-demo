// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: demo.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for AccountService service

func NewAccountServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AccountService service

type AccountService interface {
	AccountRegister(ctx context.Context, in *ReqAccountRegister, opts ...client.CallOption) (*ResAccountRegister, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) AccountRegister(ctx context.Context, in *ReqAccountRegister, opts ...client.CallOption) (*ResAccountRegister, error) {
	req := c.c.NewRequest(c.name, "AccountService.AccountRegister", in)
	out := new(ResAccountRegister)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceHandler interface {
	AccountRegister(context.Context, *ReqAccountRegister, *ResAccountRegister) error
}

func RegisterAccountServiceHandler(s server.Server, hdlr AccountServiceHandler, opts ...server.HandlerOption) error {
	type accountService interface {
		AccountRegister(ctx context.Context, in *ReqAccountRegister, out *ResAccountRegister) error
	}
	type AccountService struct {
		accountService
	}
	h := &accountServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AccountService{h}, opts...))
}

type accountServiceHandler struct {
	AccountServiceHandler
}

func (h *accountServiceHandler) AccountRegister(ctx context.Context, in *ReqAccountRegister, out *ResAccountRegister) error {
	return h.AccountServiceHandler.AccountRegister(ctx, in, out)
}
