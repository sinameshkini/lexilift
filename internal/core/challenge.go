package core

import (
	"context"
	"lexilift/pkg/endpoint"
)

func (c *Core) Challenge(ctx context.Context, req *endpoint.ChallengeRequest) (resp *endpoint.ChallengeResponse, err error) {
	eval, err := c.llm.EvaluateWriting(ctx, req.Word, req.Input)
	if err != nil {
		return nil, err
	}

	resp = &endpoint.ChallengeResponse{
		Result: eval,
	}

	return
}
