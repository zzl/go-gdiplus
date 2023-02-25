package gdip

func Color(r, g, b byte) ARGB {
	return ARGB(uint32(b) |
		(uint32(g) << 8) |
		(uint32(r) << 16) |
		(uint32(255) << 24))
}

func ColorA(r, g, b, a byte) ARGB {
	return ARGB(uint32(b) |
		(uint32(g) << 8) |
		(uint32(r) << 16) |
		(uint32(a) << 24))
}
