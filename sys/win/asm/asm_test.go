package asm

import (
	"log"
	"time"
	"reflect"
	"testing"
)

func Neg(x uint64) int64

func Foo() {

}

func TestASMNeg(t *testing.T) {
	x := Neg(42)
	if int64(-42) != Neg(42) {
		t.Fatalf("error [ %d != %d ]", x, -42)
	}
}

func TestASMNop(t *testing.T) {
	ptr := reflect.ValueOf(Foo).Pointer()
	log.Printf("foo: 0x%X", ptr)

	
	time.Sleep(time.Minute*5)
}
