# gitinit
[![Build Status](https://travis-ci.org/WindomZ/gitinit.svg?branch=master)](https://travis-ci.org/WindomZ/gitinit)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

![v0.1.1](https://img.shields.io/badge/version-v0.1.1-yellow.svg)
![status](https://img.shields.io/badge/status-beta-yellow.svg)

A cli tool, easy way to `git init` a new repository.

## Features

- [x] `gitinit --init` - create an empty Git repository or reinitialize an existing one
- [x] `gitinit --init --remote=<repo>` - same as `gitinit --init`, and manage *origin* repository to `<repo>`
- [x] `gitinit --bare` - create a bare repository

## Usage

```bash
$ gitinit -h

  A cli tool, easy way to 'git init' a new repository.

  Usage:
    gitinit (-i|--init) [--remote=<repo>]
    gitinit -b|--bare
    gitinit -h|--help
    gitinit -v|--version

  Options:
    --remote=<repo>
                  manage 'origin' repository to '<repo>'
    -h --help     show help message
    -v --version  show version
```

## License

The [MIT License](https://github.com/WindomZ/gitinit/blob/dev/LICENSE)
