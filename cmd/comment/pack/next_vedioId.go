package pack

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch             = int64(1653363363137)
	timestampBits     = int64(41)
	datacenteridBits  = int64(3)
	workeridBits      = int64(7)
	sequenceBits      = int64(12)
	timestampMax      = int64(-1 ^ (-1 << timestampBits))
	datacenteridMax   = int64(-1 ^ (-1 << datacenteridBits))
	workeridMax       = int64(-1 ^ (-1 << workeridBits))
	sequenceMax       = int64(-1 ^ (-1 << sequenceBits))
	workeridShift     = sequenceBits
	datacenteridShift = sequenceBits + workeridBits
	timestampShift    = sequenceBits + workeridBits + datacenteridBits
)

type Snowflake struct {
	sync.Mutex
	Timestamp    int64
	Workerid     int64
	Datacenterid int64
	Sequence     int64
}

func (s *Snowflake) NextVal() (int64, error) {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixNano() / 1000000
	if s.Timestamp == now {
		s.Sequence = (s.Sequence + 1) & sequenceMax
		if s.Sequence == 0 {
			for now <= s.Timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.Sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		return -1, errors.New("epoch must be less timestampMax")
	}
	s.Timestamp = now
	r := int64((t)<<timestampShift | (s.Datacenterid << datacenteridShift) | (s.Workerid << workeridShift) | (s.Sequence))
	return r, nil
}
