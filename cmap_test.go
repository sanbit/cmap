package cmap

import (
	"testing"
)

func Test_map_bucket_zero(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {

		}
	}()

	NewConcurrentMap(0)
	t.Error("bucket=0 not panic")
}

func Test_map_bucket_negative(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {

		}
	}()

	NewConcurrentMap(-8)
	t.Error("bucket<0 not panic")
}

func Test_get_exist(t *testing.T) {

	m := NewConcurrentMap(5)

	value := "aaa"
	m.Set("aaa", value)
	m.Set("aaa", value)

	o := m.Get("aaa")

	if o == nil {
		t.Fatal("already set but can not get")
	}

	v, ok := o.(string)

	if !ok {
		t.Fatal("get a wrong type")
	}

	if v != value {
		t.Fatal("get a wrong value !=", value)
	}
}

func Test_get_nil(t *testing.T) {

	m := NewConcurrentMap(5)

	value := "aaa"
	m.Set("aaa", value)
	m.Set("aaa", value)

	o := m.Get("bbb")

	if o != nil {
		t.Fatal("already set but can not get")
	}
}
