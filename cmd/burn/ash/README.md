## Introduction
Ash is scripting language that executes Burn commands in conditional loop.

## Examples
Every 5 seconds, sends text from first argument on chat channel of game character with serial ID from argument 2:
```
true {
    charman -o set -a chat @1 -t @2;
    wait(5);
}
```