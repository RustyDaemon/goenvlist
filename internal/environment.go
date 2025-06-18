package internal

import (
	"fmt"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

type Variables map[string][]string

func GetEnvironment() (Variables, error) {
	output, err := runEnvCommand()
	if err != nil {
		return nil, fmt.Errorf("failed to get environment variables: %w", err)
	}

	vars, err := parseEnvOutput(output)
	if err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return vars, nil
}

func runEnvCommand() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "-Command", "Get-ChildItem Env: | ForEach-Object { $_.Name + '=' + $_.Value }")
	case "darwin", "linux", "freebsd", "openbsd", "netbsd":
		cmd = exec.Command("env")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %w", err)
	}

	return string(output), nil
}

func parseEnvOutput(source string) (Variables, error) {
	result := make(Variables)

	separator := ":"
	if runtime.GOOS == "windows" {
		separator = ";"
	}

	for _, line := range strings.Split(source, "\n") {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key, value := parts[0], parts[1]

		if isPathVariable(key) {
			rawValues := strings.Split(value, separator)
			filtered := make([]string, 0, len(rawValues))

			for _, v := range rawValues {
				v = strings.TrimSpace(v)
				if v != "" {
					filtered = append(filtered, v)
				}
			}

			result[key] = filtered
		} else {
			result[key] = []string{value}
		}
	}

	for k, v := range result {
		if len(v) > 1 {
			sort.Strings(v)
			result[k] = v
		}
	}

	return result, nil
}

func isPathVariable(name string) bool {
	pathVars := map[string]bool{
		"PATH":            true,
		"Path":            true,
		"GOPATH":          true,
		"CLASSPATH":       true,
		"LD_LIBRARY_PATH": true,
		"LIBRARY_PATH":    true,
	}

	return pathVars[name]
}

func GetCommonVariables() map[string][]string {
	return map[string][]string{
		"darwin": {
			"PATH", "HOME", "SHELL", "USER", "LOGNAME",
			"LANG", "PWD", "OLDPWD", "SHLVL", "_", "TERM", "TMPDIR",
		},
		"linux": {
			"PATH", "HOME", "SHELL", "USER", "LOGNAME",
			"LANG", "PWD", "OLDPWD", "SHLVL", "_", "TERM", "TMPDIR",
		},
		"windows": {
			"APPDATA", "COMPUTERNAME", "HOMEDRIVE", "HOMEPATH",
			"OS", "Path", "ProgramData", "ProgramFiles",
			"ProgramFiles(x86)", "SESSIONNAME", "SystemDrive",
			"USERDOMAIN", "USERPROFILE",
		},
	}
}

func GetPathVariableName() string {
	if runtime.GOOS == "windows" {
		return "Path"
	}
	return "PATH"
}
