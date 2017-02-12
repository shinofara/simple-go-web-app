package config

import (
	"testing"
	"reflect"
)

func TestLoadDBConfig(t *testing.T) {
	actual, _ := LoadDBConfig("testdata/test.yml")

	expected := &DBConfigs{
		&DBConfig{
			Name: "test",
			Host: "localhost",
			User: "test",
			Password: "test",
			Port: 3306,
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}
}