package gorethink

import (
	"reflect"

	"github.com/the-control-group/gorethink-v1.15/encoding"
)

func init() {
	// Set encoding package
	encoding.IgnoreType(reflect.TypeOf(Term{}))
}
