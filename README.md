## Introduction
  Flame is RPG game engine written from scratch in Go.

  The main goal is to create simple, flexible, extensible and completely modular game engine.
  
  Flame parses readable text files and creates game objects, this means that all game data is easy to modify and extend.

  Easiest way to create a game with Flame is to download some graphical or textual fronted(like [Mural](https://github.com/isangeles/mural) or [Burn Shell](https://github.com/isangeles/burnsh)) and create module or modify existing one(like [Arena](https://github.com/Isangeles/arena)).

  The project idea is based on [Senlin](https://github.com/isangeles/senlin) game engine.

  Flame modules are available for download [here](http://flame.isangeles.pl/mods).

  Currently in a early development stage.

  Flame as a project consists with many different repositories, some of them are independent and can be reused in other projects:

  * [Flame](https://github.com/Isangeles/flame) - engine core
  * [Burn](https://github.com/Isangeles/burn) - commands interpreter with it's own scripting language [Ash](https://github.com/Isangeles/burn/tree/master/ash) for creating cutscenes etc.
  * [Burnsh](https://github.com/Isangeles/burnsh) - textual fronted(CLI)
  * [Mural](https://github.com/Isangeles/mural) - graphical fronted(2d GUI)
  * [MTK](https://github.com/Isangeles/mtk) - simple graphical toolkit
  * [Stone](https://github.com/Isangeles/stone) - simple library to render [Tiled](https://www.mapeditor.org) maps
  * [Arena](https://github.com/Isangeles/arena) - example Flame module

  ### Games that use Flame:
  #### Arena ####

  Description: simple demo game based on [Arena](https://github.com/isangeles/arena) module with [Mural GUI](https://github.com/isangeles/mural) support.

  Download: [Linux](https://my.opendesktop.org/s/xmxszBXyMQCK5xB), [Windows](https://my.opendesktop.org/s/gcKQmFRdTj8sBdp)

## Usage
  You can find usage examples in [example](https://github.com/Isangeles/flame/tree/master/example) package.

## Configuration
Configuration values are loaded from `.flame` file in Flame executable directory.

### Configuration values:
```
  lang:[language ID]
```
Description: specifies game language, language ID is name of directory with translation files in lang directories(e.g. `data/lang` or `data/modules/[mod]/lang`).

```
  module:[module ID]
```
Description: specifies module from `data/modules` directory to load at start, module ID is ID specified in `.module` file inside main module directory.

```
  debug:[true/false]
```
Description: enables engine debug mode(shows debug messages in engine log), 'true' enables mode, everything else sets mode disabled.

## Modules
Modules are stored by default in `data/modules` directory, different path to module can be specified in engine configuration file(`.flame`).

Modules contains all game data in form of textual files. Modules are divided into chapters, thats contains chapter-specific data.

Module data are available across all chapters, data files are placed subdirectories(`/items`, `/objects`, etc.) in module directory.

Chapter data are available only when specific chapter is active, data files are placed in subdirectories(`/npc`, `/dialogs`, etc.) in chapter directory(in `[module]/chapters`).

Translation files are placed in `/lang` directory both for modules and chapters.

Example module is available [here](https://github.com/Isangeles/arena).

## Documentation
Documentation of configuration files and data files in the form of troff pages is available under `doc/data` directory.

You can easily view documentation pages with `man` command.

For example to display documentation page for XML files with .characters extension:
```
$ man doc/data/file/xml/.characters
```

Note that documentation is still incomplete.

## Contributing
You are welcome to contribute to project development.

If you looking for things to do, then check [TODO file](https://github.com/Isangeles/flame/blob/master/TODO) or contact me(dev@isangeles.pl).

When you find something to do, create new branch for your feature.
After you finish, open pull request to merge your changes with master branch.

## Contact
* Isangeles <<dev@isangeles.pl>>

## License
Copyright 2018-2020 Dariusz Sikora <<dev@isangeles.pl>>

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
