package scramble
import (
	"testing"
	"fmt"
)

func TestScramble(t *testing.T) {
	for in := 0; in < 1000; in++ {
		scr := scrabmble(in)
		out := scrabmble(scr)
		fmt.Printf("in: %d, scr: %d, out: %d\n",in, scr, out)
		if in != out {
			t.Errorf("i != scramble(scramble(i)), in: %d, scr:%d, out: %d", in, scr, out)
		}
	}
}
