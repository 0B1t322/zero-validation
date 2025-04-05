package testdata

import (
	"github.com/0B1t322/zero-validation/internal/_testdata/subpkg1"
	"github.com/0B1t322/zero-validation/internal/optional"
	"time"
)

type Todo struct {
	ID        uint64
	Title     string
	CreatedAt time.Time
}

type EnumType int

const (
	EnumTypeFirst EnumType = iota
	EnumTypeSecond
)

type (
	ToDo struct {
		ID           uint64       `json:"id"`
		Name         string       `json:"name"`
		Number       *int64       `json:"number"`
		Sub          SubType      `json:"sub_type"`
		Enum         EnumType     `json:"enum_type"`
		Attr         Attr         `json:"attr"`
		AttrPtr      *Attr        `json:"attr_ptr"`
		Field        Interface    `json:"field"`
		UUID         subpkg1.UUID `json:"uuid"`
		CreateAt     *time.Time   `json:"create_at"`
		SliceOfUint  []uint       `json:"sliceOfUint"`
		SliceOfAttr  []Attr       `json:"sliceOfAttr"`
		Inner        Inner        `json:"inner"`
		SomeOptional optional.Optional[uint64]
	}

	Interface interface {
		Some()
	}

	SubType string

	Inner struct {
		ID uint64
	}
)
