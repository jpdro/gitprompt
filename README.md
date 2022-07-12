# gitprompt

This is a tiny little program that prints the current git branch in oh-my-zsh style. The 💩 means the repo is [dirty](https://stackoverflow.com/questions/20642980/does-git-dirty-mean-files-not-staged-or-not-committed-glossary-conflict).

![gitprompt screenshot](https://i.imgur.com/F2IcOn6.jpg)

## Installation

| OS     | Platform        | URL                                                                                                           |
| ------ | --------------- | ------------------------------------------------------------------------------------------------------------- |
| macOS  | Apple Silicon   | [Download](https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-aarch64-apple-darwin)        |
| macOS  | Intel           | [Download](https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-x86_64-apple-darwin)         |
| Linux  | x86_64          | [Download](https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-x86_64-unknown-linux-gnu)    |
| Linux  | x86_64 (musl)   | [Download](https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-x86_64-unknown-linux-musl)   |
| Linux  | ARM64           | [Download](https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-aarch64-unknown-linux-gnu)   |
| Linux  | ARM64 (musl)    | [Download](https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-aarch64-unknown-linux-musl)  |
| Linux  | RISC V          | [Download](https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-riscv64gc-unknown-linux-gnu) |

The rows marked `musl` are for Linux distros that use [musl libc](https://en.wikipedia.org/wiki/Musl).
The most popular such distro is [Alpine Linux](https://www.alpinelinux.org). If
you're not running Alpine, you almost certainly want the non-`musl` version.

Here's how to install `gitprompt` on an Apple Silicon Mac. You can adjust the
URL for your platform.

```sh
sudo curl --retry 3 --max-time 60 --output /usr/local/bin/gitprompt -sSfL https://github.com/ryboe/gitprompt/releases/latest/download/gitprompt-aarch64-apple-darwin
sudo chmod +x /usr/local/bin/gitprompt
```

## Usage

You can use it in your `.zshrc` like this:

```sh
# enable shell commands in the prompt string. required for $(gitprompt)
setopt prompt_subst

# fancy prompt formatting documented at http://zsh.sourceforge.net/Doc/Release/Prompt-Expansion.html
export PROMPT='%F{cyan}%B%f %F{yellow}$(gitprompt)%f '
```
