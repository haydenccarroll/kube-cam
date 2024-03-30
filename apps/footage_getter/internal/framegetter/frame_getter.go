package framegetter

import "context"

type FrameGetter interface {
	Get(ctx context.Context) ([]byte, error)
}

type FrameGetterFunc func(ctx context.Context) ([]byte, error)

func (f FrameGetterFunc) Get(ctx context.Context) ([]byte, error) {
	return f(ctx)
}
