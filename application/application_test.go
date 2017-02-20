package application

import (
	"testing"
)

func TestGenerateIndexKey(t *testing.T) {
	expected := "get/hoge/hoge"	
	actual := GenerateIndexKey("/hoge/hoge")

	if expected != actual {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}

	actual = GenerateIndexKey("/hoge/hoge/")

	if expected != actual {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}	
}
