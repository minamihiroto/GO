package main

import(
	"testing"
)

func TestAppName(t *testing.T){
	expect := "zoo application"
	actual := AppName()

	if expect != actual{
		t.Errorf("%s、%sというずれがあるよ",expect,actual)
	}
}