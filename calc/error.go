package calc

import "fmt"

func InvalidSymbolError(s string) error {
	return fmt.Errorf("\ninvalid symbol > %s", s)
}
