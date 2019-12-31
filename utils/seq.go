package utils

import (
	"sync"
)

// Seq is used to generate new sequence numbers for commands
type Seq interface {
	// Next will generate and return a new sequence number.
	// The ch argument will receive data once Handle() is called with the corresponding seq number
	// It's important to call the release() func when the callback is no longer needed
	// in order to prevent sequence exhaustion. Preferably via defer
	Next(ch chan<- interface{}) (seq uint8, release func())

	// Handle sends data to the corresponding waiting channel
	// Returns false if no waiting channel is found
	Handle(seq uint8, data interface{}) (found bool)
}

type seqNo struct {
	lock    sync.RWMutex
	seq     uint8
	waits   map[uint8]chan<- interface{}
	queueCh chan struct{}
}

// NewSeq returns a new Seq
func NewSeq() Seq {
	return &seqNo{
		waits:   map[uint8]chan<- interface{}{},
		queueCh: make(chan struct{}, 1),
	}
}

func (s *seqNo) Next(ch chan<- interface{}) (uint8, func()) {
	for {
		seq, ok := s.find(ch)
		if ok {
			return seq, func() { s.release(seq) }
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
	for i := 0; i < 255; i++ {
		if seq == 255 {
			seq = 0
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
	delete(s.waits, seq)
	if s.queueCh != nil {
		close(s.queueCh)
		s.queueCh = nil
	}
}
