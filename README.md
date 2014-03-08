# vpm

vpm is a lightweight Vim plugin manager intended to work alongside
[Pathogen](https://github.com/tpope/vim-pathogen).

## Usage

vpm works with GitHub repositories. `PLUGIN` is of the form
`<username>/<repository>`, e.g., `tpope/vim-fugitive`. Plugins can be specified
in `~/.vimplugins`. When `vpm bootstrap` is run these
will be inserted for you.

```
Usage:
  vpm [OPTION] [COMMAND] [PLUGIN]

Options:
  -h, --help     - Usage instructions
  -v, --version, - Print current version

Commands:
  List:
    vpm list

  Insert:
    vpm        PLUGIN
    vpm insert PLUGIN

  Remove:
    vpm remove PLUGIN

  Update:
    vpm update [PLUGIN]

  Bootstrap:
    vpm bootstrap
```

## Installation

### Homebrew

vpm can be installed with [Homebrew](http://brew.sh/).

    brew tap KevinSjoberg/formulas
    brew install vpm

## Development

vpm tests are executed using [Bats](https://github.com/sstephenson/bats)
`brew install bats`

    bats test
    bats test/<file>.bats


## License

Copyright (c) 2013 Kevin Sjöberg. vpm may be freely distributed under the MIT
license. More information can be found in the LICENSE file.
