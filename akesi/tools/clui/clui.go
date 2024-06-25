// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"strings"
)

// GetEnv returns an existing environment variable.
func GetEnv(name string) (string, error) {
	data, ok := os.LookupEnv(name)
	data = strings.TrimSpace(data)

	switch {
	case !ok:
		return "", fmt.Errorf("cannot access envvar %q: does not exist", name)
	case data == "":
		return "", fmt.Errorf("cannot access envvar %q: is blank", name)
	default:
		return data, nil
	}
}

// ParseArgs returns an argument map from a parameter and argument slice. If a
// parameter contains a colon, the text after the colon is a default argument.
func ParseArgs(paras, argus []string) (map[string]string, error) {
	var amap = make(map[string]string)

	for i, para := range paras {
		name, dflt, ok := strings.Cut(para, ":")
		switch {
		case i >= len(argus) && !ok:
			return nil, fmt.Errorf("argument %q not provided", name)
		case i >= len(argus) && ok:
			amap[name] = dflt
		default:
			amap[name] = argus[i]
		}
	}

	return amap, nil
}

// SplitArgs splits an argument slice into a task name and argument slice.
func SplitArgs(argus []string) (string, []string) {
	switch len(argus) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(argus[0]), nil
	default:
		return strings.ToLower(argus[0]), argus[1:]
	}
}
