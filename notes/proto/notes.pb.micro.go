// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: notes/proto/notes.proto

package notes

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

// Api Endpoints for Notes service

func NewNotesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Notes service

type NotesService interface {
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	UpdateStream(ctx context.Context, opts ...client.CallOption) (Notes_UpdateStreamService, error)
}

type notesService struct {
	c    client.Client
	name string
}

func NewNotesService(name string, c client.Client) NotesService {
	return &notesService{
		c:    c,
		name: name,
	}
}

func (c *notesService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Notes.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Notes.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Notes.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Notes.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesService) UpdateStream(ctx context.Context, opts ...client.CallOption) (Notes_UpdateStreamService, error) {
	req := c.c.NewRequest(c.name, "Notes.UpdateStream", &UpdateRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &notesServiceUpdateStream{stream}, nil
}

type Notes_UpdateStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseAndRecv() (*UpdateResponse, error)
	Send(*UpdateRequest) error
}

type notesServiceUpdateStream struct {
	stream client.Stream
}

func (x *notesServiceUpdateStream) CloseAndRecv() (*UpdateResponse, error) {
	if err := x.stream.Close(); err != nil {
		return nil, err
	}
	r := new(UpdateResponse)
	err := x.RecvMsg(r)
	return r, err
}

func (x *notesServiceUpdateStream) Context() context.Context {
	return x.stream.Context()
}

func (x *notesServiceUpdateStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *notesServiceUpdateStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *notesServiceUpdateStream) Send(m *UpdateRequest) error {
	return x.stream.Send(m)
}

// Server API for Notes service

type NotesHandler interface {
	List(context.Context, *ListRequest, *ListResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	UpdateStream(context.Context, Notes_UpdateStreamStream) error
}

func RegisterNotesHandler(s server.Server, hdlr NotesHandler, opts ...server.HandlerOption) error {
	type notes interface {
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		UpdateStream(ctx context.Context, stream server.Stream) error
	}
	type Notes struct {
		notes
	}
	h := &notesHandler{hdlr}
	return s.Handle(s.NewHandler(&Notes{h}, opts...))
}

type notesHandler struct {
	NotesHandler
}

func (h *notesHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.NotesHandler.List(ctx, in, out)
}

func (h *notesHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.NotesHandler.Create(ctx, in, out)
}

func (h *notesHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.NotesHandler.Delete(ctx, in, out)
}

func (h *notesHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.NotesHandler.Update(ctx, in, out)
}

func (h *notesHandler) UpdateStream(ctx context.Context, stream server.Stream) error {
	return h.NotesHandler.UpdateStream(ctx, &notesUpdateStreamStream{stream})
}

type Notes_UpdateStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	SendAndClose(*UpdateResponse) error
	Recv() (*UpdateRequest, error)
}

type notesUpdateStreamStream struct {
	stream server.Stream
}

func (x *notesUpdateStreamStream) SendAndClose(in *UpdateResponse) error {
	if err := x.SendMsg(in); err != nil {
		return err
	}
	return x.stream.Close()
}

func (x *notesUpdateStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *notesUpdateStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *notesUpdateStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *notesUpdateStreamStream) Recv() (*UpdateRequest, error) {
	m := new(UpdateRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
