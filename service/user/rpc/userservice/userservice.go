// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./userservice_mock.go -package userservice -source $GOFILE

package userservice

import (
	"context"

	"hey-go-zero/service/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	IdsReq        = user.IdsReq
	UserListReply = user.UserListReply
	UserReq       = user.UserReq
	UserReply     = user.UserReply

	UserService interface {
		//  findone
		FindOne(ctx context.Context, in *UserReq) (*UserReply, error)
		//  findByIds
		FindByIds(ctx context.Context, in *IdsReq) (*UserListReply, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

//  findone
func (m *defaultUserService) FindOne(ctx context.Context, in *UserReq) (*UserReply, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.FindOne(ctx, in)
}

//  findByIds
func (m *defaultUserService) FindByIds(ctx context.Context, in *IdsReq) (*UserListReply, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.FindByIds(ctx, in)
}
