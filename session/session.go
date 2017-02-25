package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type Session struct {
	store sessions.Store
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request, secret string) *Session {
	return &Session{
		store: sessions.NewCookieStore([]byte(secret)),
		r: r,
		w: w,
	}
}

func (s *Session) Get(name string) (*sessions.Session, error) {
	return s.store.Get(s.r, name)
}

func (s *Session) Save(sess *sessions.Session) error {
	return sess.Save(s.r, s.w)	
}

func (s *Session) GetValue(sess *sessions.Session, name string) interface{} {
	return sess.Values[name]
}

func (s *Session) SetValue(sess *sessions.Session, name string, value interface{}) {
	sess.Values[name] = value
}
