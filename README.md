# vpm
vpm is a lightweight Vim plugin manager that should work out of the box with
[Pathogen](https://github.com/tpope/vim-pathogen).

## Usage

To list your current plugins:

    vpm list

To update your current plugins:

    vpm update

To install a given plugin:

    vpm install PLUGIN

where `PLUGIN` is either a valid [Git](http://git-scm.com/) URL or GitHub
repository such as `tpope/vim-sensible`.

To uninstall a given plugin:

    VPM uninstall PLUGIN

where `PLUGIN` is the name of a currently installed plugin. See `vpm list` for
installed plugins.

## License
Copyright (c) 2013 Kevin Sjöberg. vpm may be freely distributed under the MIT
license. More information can be found in the LICENSE file.
