package option

import "strconv"

func (opt *Option) Int(n int) error {
	if !opt.Found {
		return nil
	}

	if val, err := strconv.ParseInt(opt.Environ, 10, n); err != nil {
		return opt.ParseError(err)
	} else {
		opt.Value.SetInt(val)
	}

	return nil
}
