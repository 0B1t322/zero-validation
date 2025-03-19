package errors_old

import (
	"sync"
)

type lazyErrorOperation interface {
	apply(errs Errors)
}

type errorsOperation struct {
	// one of

	set  *errorsKeyValue
	join Errors
}

func (op *errorsOperation) apply(errs Errors) {
	if set := op.set; set != nil {
		errs[set.key] = set.value
		return
	}

	if join := op.join; join != nil {
		errs.Join(join)
		return
	}
}

type errorsKeyValue struct {
	key   string
	value error
}

type LazyErrors struct {
	errors        Errors
	initOnce      sync.Once
	startLen      int
	ops           []*errorsOperation
	uniqueSetKeys map[string]struct{}
	maxJoinLen    int
}

func (l *LazyErrors) init() {
	l.initOnce.Do(
		func() {
			if len(l.ops) == 0 {
				return
			}

			startLen := max(l.maxJoinLen, len(l.uniqueSetKeys))
			l.errors = make(Errors, startLen)
			for _, op := range l.ops {
				op.apply(l.errors)
			}
		},
	)
}

func NewLazyErrors() ErrorsBuilder {
	return &LazyErrors{
		errors:        nil,
		initOnce:      sync.Once{},
		uniqueSetKeys: make(map[string]struct{}),
	}
}

func (l *LazyErrors) Get() Errors {
	l.init()

	return l.errors
}

func (l *LazyErrors) Set(key string, value error) {
	l.ops = append(l.ops, &errorsOperation{
		set: &errorsKeyValue{key, value},
	})
	l.uniqueSetKeys[key] = struct{}{}
}

func (l *LazyErrors) Join(with Errors) {
	l.ops = append(l.ops, &errorsOperation{
		join: with,
	})
	l.maxJoinLen = max(l.maxJoinLen, len(with))
}
