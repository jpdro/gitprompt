package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	branch := gitBranch(wd)
	if branch == "" {
		return
	}

	// %% is how you escape a percent sign. It translates to a single '%' char.
	colorized := fmt.Sprintf(" %%F{blue}(%%F{red}%s%%F{blue})%%f", branch)
	dirty := isDirty()
	if err != nil {
		fmt.Println(colorized)
		return
	}

	if dirty {
		colorized += "💩"
	}

	fmt.Println(colorized)
}

// gitBranch returns the name of the current branch. This is recorded in the
// .git/HEAD file. If we're not in a git repo or there's an error, it will
// return an empty string.
func gitBranch(path string) string {
	if strings.Contains(path, string(os.PathSeparator)+".git") {
		return ".git"
	}

	d := findGitDir(path)
	if d == "" {
		return ""
	}

	b, err := ioutil.ReadFile(filepath.Join(d, ".git/HEAD"))
	if err != nil {
		return ""
	}

	prefix := []byte("refs/heads/")
	i := bytes.Index(b, prefix)
	if i == -1 {
		return ""
	}
	i += len(prefix)

	return string(bytes.TrimSpace(b[i:]))
}

// findGitDir walks from the current directory to the root directory, returning
// the absolute path to the first .git directory it finds. If no .git directory
// is found, it returns an error.
func findGitDir(path string) string {
	for path != "" {
		_, err := os.Stat(filepath.Join(path, ".git"))
		if err == nil {
			return path
		}

		path = strings.TrimRight(path, string(os.PathSeparator))
		path, _ = filepath.Split(path)
	}

	return ""
}

// isDirty returns true if the repo has changed since the last commit.
func isDirty() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var buf strings.Builder
	cmd := exec.CommandContext(ctx, "git", "ls-files", "--deleted", "--modified", "--unmerged", "--killed", "--other", "--exclude-standard")
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		return false
	}

	stdout := strings.TrimSpace(buf.String())

	// If there's any output, git ls-files has printed filenames. This
	// means there are dirty or untracked files.
	return len(stdout) > 0
}
