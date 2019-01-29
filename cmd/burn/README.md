## Introduction
  Burn is command interpreter for the Flame engine.

## Syntax
  Standard Burn command syntax:
```
  [tool name] -o [option] -t [targets...] -a [arguments...]
```
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