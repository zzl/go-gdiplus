package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// font
	gdipCreateFontFromDC             *windows.LazyProc
	gdipCreateFontFromLogfontA       *windows.LazyProc
	gdipCreateFontFromLogfontW       *windows.LazyProc
	gdipCloneFont                    *windows.LazyProc
	gdipCreateFont                   *windows.LazyProc
	gdipDeleteFont                   *windows.LazyProc
	gdipGetFamily                    *windows.LazyProc
	gdipGetFontStyle                 *windows.LazyProc
	gdipGetFontSize                  *windows.LazyProc
	gdipGetFontUnit                  *windows.LazyProc
	gdipGetFontHeight                *windows.LazyProc
	gdipGetFontHeightGivenDPI        *windows.LazyProc
	gdipGetLogFontA                  *windows.LazyProc
	gdipGetLogFontW                  *windows.LazyProc
	gdipNewInstalledFontCollection   *windows.LazyProc
	gdipNewPrivateFontCollection     *windows.LazyProc
	gdipDeletePrivateFontCollection  *windows.LazyProc
	gdipGetFontCollectionFamilyCount *windows.LazyProc
	gdipGetFontCollectionFamilyList  *windows.LazyProc
	gdipPrivateAddFontFile           *windows.LazyProc
	gdipPrivateAddMemoryFont         *windows.LazyProc

	// fontfamily
	gdipDeleteFontFamily              *windows.LazyProc
	gdipCloneFontFamily               *windows.LazyProc
	gdipCreateFontFamilyFromName      *windows.LazyProc
	gdipGetGenericFontFamilySansSerif *windows.LazyProc
	gdipGetGenericFontFamilySerif     *windows.LazyProc
	gdipGetGenericFontFamilyMonospace *windows.LazyProc
	gdipGetFamilyName                 *windows.LazyProc
	gdipIsStyleAvailable              *windows.LazyProc
	gdipFontCollectionEnumerable      *windows.LazyProc
	gdipFontCollectionEnumerate       *windows.LazyProc
	gdipGetEmHeight                   *windows.LazyProc
	gdipGetCellAscent                 *windows.LazyProc
	gdipGetCellDescent                *windows.LazyProc
	gdipGetLineSpacing                *windows.LazyProc
)

func init() {

	// font
	gdipCreateFontFromDC = dll.NewProc("GdipCreateFontFromDC")
	gdipCreateFontFromLogfontA = dll.NewProc("GdipCreateFontFromLogfontA")
	gdipCreateFontFromLogfontW = dll.NewProc("GdipCreateFontFromLogfontW")
	gdipCloneFont = dll.NewProc("GdipCloneFont")
	gdipCreateFont = dll.NewProc("GdipCreateFont")
	gdipDeleteFont = dll.NewProc("GdipDeleteFont")
	gdipGetFamily = dll.NewProc("GdipGetFamily")
	gdipGetFontStyle = dll.NewProc("GdipGetFontStyle")
	gdipGetFontSize = dll.NewProc("GdipGetFontSize")
	gdipGetFontUnit = dll.NewProc("GdipGetFontUnit")
	gdipGetFontHeight = dll.NewProc("GdipGetFontHeight")
	gdipGetFontHeightGivenDPI = dll.NewProc("GdipGetFontHeightGivenDPI")
	gdipGetLogFontA = dll.NewProc("GdipGetLogFontA")
	gdipGetLogFontW = dll.NewProc("GdipGetLogFontW")
	gdipNewInstalledFontCollection = dll.NewProc("GdipNewInstalledFontCollection")
	gdipNewPrivateFontCollection = dll.NewProc("GdipNewPrivateFontCollection")
	gdipDeletePrivateFontCollection = dll.NewProc("GdipDeletePrivateFontCollection")
	gdipGetFontCollectionFamilyCount = dll.NewProc("GdipGetFontCollectionFamilyCount")
	gdipGetFontCollectionFamilyList = dll.NewProc("GdipGetFontCollectionFamilyList")
	gdipPrivateAddFontFile = dll.NewProc("GdipPrivateAddFontFile")
	gdipPrivateAddMemoryFont = dll.NewProc("GdipPrivateAddMemoryFont")

	// fontfamily
	gdipDeleteFontFamily = dll.NewProc("GdipDeleteFontFamily")
	gdipCloneFontFamily = dll.NewProc("GdipCloneFontFamily")
	gdipCreateFontFamilyFromName = dll.NewProc("GdipCreateFontFamilyFromName")
	gdipGetGenericFontFamilySansSerif = dll.NewProc("GdipGetGenericFontFamilySansSerif")
	gdipGetGenericFontFamilySerif = dll.NewProc("GdipGetGenericFontFamilySerif")
	gdipGetGenericFontFamilyMonospace = dll.NewProc("GdipGetGenericFontFamilyMonospace")
	gdipGetFamilyName = dll.NewProc("GdipGetFamilyName")
	gdipIsStyleAvailable = dll.NewProc("GdipIsStyleAvailable")
	gdipFontCollectionEnumerable = dll.NewProc("GdipFontCollectionEnumerable")
	gdipFontCollectionEnumerate = dll.NewProc("GdipFontCollectionEnumerate")
	gdipGetEmHeight = dll.NewProc("GdipGetEmHeight")
	gdipGetCellAscent = dll.NewProc("GdipGetCellAscent")
	gdipGetCellDescent = dll.NewProc("GdipGetCellDescent")
	gdipGetLineSpacing = dll.NewProc("GdipGetLineSpacing")
}

