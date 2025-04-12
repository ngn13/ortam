package option

func (opt *Option) Bool() error {
	if !opt.Found {
		return nil
	}

	switch opt.Environ {
	case "1", "true":
		opt.Value.SetBool(true)

	case "0", "false":
		opt.Value.SetBool(false)

	default:
		return opt.TypeError()
	}

	return nil
}
