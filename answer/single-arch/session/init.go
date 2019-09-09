package session

var SessionMg SessionManager

func Init() {
	//读取配置文件，初始化session
	SessionMg = NewMemoryManager()
}

func Get(id string) (session Session, err error) {
	return SessionMg.Get(id)
}
func CreateSession() (session Session, err error) {
	return SessionMg.CreateSession()
}
