package option

import (
	"fmt"
	"strings"
)

// environment variable has an invalid type
type TypeError struct {
	Env  string
	Type string
}

func (e *TypeError) Error() string {
  typ := strings.ToLower(e.Type)

  switch typ[0] {
  case 'a', 'e', 'i', 'o', 'u':
	  return fmt.Sprintf("expected an %s value for %s", typ, e.Env)
  default:
	  return fmt.Sprintf("expected a %s value for %s", typ, e.Env)
  }
}

// failed to parse an environment variable
type ParseError struct {
	Env  string
	Type string
	Err  error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("failed to parse %s as %s: %s", e.Env, strings.ToLower(e.Type), e.Err.Error())
}
