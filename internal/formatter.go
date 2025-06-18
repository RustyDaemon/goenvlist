package internal

import (
	"fmt"
	"runtime"
	"sort"
	"strings"

	"github.com/fatih/color"
)

type Formatter struct {
	options *Options
	vars    Variables
}

func NewFormatter(options *Options, vars Variables) *Formatter {
	return &Formatter{
		options: options,
		vars:    vars,
	}
}

func (f *Formatter) Display() {
	keys := f.determineKeysToShow()
	f.displaySelected(keys)
}

func (f *Formatter) determineKeysToShow() []string {
	if len(f.options.Filter) > 0 {
		return f.options.Filter
	}

	if f.options.Path {
		return []string{GetPathVariableName()}
	}

	if f.options.Simple {
		commonVars := GetCommonVariables()
		if vars, ok := commonVars[runtime.GOOS]; ok {
			return vars
		}
	}

	return f.getAllKeys()
}

func (f *Formatter) getAllKeys() []string {
	keys := make([]string, 0, len(f.vars))
	for k := range f.vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (f *Formatter) displaySelected(keys []string) {
	showAll := len(keys) == 0

	keySet := make(map[string]struct{})
	for _, k := range keys {
		keySet[k] = struct{}{}
	}

	if showAll {
		keys = f.getAllKeys()
	}

	for _, key := range keys {
		values, exists := f.vars[key]
		if !exists {
			continue
		}

		if f.options.Raw {
			f.displayRaw(key, values)
		} else {
			f.displayFormatted(key, values)
		}
	}
}

func (f *Formatter) displayFormatted(key string, values []string) {
	keyColor := color.New(color.FgGreen).SprintFunc()
	valueColor := color.New(color.FgBlue).SprintFunc()

	if len(values) > 1 {
		fmt.Println(keyColor(key + ":"))
		for _, v := range values {
			fmt.Printf("    %s\n", valueColor(v))
		}
	} else if len(values) == 1 {
		fmt.Printf("%s: %s\n", keyColor(key), valueColor(values[0]))
	}
}

func (f *Formatter) displayRaw(key string, values []string) {
	if len(values) > 1 {
		separator := ":"
		if runtime.GOOS == "windows" {
			separator = ";"
		}
		fmt.Printf("%s=%s\n", key, strings.Join(values, separator))
	} else if len(values) == 1 {
		fmt.Printf("%s=%s\n", key, values[0])
	}
}