// font
func CreateFontFromDC(hdc win32.HDC, font **Font) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFontFromDC.Addr(),
		uintptr(hdc),
		uintptr(unsafe.Pointer(font)))
	return Status(ret)
}

func CreateFontFromLogfontA(hdc win32.HDC, logfont *win32.LOGFONTA, font **Font) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFontFromLogfontA.Addr(),
		uintptr(hdc),
		uintptr(unsafe.Pointer(logfont)),
		uintptr(unsafe.Pointer(font)))
	return Status(ret)
}

func CreateFontFromLogfontW(hdc win32.HDC, logfont *win32.LOGFONT, font **Font) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFontFromLogfontW.Addr(),
		uintptr(hdc),
		uintptr(unsafe.Pointer(logfont)),
		uintptr(unsafe.Pointer(font)))
	return Status(ret)
}

func CloneFont(font *Font, cloneFont **Font) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneFont.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(cloneFont)))
	return Status(ret)
}

func CreateFont(fontFamily *FontFamily, emSize float32, style FontStyle, unit Unit, font **Font) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFont.Addr(),
		uintptr(unsafe.Pointer(fontFamily)),
		uintptr(math.Float32bits(emSize)),
		uintptr(style),
		uintptr(unit),
		uintptr(unsafe.Pointer(font)))
	return Status(ret)
}

func DeleteFont(font *Font) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteFont.Addr(),
		uintptr(unsafe.Pointer(font)))
	return Status(ret)
}

func GetFamily(font *Font, family **FontFamily) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFamily.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(family)))
	return Status(ret)
}

func GetFontStyle(font *Font, style *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFontStyle.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(style)))
	return Status(ret)
}

func GetFontSize(font *Font, size *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFontSize.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(size)))
	return Status(ret)
}

func GetFontUnit(font *Font, unit *Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFontUnit.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(unit)))
	return Status(ret)
}

func GetFontHeight(font *Font, graphics *Graphics, height *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFontHeight.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(height)))
	return Status(ret)
}

func GetFontHeightGivenDPI(font *Font, dpi float32, height *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFontHeightGivenDPI.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(math.Float32bits(dpi)),
		uintptr(unsafe.Pointer(height)))
	return Status(ret)
}

func GetLogFontA(font *Font, graphics *Graphics, logfontA *win32.LOGFONTA) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLogFontA.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(logfontA)))
	return Status(ret)
}

func GetLogFontW(font *Font, graphics *Graphics, logfontW *win32.LOGFONT) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLogFontW.Addr(),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(logfontW)))
	return Status(ret)
}

func NewInstalledFontCollection(fontCollection **FontCollection) Status {
	ret, _, _ := syscall.SyscallN(gdipNewInstalledFontCollection.Addr(),
		uintptr(unsafe.Pointer(fontCollection)))
	return Status(ret)
}

func NewPrivateFontCollection(fontCollection **FontCollection) Status {
	ret, _, _ := syscall.SyscallN(gdipNewPrivateFontCollection.Addr(),
		uintptr(unsafe.Pointer(fontCollection)))
	return Status(ret)
}

func DeletePrivateFontCollection(fontCollection **FontCollection) Status {
	ret, _, _ := syscall.SyscallN(gdipDeletePrivateFontCollection.Addr(),
		uintptr(unsafe.Pointer(fontCollection)))
	return Status(ret)
}

