package option

func (opt *Option) String() error {
	if !opt.Found {
		return nil
	}

	opt.Value.SetString(opt.Environ)
	return nil
}
