package option

import (
	"reflect"
)

func (opt *Option) Struct() error {
	var (
		err error
		typ = opt.Value.Type()
	)

	for i := 0; i < opt.Value.NumField(); i++ {
		opt_new := New(opt.Name, opt.Value.Field(i), typ.Field(i))

		switch opt_new.Value.Kind() {
		case reflect.Bool:
			err = opt_new.Bool()

		case reflect.String:
			err = opt_new.String()

		case reflect.Uint, reflect.Uint8:
			err = opt_new.Uint(8)

		case reflect.Uint16:
			err = opt_new.Uint(16)

		case reflect.Uint32:
			err = opt_new.Uint(32)

		case reflect.Uint64:
			err = opt_new.Uint(64)

		case reflect.Int, reflect.Int8:
			err = opt_new.Int(8)

		case reflect.Int16:
			err = opt_new.Int(16)

		case reflect.Int32:
			err = opt_new.Int(32)

		case reflect.Int64:
			if opt_new.Field.Type.String() == "time.Duration" {
				err = opt_new.Duration()
			} else {
				err = opt_new.Int(64)
			}

		case reflect.Float32:
			err = opt_new.Float(32)

		case reflect.Float64:
			err = opt_new.Float(64)

		case reflect.Struct:
			err = opt_new.Struct()

		case reflect.Ptr:
			if opt_new.Field.Type.String() == "*url.URL" {
				err = opt_new.Url()
			}

		default:
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}
