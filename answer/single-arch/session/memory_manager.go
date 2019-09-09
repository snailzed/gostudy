package session

import (
	"github.com/google/uuid"
	"sync"
)

//保存所有用户的session
type MemoryManager struct {
	data   map[string]Session //session_id => MemorySession
	rwlock sync.RWMutex
}

//返回内存session管理实例
func NewMemoryManager() SessionManager {
	return &MemoryManager{
		data: make(map[string]Session, 1024),
	}
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
