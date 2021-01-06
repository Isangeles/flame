## Introduction
  Flame is RPG game engine written from scratch in Go.

  The main goal is to create simple, flexible, extensible and completely modular game engine.
  
  Flame parses readable text files and creates game objects, this means that all game data is easy to modify and extend.

  Easiest way to create a game with Flame is to download some graphical or textual fronted(like [Mural](https://github.com/isangeles/mural) or [Burn Shell](https://github.com/isangeles/burnsh)) and create module or use some existing one.

  Flame modules are available for download [here](http://flame.isangeles.pl/mods).

  The project idea is based on [Senlin](https://github.com/isangeles/senlin) game engine.

  Currently in an early development stage.

  Flame as a project consists with many different repositories, some of them are independent and can be reused in other projects:

  * [Flame](https://github.com/Isangeles/flame) - engine core
  * [Burn](https://github.com/Isangeles/burn) - commands interpreter with it's own scripting language [Ash](https://github.com/Isangeles/burn/tree/master/ash) for creating cutscenes etc.
  * [Fire](https://github.com/Isangeles/fire) - TCP server that enables creating multiplayer games
  * [Burnsh](https://github.com/Isangeles/burnsh) - textual fronted(CLI)
  * [Mural](https://github.com/Isangeles/mural) - graphical fronted(2D GUI)
  * [MTK](https://github.com/Isangeles/mtk) - simple graphical toolkit
  * [Stone](https://github.com/Isangeles/stone) - simple library to render [Tiled](https://www.mapeditor.org) maps
  * [Arena](https://github.com/Isangeles/arena) - example Flame module

  ### Example games:
  #### Arena ####

  Description: simple demo game based on [Arena](https://github.com/isangeles/arena) module with [Mural GUI](https://github.com/isangeles/mural) support.

  Download: [Linux](https://my.opendesktop.org/s/xmxszBXyMQCK5xB), [Windows](https://my.opendesktop.org/s/gcKQmFRdTj8sBdp), [macOS](https://my.opendesktop.org/s/5omoYQYMHGLXkfJ)

## Usage
You can find usage examples in [example](https://github.com/Isangeles/flame/tree/master/example) package.

## Modules
Modules contain all game data in the form of textual files. Modules are divided into chapters, that's contains chapter-specific data.

Modules are stored by default in `data/modules` directory.

Module data are available across all chapters, data files are placed in sub-directories(`/items`, `/objects`, etc.) in the module directory.

Chapter data are available only when a specific chapter is active, data files are placed in sub-directories(`/characters`, `/dialogs`, etc.) in chapter directory(in `[module]/chapters`).

Translation files are placed in `/lang` directory both for modules and chapters.

The example module is available [here](https://github.com/Isangeles/arena).

## Documentation
Source code documentation can be easily browsed with `go doc` command.

Documentation of configuration files and data files in the form of Troff pages is available under `doc/data` directory.

You can easily view documentation pages with `man` command.

For example to display documentation page for XML files with .characters extension:
```
$ man doc/data/file/xml/.characters
```

Note that documentation of data files is still incomplete.

## Contributing
You are welcome to contribute to project development.

If you looking for things to do, then check [TODO file](https://github.com/Isangeles/flame/blob/master/TODO) or contact me(dev@isangeles.pl).

When you find something to do, create a new branch for your feature.
After you finish, open a pull request to merge your changes with master branch.

## Contact
* Isangeles <<dev@isangeles.pl>>

## License
Copyright 2018-2021 Dariusz Sikora <<dev@isangeles.pl>>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
MA 02110-1301, USA.
