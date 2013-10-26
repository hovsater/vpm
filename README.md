# vpm
vpm is a lightweight Vim plugin manager that should work out of the box with
[Pathogen](https://github.com/tpope/vim-pathogen).

## Usage
vpm works with GitHub repositories. `PLUGIN` is of the form
`<username>/<repository>`, e.g., `tpope/vim-fugitive`.

```
Usage:
  vpm [OPTION] [PLUGIN]

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

## License
Copyright (c) 2013 Kevin Sjöberg. vpm may be freely distributed under the MIT
license. More information can be found in the LICENSE file.
