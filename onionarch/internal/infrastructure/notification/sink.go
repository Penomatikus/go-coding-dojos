package notification

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
)

type JSONSinkWriter struct {
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
	Offset    int
	SessionID model.SessionID
	Sink      *map[model.SessionID][]model.Notification
}

func (r JSONSinkReader) Read(p []byte) (n int, err error) {
	sessionNotifications := (*r.Sink)[r.SessionID][r.Offset:]
	b, err := json.Marshal(sessionNotifications)
	if err != nil {
		return 0, err
	}
	// dirty
	p = p[:cap(b)]
	n = copy(p, b)
	return n, nil
}

type service struct{}

func PrivideService() notification.Service {
	return &service{}
}

func (s *service) Send(ctx context.Context, p []byte, w io.Writer) error {
	n, err := w.Write(p)
	return handleIOErr(n, err, "no bytes written")
}

func (s *service) Read(ctx context.Context, p []byte, r io.Reader) error {
	n, err := r.Read(p)
	return handleIOErr(n, err, "no bytes read")
}

func handleIOErr(n int, err error, msg string) error {
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New(msg)
	}
	return nil
}
