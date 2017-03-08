package config

import (
	"testing"
	"reflect"
)

func TestLoad(t *testing.T) {
	cfg, _ := Load("./testdata/all.yml")

	expected := &Config{
		CertFilePath: "cert.pem",
		KeyFilePath: "key.pem",
		HTTPPort: "80",
		DatabaseYmlPath: "database.yml",
		SMTP: &SMTP{
			Host: "localhost",
			Port: 1025,
		},
		Session: &Session{
			Salt: "xxxxx",
		},
	}

	if !reflect.DeepEqual(expected, cfg) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, cfg)
	}	
}

func TestLoadGlobalVar(t *testing.T) {
	_, _ = Load("./testdata/all.yml")

	expected := &Config{
		CertFilePath: "cert.pem",
		KeyFilePath: "key.pem",
		HTTPPort: "80",
		DatabaseYmlPath: "database.yml",
		SMTP: &SMTP{
			Host: "localhost",
			Port: 1025,
		},
		Session: &Session{
			Salt: "xxxxx",
		},
	}

	if !reflect.DeepEqual(expected, &cfg) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, cfg)
	}	
}

func TestGetSMTP(t *testing.T) {
	_, _ = Load("./testdata/all.yml")	
	smtp := GetSMTP()

	expected := &SMTP{
		Host: "localhost",
		Port: 1025,
	}

	if !reflect.DeepEqual(expected, smtp) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, smtp)
	}		
}
