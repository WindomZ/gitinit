# gitinit
[![Build Status](https://travis-ci.org/WindomZ/gitinit.svg?branch=master)](https://travis-ci.org/WindomZ/gitinit)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

![v0.2.0](https://img.shields.io/badge/version-v0.2.0-yellow.svg)
![status](https://img.shields.io/badge/status-beta-yellow.svg)

A cli tool, easy way to `git init` a new repository.

## Features

- [x] `gitinit --init` - create an empty Git repository or reinitialize an existing one
- [x] `gitinit --bare` - create a bare repository
- [x] `gitinit --origin=REPO` - manage *origin* repository to `REPO`

## Usage

```bash
$ gitinit -h

  A cli tool, easy way to 'git init' a new repository.

  Usage:
    gitinit (-i|--init) [--origin=REPO]
    gitinit -b|--bare
    gitinit --origin=REPO
    gitinit -h|--help
    gitinit -v|--version

  Options:
    -i --init     create an empty Git repository or reinitialize an existing one
    -b --bare     create a bare repository
    --origin=REPO
                  reset 'origin' repository to 'REPO'
    -h --help     show help message
    -v --version  show version
```

## License

The [MIT License](https://github.com/WindomZ/gitinit/blob/dev/LICENSE)
