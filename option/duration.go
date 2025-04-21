package option

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (opt *Option) Duration() error {
	if !opt.Found {
		return nil
	}

	var (
		dur uint64
		err error
	)

	dur_chr := opt.Environ[len(opt.Environ)-1]
	dur_num := strings.TrimSuffix(opt.Environ, string(dur_chr))

	if dur, err = strconv.ParseUint(dur_num, 10, 64); err != nil {
		return opt.ParseError(err)
	}

	switch dur_chr {
	case 's':
		opt.Value.Set(reflect.ValueOf(time.Duration(dur) * (time.Second)))

	case 'm':
		opt.Value.Set(reflect.ValueOf(time.Duration(dur) * (time.Minute)))

	case 'h':
		opt.Value.Set(reflect.ValueOf(time.Duration(dur) * (time.Hour)))

	default:
		return opt.ParseError(nil)
	}

	return nil
}
