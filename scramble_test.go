package scramble
import (
	"testing"
	"fmt"
)

func TestScramble(t *testing.T) {
	for i := 0; i < 10000; i++ {
		salt, inverse := GenerateSalt()
		for in := 1; in < 10000; in++ {
			scr, err := Scrabmble(in, salt, inverse)
			if err != nil {
				t.Errorf(fmt.Sprintf("%s",err))
			}
			out, err := Scrabmble(scr, salt, inverse)
			if err != nil {
				t.Errorf(fmt.Sprintf("%s",err))
			}
			if in != out {
				t.Errorf("i != scramble(scramble(i)),salt: %d, inverse: %d, in: %d, scr:%d, out: %d",
					salt, inverse, in, scr, out)
			}
		}
	}
}
