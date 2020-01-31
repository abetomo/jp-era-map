package jperamap

import (
	"fmt"
	"testing"
)

func TestTaisho(t *testing.T) {
	for i := 1; i <= 15; i++ {
		actual := JpEraMap[fmt.Sprintf("T%02d", i)]
		expected := 1912 + (i - 1)
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	}
}

func TestShowa(t *testing.T) {
	for i := 1; i <= 64; i++ {
		actual := JpEraMap[fmt.Sprintf("S%02d", i)]
		expected := 1926 + (i - 1)
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	}
}

func TestHeisei(t *testing.T) {
	for i := 1; i <= 31; i++ {
		actual := JpEraMap[fmt.Sprintf("H%02d", i)]
		expected := 1989 + (i - 1)
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	}
}

func TestReiwa(t *testing.T) {
	for i := 1; i <= 20; i++ {
		actual := JpEraMap[fmt.Sprintf("R%02d", i)]
		expected := 2019 + (i - 1)
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	}
}
