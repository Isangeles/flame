# Script thats sends text to tutorial character chat channel when PC aproching.
@1 = "player_lana_0"
@2 = "testchar_0"
@3 = "test"
{
	rawdis(@1, @2) < 50 {
    		charman -o set -a chat @3 -t testchar_0 ; @wait 5
	}
}
