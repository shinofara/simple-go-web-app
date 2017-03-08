package session

import (
	"time"
	"fmt"
	"encoding/gob"
	"github.com/shinofara/simple-go-web-app/http/session/core"
)

const (
	NameLoginSession = "login"

	KeyLoginSession = "login"
)

func init() {
	gob.Register(&LoginSession{})
}

// LoginSession ログインに必要なセッション情報を保持
type LoginSession struct {
	LastLoginDate time.Time
}

// CreateLoginSession ログインセッションを作成
func CreateLoginSession(s *core.Store) (*LoginSession, error) {
	sess := s.NewSession(NameLoginSession)

	login := &LoginSession{
		LastLoginDate: time.Now(),
	}

	// Set some session values.
	sess.Set(KeyLoginSession, login)
	if err := sess.Save(); err != nil {
		return nil, err
	}

	return login, nil
}


// GetLoginSession sessionからユーザのログイン情報を取得
func GetLoginSession(s *core.Store) (*LoginSession, error) {
	sess := s.NewSession(NameLoginSession)

	if login, ok := sess.Get(KeyLoginSession).(*LoginSession); ok {
		return login, nil
	}

	return nil, fmt.Errorf("Not found login session")
}
