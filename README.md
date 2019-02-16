## Introduction
  Flame is RPG game engine written from scratch in Go.
  
  The main goal is to create simple, flexible, extensible and completely modular game engine. Engine provides build-in CLI frontend to easy run and debug modules. Engine does not come with any graphical interface, instead is designed to be used with external GUI's(e.g. [Mural](https://github.com/isangeles/mural)).
  
  The project idea is based on [Senlin](https://github.com/isangeles/senlin) game engine, Senlin modules should be compatible with Flame(with some modifications).

  Flame modules are available for download [here](http://flame.isangeles.pl/mods).
  
  Currently in a very early development stage.

## Build & Run
  Get sources from git:
```
  $ go get -u github.com/isangeles/flame
```
  Build Flame CLI:
```
  $ go build github.com/isangeles/flame/cmd
```
  Run CLI:
```
  $ ./cmd
```

## Flame CLI:
Flame comes with a simple textual interface that uses [Burn](https://github.com/Isangeles/flame/tree/master/cmd/burn) commands interpreter to execute commands.

### Commands:
Load module:
```
  $engineman -o load -t module -a [module name] [module path](optional)
```
Description: loads module with the specified name(module directory name) and with a specified path,
if no path provided, the engine will search default modules directory(data/modules).

Create new character:
```
  $newchar
```
Description: starts new character creation dialog.

Start new game:
```
  $newgame
```
Description: starts new game dialog.

Load game:
```
  $loadgame
```
Description: starts load game dialog.

Export game character:
```
  $charman -o export -t [character ID]
```
Description: exports game character with specified ID to XML file in
data/modules/[module]/characters directory.

Import exported characters:
```
  $importchars
```
Description: imports all characters from XML files in
data/modules/[module]/characters directory.

Save game:
```
  $engineman -o save -t game -a [save file name]
```
Description: saves current game to 'savegames/[module]' directory.

Add item:
```
  $charman -o add -t [character serial ID] -a item [item ID]
```
Description: adds item with specified ID to inventory of game character with specified serial ID.

Equip item:
```
  $charman -o equip -t [character serial ID] -a [slot ID] [item serial ID]
```
Description: equips item with specified serial ID for game character with specified serial ID.

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
