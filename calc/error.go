package calc

import "fmt"

// InvalidSymbolError return new error in format ninvalid symbol > %s, where %s invalid symbol.
func InvalidSymbolError(s string) error {
	return fmt.Errorf("\ninvalid symbol > %s", s)
}
