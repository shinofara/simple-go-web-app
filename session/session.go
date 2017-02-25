package session

import (
	"github.com/gorilla/sessions"
	"github.com/shinofara/simple-go-web-app/context"	
	"net/http"
	"github.com/uber-go/zap"		
)

type Session struct {
	store sessions.Store
	w http.ResponseWriter
	r *http.Request
	logger zap.Logger
}

func New(w http.ResponseWriter, r *http.Request, secret string) *Session {
	return &Session{
		store: sessions.NewCookieStore([]byte(secret)),
		r: r,
		w: w,
		logger: context.MustGetLogger(r.Context()),
	}
}

func (s *Session) SetLoginData() {
	session, err := s.store.Get(s.r, "session-name")
	if err != nil {
		s.logger.Info(err.Error())
	}
	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values["foo1"] = "bar"
	session.Values["foo2"] = "bar"
	session.Values["foo3"] = "bar"
	session.Values["foo4"] = "bar"	
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(s.r, s.w)
}
