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
