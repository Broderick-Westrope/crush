package tools

import "runtime"

var safeCommands = []string{
	// Bash builtins and core utils
	"cal",
	"date",
	"df",
	"du",
	"echo",
	"env",
	"free",
	"groups",
	"hostname",
	"id",
	"kill",
	"killall",
	"ls",
	"nice",
	"nohup",
	"printenv",
	"ps",
	"pwd",
	"set",
	"time",
	"timeout",
	"top",
	"type",
	"uname",
	"unset",
	"uptime",
	"whatis",
	"whereis",
	"which",
	"whoami",

	// Git
	"git blame",
	"git branch",
	"git config --get",
	"git config --list",
	"git describe",
	"git diff",
	"git grep",
	"git log",
	"git ls-files",
	"git ls-remote",
	"git remote",
	"git rev-parse",
	"git shortlog",
	"git show",
	"git status",
	"git tag",
}

// containsCommandChaining reports whether s contains shell metacharacters
// that enable command chaining or substitution (&&, |, ||, ;, backticks,
// $(...), and standalone & not part of a redirect).
func containsCommandChaining(s string) bool {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ';', '`':
			return true
		case '|':
			return true
		case '&':
			if i+1 < len(s) && s[i+1] == '&' {
				return true // &&
			}
			// Standalone &: not a redirect (>... or ...>).
			precededByGT := i > 0 && s[i-1] == '>'
			followedByGT := i+1 < len(s) && s[i+1] == '>'
			if !precededByGT && !followedByGT {
				return true
			}
		case '$':
			if i+1 < len(s) && s[i+1] == '(' {
				return true // $(
			}
		}
	}
	return false
}

func init() {
	if runtime.GOOS == "windows" {
		safeCommands = append(
			safeCommands,
			// Windows-specific commands
			"ipconfig",
			"nslookup",
			"ping",
			"systeminfo",
			"tasklist",
			"where",
		)
	}
}
