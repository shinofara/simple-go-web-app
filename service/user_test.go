package service

import (
	"testing"
	"context"
	"reflect"
)

func TestNewUser(t *testing.T) {
	ctx := context.Background()
	user := NewUser(ctx)

	expected := &UserService{
		ctx: ctx,
	}

	if !reflect.DeepEqual(expected, user) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, user)
	}
}