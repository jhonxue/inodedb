// Copyright 2023 The CubeFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package inoder

import (
	"context"
	"github.com/cubefs/inodedb/proto"
)

type Server struct{}

func (s *Server) CreateShard(ctx *context.Context, req *proto.CreateShardRequest) (*proto.CreateShardResponse, error) {
}

func (s *Server) DeleteShard(ctx *context.Context, req *proto.DeleteShardRequest) (*proto.DeleteShardResponse, error) {
}

func (s *Server) UpsertInode(ctx *context.Context, req *proto.UpsertRequest) (*proto.UpsertResponse, error) {
}

func (s *Server) DeleteInode(ctx *context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
}
