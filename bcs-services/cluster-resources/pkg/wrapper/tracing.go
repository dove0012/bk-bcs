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
 *
 */

package wrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	grpc_codes "google.golang.org/grpc/codes"

	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/common/ctxkey"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/common/errcode"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/tracing"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/util/errorx"
)

// NewTracingWrapper new tracing wrapper
func NewTracingWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
			// 开始时间
			startTime := time.Now()
			md, ok := metadata.FromContext(ctx)
			if !ok {
				return errorx.New(errcode.General, "failed to get micro's metadata")
			}

			// 判断Header 是否有放置Transparent
			if value, ok := md.Get("traceparent"); ok {
				md["traceparent"] = value
				// 有则从上游解析Transparent
				ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.MapCarrier(md))
			} else {
				// 获取或生成 request id 注入到 context
				requestID := getOrCreateReqID(md)
				ctx = context.WithValue(ctx, ctxkey.RequestIDKey, requestID)
				ctx = tracing.ContextWithRequestID(ctx, requestID)
			}

			name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())

			tracer := otel.Tracer(req.Service())
			commonAttrs := []attribute.KeyValue{
				attribute.String("component", "gRPC"),
				attribute.String("method", req.Method()),
				attribute.String("url", req.Endpoint()),
			}
			ctx, span := tracer.Start(ctx, name, trace.WithSpanKind(trace.SpanKindServer),
				trace.WithAttributes(commonAttrs...))
			defer span.End()

			reqData, _ := json.Marshal(req.Body())

			// 返回Header添加Traceparent
			otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(md))

			err = fn(ctx, req, rsp)

			rspData, _ := json.Marshal(rsp)
			elapsedTime := time.Now().Sub(startTime)

			reqBody := string(reqData)
			if len(reqBody) > 1024 {
				reqBody = fmt.Sprintf("%s...(Total %s)", reqBody[:1024], humanize.Bytes(uint64(len(reqBody))))
			}

			respBody := string(rspData)
			if len(respBody) > 1024 {
				respBody = fmt.Sprintf("%s...(Total %s)", respBody[:1024], humanize.Bytes(uint64(len(respBody))))
			}

			// 设置额外标签
			span.SetAttributes(attribute.Key("req").String(reqBody))
			span.SetAttributes(attribute.Key("elapsed_ime").String(elapsedTime.String()))
			span.SetAttributes(attribute.Key("rsp").String(respBody))
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
				span.SetAttributes(tracing.GRPCStatusCodeKey.Int(int(codes.Error)))
			} else {
				span.SetAttributes(tracing.GRPCStatusCodeKey.Int(int(grpc_codes.OK)))
			}

			return err
		}
	}
}
