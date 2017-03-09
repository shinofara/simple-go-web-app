package router

import (
	"testing"
)

func TestGenerateIndexKey(t *testing.T) {
	expected := "get/hoge/hoge"	
	actual := generateIndexKey("/hoge/hoge")

	if expected != actual {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}

	actual = generateIndexKey("/hoge/hoge/")
	if expected != actual {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}	
}
