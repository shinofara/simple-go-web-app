// Package session セッション管理
package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

// SessionStore セッションストアとRequest/Responseを管理
type Store struct {
	store sessions.Store
	w http.ResponseWriter
	r *http.Request
}

// NewSessionStore creates a SessionStore
func NewSessionStore(w http.ResponseWriter, r *http.Request, secret string) *Store {
	return &Store{
		store: sessions.NewCookieStore([]byte(secret)),
		r: r,
		w: w,
	}
}

// NewSession creates a Session
func (s *Store) NewSession(name string) *Session {
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

// Session セッション操作に必要な物を管理
type Session struct {
	session *sessions.Session
	w http.ResponseWriter
	r *http.Request
}

// Save セッションを保存
func (s *Session) Save() error {
	return s.session.Save(s.r, s.w)
}

// Get セッションから取得
func (s *Session) Get(name string) interface{} {
	return s.session.Values[name]
}

// Set セッションに追加
func (s *Session) Set(name string, value interface{}) {
	s.session.Values[name] = value
}
