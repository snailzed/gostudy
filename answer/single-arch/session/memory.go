package session

import "sync"

//内存session数据结构
type MemorySession struct {
	id     string                 //session id
	data   map[string]interface{} //session content
	rwlock sync.RWMutex
}

//创建一个内存session
func NewMemorySession(id string) Session {
	return &MemorySession{
		id:   id,
		data: make(map[string]interface{}, 1024),
	}
}

//设置session
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}

//获取session
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	if val, ok := m.data[key]; !ok {
		return nil, ErrKeyNotExist
	} else {
		return val, nil
	}
}

//删除
func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return nil
}

//save
func (m *MemorySession) Save(key string) (err error) {
	return
}
