package log

import (
	"bytes"
	"errors"
	"io"
	"sync"
	"time"
)

const (
	FlushSize = 10240
	FlushTime = 5 * time.Millisecond
)

type AsyncBufferLogWriter struct {
	wait      sync.WaitGroup
	stopChan  chan bool
	logChan   chan []byte
	logBuffer bytes.Buffer
	FlushSize int64
	w         io.Writer
	block     bool
	once      sync.Once
}

type Config struct {
	// triggers refresh data size threshold
	FlushSize int64

	// The size of the async cache, which may block or return an error if exceeded
	BufferSize int

	// trigger refresh timeInterval
	FlushInterval time.Duration

	// if true, data overflow bufSize will make Write method blocking.
	// if false, data overflow will return error.
	Block bool
}

var defaultConfig = Config{
	FlushSize:     FlushSize,
	BufferSize:    360,
	FlushInterval: FlushTime,
	Block:         true,
}

// NewAsyncBufferWriter Write data with Buffer, this Writer and Closer is thread safety, but WriteCloser parameters not safety.
//  @param w Writer data
//  @param c Writer config
//  @return *AsyncBufferLogWriter
func NewAsyncBufferWriter(w io.Writer, c ...Config) *AsyncBufferLogWriter {
	conf := defaultConfig
	if len(c) > 0 {
		conf = c[0]
		if conf.FlushInterval == 0 {
			conf.FlushInterval = FlushTime
		}
		if conf.BufferSize == 0 {
			conf.BufferSize = BufferSize
		}
		if conf.FlushSize == 0 {
			conf.FlushSize = FlushSize
		}
	}

	l := AsyncBufferLogWriter{
		stopChan:  make(chan bool),
		logChan:   make(chan []byte, conf.BufferSize),
		FlushSize: conf.FlushSize,
		w:         w,
		block:     conf.Block,
	}
	l.wait.Add(1)
	l.logBuffer.Grow(conf.BufferSize * 10)

	go func() {
		defer l.wait.Done()
		defer func() {
			if w != nil {
				size := len(l.logChan)
				for i := 0; i < size; i++ {
					l.writeLog(<-l.logChan)
				}
				l.Flush()
			}
		}()
		ticker := time.NewTicker(conf.FlushInterval)
		defer ticker.Stop()
		for {
			select {
			case <-l.stopChan:
				return
			case <-ticker.C:
				l.Flush()
			default:
			}
			select {
			case <-l.stopChan:
				return
			case d, ok := <-l.logChan:
				if ok {
					l.writeLog(d)
				}
			}
		}
	}()
	return &l
}

// Flush data.
//  @receiver w AsyncBufferLogWriter pointer point address
//  @return error
func (w *AsyncBufferLogWriter) Flush() error {
	_, err := w.w.Write(w.logBuffer.Bytes())
	if err != nil {
		return err
	}
	w.logBuffer.Reset()
	return nil
}

// writeLog private method, write log.
//  @receiver w AsyncBufferLogWriter pointer point address
//  @param data log data
//  @return error
func (w *AsyncBufferLogWriter) writeLog(data []byte) error {
	w.logBuffer.Write(data)

	if int64(w.logBuffer.Len()) < w.FlushSize {
		return nil
	}

	return w.Flush()
}

// Close close writer buffer
//  @receiver w
//  @return error
func (w *AsyncBufferLogWriter) Close() error {
	w.once.Do(func() {
		close(w.stopChan)
		w.wait.Wait()
	})
	return nil
}

func (w *AsyncBufferLogWriter) Write(data []byte) (n int, err error) {
	if len(data) == 0 {
		return 0, nil
	}
	if w.block {
		select {
		case w.logChan <- data:
			return len(data), nil
		case <-w.stopChan:
			return 0, errors.New("writer is closed")
		}
	} else {
		select {
		case w.logChan <- data:
			return len(data), nil
		case <-w.stopChan:
			return 0, errors.New("writer is closed")
		default:
			return 0, errors.New("write log failed ")
		}
	}
}
