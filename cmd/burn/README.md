## Introduction
  Burn is command interpreter for the Flame engine.

## Syntax
  Standard Burn command syntax:
```
  [tool name] -o [option] -t [targets...] -a [arguments...]
```
  Beside many arguments Burn handles also many targets, e.g. one charman tool command can be executed on many game characters, which results with a combined output.
  
  Example command:
```
  >charman -o show -t player_test_0 player_test_1 -a position
```
  Example output:
```
  0.0x0.0 420.0x100.0
```
  Shows positions of game characters with serial IDs 'player_test_0' and 'player_test_1'.

  Commands can also be joined into expressions.

  Target pipe expression:
```
  [command] |t [command]
```
  Executes the first command and uses the output as target arguments to execute next command.

  Example expression:
```
  >moduleman -o show -a area-chars -t area1_test |t charman -o show -a position
```
  Example output:
```
  0.0x0.0 12.0x131.0 130.0x201.0
```
  Shows positions of all game characters in the area with ID 'area1_test'.

## Commands
Set target:
```
  $charman -o set -t [ID]_[serial] -a target [ID]_[serial]
```
Description: sets object with specified serial ID(-a) as target of character with specified serial ID(-t).

Export game character:
```
  $charman -o export -t [character ID]
```
Description: exports game character with specified ID to XML file in
data/modules/[module]/characters directory.

Load module:
```
  $engineman -o load -t module -a [module name] [module path](optional)
```
Description: loads module with the specified name(module directory name) and with a specified path,
if no path provided, the engine will search default modules directory(data/modules).

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
Description: equips item with specified ID for game character with specified serial ID.

Add effect:
```
  $charman -o add -t [character serial ID] -a effect [effect ID]
```
Description: puts effect with specified ID on game character with specified serial ID

Spawn NPC:
```
  $moduleman -o add -t character -a [character ID] [scenario ID] [areaID] [posX](optional) [posY](optional)
```
Description: spawns new chapter NPC with specified ID in specified scenario area at given position(0, 0 if not specified).