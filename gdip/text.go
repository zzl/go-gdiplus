package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// stringformat
	gdipCreateStringFormat	*windows.LazyProc
	gdipStringFormatGetGenericDefault	*windows.LazyProc
	gdipStringFormatGetGenericTypographic	*windows.LazyProc
	gdipDeleteStringFormat	*windows.LazyProc
	gdipCloneStringFormat	*windows.LazyProc
	gdipSetStringFormatFlags	*windows.LazyProc
	gdipGetStringFormatFlags	*windows.LazyProc
	gdipSetStringFormatAlign	*windows.LazyProc
	gdipGetStringFormatAlign	*windows.LazyProc
	gdipSetStringFormatLineAlign	*windows.LazyProc
	gdipGetStringFormatLineAlign	*windows.LazyProc
	gdipSetStringFormatTrimming	*windows.LazyProc
	gdipGetStringFormatTrimming	*windows.LazyProc
	gdipSetStringFormatHotkeyPrefix	*windows.LazyProc
	gdipGetStringFormatHotkeyPrefix	*windows.LazyProc
	gdipSetStringFormatTabStops	*windows.LazyProc
	gdipGetStringFormatTabStops	*windows.LazyProc
	gdipGetStringFormatTabStopCount	*windows.LazyProc
	gdipSetStringFormatDigitSubstitution	*windows.LazyProc
	gdipGetStringFormatDigitSubstitution	*windows.LazyProc
	gdipGetStringFormatMeasurableCharacterRangeCount	*windows.LazyProc
	gdipSetStringFormatMeasurableCharacterRanges	*windows.LazyProc

	// text
	gdipDrawString	*windows.LazyProc
	gdipMeasureString	*windows.LazyProc
	gdipMeasureCharacterRanges	*windows.LazyProc
	gdipDrawDriverString	*windows.LazyProc
	gdipMeasureDriverString	*windows.LazyProc
)

func init() {

	// stringformat
	gdipCreateStringFormat = dll.NewProc("GdipCreateStringFormat")
	gdipStringFormatGetGenericDefault = dll.NewProc("GdipStringFormatGetGenericDefault")
	gdipStringFormatGetGenericTypographic = dll.NewProc("GdipStringFormatGetGenericTypographic")
	gdipDeleteStringFormat = dll.NewProc("GdipDeleteStringFormat")
	gdipCloneStringFormat = dll.NewProc("GdipCloneStringFormat")
	gdipSetStringFormatFlags = dll.NewProc("GdipSetStringFormatFlags")
	gdipGetStringFormatFlags = dll.NewProc("GdipGetStringFormatFlags")
	gdipSetStringFormatAlign = dll.NewProc("GdipSetStringFormatAlign")
	gdipGetStringFormatAlign = dll.NewProc("GdipGetStringFormatAlign")
	gdipSetStringFormatLineAlign = dll.NewProc("GdipSetStringFormatLineAlign")
	gdipGetStringFormatLineAlign = dll.NewProc("GdipGetStringFormatLineAlign")
	gdipSetStringFormatTrimming = dll.NewProc("GdipSetStringFormatTrimming")
	gdipGetStringFormatTrimming = dll.NewProc("GdipGetStringFormatTrimming")
	gdipSetStringFormatHotkeyPrefix = dll.NewProc("GdipSetStringFormatHotkeyPrefix")
	gdipGetStringFormatHotkeyPrefix = dll.NewProc("GdipGetStringFormatHotkeyPrefix")
	gdipSetStringFormatTabStops = dll.NewProc("GdipSetStringFormatTabStops")
	gdipGetStringFormatTabStops = dll.NewProc("GdipGetStringFormatTabStops")
	gdipGetStringFormatTabStopCount = dll.NewProc("GdipGetStringFormatTabStopCount")
	gdipSetStringFormatDigitSubstitution = dll.NewProc("GdipSetStringFormatDigitSubstitution")
	gdipGetStringFormatDigitSubstitution = dll.NewProc("GdipGetStringFormatDigitSubstitution")
	gdipGetStringFormatMeasurableCharacterRangeCount = dll.NewProc("GdipGetStringFormatMeasurableCharacterRangeCount")
	gdipSetStringFormatMeasurableCharacterRanges = dll.NewProc("GdipSetStringFormatMeasurableCharacterRanges")

	// text
	gdipDrawString = dll.NewProc("GdipDrawString")
	gdipMeasureString = dll.NewProc("GdipMeasureString")
	gdipMeasureCharacterRanges = dll.NewProc("GdipMeasureCharacterRanges")
	gdipDrawDriverString = dll.NewProc("GdipDrawDriverString")
	gdipMeasureDriverString = dll.NewProc("GdipMeasureDriverString")
}

	// stringformat
