package scramble
import (
	"math/rand"
	"time"
	"fmt"
	"errors"
)

func GenerateSalt() (salt, inverse int) {
	salt = generateSalt()
	inverse = modInverse(salt, 0x100000000)
	return
}

func generateSalt() int {
	rand.Seed(time.Now().Unix())
	salt := rand.Intn(0xFFFFFFFF - 3) + 3
	if salt % 2 == 0 {
		salt += 1
	}
	return salt
}

func egcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	} else {
		g, y, x := egcd(b % a, a)
		return g, (x - (b/a) * y), y
	}

}

func modInverse(a, m int) int {
	g, x, _ := egcd(a, m)
	if g != 1 {
		panic("inverse does not exist.")
	}
	return x % m
}

func assertNumber(num int) bool {
	return (1 <= num && num <= 0xFFFFFFFF)
}

func assertSalt(salt, inverse int) bool {
	return trim(salt * inverse) == 1
}

func trim(num int) int {
	return num & 0xFFFFFFFF
}

func reverse32bit(num int) int {
	num = ((num >> 1) & 0x55555555) | trim((num & 0x55555555) << 1)
	num = ((num >> 2) & 0x33333333) | trim((num & 0x33333333) << 2)
	num = ((num >> 4) & 0x0F0F0F0F) | trim((num & 0x0F0F0F0F) << 4)
	num = ((num >> 8) & 0x00FF00FF) | trim((num & 0x00FF00FF) << 8)
	num = ( num >> 16)              | trim( num               << 16)
	return num
}

func Scrabmble(num, salt, inverse int) (int, error) {
	if !assertNumber(num) {
		e := fmt.Sprintf("invalid number. num = %d", num)
		return -1, errors.New(e)
	}
	if !assertSalt(salt, inverse) {
		e := fmt.Sprintf("invalid salt. salt = %d, inverse: %d",salt ,inverse)
		return -1, errors.New(e)
	}

	num ^= salt
	num = trim(num * salt)
	num = reverse32bit(num)
	num *= inverse
	num = trim(num)
	num ^= salt
	return num, nil
}
