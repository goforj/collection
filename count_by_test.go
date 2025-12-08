package collection

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountByValue(t *testing.T) {
	nums := New([]int{1, 2, 2, 2, 3})

	counted := CountByValue(nums)

	expected := map[int]int{
		1: 1,
		2: 3,
		3: 1,
	}

	if !reflect.DeepEqual(counted, expected) {
		t.Fatalf("expected %v, got %v", expected, counted)
	}
}

func TestCountBy_Callback(t *testing.T) {
	emails := New([]string{
		"alice@gmail.com",
		"bob@yahoo.com",
		"carlos@gmail.com",
	})

	counted := CountBy(emails, func(email string) string {
		return email[strings.LastIndex(email, "@")+1:]
	})

	expected := map[string]int{
		"gmail.com": 2,
		"yahoo.com": 1,
	}

	if !reflect.DeepEqual(counted, expected) {
		t.Fatalf("expected %v, got %v", expected, counted)
	}
}
