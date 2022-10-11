package logx

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// WriteSyncer output interface
type WriteSyncer interface {
	io.Writer
	Sync() error
}

// NewRollingFile split writer
func NewRollingFile(dir, filename string, maxSize, maxAge, maxBackups int) WriteSyncer {
	s, err := os.Stat(dir)
	if err != nil || !s.IsDir() {
		os.RemoveAll(dir)
		if err = os.MkdirAll(dir, 0766); err != nil {
			panic(err)
		}
	}
	if maxSize < 1 {
		maxSize = 1024
	}
	if maxAge < 1 {
		maxAge = 1
	}
	return NewSyncer(&lumberjack.Logger{
		Filename:   filepath.Join(dir, filename),
		MaxSize:    maxSize, // megabytes, MB
		MaxAge:     maxAge,  // days
		MaxBackups: maxBackups,
		LocalTime:  true,
		Compress:   true,
	})
}

type Syncer struct {
	*lumberjack.Logger
	buf       *bytes.Buffer
	logChan   chan []byte
	closeChan chan interface{}
	maxSize   int
}

func NewSyncer(l *lumberjack.Logger) *Syncer {
	ws := &Syncer{
		Logger:    l,
		buf:       bytes.NewBuffer([]byte{}),
		logChan:   make(chan []byte, 5000),
		closeChan: make(chan interface{}),
		maxSize:   1024,
	}
	go ws.run()
	return ws
}

func (s *Syncer) run() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			if s.buf.Len() > 0 {
				s.Sync()
			}
		case bs := <-s.logChan:
			if _, err := s.buf.Write(bs); err != nil {
				continue
			}
			if s.buf.Len() > s.maxSize {
				s.Sync()
			}
		case <-s.closeChan:
			s.Sync()
			return
		}
	}
}

func (s *Syncer) Stop() {
	close(s.closeChan)
}

func (s *Syncer) Write(bs []byte) (int, error) {
	b := make([]byte, len(bs))
	for i, c := range bs {
		b[i] = c
	}
	s.logChan <- b
	return 0, nil
}

func (s *Syncer) Sync() error {
	_, err := s.Logger.Write(s.buf.Bytes())
	s.buf.Reset()
	return err
}
