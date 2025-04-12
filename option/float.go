package option

import "strconv"

func (opt *Option) Float(n int) error {
	if !opt.Found {
		return nil
	}

	if val, err := strconv.ParseFloat(opt.Environ, n); err != nil {
		return opt.ParseError(err)
	} else {
		opt.Value.SetFloat(val)
	}

	return nil
}
