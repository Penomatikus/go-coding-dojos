package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
)

type JSONSinkWriter struct {
	Data []byte
	Sink *map[model.SessionID][]model.Notification
}

func (w JSONSinkWriter) Write(p []byte) (n int, err error) {
	var note model.Notification
	if err = json.Unmarshal(p, &note); err != nil {
		return
	}
	note.CreatedAt = time.Now()

	sessionNotifications := (*w.Sink)[note.SessionId]
	sessionNotifications = append(sessionNotifications, note)
	(*w.Sink)[note.SessionId] = sessionNotifications

	return len(p), err
}

type JSONSinkReader struct {
	Notifications         []model.Notification
	buffer                *bytes.Buffer
	bytesRead, bytesTotal int
}

func (r *JSONSinkReader) Read(p []byte) (n int, err error) {
	if r.buffer == nil {
		data, err := json.Marshal(r.Notifications)
		if err != nil {
			return 0, err
		}
		r.buffer = bytes.NewBuffer(data)
		r.bytesTotal = len(data)
	}

	n, err = r.buffer.Read(p)
	if err != nil {
		return
	}

	r.bytesRead += n
	if r.bytesRead > r.bytesTotal {
		err = io.EOF
	}
	return n, err
}

type service struct{}

func PrivideService() notification.Service {
	return &service{}
}

func (s *service) Send(ctx context.Context, p []byte, w io.Writer) error {
	n, err := w.Write(p)
	if err != nil || n == 0 {
		return fmt.Errorf("%s: %d bytes written", err, n)
	}
	return nil
}

func (s *service) Read(ctx context.Context, r io.Reader) (out []byte, err error) {
	buf := make([]byte, 8)
	for {
		_, err = r.Read(buf)
		switch err {
		case io.EOF:
			return out, nil
		case nil:
			out = append(out, buf...)
		default:
			return nil, err
		}
	}
}
