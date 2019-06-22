# Script that sends arg3 text on chat channel of character with
# arg 2 serial ID if raw distance between him and character with
# arg 1 serial ID is less than 50, after that waits 5 secs.
true {
     rawdis(@1, @2) < 50 {
     	charman -o set -a chat @3 -t @2;
	wait(5);
     };
}