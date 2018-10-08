package signature

import (
	"testing"
)

func TestFromAndToString(t *testing.T) {
	sig := "FA 3A ?? C3 02"
	p, err := FromString(sig)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	toString := p.String()
	if toString != sig {
		t.Fatalf("pattern sig not matching [ %s != %s ]", toString, sig)
		return
	}
}

func assertMatch(t *testing.T, p *Pattern, index int, value byte) {
	match := p.MatchAt(index, value)
	if !match {
		t.Fatalf("should match for value [ %X ] at [ %d ] in signature [ %s ]", value, index, p)
	}
}

func assertNotMatch(t *testing.T, p *Pattern, index int, value byte) {
	match := p.MatchAt(index, value)
	if match {
		t.Fatalf("should not match for value [ %X ] at [ %d ] in signature [ %s ]", value, index, p)
	}
}

func TestMatchAt(t *testing.T) {
	sig := "FA 3A ?? C3 02"
	p, err := FromString(sig)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	assertMatch(t, p, 1, 58 /*3A*/)
	assertMatch(t, p, 2, 123 /*7B*/)
	assertNotMatch(t, p, 0, 58 /*3A*/)
	assertNotMatch(t, p, 23, 123 /*7B*/)
}
