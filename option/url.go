package option

import (
	"net/url"
	"reflect"
)

func (opt *Option) Url() error {
	var (
		urlp *url.URL
		err  error
	)

	if !opt.Found {
		return nil
	}

	if urlp, err = url.Parse(opt.Environ); err != nil {
		return opt.ParseError(err)
	}

	if urlp.Scheme == "" || urlp.Host == "" {
		return opt.ParseError(nil)
	}

	opt.Value.Set(reflect.ValueOf(urlp))
	return nil
}
