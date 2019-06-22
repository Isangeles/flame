# Script that sets position of character with serial ID from arg 1
# to 0x0 if range between him and character with serial ID from arg 2
# is less than 50.
true {
     rawdis(@1, @2) < 50 {
     	charman -o set -a position 0 0 -t @1;
	wait(5);
     };
}