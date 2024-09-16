package couriorproviders

import (
	"bytes"
	"context"
)

type CouriorSender interface {
	SendCourior(ctx context.Context, path string, buf *bytes.Buffer) error
}