func CreateStringFormat(formatAttributes int32, language win32.LANGID, format **StringFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateStringFormat.Addr(), 
		uintptr(formatAttributes),
		uintptr(language),
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func StringFormatGetGenericDefault(format **StringFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipStringFormatGetGenericDefault.Addr(), 
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func StringFormatGetGenericTypographic(format **StringFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipStringFormatGetGenericTypographic.Addr(), 
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func DeleteStringFormat(format *StringFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteStringFormat.Addr(), 
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func CloneStringFormat(format *StringFormat, newFormat **StringFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneStringFormat.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(newFormat)))
	return Status(ret)
}

func SetStringFormatFlags(format *StringFormat, flags int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatFlags.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(flags))
	return Status(ret)
}

func GetStringFormatFlags(format *StringFormat, flags *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatFlags.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(flags)))
	return Status(ret)
}

func SetStringFormatAlign(format *StringFormat, align StringAlignment) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatAlign.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(align))
	return Status(ret)
}

func GetStringFormatAlign(format *StringFormat, align *StringAlignment) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatAlign.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(align)))
	return Status(ret)
}

func SetStringFormatLineAlign(format *StringFormat, align StringAlignment) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatLineAlign.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(align))
	return Status(ret)
}

func GetStringFormatLineAlign(format *StringFormat, align *StringAlignment) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatLineAlign.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(align)))
	return Status(ret)
}

func SetStringFormatTrimming(format *StringFormat, trimming StringTrimming) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatTrimming.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(trimming))
	return Status(ret)
}

func GetStringFormatTrimming(format *StringFormat, trimming *StringTrimming) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatTrimming.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(trimming)))
	return Status(ret)
}

func SetStringFormatHotkeyPrefix(format *StringFormat, hotkeyPrefix int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatHotkeyPrefix.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(hotkeyPrefix))
	return Status(ret)
}

func GetStringFormatHotkeyPrefix(format *StringFormat, hotkeyPrefix *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatHotkeyPrefix.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(hotkeyPrefix)))
	return Status(ret)
}

func SetStringFormatTabStops(format *StringFormat, firstTabOffset float32, count int32, tabStops *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatTabStops.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(math.Float32bits(firstTabOffset)),
		uintptr(count),
		uintptr(unsafe.Pointer(tabStops)))
	return Status(ret)
}

func GetStringFormatTabStops(format *StringFormat, count int32, firstTabOffset *float32, tabStops *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatTabStops.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(count),
		uintptr(unsafe.Pointer(firstTabOffset)),
		uintptr(unsafe.Pointer(tabStops)))
	return Status(ret)
}

func GetStringFormatTabStopCount(format *StringFormat, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatTabStopCount.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func SetStringFormatDigitSubstitution(format *StringFormat, language win32.LANGID, substitute StringDigitSubstitute) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatDigitSubstitution.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(language),
		uintptr(substitute))
	return Status(ret)
}

func GetStringFormatDigitSubstitution(format *StringFormat, language *win32.LANGID, substitute *StringDigitSubstitute) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatDigitSubstitution.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(language)),
		uintptr(unsafe.Pointer(substitute)))
	return Status(ret)
}

func GetStringFormatMeasurableCharacterRangeCount(format *StringFormat, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetStringFormatMeasurableCharacterRangeCount.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func SetStringFormatMeasurableCharacterRanges(format *StringFormat, rangeCount int32, ranges *CharacterRange) Status {
	ret, _, _ := syscall.SyscallN(gdipSetStringFormatMeasurableCharacterRanges.Addr(), 
		uintptr(unsafe.Pointer(format)),
		uintptr(rangeCount),
		uintptr(unsafe.Pointer(ranges)))
	return Status(ret)
}


	// text
func DrawString(graphics *Graphics, string *uint16, length int32, font *Font, layoutRect *RectF, stringFormat *StringFormat, brush *Brush) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawString.Addr(), 
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(string)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(stringFormat)),
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func MeasureString(graphics *Graphics, string *uint16, length int32, font *Font, layoutRect *RectF, stringFormat *StringFormat, boundingBox *RectF, codepointsFitted *int32, linesFilled *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipMeasureString.Addr(), 
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(string)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(stringFormat)),
		uintptr(unsafe.Pointer(boundingBox)),
		uintptr(unsafe.Pointer(codepointsFitted)),
		uintptr(unsafe.Pointer(linesFilled)))
	return Status(ret)
}

func MeasureCharacterRanges(graphics *Graphics, string *uint16, length int32, font *Font, layoutRect *RectF, stringFormat *StringFormat, regionCount int32, regions **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipMeasureCharacterRanges.Addr(), 
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(string)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(stringFormat)),
		uintptr(regionCount),
		uintptr(unsafe.Pointer(regions)))
	return Status(ret)
}

func DrawDriverString(graphics *Graphics, text *win32.UINT16, length int32, font *Font, brush *Brush, positions *PointF, flags int32, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawDriverString.Addr(), 
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(text)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(flags),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func MeasureDriverString(graphics *Graphics, text *win32.UINT16, length int32, font *Font, positions *PointF, flags int32, matrix *Matrix, boundingBox *RectF) Status {
	ret, _, _ := syscall.SyscallN(gdipMeasureDriverString.Addr(), 
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(text)),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(flags),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(boundingBox)))
	return Status(ret)
}

