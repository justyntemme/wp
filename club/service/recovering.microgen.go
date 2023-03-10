// Code generated by microgen 1.0.4. DO NOT EDIT.

package service

import (
	"context"
	"fmt"
	log "github.com/go-kit/kit/log"
	service "github.com/justyntemme/wp/club"
)

// RecoveringMiddleware recovers panics from method calls, writes to provided logger and returns the error of panic as method error.
func RecoveringMiddleware(logger log.Logger) Middleware {
	return func(next service.ClubService) service.ClubService {
		return &recoveringMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type recoveringMiddleware struct {
	logger log.Logger
	next   service.ClubService
}

func (M recoveringMiddleware) GetClubById(ctx context.Context, id string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetClubById", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetClubById(ctx, id)
}

func (M recoveringMiddleware) GetAllClubsNearMe(ctx context.Context, limit int32) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetAllClubsNearMe", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetAllClubsNearMe(ctx, limit)
}
