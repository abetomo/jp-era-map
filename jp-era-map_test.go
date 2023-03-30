package jperamap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.org/x/text/width"
)

func TestMeiji(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 45; i++ {
			actual := JpEraMap[fmt.Sprintf("M%02d", i)]
			expected := 1868 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["M46"]
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestFuncMeiji(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 45; i++ {
			actual := JpEra(fmt.Sprintf("M%02d", i))
			expected := 1868 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 45; i++ {
			actual := JpEra(fmt.Sprintf("1%02d", i))
			expected := 1868 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): JP text", func(t *testing.T) {
		for i := 1; i <= 45; i++ {
			actual := JpEra(fmt.Sprintf("明治%02d年", i))
			expected := 1868 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 45; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("明治%02d年", i)))
			expected := 1868 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 45; i++ {
			actual := JpEra(fmt.Sprintf("明治%d年", i))
			expected := 1868 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 45; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("明治%d年", i)))
			expected := 1868 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEra("146")
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestTaisho(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 15; i++ {
			actual := JpEraMap[fmt.Sprintf("T%02d", i)]
			expected := 1912 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["T16"]
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestFuncTaisho(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 15; i++ {
			actual := JpEra(fmt.Sprintf("T%02d", i))
			expected := 1912 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 15; i++ {
			actual := JpEra(fmt.Sprintf("2%02d", i))
			expected := 1912 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): JP text", func(t *testing.T) {
		for i := 1; i <= 15; i++ {
			actual := JpEra(fmt.Sprintf("大正%02d年", i))
			expected := 1912 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 15; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("大正%02d年", i)))
			expected := 1912 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 15; i++ {
			actual := JpEra(fmt.Sprintf("大正%d年", i))
			expected := 1912 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 15; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("大正%d年", i)))
			expected := 1912 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEra("216")
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestShowa(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 64; i++ {
			actual := JpEraMap[fmt.Sprintf("S%02d", i)]
			expected := 1926 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["S65"]
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestFuncShowa(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 64; i++ {
			actual := JpEra(fmt.Sprintf("S%02d", i))
			expected := 1926 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 64; i++ {
			actual := JpEra(fmt.Sprintf("3%02d", i))
			expected := 1926 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): JP text", func(t *testing.T) {
		for i := 1; i <= 64; i++ {
			actual := JpEra(fmt.Sprintf("昭和%02d年", i))
			expected := 1926 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 64; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("昭和%02d年", i)))
			expected := 1926 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 64; i++ {
			actual := JpEra(fmt.Sprintf("昭和%d年", i))
			expected := 1926 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 64; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("昭和%d年", i)))
			expected := 1926 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEraMap["365"]
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestHeisei(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 31; i++ {
			actual := JpEraMap[fmt.Sprintf("H%02d", i)]
			expected := 1989 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["H32"]
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestFuncHeisei(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 31; i++ {
			actual := JpEra(fmt.Sprintf("H%02d", i))
			expected := 1989 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 31; i++ {
			actual := JpEra(fmt.Sprintf("4%02d", i))
			expected := 1989 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): JP text", func(t *testing.T) {
		for i := 1; i <= 31; i++ {
			actual := JpEra(fmt.Sprintf("平成%02d年", i))
			expected := 1989 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 31; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("平成%02d年", i)))
			expected := 1989 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 31; i++ {
			actual := JpEra(fmt.Sprintf("平成%d年", i))
			expected := 1989 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 31; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("平成%d年", i)))
			expected := 1989 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): no data exists", func(t *testing.T) {
		actual := JpEra("432")
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestReiwa(t *testing.T) {
	t.Run("mapping data", func(t *testing.T) {
		for i := 1; i <= 50; i++ {
			actual := JpEraMap[fmt.Sprintf("R%02d", i)]
			expected := 2019 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEraMap["R99"]
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func TestFuncReiwa(t *testing.T) {
	t.Run("JpEra(): prefix is a letter", func(t *testing.T) {
		for i := 1; i <= 50; i++ {
			actual := JpEra(fmt.Sprintf("R%02d", i))
			expected := 2019 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): prefix is a number", func(t *testing.T) {
		for i := 1; i <= 50; i++ {
			actual := JpEra(fmt.Sprintf("5%02d", i))
			expected := 2019 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("JpEra(): JP text", func(t *testing.T) {
		for i := 1; i <= 50; i++ {
			actual := JpEra(fmt.Sprintf("令和%02d年", i))
			expected := 2019 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 50; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("令和%02d年", i)))
			expected := 2019 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 50; i++ {
			actual := JpEra(fmt.Sprintf("令和%d年", i))
			expected := 2019 + (i - 1)
			assert.Equal(t, expected, actual)
		}

		for i := 1; i <= 50; i++ {
			actual := JpEra(width.Widen.String(fmt.Sprintf("令和%d年", i)))
			expected := 2019 + (i - 1)
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("mapping data: no data exists", func(t *testing.T) {
		actual := JpEra("599")
		expected := 0
		assert.Equal(t, expected, actual)
	})
}

func ExampleJpEraMap() {
	fmt.Printf("%d\n", JpEraMap["S24"])
	fmt.Printf("%d\n", JpEraMap["H24"])
	// Output:
	// 1949
	// 2012
}

func ExampleJpEra() {
	fmt.Println(JpEra("大正3年"))
	fmt.Println(JpEra("R01"))
	// Output:
	// 1914
	// 2019
}
