package asm

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func Neg(x uint64) int64

func Foo() {
	log.Println("Test")
}

func Bar() {}

func TestASMNeg(t *testing.T) {
	x := Neg(42)
	if int64(-42) != Neg(42) {
		t.Fatalf("error [ %d != %d ]", x, -42)
	}
}

func TestASMNop(t *testing.T) {
	fooPtr := reflect.ValueOf(Foo).Pointer()
	log.Printf("foo: 0x%X", fooPtr)

	err := Nop(fooPtr)
	if err != nil {
		t.Fatal(err)
	}

	err = Return(fooPtr + uintptr(1))
	if err != nil {
		t.Fatal(err)
	}

	//time.Sleep(time.Minute * 5)
}

func TestASMCall(t *testing.T) {
	fooPtr := reflect.ValueOf(Foo).Pointer()
	log.Printf("foo: 0x%X", fooPtr)

	barPtr := reflect.ValueOf(Bar).Pointer()
	log.Printf("bar : 0x%X", barPtr)

	err := Jump(fooPtr, barPtr)
	if err != nil {
		t.Fatal(err)
	}

	err = Return(fooPtr + uintptr(9))
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Minute * 5)
}

func TestASMMovEAX(t *testing.T) {
	fooPtr := reflect.ValueOf(Foo).Pointer()
	log.Printf("foo: 0x%X", fooPtr)

	value := uint64(234)
	offset, err := MovEAX(fooPtr, value)
	if err != nil {
		t.Fatal(err)
	}

	err = Return(fooPtr + uintptr(offset))
	if err != nil {
		t.Fatal(err)
	}

	for {
		Foo()
		time.Sleep(time.Second * 1)
	}

}
