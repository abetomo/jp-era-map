package jperamap

import (
	"fmt"
	"testing"
)

func TestMeiji(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 45; i++ {
			actual := JpEraMap[fmt.Sprintf("M%02d", i)]
			expected := 1868 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["M46"]
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestFuncMeiji(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 45; i++ {
			actual := JpEra(fmt.Sprintf("M%02d", i))
			expected := 1868 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 45; i++ {
			actual := JpEra(fmt.Sprintf("1%02d", i))
			expected := 1868 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEra("146")
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestTaisho(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 15; i++ {
			actual := JpEraMap[fmt.Sprintf("T%02d", i)]
			expected := 1912 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["T16"]
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestFuncTaisho(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 15; i++ {
			actual := JpEra(fmt.Sprintf("T%02d", i))
			expected := 1912 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 15; i++ {
			actual := JpEra(fmt.Sprintf("2%02d", i))
			expected := 1912 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEra("216")
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestShowa(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 64; i++ {
			actual := JpEraMap[fmt.Sprintf("S%02d", i)]
			expected := 1926 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["S65"]
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestFuncShowa(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 64; i++ {
			actual := JpEra(fmt.Sprintf("S%02d", i))
			expected := 1926 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 64; i++ {
			actual := JpEra(fmt.Sprintf("3%02d", i))
			expected := 1926 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEraMap["365"]
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestHeisei(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 31; i++ {
			actual := JpEraMap[fmt.Sprintf("H%02d", i)]
			expected := 1989 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["H32"]
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestFuncHeisei(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 31; i++ {
			actual := JpEra(fmt.Sprintf("H%02d", i))
			expected := 1989 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 31; i++ {
			actual := JpEra(fmt.Sprintf("4%02d", i))
			expected := 1989 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEra("432")
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestReiwa(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 50; i++ {
			actual := JpEraMap[fmt.Sprintf("R%02d", i)]
			expected := 2019 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["R99"]
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func TestFuncReiwa(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 50; i++ {
			actual := JpEra(fmt.Sprintf("R%02d", i))
			expected := 2019 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 20; i++ {
			actual := JpEra(fmt.Sprintf("5%02d", i))
			expected := 2019 + (i - 1)
			if actual != expected {
				t.Fatalf("%v not match %v", actual, expected)
			}
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEra("599")
		expected := 0
		if actual != expected {
			t.Fatalf("%v not match %v", actual, expected)
		}
	})
}

func ExampleJpEra() {
	year := JpEra("R01")
	fmt.Println(year)
	// Output: 2019
}
