## Introduction
  Burn Shell is command line interface for [Flame engine](https://github.com/isangeles/flame).

  CLI uses [Burn](https://github.com/Isangeles/flame/tree/master/cmd/burn) to handle user input and communicate with engine.
  
  All commands must be prefixed with '$' character.

## Usage
  
## Commands
To run Burn or bsh command use '$' character as prefix.
Without prefix, command will be treated as text and printed to out or sent to active player
chat channel if game was started

### Burn Shell build-in commands:

Create module:
```
  $newmod
```
Description: Starts new module creation dialog. New module will be created in 'data/modules' directory. New module contains one chapter and start area.

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

Import exported characters:
```
  $importchars
```
Description: imports all characters from XML files in
data/modules/[module]/characters directory.

Set target:
```
  $target
```
Description: searches current area for nearby targets to set for active PC.

Target information:
```
  $tarinfo
```
Description: prints informations about active PC target.

Loot target:
```
  $loot
```
Description: transfers all items from current dead target to active PC.

Talk with with target:
```
  $talk
```
Description: starts dialog with current PC target.

Show quests journal:
```
  $quests
```
Description: shows active PC quests.

Use character skill:
```
  $useskill
```
Description: starts dialog to use one of active PC skills.

Exit program:
```
  $close
```
Description: terminates program.

## Scripts
To run Ash script use '%' prefix, scripts are executed from 'data/scripts' directory.
Use '&' suffix to run script in background.
