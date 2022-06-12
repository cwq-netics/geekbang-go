package limiter

import (
	"sync"
	"time"
)

type slot struct {
	timestamp time.Time
	count     int
}

type SlidingWindow struct {
	mu              sync.Mutex // 互斥锁保护其他字段
	SlotDuration    time.Duration
	WindowsDuration time.Duration
	slotsCount      int
	slots           []*slot
	maxReqNum       int
}

//清理过期小窗口
func (l *SlidingWindow) cleanExpiredSlot(now time.Time) {
	offset := -1
	needClean := false
	for i, v := range l.slots {
		if v.timestamp.Add(l.WindowsDuration).After(now) {
			needClean = true
			break
		}
		offset = i
	}
	if needClean {
		l.slots = l.slots[offset+1:]
	}
}

func (l *SlidingWindow) countReq() (count int) {
	for _, v := range l.slots {
		count += v.count
	}
	return
}

func (l *SlidingWindow) AllowRequest() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	// 清理时间窗中过期slot
	l.cleanExpiredSlot(now)
	// 判断请求是否超限
	if l.countReq() >= l.maxReqNum {
		return false
	}
	var last *slot
	if len(l.slots) > 0 {
		last = l.slots[len(l.slots)-1]
		if last.timestamp.Add(l.SlotDuration).Before(now) {
			// 如果当前时间已经超过这个时间插槽的跨度，那么新建一个时间插槽
			last = &slot{timestamp: now, count: 1}
			l.slots = append(l.slots, last)
		} else {
			last.count++
		}
	} else {
		last = &slot{timestamp: now, count: 1}
		l.slots = append(l.slots, last)
	}
	return true
}

func NewSlidingWindow(slot time.Duration, windows time.Duration, maxReqNum int) *SlidingWindow {
	return &SlidingWindow{
		SlotDuration:    slot,
		WindowsDuration: windows,
		slotsCount:      int(windows / slot),
		maxReqNum:       maxReqNum,
	}
}
