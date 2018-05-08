package evaluator

import (
	"testing"
)

func Test(t *testing.T) {
	r := newReader("  A|Bac12&  !(1|A)")

	if r == nil {
		t.Errorf("newReader shouldn't be nil")
	}

	if r.len() != 18 {
		t.Errorf("Length should be 18, nothing is read currently")
	}

	if err := r.clear(' '); err != nil {
		t.Errorf("Clear shouldn't had to return nil")
	}

	if r.len() != 16 {
		t.Errorf("After clearing the whitespaces, len should be 16")
	}

	if ch, err := r.read(); err != nil {
		t.Errorf("Read shouldn't return error")
	} else if ch != 'A' {
		t.Errorf("The read character is not A")
	}

	if r.len() != 15 {
		t.Errorf("After reading the first rune, len should be 15")
	}

	if ch, err := r.seek(); err != nil {
		t.Errorf("Seek shouldn't return error")
	} else if ch != '|' {
		t.Errorf("Seek returned wrong character")
	}

	if r.len() != 15 {
		t.Errorf("Seek should only check the char, but not mark it as read")
	}

	if op, err := r.readOperator(); err != nil {
		t.Errorf("read operator should be successful")
	} else if !op.equal(or) {
		t.Errorf("wrong operator returned")
	}

	if p, err := r.readParameter(); err != nil {
		t.Errorf("read parameter should be successful")
	} else if p != "Bac12" {
		t.Errorf("wrong parameter returned")
	}

	if op, _ := r.readOperator(); !op.equal(and) {
		t.Errorf("wrong operator returned")
	}

	r.clear(' ')
	if op, _ := r.readOperator(); !op.equal(not) {
		t.Errorf("wrong operator returned")
	}

	if op, _ := r.readOperator(); !op.equal(leftBracket) {
		t.Errorf("wrong operator returned")
	}

	if _, err := r.readParameter(); err == nil {
		t.Errorf("error should be returned")
	}

	if _, err := r.readOperator(); err == nil {
		t.Errorf("error should be returned")
	}

	r.read()
	r.readOperator()
	r.readParameter()

	if op, _ := r.readOperator(); !op.equal(rightBracket) {
		t.Errorf("wrong operator returned")
	}

	if r.len() != 0 {
		t.Errorf("data has ended. 0 should be returned")
	}

	_, err0 := r.read()
	_, err1 := r.seek()
	_, err2 := r.readOperator()
	_, err3 := r.readParameter()

	if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("data has ended, error should be returned")
	}
}
