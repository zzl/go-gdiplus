package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// region
	gdipCreateRegion	*windows.LazyProc
	gdipCreateRegionRect	*windows.LazyProc
	gdipCreateRegionRectI	*windows.LazyProc
	gdipCreateRegionPath	*windows.LazyProc
	gdipCreateRegionRgnData	*windows.LazyProc
	gdipCreateRegionHrgn	*windows.LazyProc
	gdipCloneRegion	*windows.LazyProc
	gdipDeleteRegion	*windows.LazyProc
	gdipSetInfinite	*windows.LazyProc
	gdipSetEmpty	*windows.LazyProc
	gdipCombineRegionRect	*windows.LazyProc
	gdipCombineRegionRectI	*windows.LazyProc
	gdipCombineRegionPath	*windows.LazyProc
	gdipCombineRegionRegion	*windows.LazyProc
	gdipTranslateRegion	*windows.LazyProc
	gdipTranslateRegionI	*windows.LazyProc
	gdipTransformRegion	*windows.LazyProc
	gdipGetRegionBounds	*windows.LazyProc
	gdipGetRegionBoundsI	*windows.LazyProc
	gdipGetRegionHRgn	*windows.LazyProc
	gdipIsEmptyRegion	*windows.LazyProc
	gdipIsInfiniteRegion	*windows.LazyProc
	gdipIsEqualRegion	*windows.LazyProc
	gdipGetRegionDataSize	*windows.LazyProc
	gdipGetRegionData	*windows.LazyProc
	gdipIsVisibleRegionPoint	*windows.LazyProc
	gdipIsVisibleRegionPointI	*windows.LazyProc
	gdipIsVisibleRegionRect	*windows.LazyProc
	gdipIsVisibleRegionRectI	*windows.LazyProc
	gdipGetRegionScansCount	*windows.LazyProc
	gdipGetRegionScans	*windows.LazyProc
	gdipGetRegionScansI	*windows.LazyProc
)

func init() {

	// region
	gdipCreateRegion = dll.NewProc("GdipCreateRegion")
	gdipCreateRegionRect = dll.NewProc("GdipCreateRegionRect")
	gdipCreateRegionRectI = dll.NewProc("GdipCreateRegionRectI")
	gdipCreateRegionPath = dll.NewProc("GdipCreateRegionPath")
	gdipCreateRegionRgnData = dll.NewProc("GdipCreateRegionRgnData")
	gdipCreateRegionHrgn = dll.NewProc("GdipCreateRegionHrgn")
	gdipCloneRegion = dll.NewProc("GdipCloneRegion")
	gdipDeleteRegion = dll.NewProc("GdipDeleteRegion")
	gdipSetInfinite = dll.NewProc("GdipSetInfinite")
	gdipSetEmpty = dll.NewProc("GdipSetEmpty")
	gdipCombineRegionRect = dll.NewProc("GdipCombineRegionRect")
	gdipCombineRegionRectI = dll.NewProc("GdipCombineRegionRectI")
	gdipCombineRegionPath = dll.NewProc("GdipCombineRegionPath")
	gdipCombineRegionRegion = dll.NewProc("GdipCombineRegionRegion")
	gdipTranslateRegion = dll.NewProc("GdipTranslateRegion")
	gdipTranslateRegionI = dll.NewProc("GdipTranslateRegionI")
	gdipTransformRegion = dll.NewProc("GdipTransformRegion")
	gdipGetRegionBounds = dll.NewProc("GdipGetRegionBounds")
	gdipGetRegionBoundsI = dll.NewProc("GdipGetRegionBoundsI")
	gdipGetRegionHRgn = dll.NewProc("GdipGetRegionHRgn")
	gdipIsEmptyRegion = dll.NewProc("GdipIsEmptyRegion")
	gdipIsInfiniteRegion = dll.NewProc("GdipIsInfiniteRegion")
	gdipIsEqualRegion = dll.NewProc("GdipIsEqualRegion")
	gdipGetRegionDataSize = dll.NewProc("GdipGetRegionDataSize")
	gdipGetRegionData = dll.NewProc("GdipGetRegionData")
	gdipIsVisibleRegionPoint = dll.NewProc("GdipIsVisibleRegionPoint")
	gdipIsVisibleRegionPointI = dll.NewProc("GdipIsVisibleRegionPointI")
	gdipIsVisibleRegionRect = dll.NewProc("GdipIsVisibleRegionRect")
	gdipIsVisibleRegionRectI = dll.NewProc("GdipIsVisibleRegionRectI")
	gdipGetRegionScansCount = dll.NewProc("GdipGetRegionScansCount")
	gdipGetRegionScans = dll.NewProc("GdipGetRegionScans")
	gdipGetRegionScansI = dll.NewProc("GdipGetRegionScansI")
}

	// region
