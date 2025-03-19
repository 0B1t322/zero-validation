package syncx

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSlice_Elements_BlockAppend(t *testing.T) {
	t.Parallel()

	s := new(Slice[int])

	s.Append(1)
	s.Append(2)

	go func() {
		for idx, elem := range s.Elements() {
			time.Sleep(10 * time.Millisecond)
			switch idx {
			case 0:
				assert.Equal(t, 1, elem)
			case 1:
				assert.Equal(t, 2, elem)
			default:
				assert.Fail(t, "unexpected element")
			}
		}
	}()

	go func() {
		for idx, elem := range s.Elements() {
			time.Sleep(10 * time.Millisecond)
			switch idx {
			case 0:
				assert.Equal(t, 1, elem)
			case 1:
				assert.Equal(t, 2, elem)
			default:
				assert.Fail(t, "unexpected element")
			}
		}
	}()

	s.Append(3)
}
