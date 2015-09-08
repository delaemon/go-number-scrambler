package scramble

func trim(v int) int {
	return v & 0xffffffff
}

func scrabmble(v int) int {
	v ^= 0x1ca7bc5b
	v *= 0x1ca7bc5b
	v = trim(v)

	v = ((v >> 1) & 0x55555555) | trim((v & 0x55555555) << 1)
	v = ((v >> 2) & 0x33333333) | trim((v & 0x33333333) << 2)
	v = ((v >> 4) & 0x0F0F0F0F) | trim((v & 0x0F0F0F0F) << 4)
	v = ((v >> 8) & 0x00FF00FF) | trim((v & 0x00FF00FF) << 8)
	v = ( v >> 16)              | trim(v                << 16)

	v *= 0x6b5f13d3
	v = trim(v)
	v ^= 0x1ca7bc5b
	return v
}
