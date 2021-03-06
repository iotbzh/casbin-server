// Copyright 2018 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"testing"

	pb "github.com/iotbzh/casbin-server/proto"
)

type testEngine struct {
	s   *Server
	ctx context.Context
	h   int32
}

func newTestEngine(t *testing.T, from, connectStr string, modelLoc string) *testEngine {
	s := NewServer(from, connectStr, false, modelLoc)
	ctx := context.Background()

	resp, err := s.NewEnforcer(ctx, &pb.NewEnforcerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	return &testEngine{s: s, ctx: ctx, h: resp.Handler}
}
