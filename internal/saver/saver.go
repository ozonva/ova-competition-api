package saver

import (
	"context"
	"errors"
	"fmt"
	"ozonva/ova-competition-api/internal/flusher"
	"ozonva/ova-competition-api/internal/models"
	"sync"
	"time"
)

type saver struct {
	m              sync.Mutex
	ticker         *time.Ticker
	interval       time.Duration
	buffer         []models.Competition
	flusher        flusher.Flusher
	capacity       uint
	tickerStopChan chan struct{}
	ctx            context.Context
}

// Saver в фоновом режиме периодически сохраняет соревнования в хранилище
type Saver interface {
	// Save асинхронно сохраняет соревнование
	Save(competition models.Competition) error
	// Close завершает работу Saver, сохраняя уже накопленные данные
	Close() error
}

// NewSaver возвращает Saver с поддержкой периодического сохранения
func NewSaver(
	capacity uint,
	interval time.Duration,
	flusher flusher.Flusher,
	ctx context.Context,
) Saver {
	s := &saver{
		capacity: capacity,
		interval: interval,
		buffer:   make([]models.Competition, 0, capacity),
		flusher:  flusher,
		ctx:      ctx,
	}
	s.init()
	return s
}

func (s *saver) init() {
	s.ticker = time.NewTicker(s.interval)
	s.tickerStopChan = make(chan struct{})

	var wg *sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case _, ok := <-s.tickerStopChan:
				if !ok {
					s.ticker.Stop()
					return
				}
			case <-s.ticker.C:
				err := s.flush()
				if err != nil {
					return
				}
				break
			}
		}
	}()

	wg.Wait()
}

func (s *saver) flush() error {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.buffer) > 0 {
		err := s.flusher.Flush(s.ctx, s.buffer)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to flush buffer: %v", err))
		}
		s.buffer = make([]models.Competition, 0, s.capacity)
	}

	return nil
}

func (s *saver) Save(competition models.Competition) error {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.buffer) >= int(s.capacity) {
		return errors.New(fmt.Sprintf("error saving competition %v", competition))
	}

	s.buffer = append(s.buffer, competition)
	return nil
}

func (s *saver) Close() error {
	err := s.flush()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to close saver: %v", err))
	}

	s.tickerStopChan <- struct{}{}
	close(s.tickerStopChan)
	return nil
}
