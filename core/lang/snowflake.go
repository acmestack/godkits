/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lang

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch             = int64(1640966400000)                           // Set start time (timestamp / MS): 2022-01-01 00:00:00, valid for 69 years
	timestampBits     = uint(41)                                       // Time stamp occupied digits
	dataCenterIdBits  = uint(5)                                        // Bytes occupied by data center ID
	workerIdBits      = uint(5)                                        // Number of bytes occupied by machine ID
	sequenceBits      = uint(12)                                       // Number of bytes occupied by the sequence
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // Timestamp maximum
	dataCenterIdMax   = int64(-1 ^ (-1 << dataCenterIdBits))           // Maximum number of data center IDS supported
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))               // Maximum number of machine IDS supported
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // Maximum number of sequence ids supported
	workerIdShift     = sequenceBits                                   // machine id left shift number
	dataCenterIdShift = sequenceBits + workerIdBits                    // Data center id left shift number
	timestampShift    = sequenceBits + workerIdBits + dataCenterIdBits // Timestamp left shift
)

// Snowflake Snowflake
type Snowflake struct {
	sync.Mutex
	timestamp    int64
	workerId     int64
	dataCenterId int64
	sequence     int64
}

// NewSnowflake NewSnowflake
//  @param dataCenterId
//  @param workerId
//  @return *Snowflake
//  @return error
func NewSnowflake(dataCenterId, workerId int64) (*Snowflake, error) {
	if dataCenterId < 0 || dataCenterId > dataCenterIdMax {
		return nil, fmt.Errorf("dataCenterId must be between 0 and %d", dataCenterIdMax-1)
	}
	if workerId < 0 || workerId > workerIdMax {
		return nil, fmt.Errorf("workerId must be between 0 and %d", workerIdMax-1)
	}
	return &Snowflake{
		timestamp:    0,
		dataCenterId: dataCenterId,
		workerId:     workerId,
		sequence:     0,
	}, nil
}

// NextVal
//  @receiver s
//  @return int64
func (s *Snowflake) NextVal() int64 {
	s.Lock()
	now := time.Now().UnixNano() / 1000000 // 转毫秒
	if s.timestamp == now {
		// The same timestamp generates different data
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// Exceeded 12bit length, need to wait for the next millisecond
			// next millisecond will use sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// different timestamps are recounted using sequence:0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		fmt.Println(fmt.Sprintf("epoch must be between 0 and %d", timestampMax-1))
		return 0
	}
	s.timestamp = now
	r := int64((t)<<timestampShift | (s.dataCenterId << dataCenterIdShift) | (s.workerId << workerIdShift) | (s.sequence))
	s.Unlock()
	return r
}

// GetDeviceID
//  @param sid
//  @return dataCenterId
//  @return workerId
func GetDeviceID(sid int64) (dataCenterId, workerId int64) {
	dataCenterId = (sid >> dataCenterIdShift) & dataCenterIdMax
	workerId = (sid >> workerIdShift) & workerIdMax
	return
}

// GetTimestamp
//  @param sid
//  @return timestamp
func GetTimestamp(sid int64) (timestamp int64) {
	timestamp = (sid >> timestampShift) & timestampMax
	return
}

// GetGenTimestamp
//  @param sid
//  @return timestamp
func GetGenTimestamp(sid int64) (timestamp int64) {
	timestamp = GetTimestamp(sid) + epoch
	return
}

// GetGenTime
//  @param sid
//  @return t
func GetGenTime(sid int64) (t string) {
	// The timestamp/1000 obtained by GetGenTimestamp needs to be converted into seconds
	t = time.Unix(GetGenTimestamp(sid)/1000, 0).Format("2006-01-02 15:04:05")
	return
}

// GetTimestampStatus Get the percentage of timestamps used: range (0.0 - 1.0)
//  @return state
func GetTimestampStatus() (state float64) {
	state = float64(time.Now().UnixNano()/1000000-epoch) / float64(timestampMax)
	return
}
