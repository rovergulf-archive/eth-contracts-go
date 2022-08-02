package ethutils

import (
	"regexp"
)

var AddressRegexp = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func IsStringAddress(address string) bool {
	return AddressRegexp.MatchString(address)
}
