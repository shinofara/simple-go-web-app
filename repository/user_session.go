package repository

import (
	"fmt"
	"time"
	"github.com/shinofara/simple-go-web-app/session"
	"github.com/shinofara/simple-go-web-app/entity"
)

func CreateLoginSession(s *session.Session) (*entity.LoginSession, error) {
	sess, err := s.Get("session-name")
	if err != nil {
		return nil, err
	}

	login := &entity.LoginSession{
		LastLoginDate: time.Now(),
	}
	
	// Set some session values.
	s.SetValue(sess, "login", 123)
	if err := s.Save(sess); err != nil {
		return nil, err
	}
	
	return login, nil
}

func GetLoginSession(s *session.Session) (*entity.LoginSession, error) {
	sess, err := s.Get("session-name")
	if err != nil {
		return nil, err
	}

	if login, ok := s.GetValue(sess, "login").(*entity.LoginSession); ok {
		return login, nil
	}

	return nil, fmt.Errorf("Not found login session")
}
