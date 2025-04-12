package option

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"
)

type Option struct {
	Name    string
	Environ string
	Found   bool

	Field reflect.StructField
	Value reflect.Value
}

func toEnvName(name string) string {
	name_rune := []rune(name)
	env_name := ""

	for indx, char := range name {
		if indx > 0 && unicode.IsUpper(char) && unicode.IsLower(name_rune[indx-1]) {
			env_name += fmt.Sprintf("_%c", char)
		} else {
			env_name += string(char)
		}
	}

	return strings.ToUpper(env_name)
}

func New(prefix string, value reflect.Value, field ...reflect.StructField) Option {
	var opt Option

	if len(field) > 0 {
		opt.Field = field[0]
		name := ""

		if name = opt.Field.Tag.Get("ortam"); name == "" {
			name = toEnvName(opt.Field.Name)
		}

		if prefix != "" {
			opt.Name = fmt.Sprintf("%s_%s", prefix, name)
		} else {
			opt.Name = name
		}
	} else {
		opt.Name = prefix
	}

	opt.Value = value
	opt.Environ, opt.Found = os.LookupEnv(opt.Name)

	return opt
}

func (opt *Option) TypeError() error {
	return &TypeError{
		Env:  opt.Name,
		Type: opt.Field.Type.Name(),
	}
}

func (opt *Option) ParseError(err error) error {
	return &ParseError{
		Env:  opt.Name,
		Err:  err,
		Type: opt.Field.Type.Name(),
	}
}
