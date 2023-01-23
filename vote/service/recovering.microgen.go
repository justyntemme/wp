// Code generated by microgen 1.0.4. DO NOT EDIT.

package service

import (
	"context"
	"fmt"
	log "github.com/go-kit/kit/log"
	service "github.com/justyntemme/wp/vote"
)

// RecoveringMiddleware recovers panics from method calls, writes to provided logger and returns the error of panic as method error.
func RecoveringMiddleware(logger log.Logger) Middleware {
	return func(next service.VoteService) service.VoteService {
		return &recoveringMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type recoveringMiddleware struct {
	logger log.Logger
	next   service.VoteService
}

func (M recoveringMiddleware) GetVoteById(ctx context.Context, id string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetVoteById", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetVoteById(ctx, id)
}

func (M recoveringMiddleware) GetVotesByUserId(ctx context.Context, id string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetVotesByUserId", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetVotesByUserId(ctx, id)
}

func (M recoveringMiddleware) GetVotesByClubId(ctx context.Context, id string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetVotesByClubId", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetVotesByClubId(ctx, id)
}
