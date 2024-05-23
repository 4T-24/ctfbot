package app

import (
	"context"
	"time"

	"ctfbot/internal/values"

	"github.com/sirupsen/logrus"
)

func (a *App) Start() {
	go func() {
		err := a.client.OpenGateway(context.TODO())
		if err == nil {
			return
		}

		select {
		case a.errorChannel <- err:
		default:
		}
	}()

	select {
	case <-a.shutdown:
		if a.client != nil {
			a.client.Close(context.TODO())
		}
	case err := <-a.errorChannel:
		logrus.WithField("error", err).Error("An error stopped execution")
	}

	time.Sleep(2 * time.Second)
}

func (a *App) Shutdown() error {
	select {
	case a.shutdown <- struct{}{}:
		return nil
	default:
		return values.ErrAppAlreadyClosed
	}
}
