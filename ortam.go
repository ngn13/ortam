package ortam

import (
	"reflect"

	"github.com/ngn13/ortam/option"
)

// provided argument is invalid
type ArgError struct{}

func (e *ArgError) Error() string {
	return "config argument is not a valid structure pointer"
}

func Load(config any, prefix ...string) error {
	var (
		opt option.Option
		val reflect.Value
	)

	if nil == config {
		return &ArgError{}
	}

	val = reflect.ValueOf(config)

	if val.Kind() != reflect.Ptr && val.Kind() != reflect.Interface {
		return &ArgError{}
	}

	if len(prefix) > 0 {
		opt = option.New(prefix[0], val.Elem())
	} else {
		opt = option.New("", val.Elem())
	}

	return opt.Struct()
}