func CreateRegion(region **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateRegion.Addr(), 
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func CreateRegionRect(rect *RectF, region **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateRegionRect.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func CreateRegionRectI(rect *Rect, region **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateRegionRectI.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func CreateRegionPath(path *Path, region **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateRegionPath.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func CreateRegionRgnData(regionData *win32.BYTE, size int32, region **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateRegionRgnData.Addr(), 
		uintptr(unsafe.Pointer(regionData)),
		uintptr(size),
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func CreateRegionHrgn(hRgn win32.HRGN, region **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateRegionHrgn.Addr(), 
		uintptr(hRgn),
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func CloneRegion(region *Region, cloneRegion **Region) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneRegion.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(cloneRegion)))
	return Status(ret)
}

func DeleteRegion(region *Region) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteRegion.Addr(), 
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func SetInfinite(region *Region) Status {
	ret, _, _ := syscall.SyscallN(gdipSetInfinite.Addr(), 
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func SetEmpty(region *Region) Status {
	ret, _, _ := syscall.SyscallN(gdipSetEmpty.Addr(), 
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func CombineRegionRect(region *Region, rect *RectF, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipCombineRegionRect.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(combineMode))
	return Status(ret)
}

func CombineRegionRectI(region *Region, rect *Rect, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipCombineRegionRectI.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(combineMode))
	return Status(ret)
}

func CombineRegionPath(region *Region, path *Path, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipCombineRegionPath.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(path)),
		uintptr(combineMode))
	return Status(ret)
}

func CombineRegionRegion(region *Region, region2 *Region, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipCombineRegionRegion.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(region2)),
		uintptr(combineMode))
	return Status(ret)
}

func TranslateRegion(region *Region, dx float32, dy float32) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateRegion.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)))
	return Status(ret)
}

func TranslateRegionI(region *Region, dx int32, dy int32) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateRegionI.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(dx),
		uintptr(dy))
	return Status(ret)
}

func TransformRegion(region *Region, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipTransformRegion.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func GetRegionBounds(region *Region, graphics *Graphics, rect *RectF) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionBounds.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func GetRegionBoundsI(region *Region, graphics *Graphics, rect *Rect) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionBoundsI.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func GetRegionHRgn(region *Region, graphics *Graphics, hRgn *win32.HRGN) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionHRgn.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(hRgn)))
	return Status(ret)
}

func IsEmptyRegion(region *Region, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsEmptyRegion.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsInfiniteRegion(region *Region, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsInfiniteRegion.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsEqualRegion(region *Region, region2 *Region, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsEqualRegion.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(region2)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func GetRegionDataSize(region *Region, bufferSize *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionDataSize.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(bufferSize)))
	return Status(ret)
}

func GetRegionData(region *Region, buffer *win32.BYTE, bufferSize uint32, sizeFilled *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionData.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(bufferSize),
		uintptr(unsafe.Pointer(sizeFilled)))
	return Status(ret)
}

func IsVisibleRegionPoint(region *Region, x float32, y float32, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisibleRegionPoint.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisibleRegionPointI(region *Region, x int32, y int32, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisibleRegionPointI.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisibleRegionRect(region *Region, x float32, y float32, width float32, height float32, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisibleRegionRect.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisibleRegionRectI(region *Region, x int32, y int32, width int32, height int32, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisibleRegionRectI.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func GetRegionScansCount(region *Region, count *uint32, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionScansCount.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(count)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func GetRegionScans(region *Region, rects *RectF, count *int32, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionScans.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(unsafe.Pointer(count)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func GetRegionScansI(region *Region, rects *Rect, count *int32, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRegionScansI.Addr(), 
		uintptr(unsafe.Pointer(region)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(unsafe.Pointer(count)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

