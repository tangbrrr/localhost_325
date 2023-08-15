package mongodb

import (
	"context"

	merr "github.com/tangbo/twatt/mond/wind/err"
	"github.com/tangbo/twatt/mond/wind/utils/endpoint"
)

func ctxDead(endpoint endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		select {
		case <-ctx.Done():
			return nil, merr.MongoContextTimeoutErr
		default:
			return endpoint(ctx, req)
		}
	}
}
