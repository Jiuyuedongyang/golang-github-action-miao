package main

import (
	"GitHub_cicd_test/pets"
	"testing"
)

func TestCat(t *testing.T)  {
	saying:=pets.Cat()
	if saying!="Miao~~~"{
		t.Errorf("cat say %s",saying)
	}

}