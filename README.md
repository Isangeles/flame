## Introduction
  Flame is RPG game engine written from scratch in Go.

  The main goal is to create simple, flexible, extensible and completely modular game engine.
  Flame parses readable text files and creates game objects, this means that all game data is easy to modify and extend.

  Easiest way to use Flame is to download some kind graphical or textual interface, like [Mural](https://github.com/isangeles/mural)(GUI) or [Burn Shell](https://github.com/isangeles/tree/master/cmd)(CLI).

  The project idea is based on [Senlin](https://github.com/isangeles/senlin) game engine.

  Flame modules are available for download [here](http://flame.isangeles.pl/mods).

  Currently in a very early development stage.


  ### List of games that use Flame:
  #### Arena ####

  Description: simple demo game that presents engine and [Mural GUI](https://github.com/isangeles/mural) features.

  Download: [Linux](https://drive.google.com/open?id=1CAUiHdGq8sxrrNWkRwF1QSaNSVWLKDVg), [Windows](https://drive.google.com/open?id=1rR_k_39o-hqTywUZO628ggA3iN7ZBZTJ)

## Usage
  You can find usage examples in [example](https://github.com/Isangeles/flame/tree/master/exmaple) package.

## Build & Run
  Get sources from git:
```
  $ go get -u github.com/isangeles/flame
```
  Build CLI:
```
  $ go build github.com/isangeles/flame/cmd
```
  Run CLI:
```
  $ ./cmd
```

## Configuration
Configuration values are loaded from '.flame' file in Flame executable directory.

### Configuration values:
```
  lang:[language ID];
```
Description: specifies game language, language ID is name of directory with translation files in lang directories(e.g. 'data/lang' or 'data/modules/[mod]/lang').

```
  module:[module ID];[module path](optional);
```
Description: specifies module to load at start, module ID is ID specified in 'mod.conf' file inside main module directory, module path is optional and stands for module directory path, if not provided engine will search default modules path('data/modules').

```
  debug:[true/false];
```
Description: enables engine debug mode(shows debug messages in engine log), 'true' enables mode, everything else sets mode disabled.

## Burn Shell
Flame comes with [Burn Shell](https://github.com/Isangeles/flame/tree/master/cmd), a simple textual interface that uses [Burn](https://github.com/Isangeles/flame/tree/master/cmd/burn) commands interpreter to execute commands.

## Contact
* Isangeles <<dev@isangeles.pl>>

## License
Copyright 2018-2019 Dariusz Sikora <<dev@isangeles.pl>>

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
