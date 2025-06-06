## Introduction
  Flame is RPG game engine written from scratch in Go.

  The main goal is to create simple, flexible, extensible and completely modular game engine.
  
  Flame is able to create all game objects from textual data, this guarantees that game data is easy to modify and extend.
  
  This repository contains engine core API, which allows to load game module data from textual format, create and modify game objects, and export game module data back into textual format(i.e. basic game flow load game -> play -> save game).

  Easiest way to create a game with Flame is to use some graphical or textual front-end(like [Mural](https://github.com/isangeles/mural) or [Burn Shell](https://github.com/isangeles/burnsh)) and create flame module containg game data files.

  Example Flame modules are available for download [here](https://flame.isangeles.dev/mods).

  The project idea is based on [Senlin](https://github.com/isangeles/senlin) game engine.

  Flame as a project consists with many different repositories, some of them are independent and can be reused in other projects:

  * [Flame](https://github.com/Isangeles/flame) - engine core API
  * [Burn](https://github.com/Isangeles/burn) - commands interpreter with it's own scripting language [Ash](https://github.com/Isangeles/burn/tree/master/ash) for creating cutscenes etc.
  * [Fire](https://github.com/Isangeles/fire) - TCP server that enables creating multiplayer games
  * [Ignite](https://github.com/Isangeles/ignite) - AI program for the Fire server
  * [Burnsh](https://github.com/Isangeles/burnsh) - textual fronted(CLI)
  * [Mural](https://github.com/Isangeles/mural) - graphical fronted(2D GUI)
  * [MTK](https://github.com/Isangeles/mtk) - simple graphical toolkit
  * [Stone](https://github.com/Isangeles/stone) - simple library to render [Tiled](https://www.mapeditor.org) maps
  * [Arena](https://github.com/Isangeles/arena) - example Flame module

  ### Example games:
  #### Arena ####

  Simple demo game based on [Arena](https://github.com/isangeles/arena) module with [Mural GUI](https://github.com/isangeles/mural) support.

  Download: [Linux](https://my.opendesktop.org/s/xmxszBXyMQCK5xB), [macOS](https://my.opendesktop.org/s/5omoYQYMHGLXkfJ), [Windows](https://my.opendesktop.org/s/gcKQmFRdTj8sBdp)
  #### OpenElwynn ####

  2D Game that recreates the Elwynn Forest area from WoW, with multiplayer support.

  Download: [Linux](https://my.opendesktop.org/s/ctjfGeFAtjBHEXa), [macOS](https://my.opendesktop.org/s/FXyfCYqndaLPCf3), [Windows](https://my.opendesktop.org/s/q52jJCZtpJdy3bb)

  Repository: [GitHub](https://github.com/Isangeles-Softworks/openelwynn)

## Usage
You can find usage examples in [example](https://github.com/Isangeles/flame/tree/master/example) package.

## Modules
Modules contain all game data in the form of text files. Modules are divided into chapters containing chapter-specific data.

Modules are stored by default in `data/modules` directory.

Module data are available across all chapters, data files are placed in sub-directories(`/items`, `/characters`, etc.) in the module directory.

Chapter data are available only when a specific chapter is active, data files are placed in sub-directories(`/characters`, `/dialogs`, etc.) in chapter directory(in `[module]/chapters`).

Translation files are placed in `/lang` directory both for modules and chapters.

The example module is available [here](https://github.com/Isangeles/arena).

## Documentation
Source code documentation can be easily browsed with `go doc` command.

Documentation of configuration files and data structures in the form of Troff pages is available under `doc` directory.

You can easily view documentation pages with `man` command.

For example to display documentation page for character data structure:
```
man doc/characters
```

Note that documentation of data structures is still incomplete.

## Contributing
You are welcome to contribute to project development.

If you looking for things to do, then check [TODO file](https://github.com/Isangeles/flame/blob/master/TODO) or contact maintainer(ds@isangeles.dev).

When you find something to do, create a new branch for your feature.
After you finish, open a pull request to merge your changes with master branch.

## Contact
* Isangeles <<ds@isangeles.dev>>

## License
Copyright 2018-2025 Dariusz Sikora <<ds@isangeles.dev>>

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
