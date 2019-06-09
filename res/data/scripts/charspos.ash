# Script thats show position of all character in area with ID from first argument.
{
    echo(moduleman -o show -a area-chars -t @1 |t charman -o show -a position);
}