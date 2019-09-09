package session

//sesion接口操作
type Session interface {
	Id() string
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save(key string) error
}
