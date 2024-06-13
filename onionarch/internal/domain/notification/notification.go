package notification

import (
	"context"
	"io"
)

type Service interface {
	Send(ctx context.Context, p []byte, w io.Writer) error
	Read(ctx context.Context, p []byte, r io.Reader) error
}
