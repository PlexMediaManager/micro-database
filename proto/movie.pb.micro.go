// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: movie.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MovieService service

type MovieService interface {
	FindDownloaded(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error)
}

type movieService struct {
	c    client.Client
	name string
}

func NewMovieService(name string, c client.Client) MovieService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &movieService{
		c:    c,
		name: name,
	}
}

func (c *movieService) FindDownloaded(ctx context.Context, in *DatabaseEmpty, opts ...client.CallOption) (*DatabaseResponse, error) {
	req := c.c.NewRequest(c.name, "MovieService.FindDownloaded", in)
	out := new(DatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieService service

type MovieServiceHandler interface {
	FindDownloaded(context.Context, *DatabaseEmpty, *DatabaseResponse) error
}

func RegisterMovieServiceHandler(s server.Server, hdlr MovieServiceHandler, opts ...server.HandlerOption) error {
	type movieService interface {
		FindDownloaded(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error
	}
	type MovieService struct {
		movieService
	}
	h := &movieServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MovieService{h}, opts...))
}

type movieServiceHandler struct {
	MovieServiceHandler
}

func (h *movieServiceHandler) FindDownloaded(ctx context.Context, in *DatabaseEmpty, out *DatabaseResponse) error {
	return h.MovieServiceHandler.FindDownloaded(ctx, in, out)
}
