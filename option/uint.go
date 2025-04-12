package option

import "strconv"

func (opt *Option) Uint(n int) error {
	if !opt.Found {
		return nil
	}

	if val, err := strconv.ParseUint(opt.Environ, 10, n); err != nil {
		return opt.ParseError(err)
	} else {
		opt.Value.SetUint(val)
	}

	return nil
}
