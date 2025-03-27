package matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructMatcher_Match(t *testing.T) {
	t.Parallel()

	type args struct {
		structName string
	}
	type testCase struct {
		name     string
		args     args
		init     func() StructMatcher
		wantBool assert.BoolAssertionFunc
	}

	tests := []testCase{
		{
			name: "empty builder call — always true",
			args: args{
				structName: "Some",
			},
			init:     func() StructMatcher { return NewBuilder().Build() },
			wantBool: assert.True,
		},
		{
			name: "only match",
			args: args{
				structName: "Some",
			},
			init:     func() StructMatcher { return NewBuilder().AddRegexpMatches("Some").Build() },
			wantBool: assert.True,
		},
		{
			name: "only match",
			args: args{
				structName: "Req",
			},
			init: func() StructMatcher {
				return NewBuilder().AddRegexpMatches("Some").Build()
			},
			wantBool: assert.False,
		},
		{
			name: "only exclude",
			args: args{
				structName: "Req",
			},
			init:     func() StructMatcher { return NewBuilder().AddRegexpExcludes("Some").Build() },
			wantBool: assert.True,
		},
		{
			name: "only exclude",
			args: args{
				structName: "Some",
			},
			init:     func() StructMatcher { return NewBuilder().AddRegexpExcludes("Some").Build() },
			wantBool: assert.False,
		},
		{
			name: "exclude and match",
			args: args{
				structName: "Some",
			},
			init: func() StructMatcher {
				return NewBuilder().AddRegexpMatches("Some").AddRegexpExcludes("Response").Build()
			},
			wantBool: assert.True,
		},
		{
			name: "exclude and match",
			args: args{
				structName: "Response",
			},
			init: func() StructMatcher {
				return NewBuilder().AddRegexpMatches("Some").AddRegexpExcludes("Response").Build()
			},
			wantBool: assert.False,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			matcher := tt.init()
			gotBool := matcher.Match(tt.args.structName)

			tt.wantBool(t, gotBool)
		})
	}
}
