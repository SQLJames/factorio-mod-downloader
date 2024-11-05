package semver

import (
	"fmt"
	"strings"

	xsemver "golang.org/x/mod/semver"
)

func GetMajorMinor(v string) (string, error) {
	if !strings.HasPrefix(v, "v") {
		v = fmt.Sprintf("v%s", v)
	}
	if !xsemver.IsValid(v) {
		return "", fmt.Errorf("string is not a valid semver: %s", v)
	}
	return xsemver.MajorMinor(v), nil
}
