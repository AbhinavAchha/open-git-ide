package main

import (
	"fmt"
	"os/exec"
	"strings"

	"golang.org/x/exp/slog"
)

func main() {
	p, err := getCWDGhRepo()
	if err != nil {
		slog.Error(err.Error())
	}

	if err = exec.Command("xdg-open", p).Run(); err != nil {
		slog.Error(err.Error())
	}

}

func getCWDGhRepo() (string, error) {
	output, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		err = fmt.Errorf("failed to get git remote url %w", err)
		slog.Error(err.Error())
		return "", err
	}

	domain, p := getDomainPath(string(output))

	p = p[:len(p)-5]

	switch domain {
	case "github":
		p = fmt.Sprintf("github.dev/%s", p)
	case "gitlab":
		p = fmt.Sprintf("gitlab.com/-/ide/project/%s", p)
	default:
		err := fmt.Errorf("unknown git remote domain %s", domain)
		slog.Error(err.Error())
		return "", err
	}

	return p, nil
}

func getDomainPath(s string) (string, string) {
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "git@")
	return s[:6], s[11:]
}
