// Code generated by microgen 1.0.4. DO NOT EDIT.

package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	service "github.com/justyntemme/wp/user"
	"time"
)

// LoggingMiddleware writes params, results and working time of method call to provided logger after its execution.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.UserService) service.UserService {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   service.UserService
}

func (M loggingMiddleware) GetUserById(arg0 context.Context, arg1 string) (res0 string, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUserById",
			"message", "GetUserById called",
			"request", logGetUserByIdRequest{Id: arg1},
			"response", logGetUserByIdResponse{Result: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetUserById(arg0, arg1)
}

type (
	logGetUserByIdRequest struct {
		Id string
	}
	logGetUserByIdResponse struct {
		Result string
	}
)
