package gdip

import (
	"syscall"
)

type Size struct {
	Width  int32
	Height int32
}

type SizeF struct {
	Width  float32
	Height float32
}

type Point struct {
	X int32
	Y int32
}

type PointF struct {
	X float32
	Y float32
}

type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type RectF struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type PathData struct {
	Count int32
	Points *PointF
}

type CharacterRange struct {
	First  int32
	Length int32
}

//
type Graphics struct{}
type Brush struct{}
type Texture struct{ Brush }
type SolidFill struct{ Brush }
type LineGradient struct{ Brush }
type PathGradient struct{ Brush }
type Hatch struct{ Brush }

type Pen struct{}
type CustomLineCap struct{}
type AdjustableArrowCap struct{ CustomLineCap }

type Image struct{}
type Bitmap struct{ Image }

type ImageAttributes struct{}

type Path struct{}
type Region struct{}
type PathIterator struct{}

type FontFamily struct{}
type Font struct{}
type StringFormat struct{}
type FontCollection struct{}
type InstalledFontCollection struct{ FontCollection }
type PrivateFontCollection struct{ FontCollection }

type CachedBitmap struct{}

//
type ARGB uint32
type Matrix struct{}

type BitmapData struct {
	Width       uint32
	Height      uint32
	Stride      int32
	PixelFormat PixelFormat
	Scan0       uintptr
	Reserved    uintptr
}

// Windows Types
type HANDLE uintptr
type HPALETTE uintptr
type HBITMAP uintptr
type HDC uintptr
type HWND uintptr

type GraphicsState uint32
type GraphicsContainer uint32


type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread int32
	SuppressExternalCodecs   int32
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}


type EncoderParameter struct {
	Guid           syscall.GUID
	NumberOfValues uint32
	TypeAPI        uint32
	Value          uintptr
}

type EncoderParameters struct {
	Count     uint32
	Parameter [1]EncoderParameter
}
type PropertyItem struct {
	Id     uint32
	Length uint32
	Type   uint16
	Value  uintptr
}

type ColorPalette struct {
	Flags   uint32
	Count   uint32
	Entries [1]ARGB
}

type CGpEffect struct{}

type ImageItemData struct{}
type GdiplusAbort struct{}
type Metafile struct{}

type ColorMap struct{}
type ColorMatrix struct{}
type MetafileHeader struct{}
type WmfPlaceableFileHeader struct{}
type ImageCodecInfo struct{}

type GpPathData struct{}

func NewRect(x, y, width, height int32) *Rect {
	return &Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func NewRectF(x, y, width, height float32) *RectF {
	return &RectF{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (rect *Rect) Left() int32 {
	return rect.X
}

func (rect *Rect) Top() int32 {
	return rect.Y
}

func (rect *Rect) Right() int32 {
	return rect.X + rect.Width
}

func (rect *Rect) Bottom() int32 {
	return rect.Y + rect.Height
}

func (rect *RectF) Left() float32 {
	return rect.X
}

func (rect *RectF) Top() float32 {
	return rect.Y
}

func (rect *RectF) Right() float32 {
	return rect.X + rect.Width
}

func (rect *RectF) Bottom() float32 {
	return rect.Y + rect.Height
}

func (s Status) String() string {
	switch s {
	case Ok:
		return "Ok"

	case GenericError:
		return "GenericError"

	case InvalidParameter:
		return "InvalidParameter"

	case OutOfMemory:
		return "OutOfMemory"

	case ObjectBusy:
		return "ObjectBusy"

	case InsufficientBuffer:
		return "InsufficientBuffer"

	case NotImplemented:
		return "NotImplemented"

	case Win32Error:
		return "Win32Error"

	case WrongState:
		return "WrongState"

	case Aborted:
		return "Aborted"

	case FileNotFound:
		return "FileNotFound"

	case ValueOverflow:
		return "ValueOverflow"

	case AccessDenied:
		return "AccessDenied"

	case UnknownImageFormat:
		return "UnknownImageFormat"

	case FontFamilyNotFound:
		return "FontFamilyNotFound"

	case FontStyleNotFound:
		return "FontStyleNotFound"

	case NotTrueTypeFont:
		return "NotTrueTypeFont"

	case UnsupportedGdiplusVersion:
		return "UnsupportedGdiplusVersion"

	case GdiplusNotInitialized:
		return "GdiplusNotInitialized"

	case PropertyNotFound:
		return "PropertyNotFound"

	case PropertyNotSupported:
		return "PropertyNotSupported"

	case ProfileNotFound:
		return "ProfileNotFound"
	}

	return "Unknown Status Value"
}
