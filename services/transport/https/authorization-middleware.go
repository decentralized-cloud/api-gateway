// Package https implements functions to expose api-gateway service endpoint using HTTPS/GraphQL protocol.
package https

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	gocorefasthttp "github.com/micro-business/go-core/jwt/fasthttp"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (service *transportService) createAuthMiddleware(endpointName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			_, err = gocorefasthttp.ParseAndVerifyToken(ctx, service.jwksURL, true)
			if err != nil {
				return nil, err
			}

			convertedCtx, ok := ctx.(*fasthttp.RequestCtx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "Failed to cast ctx to fasthttp.RequestCtx")
			}

			if bearerToken := string(convertedCtx.Request.Header.Peek(fasthttp.HeaderAuthorization)); len(bearerToken) != 0 {
				ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(fasthttp.HeaderAuthorization, bearerToken))
			}

			return next(ctx, request)
		}
	}
}
