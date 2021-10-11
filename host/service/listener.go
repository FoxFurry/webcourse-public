package service

import (
	"context"
	"encoding/json"
	"github.com/foxfurry/university/webcourse/host/model"
	"log"
	"net/http"
	"sync"
	"time"
)

type IService interface {
	Start(ctx context.Context)
}

type listenerService struct {
	currentSessionMutex sync.Mutex
	currentSession      *model.Session
}

func NewListenerService() IService {
	service := &listenerService{
		currentSessionMutex: sync.Mutex{},
		currentSession:      nil,
	}
	return service
}

func (m *listenerService) Start(ctx context.Context) {
	log.Println("Starting listener service")
	go m.listen(ctx)
}

func (m *listenerService) listen(ctx context.Context) {
	checkListener := time.Tick(time.Second * 3)

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping service listening")
			return
		case <-checkListener:
			res, err := http.Get("http://client:8080/status")
			if err != nil {
				log.Printf("Error while getting listener status: %v", err)
				continue
			}

			if res.StatusCode == 200 {
				var sessionHolder model.Session

				if err = json.NewDecoder(res.Body).Decode(&sessionHolder); err != nil {
					log.Println("Could not decode")
					continue
				}

				log.Printf("Current user listens to %v at %d second", sessionHolder.CurrentSong.Title, sessionHolder.Position)
			} else if res.StatusCode == 404 {
				log.Println("Current user doesn't listen to anything")
			}
		}
	}
}
