package repository

import (
	"fmt"
	"time"
	"encoding/gob"
	"github.com/shinofara/simple-go-web-app/session"
	"github.com/shinofara/simple-go-web-app/entity"
)

func init() {
	//sessionで利用したいstructはここでせってい
	gob.Register(&entity.LoginSession{})
}

func CreateLoginSession(s *session.SessionStore) (*entity.LoginSession, error) {
	sess := s.NewSession("session-name")

	login := &entity.LoginSession{
		LastLoginDate: time.Now(),
	}
	
	// Set some session values.
	sess.Set("login", login)
	if err := sess.Save(); err != nil {
		return nil, err
	}
	
	return login, nil
}

func GetLoginSession(s *session.SessionStore) (*entity.LoginSession, error) {
	sess := s.NewSession("session-name")

	if login, ok := sess.Get("login").(*entity.LoginSession); ok {
		return login, nil
	}

	return nil, fmt.Errorf("Not found login session")
}
