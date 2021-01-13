// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/feeds.proto

package feeds

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Api Endpoints for Feeds service

func NewFeedsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Feeds service

type FeedsService interface {
	New(ctx context.Context, in *NewRequest, opts ...client.CallOption) (*NewResponse, error)
	Entries(ctx context.Context, in *EntriesRequest, opts ...client.CallOption) (*EntriesResponse, error)
}

type feedsService struct {
	c    client.Client
	name string
}

func NewFeedsService(name string, c client.Client) FeedsService {
	return &feedsService{
		c:    c,
		name: name,
	}
}

func (c *feedsService) New(ctx context.Context, in *NewRequest, opts ...client.CallOption) (*NewResponse, error) {
	req := c.c.NewRequest(c.name, "Feeds.New", in)
	out := new(NewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsService) Entries(ctx context.Context, in *EntriesRequest, opts ...client.CallOption) (*EntriesResponse, error) {
	req := c.c.NewRequest(c.name, "Feeds.Entries", in)
	out := new(EntriesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Feeds service

type FeedsHandler interface {
	New(context.Context, *NewRequest, *NewResponse) error
	Entries(context.Context, *EntriesRequest, *EntriesResponse) error
}

func RegisterFeedsHandler(s server.Server, hdlr FeedsHandler, opts ...server.HandlerOption) error {
	type feeds interface {
		New(ctx context.Context, in *NewRequest, out *NewResponse) error
		Entries(ctx context.Context, in *EntriesRequest, out *EntriesResponse) error
	}
	type Feeds struct {
		feeds
	}
	h := &feedsHandler{hdlr}
	return s.Handle(s.NewHandler(&Feeds{h}, opts...))
}

type feedsHandler struct {
	FeedsHandler
}

func (h *feedsHandler) New(ctx context.Context, in *NewRequest, out *NewResponse) error {
	return h.FeedsHandler.New(ctx, in, out)
}

func (h *feedsHandler) Entries(ctx context.Context, in *EntriesRequest, out *EntriesResponse) error {
	return h.FeedsHandler.Entries(ctx, in, out)
}
