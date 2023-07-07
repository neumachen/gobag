package gobag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	cc := [][]string{
		[]string{"testCase", "test_case"},
		[]string{"TestCase", "test_case"},
		[]string{"Test Case", "test_case"},
		[]string{" Test Case", "test_case"},
		[]string{"Test Case ", "test_case"},
		[]string{" Test Case ", "test_case"},
		[]string{"test", "test"},
		[]string{"test_case", "test_case"},
		[]string{"Test", "test"},
		[]string{"", ""},
		[]string{"ManyManyWords", "many_many_words"},
		[]string{"manyManyWords", "many_many_words"},
		[]string{"AnyKind of_string", "any_kind_of_string"},
		[]string{"numbers2and55with000", "numbers_2_and_55_with_000"},
		[]string{"Foo/Boo", "foo_or_boo"},
		[]string{"Foo/Boo/Moo", "foo_or_boo_or_moo"},
	}
	for _, c := range cc {
		assert.Equal(t, c[1], SnakeCase(c[0]))
	}
}
