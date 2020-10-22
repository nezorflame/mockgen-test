package client

import "context"

type Client interface {
	Produce(ctx context.Context, objects []string) error
}
