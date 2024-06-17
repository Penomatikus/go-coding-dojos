package notification

import (
	"context"
	"io"
)

type Service interface {
	Send(ctx context.Context, p []byte, w io.Writer) error
	Read(ctx context.Context, r io.Reader) ([]byte, error)
}
