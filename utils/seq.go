package utils

import (
	"context"
	"sync"
)

// Seq is used to generate new sequence numbers for commands
type Seq interface {
	// Next will generate and return a new sequence number.
	// The ch argument will receive data once Handle() is called with the corresponding seq number
	// It's important to call the release() func when the callback is no longer needed
	// in order to prevent sequence exhaustion. Preferably via defer
	Next(ch chan<- interface{}) (seq uint8, release func(), err error)

	// Handle sends data to the corresponding waiting channel
	// Returns false if no waiting channel is found
	Handle(seq uint8, data interface{}) (found bool)
}

type Err string

func (e Err) Error() string { return string(e) }

const (
	ErrContextDone = Err("context is expired")
)

type seqNo struct {
	lock    sync.RWMutex
	ctx     context.Context
	seq     uint8
	min     uint8
	max     uint8
	waits   map[uint8]chan<- interface{}
	queueCh chan struct{}
}

// NewSeq returns a new Seq
func NewSeq(ctx context.Context) Seq {
	s := &seqNo{
		waits:   map[uint8]chan<- interface{}{},
		queueCh: make(chan struct{}, 1),
		min:     0,
		max:     255,
		ctx:     ctx,
	}
	if ctx != nil {
		go func() {
			<-ctx.Done()
			s.lock.Lock()
			defer s.lock.Unlock()
			for _, ch := range s.waits {
				close(ch)
			}
			s.waits = map[uint8]chan<- interface{}{}
		}()
	}
	return s
}

// NewSeq returns a new Seq
func NewSeqRange(ctx context.Context, min, max uint8) Seq {
	s := NewSeq(ctx).(*seqNo)
	s.min = min
	s.max = max
	s.seq = min
	return s
}

func (s *seqNo) isDone() bool {
	if s.ctx != nil {
		select {
		case <-s.ctx.Done():
			return true
		default:
			return false
		}
	}
	return false
}

func (s *seqNo) Next(ch chan<- interface{}) (uint8, func(), error) {
	for {
		if s.isDone() {
			return 0, nil, ErrContextDone
		}
		seq, ok := s.find(ch)
		if ok {
			return seq, func() { s.release(seq) }, nil
		}
		<-s.waitCh()
	}
}

func (s *seqNo) Handle(seq uint8, o interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if ch, ok := s.waits[seq]; ok {
		ch <- o
		return true
	}
	return false
}

func (s *seqNo) waitCh() <-chan struct{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.queueCh == nil {
		s.queueCh = make(chan struct{})
	}
	return s.queueCh
}

func (s *seqNo) find(ch chan<- interface{}) (uint8, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	seq := s.seq
	for i := s.min; i < s.max; i++ {
		if seq == s.max {
			seq = s.min
		} else {
			seq = seq + 1
		}

		if _, ok := s.waits[seq]; ok {
			continue
		}

		s.seq = seq
		s.waits[seq] = ch
		return seq, true
	}
	return 0, false
}

func (s *seqNo) release(seq uint8) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if ch, ok := s.waits[seq]; ok {
		close(ch)
		delete(s.waits, seq)
		if s.queueCh != nil {
			close(s.queueCh)
			s.queueCh = nil
		}
	}
}
