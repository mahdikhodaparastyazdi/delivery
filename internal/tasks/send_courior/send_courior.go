package tasks

import (
	"context"
	"delivery/internal/constants"
	"delivery/internal/dto"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "delivery/pkg/logger"

	"github.com/hibiken/asynq"
)

func NewTask(msg dto.SendCourior, maxRetry int, timeout time.Duration) (*asynq.Task, error) {
	p, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(constants.SEND_COURIOR, p, asynq.MaxRetry(maxRetry), asynq.Timeout(timeout)), nil
}

type Queue struct {
	client   *asynq.Client
	maxRetry int
	timeout  time.Duration
}

func NewQueue3PL(c *asynq.Client, maxRetry int, timout time.Duration) *Queue {
	return &Queue{
		client:   c,
		maxRetry: maxRetry,
		timeout:  timout,
	}
}

func (q *Queue) Enqueue(msg dto.SendCourior) error {
	t, err := NewTask(msg, q.maxRetry, q.timeout)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	_, err = q.client.Enqueue(t, asynq.Queue(constants.SEND_COURIOR))
	if err != nil {
		return fmt.Errorf("asynq enqueue failed: %w", err)
	}
	return nil
}

type Consumer interface {
	Consume(ctx context.Context, msg dto.SendCourior, retry, maxRetry int) error
}

type Worker struct {
	server   *asynq.Server
	consumer Consumer
	logger   log.Logger
}

func NewWorker(s *asynq.Server, c Consumer, logger log.Logger) *Worker {
	return &Worker{
		server:   s,
		consumer: c,
		logger:   logger,
	}
}

func (w *Worker) StartWorker(name string) error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(name, w.HandleTask)

	return w.server.Run(mux)
}

func (w *Worker) HandleTask(ctx context.Context, t *asynq.Task) error {
	retry, _ := asynq.GetRetryCount(ctx)
	maxRetry, _ := asynq.GetMaxRetry(ctx)

	var p dto.SendCourior
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshalling json: %w", err)
	}
	if err := w.consumer.Consume(ctx, p, retry, maxRetry); err != nil {
		if errors.Is(err, constants.ErrExpiryReached) ||
			errors.Is(err, constants.ErrWrongStatus) ||
			errors.Is(err, constants.ErrBackOffRetry) {
			return nil
		}
		w.logger.Error("failed send courior", log.J{
			"error":   err.Error(),
			"payload": p,
		})
		return fmt.Errorf("courior task failed: %w", err)
	}

	return nil
}
