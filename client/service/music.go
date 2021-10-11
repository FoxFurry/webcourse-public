package service

import (
	"context"
	"encoding/json"
	"github.com/foxfurry/university/webcourse/client/model"
	"log"
	"net/http"
	"sync"
	"time"
)

type IService interface {
	Session() *model.Session
}

type musicService struct {
	currentSessionMutex sync.Mutex
	currentSession      *model.Session
}

func NewMusicService(ctx context.Context) IService {
	service := &musicService{
		currentSessionMutex: sync.Mutex{},
		currentSession:      nil,
	}
	service.Start(ctx)
	return service
}

func (m *musicService) Session() *model.Session {
	m.currentSessionMutex.Lock()
	var tmp = m.currentSession
	m.currentSessionMutex.Unlock()
	return tmp
}

func (m *musicService) Start(ctx context.Context) {
	log.Println("Starting music service")
	go m.listen(ctx)
}

func (m *musicService) listen(ctx context.Context) {
	lookNext := time.Tick(time.Second * 10)
	songListen := time.Tick(time.Second * 1)

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping service listening")
			return
		case <-lookNext:
			nextSong := m.next()
			if nextSong != nil {
				log.Printf("Now listening: %s by %s", nextSong.Title, nextSong.Author)
				m.currentSessionMutex.Lock()
				m.currentSession = &model.Session{
					CurrentSong: *nextSong,
					Position:    0,
				}
				m.currentSessionMutex.Unlock()
			}
		case <-songListen:
			if m.Session() != nil {
				m.currentSessionMutex.Lock()
				m.currentSession.Position++
				m.currentSessionMutex.Unlock()
			}
		}
	}
}

func (m *musicService) next() *model.Song {
	res, err := http.Get("http://host:8081/next")
	if err != nil {
		log.Println("Could not get next song response")
		return nil
	}

	var nextSong model.Song
	if err = json.NewDecoder(res.Body).Decode(&nextSong); err != nil {
		log.Println("Could not decode next song body")
		return nil
	}
	return &nextSong
}
