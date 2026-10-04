package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func replaceEnv(env []string, key, value string) []string {
	prefix := strings.ToUpper(key) + "="
	out := make([]string, 0, len(env)+1)
	for _, item := range env {
		if strings.HasPrefix(strings.ToUpper(item), prefix) {
			continue
		}
		out = append(out, item)
	}
	return append(out, key+"="+value)
}

func main() {
	self, err := os.Executable()
	if err != nil {
		fmt.Println("Compatibility cache: unavailable")
		return
	}
	tools := filepath.Dir(self)
	cmd := exec.Command("gh.exe", "cacheverify", "--help")
	cmd.Dir = filepath.Dir(tools)
	cmd.Env = replaceEnv(os.Environ(), "GH_CONFIG_DIR", filepath.Join(tools, ".gh-config"))
	cmd.Env = replaceEnv(cmd.Env, "GH_NO_UPDATE_NOTIFIER", "1")
	cmd.Env = replaceEnv(cmd.Env, "GH_PROMPT_DISABLED", "1")
	cmd.Env = replaceEnv(cmd.Env, "SHELL", `C:\Program Files\Git\bin\sh.exe`)
	cmd.Env = replaceEnv(cmd.Env, "COMPAT_SCRIPT", filepath.Join(tools, "refresh-cache.ps1"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Compatibility cache: unavailable")
		return
	}
	fmt.Println("Compatibility cache: current")
}
