/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package projectmanager xxx
package projectmanager

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/bcsapi/bcsproject"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// Client xxx
type Client struct {
	ctx  context.Context
	conn *grpc.ClientConn
	bcsproject.BCSProjectClient
}

// Options for project manager client
type Options struct {
	Target   string
	Token    string
	Username string
}

var options Options

// SetProjectManagerClient set project manager client
func SetProjectManagerClient(o Options) {
	options = o
}

// NewProjectManagerClient create client for bcs-Cluster
func NewProjectManagerClient(ctx context.Context) (*Client, error) {
	header := map[string]string{
		"x-content-type":     "application/grpc+proto",
		"Content-Type":       "application/grpc",
		"X-Project-Username": options.Username, // bcs_project 部分接口要求有这个header
	}

	if len(options.Token) != 0 {
		header["Authorization"] = fmt.Sprintf("Bearer %s", options.Token)
	}
	md := metadata.New(header)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.Header(&md)))
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))) // nolint

	conn, err := grpc.NewClient(options.Target, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "create grpc client with '%s' failed", options.Target)
	}

	if conn == nil {
		return nil, fmt.Errorf("conn is nil")
	}

	return &Client{
		ctx:              metadata.NewOutgoingContext(ctx, md),
		conn:             conn,
		BCSProjectClient: bcsproject.NewBCSProjectClient(conn),
	}, nil
}

// Close close client
func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		blog.Errorf("grpc client close failed: %s", err)
	}
}
