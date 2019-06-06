# Script thats sends text to tutorial character chat channel when PC aproching.
@pc near testchar_0 {
    charman -o set -a chat hay -t testchar_0 ; @wait 5
}