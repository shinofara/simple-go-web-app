package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type SessionStore struct {
	store sessions.Store
	w http.ResponseWriter
	r *http.Request
}

func NewSessionStore(w http.ResponseWriter, r *http.Request, secret string) *SessionStore {
	return &SessionStore{
		store: sessions.NewCookieStore([]byte(secret)),
		r: r,
		w: w,
	}
}

func (s *SessionStore) NewSession(name string) *Session {
	session, err := s.store.Get(s.r, name)
	if err != nil {
		panic(err)
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: false,
		Secure: false,
	}
	
	return &Session{
		session: session,
		r: s.r,
		w: s.w,
	}
}

type Session struct {
	session *sessions.Session
	w http.ResponseWriter
	r *http.Request
}

func (s *Session) Save() error {
	return s.session.Save(s.r, s.w)
}

func (s *Session) Get(name string) interface{} {
	return s.session.Values[name]
}

func (s *Session) Set(name string, value interface{}) {
	s.session.Values[name] = value
}
