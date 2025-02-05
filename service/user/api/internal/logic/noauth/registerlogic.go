//  Copyright [2020] [hey-go-zero]
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package logic

import (
	"context"

	"hey-go-zero/common/errorx"
	"hey-go-zero/common/regex"
	"hey-go-zero/service/user/api/internal/logic"
	"hey-go-zero/service/user/api/internal/svc"
	"hey-go-zero/service/user/api/internal/types"
	"hey-go-zero/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.UserRegisterReq) error {
	if !regex.Match(req.Username, regex.Username) {
		return logic.InvalidUsername
	}

	if !regex.Match(req.Passowrd, regex.Password) {
		return logic.InvalidPassword
	}

	_, err := l.svcCtx.UserModel.FindOneByUsername(req.Username)
	switch err {
	case nil:
		return errorx.NewDescriptionError("用户名已存在")
	case model.ErrNotFound:
		_, err = l.svcCtx.UserModel.Insert(model.User{
			Username: req.Username,
			Password: req.Passowrd,
			Role:     req.Role,
		})
		return err
	default:
		return err
	}
}
