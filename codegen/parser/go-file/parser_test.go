package parser

import (
	"github.com/0B1t322/zero-validation/codegen/matcher"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	t.Parallel()

	const path = "/Users/danademin/Projects/Golang/zero-validation/internal/_testdata"

	p := NewParser(
		WithStructMatcherBuilder(
			matcher.
				NewBuilder().
				AddFullMatches("ToDo"),
		),
	)
	err := p.ParsePackage(path)
	require.NoError(t, err)
}
