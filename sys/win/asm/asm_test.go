package asm

import (
	"log"
	"reflect"
	"testing"
	"time"
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
	Foo()
	ptr := reflect.ValueOf(Foo).Pointer()
	log.Printf("foo: 0x%X", ptr)

	err := Nop(ptr)
	if err != nil {
		t.Fatal(err)
	}

	err = Return(ptr + uintptr(1))
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Minute * 5)
}
