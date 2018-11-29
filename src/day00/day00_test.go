package day00

import "testing"

func TestCase01(t *testing.T){
	expected := int64(4)
	actual := ProduceSum("1111")
	if actual != expected {
		t.Error("Test failed")
	}
}

func TestCase02(t *testing.T){
	expected := int64(9)
	actual := ProduceSum("91212129")
	if actual != expected {
		t.Error("Test failed")
	}
}