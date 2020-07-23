// https://github.com/golang/go/issues/17604
// 无锁队列的实现 https://coolshell.cn/articles/8239.html
// @source go-lib/goroutine/cas/atom-cas.go
package queue

import (
    "sync"
    "sync/atomic"
)

const sizeSection int = 10000

type Message struct {
    id int
}

type Queue struct {
    mu   *sync.Mutex
    db   []Message
    hasp int32
    tail int
    size int
}

func New() Queue {
    q := Queue{
        new(sync.Mutex), make([]Message, sizeSection, sizeSection), 0, 0, sizeSection,
    }
    return q
}

func (q Queue) PushTailCAS(n Message) {
    for {
        if atomic.CompareAndSwapInt32(&q.hasp, 0, 1) {
            break
        }
    }
    q.db[q.tail] = n
    q.tail++
    if q.tail > q.size {
        q.db = append(q.db, make([]Message, sizeSection, sizeSection)...)
        q.size += sizeSection
    }
    atomic.StoreInt32(&q.hasp, 0)
}
func (q Queue) PushTailMutex(n Message) {
    q.mu.Lock()
    q.db[q.tail] = n
    q.tail++
    if q.tail > q.size {
        q.db = append(q.db, make([]Message, sizeSection, sizeSection)...)
        q.size += sizeSection
    }
    q.mu.Unlock()
}
