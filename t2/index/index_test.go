package index

import (
	"reflect"
	"testing"
)

var idx Index

func init() {
	// a: hello hello cat
	// b: mother wash cat
	// c: mother wash window

	idx = Index{
		"hello": occurrences{
			"a": 2,
		},
		"cat": occurrences{
			"a": 1,
			"b": 1,
		},
		"mother": occurrences{
			"b": 1,
			"c": 1,
		},
		"wash": occurrences{
			"b": 1,
			"c": 1,
		},
		"window": occurrences{
			"c": 1,
		},
	}
}

func TestSearch1(t *testing.T) {
	q := "hello cat"
	expected := []Result{
		Result{
			File:  "a",
			Count: 3,
		},
	}

	actual, err := idx.Search(q)
	if err != nil {
		t.Errorf("%q leads to an error %q", q, err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v is not equal to %#v", expected, actual)
	}
}

func TestSearch2(t *testing.T) {
	q := "mother washs"
	expected := []Result{
		Result{
			File:  "b",
			Count: 2,
		},
		Result{
			File:  "c",
			Count: 2,
		},
	}

	actual, err := idx.Search(q)
	if err != nil {
		t.Errorf("%q leads to an error %q", q, err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v is not equal to %#v", expected, actual)
	}
}

func TestSearch3(t *testing.T) {
	q := "mother wash house"
	expected := []Result{}

	actual, err := idx.Search(q)
	if err != nil {
		t.Errorf("%q leads to an error %q", q, err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v is not equal to %#v", expected, actual)
	}
}

func TestAdd(t *testing.T) {
	a := "hello, hello cat"
	b := "mother wash cats"
	c := "mother washed window"

	actual := New()
	err := actual.Add(a, "a")
	if err != nil {
		t.Errorf("adding %q leads to an error %q", a, err)
	}
	err = actual.Add(b, "b")
	if err != nil {
		t.Errorf("adding %q leads to an error %q", b, err)
	}
	err = actual.Add(c, "c")
	if err != nil {
		t.Errorf("adding %q leads to an error %q", c, err)
	}

	if !reflect.DeepEqual(idx, actual) {
		t.Errorf("expected %#v is not equal to %#v", idx, actual)
	}
}

func TestMerge(t *testing.T) {
	a := "hello, hello cat"
	b := "mother wash cats"
	c := "mother washed window"

	actual := New()
	err := actual.Add(a, "a")
	if err != nil {
		t.Errorf("adding %q leads to an error %q", a, err)
	}

	append1 := New()
	err = append1.Add(b, "b")
	if err != nil {
		t.Errorf("adding %q leads to an error %q", b, err)
	}
	append2 := New()
	err = append2.Add(c, "c")
	if err != nil {
		t.Errorf("adding %q leads to an error %q", c, err)
	}

	actual.Merge(append1)
	actual.Merge(append2)

	if !reflect.DeepEqual(idx, actual) {
		t.Errorf("expected %#v is not equal to %#v", idx, actual)
	}
}
