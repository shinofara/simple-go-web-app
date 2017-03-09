package session

import (
	"time"
	"fmt"
	"encoding/gob"
	"net/http"
	"github.com/shinofara/simple-go-web-app/http/context"
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
func CreateLoginSession(r *http.Request) (*LoginSession, error) {
	sessionStore := context.MustGetSessionStore(r.Context())
	s := sessionStore.NewSession(NameLoginSession)

	login := &LoginSession{
		LastLoginDate: time.Now(),
	}

	// Set some session values.
	s.Set(KeyLoginSession, login)
	if err := s.Save(); err != nil {
		return nil, err
	}

	return login, nil
}


// GetLoginSession sessionからユーザのログイン情報を取得
func GetLoginSession(r *http.Request) (*LoginSession, error) {
	sessionStore := context.MustGetSessionStore(r.Context())
	s := sessionStore.NewSession(NameLoginSession)

	if login, ok := s.Get(KeyLoginSession).(*LoginSession); ok {
		return login, nil
	}

	return nil, fmt.Errorf("Not found login session")
}