func GetFontCollectionFamilyCount(fontCollection *FontCollection, numFound *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFontCollectionFamilyCount.Addr(),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(numFound)))
	return Status(ret)
}

func GetFontCollectionFamilyList(fontCollection *FontCollection, numSought int32, gpfamilies []*FontFamily, numFound *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFontCollectionFamilyList.Addr(),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(numSought),
		uintptr(unsafe.Pointer(&gpfamilies[0])),
		uintptr(unsafe.Pointer(numFound)))
	return Status(ret)
}

func PrivateAddFontFile(fontCollection *FontCollection, filename *uint16) Status {
	ret, _, _ := syscall.SyscallN(gdipPrivateAddFontFile.Addr(),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(filename)))
	return Status(ret)
}

func PrivateAddMemoryFont(fontCollection *FontCollection, memory *byte, length int32) Status {
	ret, _, _ := syscall.SyscallN(gdipPrivateAddMemoryFont.Addr(),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(memory)),
		uintptr(length))
	return Status(ret)
}

// fontfamily
func DeleteFontFamily(fontFamily *FontFamily) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteFontFamily.Addr(),
		uintptr(unsafe.Pointer(fontFamily)))
	return Status(ret)
}

func CloneFontFamily(fontFamily *FontFamily, clonedFontFamily **FontFamily) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneFontFamily.Addr(),
		uintptr(unsafe.Pointer(fontFamily)),
		uintptr(unsafe.Pointer(clonedFontFamily)))
	return Status(ret)
}

func CreateFontFamilyFromName(name *uint16, fontCollection *FontCollection, FontFamily **FontFamily) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFontFamilyFromName.Addr(),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(FontFamily)))
	return Status(ret)
}

func GetGenericFontFamilySansSerif(nativeFamily **FontFamily) Status {
	ret, _, _ := syscall.SyscallN(gdipGetGenericFontFamilySansSerif.Addr(),
		uintptr(unsafe.Pointer(nativeFamily)))
	return Status(ret)
}

func GetGenericFontFamilySerif(nativeFamily **FontFamily) Status {
	ret, _, _ := syscall.SyscallN(gdipGetGenericFontFamilySerif.Addr(),
		uintptr(unsafe.Pointer(nativeFamily)))
	return Status(ret)
}

func GetGenericFontFamilyMonospace(nativeFamily **FontFamily) Status {
	ret, _, _ := syscall.SyscallN(gdipGetGenericFontFamilyMonospace.Addr(),
		uintptr(unsafe.Pointer(nativeFamily)))
	return Status(ret)
}

func GetFamilyName(family *FontFamily, name []uint16, language win32.LANGID) Status {
	ret, _, _ := syscall.SyscallN(gdipGetFamilyName.Addr(),
		uintptr(unsafe.Pointer(family)),
		uintptr(unsafe.Pointer(&name[0])),
		uintptr(language))
	return Status(ret)
}

func IsStyleAvailable(family *FontFamily, style int32, IsStyleAvailable *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsStyleAvailable.Addr(),
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(IsStyleAvailable)))
	return Status(ret)
}

func FontCollectionEnumerable(fontCollection *FontCollection, graphics *Graphics, numFound *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFontCollectionEnumerable.Addr(),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(numFound)))
	return Status(ret)
}

func FontCollectionEnumerate(fontCollection *FontCollection, numSought int32, gpfamilies []*FontFamily, numFound *int32, graphics *Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipFontCollectionEnumerate.Addr(),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(numSought),
		uintptr(unsafe.Pointer(&gpfamilies[0])),
		uintptr(unsafe.Pointer(numFound)),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func GetEmHeight(family *FontFamily, style int32, EmHeight *win32.UINT16) Status {
	ret, _, _ := syscall.SyscallN(gdipGetEmHeight.Addr(),
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(EmHeight)))
	return Status(ret)
}

func GetCellAscent(family *FontFamily, style int32, CellAscent *win32.UINT16) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCellAscent.Addr(),
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(CellAscent)))
	return Status(ret)
}

func GetCellDescent(family *FontFamily, style int32, CellDescent *win32.UINT16) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCellDescent.Addr(),
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(CellDescent)))
	return Status(ret)
}

func GetLineSpacing(family *FontFamily, style int32, LineSpacing *win32.UINT16) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineSpacing.Addr(),
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(unsafe.Pointer(LineSpacing)))
	return Status(ret)
}
