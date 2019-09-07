package session

import (
	"github.com/google/uuid"
	"sync"
)

type MemoryManager struct {
	data   map[string]Session //session_id => MemorySession
	rwlock sync.RWMutex
}

//获取session
func (m *MemoryManager) Get(id string) (session Session, err error) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	session, ok := m.data[id]
	if !ok {
		err = ErrSessionNotExist
		return
	}
	return
}

//创建session
func (m *MemoryManager) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	id := uuid.New().String()
	m.data[id] = NewMemorySession(id)
	return m.data[id], nil
}
