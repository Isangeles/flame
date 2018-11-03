## Introduction
  Flame is RPG game engine written from scratch in Go.
  
  The main goal is to create simple, flexible, extensible and completely modular game engine. Engine provides build-in CLI frontend to easy run and debug modules.
  
  The project idea is based on [Senlin](https://github.com/isangeles/senlin) game engine, senlin modules should be compatible with Flame.
  
  Currently in a very early development stage.

## Install & Run
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

## Flame CLI commands
Load module:
```
  $engineman -o load -t module -a [module name] [module path](optional)
```
Loads module with the specified name(module directory name) and with a specified path,
if no path provided, the engine will search default modules directory(data/modules).

Create new character:
```
  $newchar
```

Start new game:
```
  $newgame
```

## Contact
* Isangeles <dev@isangeles.pl>

## License
Copyright 2018 Dariusz Sikora <<dev@isangeles.pl>>
 
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
