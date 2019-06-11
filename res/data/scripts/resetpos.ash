# Script thats sets position of character with from first arg to 0x0.
{
    charman -o set -a position 0 0 -t @1;
    wait(5);  
    echo(charman -o show -a position -t @1);
}