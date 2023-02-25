package gdip

import (
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
)

const LANG_NEUTRAL = 0x00

const (
	AlphaShift = 24
	RedShift   = 16
	GreenShift = 8
	BlueShift  = 0
)

const (
	AlphaMask = 0xff000000
	RedMask   = 0x00ff0000
	GreenMask = 0x0000ff00
	BlueMask  = 0x000000ff
)

// In-memory pixel data formats:
// bits 0-7 = format index
// bits 8-15 = pixel size (in bits)
// bits 16-23 = flags
// bits 24-31 = reserved
type PixelFormat int32

const (
	PixelFormatIndexed   = 0x00010000 // Indexes into a palette
	PixelFormatGDI       = 0x00020000 // Is a GDI-supported format
	PixelFormatAlpha     = 0x00040000 // Has an alpha component
	PixelFormatPAlpha    = 0x00080000 // Pre-multiplied alpha
	PixelFormatExtended  = 0x00100000 // Extended color 16 bits/channel
	PixelFormatCanonical = 0x00200000

	PixelFormatUndefined = 0
	PixelFormatDontCare  = 0

	PixelFormat1bppIndexed    = 1 | (1 << 8) | PixelFormatIndexed | PixelFormatGDI
	PixelFormat4bppIndexed    = 2 | (4 << 8) | PixelFormatIndexed | PixelFormatGDI
	PixelFormat8bppIndexed    = 3 | (8 << 8) | PixelFormatIndexed | PixelFormatGDI
	PixelFormat16bppGrayScale = 4 | (16 << 8) | PixelFormatExtended
	PixelFormat16bppRGB555    = 5 | (16 << 8) | PixelFormatGDI
	PixelFormat16bppRGB565    = 6 | (16 << 8) | PixelFormatGDI
	PixelFormat16bppARGB1555  = 7 | (16 << 8) | PixelFormatAlpha | PixelFormatGDI
	PixelFormat24bppRGB       = 8 | (24 << 8) | PixelFormatGDI
	PixelFormat32bppRGB       = 9 | (32 << 8) | PixelFormatGDI
	PixelFormat32bppARGB      = 10 | (32 << 8) | PixelFormatAlpha | PixelFormatGDI | PixelFormatCanonical
	PixelFormat32bppPARGB     = 11 | (32 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatGDI
	PixelFormat48bppRGB       = 12 | (48 << 8) | PixelFormatExtended
	PixelFormat64bppARGB      = 13 | (64 << 8) | PixelFormatAlpha | PixelFormatCanonical | PixelFormatExtended
	PixelFormat64bppPARGB     = 14 | (64 << 8) | PixelFormatAlpha | PixelFormatPAlpha | PixelFormatExtended
	PixelFormat32bppCMYK      = 15 | (32 << 8)
	PixelFormatMax            = 16
)

// Predefined multi-frame dimension IDs
var (
	FrameDimensionTime = syscall.GUID{0x6aedbd6d, 0x3fb5, 0x418a,
		[8]byte{0x83, 0xa6, 0x7f, 0x45, 0x22, 0x9d, 0xc8, 0x72}}
)

var (
	BmpEncoderId = win32.CLSID{0x557cf400, 0x1a04, 0x11d3,
		[8]byte{0x9a, 0x73, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e}}

	JpgEncoderId = win32.CLSID{0x557cf401, 0x1a04, 0x11d3,
		[8]byte{0x9a, 0x73, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e}}

	PngEncoderId = win32.CLSID{0x557cf406, 0x1a04, 0x11d3,
		[8]byte{0x9a, 0x73, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e}}

	GifEncoderId = win32.CLSID{0x557cf402, 0x1a04, 0x11d3,
		[8]byte{0x9a, 0x73, 0x00, 0x00, 0xf8, 0x1e, 0xf3, 0x2e}}
)
