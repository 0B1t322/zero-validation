package parser

import (
	"github.com/0B1t322/zero-validaton/codegen/parser"
	"github.com/0B1t322/zero-validaton/internal/set"
	"maps"
)

type importSet struct {
	set.Set[parser.Import]

	importByAlias map[string]parser.Import
}

func newImportSet() *importSet {
	return &importSet{
		Set:           set.New[parser.Import](),
		importByAlias: make(map[string]parser.Import),
	}
}

func (s *importSet) Add(value parser.Import) {
	s.Set.Add(value)

	s.importByAlias[value.GetUsedPackageName()] = value
}

func (s *importSet) AddMany(values ...parser.Import) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *importSet) Clean() {
	s.Set.Clean()
	maps.DeleteFunc(s.importByAlias, func(_ string, _ parser.Import) bool {
		return true
	})
}
