package session

type SessionManager interface {
	Get(id string) (session Session, err error)
	CreateSession() (session Session, err error)
}
