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

	"hey-go-zero/service/course/api/internal/svc"
	"hey-go-zero/service/course/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCourseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCourseListLogic {
	return GetCourseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseListLogic) GetCourseList(req types.CourseListReq) (*types.CourseListReply, error) {
	// todo: add your logic here and delete this line

	return &types.CourseListReply{}, nil
}
