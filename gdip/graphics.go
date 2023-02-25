package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// graphics
	gdipFlush                                  *windows.LazyProc
	gdipCreateFromHDC                          *windows.LazyProc
	gdipCreateFromHDC2                         *windows.LazyProc
	gdipCreateFromHWND                         *windows.LazyProc
	gdipCreateFromHWNDICM                      *windows.LazyProc
	gdipDeleteGraphics                         *windows.LazyProc
	gdipGetDC                                  *windows.LazyProc
	gdipReleaseDC                              *windows.LazyProc
	gdipSetCompositingMode                     *windows.LazyProc
	gdipGetCompositingMode                     *windows.LazyProc
	gdipSetRenderingOrigin                     *windows.LazyProc
	gdipGetRenderingOrigin                     *windows.LazyProc
	gdipSetCompositingQuality                  *windows.LazyProc
	gdipGetCompositingQuality                  *windows.LazyProc
	gdipSetSmoothingMode                       *windows.LazyProc
	gdipGetSmoothingMode                       *windows.LazyProc
	gdipSetPixelOffsetMode                     *windows.LazyProc
	gdipGetPixelOffsetMode                     *windows.LazyProc
	gdipSetTextRenderingHint                   *windows.LazyProc
	gdipGetTextRenderingHint                   *windows.LazyProc
	gdipSetTextContrast                        *windows.LazyProc
	gdipGetTextContrast                        *windows.LazyProc
	gdipSetInterpolationMode                   *windows.LazyProc
	gdipGraphicsSetAbort                       *windows.LazyProc
	gdipGetInterpolationMode                   *windows.LazyProc
	gdipSetWorldTransform                      *windows.LazyProc
	gdipResetWorldTransform                    *windows.LazyProc
	gdipMultiplyWorldTransform                 *windows.LazyProc
	gdipTranslateWorldTransform                *windows.LazyProc
	gdipScaleWorldTransform                    *windows.LazyProc
	gdipRotateWorldTransform                   *windows.LazyProc
	gdipGetWorldTransform                      *windows.LazyProc
	gdipResetPageTransform                     *windows.LazyProc
	gdipGetPageUnit                            *windows.LazyProc
	gdipGetPageScale                           *windows.LazyProc
	gdipSetPageUnit                            *windows.LazyProc
	gdipSetPageScale                           *windows.LazyProc
	gdipGetDpiX                                *windows.LazyProc
	gdipGetDpiY                                *windows.LazyProc
	gdipTransformPoints                        *windows.LazyProc
	gdipTransformPointsI                       *windows.LazyProc
	gdipGetNearestColor                        *windows.LazyProc
	gdipCreateHalftonePalette                  *windows.LazyProc
	gdipDrawLine                               *windows.LazyProc
	gdipDrawLineI                              *windows.LazyProc
	gdipDrawLines                              *windows.LazyProc
	gdipDrawLinesI                             *windows.LazyProc
	gdipDrawArc                                *windows.LazyProc
	gdipDrawArcI                               *windows.LazyProc
	gdipDrawBezier                             *windows.LazyProc
	gdipDrawBezierI                            *windows.LazyProc
	gdipDrawBeziers                            *windows.LazyProc
	gdipDrawBeziersI                           *windows.LazyProc
	gdipDrawRectangle                          *windows.LazyProc
	gdipDrawRectangleI                         *windows.LazyProc
	gdipDrawRectangles                         *windows.LazyProc
	gdipDrawRectanglesI                        *windows.LazyProc
	gdipDrawEllipse                            *windows.LazyProc
	gdipDrawEllipseI                           *windows.LazyProc
	gdipDrawPie                                *windows.LazyProc
	gdipDrawPieI                               *windows.LazyProc
	gdipDrawPolygon                            *windows.LazyProc
	gdipDrawPolygonI                           *windows.LazyProc
	gdipDrawPath                               *windows.LazyProc
	gdipDrawCurve                              *windows.LazyProc
	gdipDrawCurveI                             *windows.LazyProc
	gdipDrawCurve2                             *windows.LazyProc
	gdipDrawCurve2I                            *windows.LazyProc
	gdipDrawCurve3                             *windows.LazyProc
	gdipDrawCurve3I                            *windows.LazyProc
	gdipDrawClosedCurve                        *windows.LazyProc
	gdipDrawClosedCurveI                       *windows.LazyProc
	gdipDrawClosedCurve2                       *windows.LazyProc
	gdipDrawClosedCurve2I                      *windows.LazyProc
	gdipGraphicsClear                          *windows.LazyProc
	gdipFillRectangle                          *windows.LazyProc
	gdipFillRectangleI                         *windows.LazyProc
	gdipFillRectangles                         *windows.LazyProc
	gdipFillRectanglesI                        *windows.LazyProc
	gdipFillPolygon                            *windows.LazyProc
	gdipFillPolygonI                           *windows.LazyProc
	gdipFillPolygon2                           *windows.LazyProc
	gdipFillPolygon2I                          *windows.LazyProc
	gdipFillEllipse                            *windows.LazyProc
	gdipFillEllipseI                           *windows.LazyProc
	gdipFillPie                                *windows.LazyProc
	gdipFillPieI                               *windows.LazyProc
	gdipFillPath                               *windows.LazyProc
	gdipFillClosedCurve                        *windows.LazyProc
	gdipFillClosedCurveI                       *windows.LazyProc
	gdipFillClosedCurve2                       *windows.LazyProc
	gdipFillClosedCurve2I                      *windows.LazyProc
	gdipFillRegion                             *windows.LazyProc
	gdipDrawImage                              *windows.LazyProc
	gdipDrawImageI                             *windows.LazyProc
	gdipDrawImageRect                          *windows.LazyProc
	gdipDrawImageRectI                         *windows.LazyProc
	gdipDrawImagePoints                        *windows.LazyProc
	gdipDrawImagePointsI                       *windows.LazyProc
	gdipDrawImagePointRect                     *windows.LazyProc
	gdipDrawImagePointRectI                    *windows.LazyProc
	gdipDrawImageRectRect                      *windows.LazyProc
	gdipDrawImageRectRectI                     *windows.LazyProc
	gdipDrawImagePointsRect                    *windows.LazyProc
	gdipDrawImagePointsRectI                   *windows.LazyProc
	gdipDrawImageFX                            *windows.LazyProc
	gdipEnumerateMetafileDestPoint             *windows.LazyProc
	gdipEnumerateMetafileDestPointI            *windows.LazyProc
	gdipEnumerateMetafileDestRect              *windows.LazyProc
	gdipEnumerateMetafileDestRectI             *windows.LazyProc
	gdipEnumerateMetafileDestPoints            *windows.LazyProc
	gdipEnumerateMetafileDestPointsI           *windows.LazyProc
	gdipEnumerateMetafileSrcRectDestPoint      *windows.LazyProc
	gdipEnumerateMetafileSrcRectDestPointI     *windows.LazyProc
	gdipEnumerateMetafileSrcRectDestRect       *windows.LazyProc
	gdipEnumerateMetafileSrcRectDestRectI      *windows.LazyProc
	gdipEnumerateMetafileSrcRectDestPoints     *windows.LazyProc
	gdipEnumerateMetafileSrcRectDestPointsI    *windows.LazyProc
	gdipPlayMetafileRecord                     *windows.LazyProc
	gdipSetClipGraphics                        *windows.LazyProc
	gdipSetClipRect                            *windows.LazyProc
	gdipSetClipRectI                           *windows.LazyProc
	gdipSetClipPath                            *windows.LazyProc
	gdipSetClipRegion                          *windows.LazyProc
	gdipSetClipHrgn                            *windows.LazyProc
	gdipResetClip                              *windows.LazyProc
	gdipTranslateClip                          *windows.LazyProc
	gdipTranslateClipI                         *windows.LazyProc
	gdipGetClip                                *windows.LazyProc
	gdipGetClipBounds                          *windows.LazyProc
	gdipGetClipBoundsI                         *windows.LazyProc
	gdipIsClipEmpty                            *windows.LazyProc
	gdipGetVisibleClipBounds                   *windows.LazyProc
	gdipGetVisibleClipBoundsI                  *windows.LazyProc
	gdipIsVisibleClipEmpty                     *windows.LazyProc
	gdipIsVisiblePoint                         *windows.LazyProc
	gdipIsVisiblePointI                        *windows.LazyProc
	gdipIsVisibleRect                          *windows.LazyProc
	gdipIsVisibleRectI                         *windows.LazyProc
	gdipSaveGraphics                           *windows.LazyProc
	gdipRestoreGraphics                        *windows.LazyProc
	gdipBeginContainer                         *windows.LazyProc
	gdipBeginContainerI                        *windows.LazyProc
	gdipBeginContainer2                        *windows.LazyProc
	gdipEndContainer                           *windows.LazyProc
	gdipGetMetafileHeaderFromEmf               *windows.LazyProc
	gdipGetMetafileHeaderFromFile              *windows.LazyProc
	gdipGetMetafileHeaderFromStream            *windows.LazyProc
	gdipGetMetafileHeaderFromMetafile          *windows.LazyProc
	gdipGetHemfFromMetafile                    *windows.LazyProc
	gdipCreateStreamOnFile                     *windows.LazyProc
	gdipCreateMetafileFromWmf                  *windows.LazyProc
	gdipCreateMetafileFromEmf                  *windows.LazyProc
	gdipCreateMetafileFromFile                 *windows.LazyProc
	gdipCreateMetafileFromWmfFile              *windows.LazyProc
	gdipCreateMetafileFromStream               *windows.LazyProc
	gdipRecordMetafile                         *windows.LazyProc
	gdipRecordMetafileI                        *windows.LazyProc
	gdipRecordMetafileFileName                 *windows.LazyProc
	gdipRecordMetafileFileNameI                *windows.LazyProc
	gdipRecordMetafileStream                   *windows.LazyProc
	gdipRecordMetafileStreamI                  *windows.LazyProc
	gdipSetMetafileDownLevelRasterizationLimit *windows.LazyProc
	gdipGetMetafileDownLevelRasterizationLimit *windows.LazyProc
	gdipGetImageDecodersSize                   *windows.LazyProc
	gdipGetImageDecoders                       *windows.LazyProc
	gdipGetImageEncodersSize                   *windows.LazyProc
	gdipGetImageEncoders                       *windows.LazyProc
	gdipComment                                *windows.LazyProc
)

func init() {

	// graphics
	gdipFlush = dll.NewProc("GdipFlush")
	gdipCreateFromHDC = dll.NewProc("GdipCreateFromHDC")
	gdipCreateFromHDC2 = dll.NewProc("GdipCreateFromHDC2")
	gdipCreateFromHWND = dll.NewProc("GdipCreateFromHWND")
	gdipCreateFromHWNDICM = dll.NewProc("GdipCreateFromHWNDICM")
	gdipDeleteGraphics = dll.NewProc("GdipDeleteGraphics")
	gdipGetDC = dll.NewProc("GdipGetDC")
	gdipReleaseDC = dll.NewProc("GdipReleaseDC")
	gdipSetCompositingMode = dll.NewProc("GdipSetCompositingMode")
	gdipGetCompositingMode = dll.NewProc("GdipGetCompositingMode")
	gdipSetRenderingOrigin = dll.NewProc("GdipSetRenderingOrigin")
	gdipGetRenderingOrigin = dll.NewProc("GdipGetRenderingOrigin")
	gdipSetCompositingQuality = dll.NewProc("GdipSetCompositingQuality")
	gdipGetCompositingQuality = dll.NewProc("GdipGetCompositingQuality")
	gdipSetSmoothingMode = dll.NewProc("GdipSetSmoothingMode")
	gdipGetSmoothingMode = dll.NewProc("GdipGetSmoothingMode")
	gdipSetPixelOffsetMode = dll.NewProc("GdipSetPixelOffsetMode")
	gdipGetPixelOffsetMode = dll.NewProc("GdipGetPixelOffsetMode")
	gdipSetTextRenderingHint = dll.NewProc("GdipSetTextRenderingHint")
	gdipGetTextRenderingHint = dll.NewProc("GdipGetTextRenderingHint")
	gdipSetTextContrast = dll.NewProc("GdipSetTextContrast")
	gdipGetTextContrast = dll.NewProc("GdipGetTextContrast")
	gdipSetInterpolationMode = dll.NewProc("GdipSetInterpolationMode")
	gdipGraphicsSetAbort = dll.NewProc("GdipGraphicsSetAbort")
	gdipGetInterpolationMode = dll.NewProc("GdipGetInterpolationMode")
	gdipSetWorldTransform = dll.NewProc("GdipSetWorldTransform")
	gdipResetWorldTransform = dll.NewProc("GdipResetWorldTransform")
	gdipMultiplyWorldTransform = dll.NewProc("GdipMultiplyWorldTransform")
	gdipTranslateWorldTransform = dll.NewProc("GdipTranslateWorldTransform")
	gdipScaleWorldTransform = dll.NewProc("GdipScaleWorldTransform")
	gdipRotateWorldTransform = dll.NewProc("GdipRotateWorldTransform")
	gdipGetWorldTransform = dll.NewProc("GdipGetWorldTransform")
	gdipResetPageTransform = dll.NewProc("GdipResetPageTransform")
	gdipGetPageUnit = dll.NewProc("GdipGetPageUnit")
	gdipGetPageScale = dll.NewProc("GdipGetPageScale")
	gdipSetPageUnit = dll.NewProc("GdipSetPageUnit")
	gdipSetPageScale = dll.NewProc("GdipSetPageScale")
	gdipGetDpiX = dll.NewProc("GdipGetDpiX")
	gdipGetDpiY = dll.NewProc("GdipGetDpiY")
	gdipTransformPoints = dll.NewProc("GdipTransformPoints")
	gdipTransformPointsI = dll.NewProc("GdipTransformPointsI")
	gdipGetNearestColor = dll.NewProc("GdipGetNearestColor")
	gdipCreateHalftonePalette = dll.NewProc("GdipCreateHalftonePalette")
	gdipDrawLine = dll.NewProc("GdipDrawLine")
	gdipDrawLineI = dll.NewProc("GdipDrawLineI")
	gdipDrawLines = dll.NewProc("GdipDrawLines")
	gdipDrawLinesI = dll.NewProc("GdipDrawLinesI")
	gdipDrawArc = dll.NewProc("GdipDrawArc")
	gdipDrawArcI = dll.NewProc("GdipDrawArcI")
	gdipDrawBezier = dll.NewProc("GdipDrawBezier")
	gdipDrawBezierI = dll.NewProc("GdipDrawBezierI")
	gdipDrawBeziers = dll.NewProc("GdipDrawBeziers")
	gdipDrawBeziersI = dll.NewProc("GdipDrawBeziersI")
	gdipDrawRectangle = dll.NewProc("GdipDrawRectangle")
	gdipDrawRectangleI = dll.NewProc("GdipDrawRectangleI")
	gdipDrawRectangles = dll.NewProc("GdipDrawRectangles")
	gdipDrawRectanglesI = dll.NewProc("GdipDrawRectanglesI")
	gdipDrawEllipse = dll.NewProc("GdipDrawEllipse")
	gdipDrawEllipseI = dll.NewProc("GdipDrawEllipseI")
	gdipDrawPie = dll.NewProc("GdipDrawPie")
	gdipDrawPieI = dll.NewProc("GdipDrawPieI")
	gdipDrawPolygon = dll.NewProc("GdipDrawPolygon")
	gdipDrawPolygonI = dll.NewProc("GdipDrawPolygonI")
	gdipDrawPath = dll.NewProc("GdipDrawPath")
	gdipDrawCurve = dll.NewProc("GdipDrawCurve")
	gdipDrawCurveI = dll.NewProc("GdipDrawCurveI")
	gdipDrawCurve2 = dll.NewProc("GdipDrawCurve2")
	gdipDrawCurve2I = dll.NewProc("GdipDrawCurve2I")
	gdipDrawCurve3 = dll.NewProc("GdipDrawCurve3")
	gdipDrawCurve3I = dll.NewProc("GdipDrawCurve3I")
	gdipDrawClosedCurve = dll.NewProc("GdipDrawClosedCurve")
	gdipDrawClosedCurveI = dll.NewProc("GdipDrawClosedCurveI")
	gdipDrawClosedCurve2 = dll.NewProc("GdipDrawClosedCurve2")
	gdipDrawClosedCurve2I = dll.NewProc("GdipDrawClosedCurve2I")
	gdipGraphicsClear = dll.NewProc("GdipGraphicsClear")
	gdipFillRectangle = dll.NewProc("GdipFillRectangle")
	gdipFillRectangleI = dll.NewProc("GdipFillRectangleI")
	gdipFillRectangles = dll.NewProc("GdipFillRectangles")
	gdipFillRectanglesI = dll.NewProc("GdipFillRectanglesI")
	gdipFillPolygon = dll.NewProc("GdipFillPolygon")
	gdipFillPolygonI = dll.NewProc("GdipFillPolygonI")
	gdipFillPolygon2 = dll.NewProc("GdipFillPolygon2")
	gdipFillPolygon2I = dll.NewProc("GdipFillPolygon2I")
	gdipFillEllipse = dll.NewProc("GdipFillEllipse")
	gdipFillEllipseI = dll.NewProc("GdipFillEllipseI")
	gdipFillPie = dll.NewProc("GdipFillPie")
	gdipFillPieI = dll.NewProc("GdipFillPieI")
	gdipFillPath = dll.NewProc("GdipFillPath")
	gdipFillClosedCurve = dll.NewProc("GdipFillClosedCurve")
	gdipFillClosedCurveI = dll.NewProc("GdipFillClosedCurveI")
	gdipFillClosedCurve2 = dll.NewProc("GdipFillClosedCurve2")
	gdipFillClosedCurve2I = dll.NewProc("GdipFillClosedCurve2I")
	gdipFillRegion = dll.NewProc("GdipFillRegion")
	gdipDrawImage = dll.NewProc("GdipDrawImage")
	gdipDrawImageI = dll.NewProc("GdipDrawImageI")
	gdipDrawImageRect = dll.NewProc("GdipDrawImageRect")
	gdipDrawImageRectI = dll.NewProc("GdipDrawImageRectI")
	gdipDrawImagePoints = dll.NewProc("GdipDrawImagePoints")
	gdipDrawImagePointsI = dll.NewProc("GdipDrawImagePointsI")
	gdipDrawImagePointRect = dll.NewProc("GdipDrawImagePointRect")
	gdipDrawImagePointRectI = dll.NewProc("GdipDrawImagePointRectI")
	gdipDrawImageRectRect = dll.NewProc("GdipDrawImageRectRect")
	gdipDrawImageRectRectI = dll.NewProc("GdipDrawImageRectRectI")
	gdipDrawImagePointsRect = dll.NewProc("GdipDrawImagePointsRect")
	gdipDrawImagePointsRectI = dll.NewProc("GdipDrawImagePointsRectI")
	gdipDrawImageFX = dll.NewProc("GdipDrawImageFX")
	gdipEnumerateMetafileDestPoint = dll.NewProc("GdipEnumerateMetafileDestPoint")
	gdipEnumerateMetafileDestPointI = dll.NewProc("GdipEnumerateMetafileDestPointI")
	gdipEnumerateMetafileDestRect = dll.NewProc("GdipEnumerateMetafileDestRect")
	gdipEnumerateMetafileDestRectI = dll.NewProc("GdipEnumerateMetafileDestRectI")
	gdipEnumerateMetafileDestPoints = dll.NewProc("GdipEnumerateMetafileDestPoints")
	gdipEnumerateMetafileDestPointsI = dll.NewProc("GdipEnumerateMetafileDestPointsI")
	gdipEnumerateMetafileSrcRectDestPoint = dll.NewProc("GdipEnumerateMetafileSrcRectDestPoint")
	gdipEnumerateMetafileSrcRectDestPointI = dll.NewProc("GdipEnumerateMetafileSrcRectDestPointI")
	gdipEnumerateMetafileSrcRectDestRect = dll.NewProc("GdipEnumerateMetafileSrcRectDestRect")
	gdipEnumerateMetafileSrcRectDestRectI = dll.NewProc("GdipEnumerateMetafileSrcRectDestRectI")
	gdipEnumerateMetafileSrcRectDestPoints = dll.NewProc("GdipEnumerateMetafileSrcRectDestPoints")
	gdipEnumerateMetafileSrcRectDestPointsI = dll.NewProc("GdipEnumerateMetafileSrcRectDestPointsI")
	gdipPlayMetafileRecord = dll.NewProc("GdipPlayMetafileRecord")
	gdipSetClipGraphics = dll.NewProc("GdipSetClipGraphics")
	gdipSetClipRect = dll.NewProc("GdipSetClipRect")
	gdipSetClipRectI = dll.NewProc("GdipSetClipRectI")
	gdipSetClipPath = dll.NewProc("GdipSetClipPath")
	gdipSetClipRegion = dll.NewProc("GdipSetClipRegion")
	gdipSetClipHrgn = dll.NewProc("GdipSetClipHrgn")
	gdipResetClip = dll.NewProc("GdipResetClip")
	gdipTranslateClip = dll.NewProc("GdipTranslateClip")
	gdipTranslateClipI = dll.NewProc("GdipTranslateClipI")
	gdipGetClip = dll.NewProc("GdipGetClip")
	gdipGetClipBounds = dll.NewProc("GdipGetClipBounds")
	gdipGetClipBoundsI = dll.NewProc("GdipGetClipBoundsI")
	gdipIsClipEmpty = dll.NewProc("GdipIsClipEmpty")
	gdipGetVisibleClipBounds = dll.NewProc("GdipGetVisibleClipBounds")
	gdipGetVisibleClipBoundsI = dll.NewProc("GdipGetVisibleClipBoundsI")
	gdipIsVisibleClipEmpty = dll.NewProc("GdipIsVisibleClipEmpty")
	gdipIsVisiblePoint = dll.NewProc("GdipIsVisiblePoint")
	gdipIsVisiblePointI = dll.NewProc("GdipIsVisiblePointI")
	gdipIsVisibleRect = dll.NewProc("GdipIsVisibleRect")
	gdipIsVisibleRectI = dll.NewProc("GdipIsVisibleRectI")
	gdipSaveGraphics = dll.NewProc("GdipSaveGraphics")
	gdipRestoreGraphics = dll.NewProc("GdipRestoreGraphics")
	gdipBeginContainer = dll.NewProc("GdipBeginContainer")
	gdipBeginContainerI = dll.NewProc("GdipBeginContainerI")
	gdipBeginContainer2 = dll.NewProc("GdipBeginContainer2")
	gdipEndContainer = dll.NewProc("GdipEndContainer")
	gdipGetMetafileHeaderFromEmf = dll.NewProc("GdipGetMetafileHeaderFromEmf")
	gdipGetMetafileHeaderFromFile = dll.NewProc("GdipGetMetafileHeaderFromFile")
	gdipGetMetafileHeaderFromStream = dll.NewProc("GdipGetMetafileHeaderFromStream")
	gdipGetMetafileHeaderFromMetafile = dll.NewProc("GdipGetMetafileHeaderFromMetafile")
	gdipGetHemfFromMetafile = dll.NewProc("GdipGetHemfFromMetafile")
	gdipCreateStreamOnFile = dll.NewProc("GdipCreateStreamOnFile")
	gdipCreateMetafileFromWmf = dll.NewProc("GdipCreateMetafileFromWmf")
	gdipCreateMetafileFromEmf = dll.NewProc("GdipCreateMetafileFromEmf")
	gdipCreateMetafileFromFile = dll.NewProc("GdipCreateMetafileFromFile")
	gdipCreateMetafileFromWmfFile = dll.NewProc("GdipCreateMetafileFromWmfFile")
	gdipCreateMetafileFromStream = dll.NewProc("GdipCreateMetafileFromStream")
	gdipRecordMetafile = dll.NewProc("GdipRecordMetafile")
	gdipRecordMetafileI = dll.NewProc("GdipRecordMetafileI")
	gdipRecordMetafileFileName = dll.NewProc("GdipRecordMetafileFileName")
	gdipRecordMetafileFileNameI = dll.NewProc("GdipRecordMetafileFileNameI")
	gdipRecordMetafileStream = dll.NewProc("GdipRecordMetafileStream")
	gdipRecordMetafileStreamI = dll.NewProc("GdipRecordMetafileStreamI")
	gdipSetMetafileDownLevelRasterizationLimit = dll.NewProc("GdipSetMetafileDownLevelRasterizationLimit")
	gdipGetMetafileDownLevelRasterizationLimit = dll.NewProc("GdipGetMetafileDownLevelRasterizationLimit")
	gdipGetImageDecodersSize = dll.NewProc("GdipGetImageDecodersSize")
	gdipGetImageDecoders = dll.NewProc("GdipGetImageDecoders")
	gdipGetImageEncodersSize = dll.NewProc("GdipGetImageEncodersSize")
	gdipGetImageEncoders = dll.NewProc("GdipGetImageEncoders")
	gdipComment = dll.NewProc("GdipComment")
}

// graphics
func Flush(graphics *Graphics, intention FlushIntention) Status {
	ret, _, _ := syscall.SyscallN(gdipFlush.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(intention))
	return Status(ret)
}

func CreateFromHDC(hdc win32.HDC, graphics **Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFromHDC.Addr(),
		uintptr(hdc),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func CreateFromHDC2(hdc win32.HDC, hDevice win32.HANDLE, graphics **Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFromHDC2.Addr(),
		uintptr(hdc),
		uintptr(hDevice),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func CreateFromHWND(hwnd win32.HWND, graphics **Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFromHWND.Addr(),
		uintptr(hwnd),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func CreateFromHWNDICM(hwnd win32.HWND, graphics **Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateFromHWNDICM.Addr(),
		uintptr(hwnd),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func DeleteGraphics(graphics *Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteGraphics.Addr(),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func GetDC(graphics *Graphics, hdc *win32.HDC) Status {
	ret, _, _ := syscall.SyscallN(gdipGetDC.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(hdc)))
	return Status(ret)
}

func ReleaseDC(graphics *Graphics, hdc win32.HDC) Status {
	ret, _, _ := syscall.SyscallN(gdipReleaseDC.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(hdc))
	return Status(ret)
}

func SetCompositingMode(graphics *Graphics, compositingMode CompositingMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetCompositingMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(compositingMode))
	return Status(ret)
}

func GetCompositingMode(graphics *Graphics, compositingMode *CompositingMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCompositingMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(compositingMode)))
	return Status(ret)
}

func SetRenderingOrigin(graphics *Graphics, x int32, y int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetRenderingOrigin.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(x),
		uintptr(y))
	return Status(ret)
}

func GetRenderingOrigin(graphics *Graphics, x *int32, y *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetRenderingOrigin.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(y)))
	return Status(ret)
}

func SetCompositingQuality(graphics *Graphics, compositingQuality CompositingQuality) Status {
	ret, _, _ := syscall.SyscallN(gdipSetCompositingQuality.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(compositingQuality))
	return Status(ret)
}

func GetCompositingQuality(graphics *Graphics, compositingQuality *CompositingQuality) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCompositingQuality.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(compositingQuality)))
	return Status(ret)
}

func SetSmoothingMode(graphics *Graphics, smoothingMode SmoothingMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetSmoothingMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(smoothingMode))
	return Status(ret)
}

func GetSmoothingMode(graphics *Graphics, smoothingMode *SmoothingMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetSmoothingMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(smoothingMode)))
	return Status(ret)
}

func SetPixelOffsetMode(graphics *Graphics, pixelOffsetMode PixelOffsetMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPixelOffsetMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(pixelOffsetMode))
	return Status(ret)
}

func GetPixelOffsetMode(graphics *Graphics, pixelOffsetMode *PixelOffsetMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPixelOffsetMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pixelOffsetMode)))
	return Status(ret)
}

func SetTextRenderingHint(graphics *Graphics, mode TextRenderingHint) Status {
	ret, _, _ := syscall.SyscallN(gdipSetTextRenderingHint.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(mode))
	return Status(ret)
}

func GetTextRenderingHint(graphics *Graphics, mode *TextRenderingHint) Status {
	ret, _, _ := syscall.SyscallN(gdipGetTextRenderingHint.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(mode)))
	return Status(ret)
}

func SetTextContrast(graphics *Graphics, contrast uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetTextContrast.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(contrast))
	return Status(ret)
}

func GetTextContrast(graphics *Graphics, contrast *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetTextContrast.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(contrast)))
	return Status(ret)
}

func SetInterpolationMode(graphics *Graphics, interpolationMode InterpolationMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetInterpolationMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(interpolationMode))
	return Status(ret)
}

func GraphicsSetAbort(pGraphics *Graphics, pIAbort *GdiplusAbort) Status {
	ret, _, _ := syscall.SyscallN(gdipGraphicsSetAbort.Addr(),
		uintptr(unsafe.Pointer(pGraphics)),
		uintptr(unsafe.Pointer(pIAbort)))
	return Status(ret)
}

func GetInterpolationMode(graphics *Graphics, interpolationMode *InterpolationMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetInterpolationMode.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(interpolationMode)))
	return Status(ret)
}

func SetWorldTransform(graphics *Graphics, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipSetWorldTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func ResetWorldTransform(graphics *Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipResetWorldTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func MultiplyWorldTransform(graphics *Graphics, matrix *Matrix, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipMultiplyWorldTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
	return Status(ret)
}

func TranslateWorldTransform(graphics *Graphics, dx float32, dy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateWorldTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)),
		uintptr(order))
	return Status(ret)
}

func ScaleWorldTransform(graphics *Graphics, sx float32, sy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipScaleWorldTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(sx)),
		uintptr(math.Float32bits(sy)),
		uintptr(order))
	return Status(ret)
}

func RotateWorldTransform(graphics *Graphics, angle float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipRotateWorldTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(angle)),
		uintptr(order))
	return Status(ret)
}

func GetWorldTransform(graphics *Graphics, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetWorldTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func ResetPageTransform(graphics *Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipResetPageTransform.Addr(),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func GetPageUnit(graphics *Graphics, unit *Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPageUnit.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(unit)))
	return Status(ret)
}

func GetPageScale(graphics *Graphics, scale *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPageScale.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(scale)))
	return Status(ret)
}

func SetPageUnit(graphics *Graphics, unit Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPageUnit.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unit))
	return Status(ret)
}

func SetPageScale(graphics *Graphics, scale float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPageScale.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(scale)))
	return Status(ret)
}

func GetDpiX(graphics *Graphics, dpi *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetDpiX.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(dpi)))
	return Status(ret)
}

func GetDpiY(graphics *Graphics, dpi *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetDpiY.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(dpi)))
	return Status(ret)
}

func TransformPoints(graphics *Graphics, destSpace CoordinateSpace, srcSpace CoordinateSpace, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipTransformPoints.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(destSpace),
		uintptr(srcSpace),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func TransformPointsI(graphics *Graphics, destSpace CoordinateSpace, srcSpace CoordinateSpace, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipTransformPointsI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(destSpace),
		uintptr(srcSpace),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func GetNearestColor(graphics *Graphics, argb *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGetNearestColor.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(argb)))
	return Status(ret)
}

func CreateHalftonePalette() win32.HPALETTE {
	ret, _, _ := syscall.SyscallN(gdipCreateHalftonePalette.Addr())
	return win32.HPALETTE(ret)
}

func DrawLine(graphics *Graphics, pen *Pen, x1 float32, y1 float32, x2 float32, y2 float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawLine.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x1)),
		uintptr(math.Float32bits(y1)),
		uintptr(math.Float32bits(x2)),
		uintptr(math.Float32bits(y2)))
	return Status(ret)
}

func DrawLineI(graphics *Graphics, pen *Pen, x1 int32, y1 int32, x2 int32, y2 int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawLineI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
	return Status(ret)
}

func DrawLines(graphics *Graphics, pen *Pen, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawLines.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawLinesI(graphics *Graphics, pen *Pen, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawLinesI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawArc(graphics *Graphics, pen *Pen, x float32, y float32, width float32, height float32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawArc.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func DrawArcI(graphics *Graphics, pen *Pen, x int32, y int32, width int32, height int32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawArcI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func DrawBezier(graphics *Graphics, pen *Pen, x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32, x4 float32, y4 float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawBezier.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x1)),
		uintptr(math.Float32bits(y1)),
		uintptr(math.Float32bits(x2)),
		uintptr(math.Float32bits(y2)),
		uintptr(math.Float32bits(x3)),
		uintptr(math.Float32bits(y3)),
		uintptr(math.Float32bits(x4)),
		uintptr(math.Float32bits(y4)))
	return Status(ret)
}

func DrawBezierI(graphics *Graphics, pen *Pen, x1 int32, y1 int32, x2 int32, y2 int32, x3 int32, y3 int32, x4 int32, y4 int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawBezierI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(x3),
		uintptr(y3),
		uintptr(x4),
		uintptr(y4))
	return Status(ret)
}

func DrawBeziers(graphics *Graphics, pen *Pen, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawBeziers.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawBeziersI(graphics *Graphics, pen *Pen, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawBeziersI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawRectangle(graphics *Graphics, pen *Pen, x float32, y float32, width float32, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawRectangle.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func DrawRectangleI(graphics *Graphics, pen *Pen, x int32, y int32, width int32, height int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawRectangleI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return Status(ret)
}

func DrawRectangles(graphics *Graphics, pen *Pen, rects *RectF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawRectangles.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count))
	return Status(ret)
}

func DrawRectanglesI(graphics *Graphics, pen *Pen, rects *Rect, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawRectanglesI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count))
	return Status(ret)
}

func DrawEllipse(graphics *Graphics, pen *Pen, x float32, y float32, width float32, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawEllipse.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func DrawEllipseI(graphics *Graphics, pen *Pen, x int32, y int32, width int32, height int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawEllipseI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return Status(ret)
}

func DrawPie(graphics *Graphics, pen *Pen, x float32, y float32, width float32, height float32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawPie.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func DrawPieI(graphics *Graphics, pen *Pen, x int32, y int32, width int32, height int32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawPieI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func DrawPolygon(graphics *Graphics, pen *Pen, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawPolygon.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawPolygonI(graphics *Graphics, pen *Pen, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawPolygonI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawPath(graphics *Graphics, pen *Pen, path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawPath.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func DrawCurve(graphics *Graphics, pen *Pen, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawCurve.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawCurveI(graphics *Graphics, pen *Pen, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawCurveI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawCurve2(graphics *Graphics, pen *Pen, points *PointF, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawCurve2.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func DrawCurve2I(graphics *Graphics, pen *Pen, points *Point, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawCurve2I.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func DrawCurve3(graphics *Graphics, pen *Pen, points *PointF, count int32, offset int32, numberOfSegments int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawCurve3.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(offset),
		uintptr(numberOfSegments),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func DrawCurve3I(graphics *Graphics, pen *Pen, points *Point, count int32, offset int32, numberOfSegments int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawCurve3I.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(offset),
		uintptr(numberOfSegments),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func DrawClosedCurve(graphics *Graphics, pen *Pen, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawClosedCurve.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawClosedCurveI(graphics *Graphics, pen *Pen, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawClosedCurveI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func DrawClosedCurve2(graphics *Graphics, pen *Pen, points *PointF, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawClosedCurve2.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func DrawClosedCurve2I(graphics *Graphics, pen *Pen, points *Point, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawClosedCurve2I.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func GraphicsClear(graphics *Graphics, color ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGraphicsClear.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(color))
	return Status(ret)
}

func FillRectangle(graphics *Graphics, brush *Brush, x float32, y float32, width float32, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillRectangle.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func FillRectangleI(graphics *Graphics, brush *Brush, x int32, y int32, width int32, height int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillRectangleI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return Status(ret)
}

func FillRectangles(graphics *Graphics, brush *Brush, rects *RectF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillRectangles.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count))
	return Status(ret)
}

func FillRectanglesI(graphics *Graphics, brush *Brush, rects *Rect, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillRectanglesI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count))
	return Status(ret)
}

func FillPolygon(graphics *Graphics, brush *Brush, points *PointF, count int32, fillMode FillMode) Status {
	ret, _, _ := syscall.SyscallN(gdipFillPolygon.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(fillMode))
	return Status(ret)
}

func FillPolygonI(graphics *Graphics, brush *Brush, points *Point, count int32, fillMode FillMode) Status {
	ret, _, _ := syscall.SyscallN(gdipFillPolygonI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(fillMode))
	return Status(ret)
}

func FillPolygon2(graphics *Graphics, brush *Brush, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillPolygon2.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func FillPolygon2I(graphics *Graphics, brush *Brush, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillPolygon2I.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func FillEllipse(graphics *Graphics, brush *Brush, x float32, y float32, width float32, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillEllipse.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func FillEllipseI(graphics *Graphics, brush *Brush, x int32, y int32, width int32, height int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillEllipseI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return Status(ret)
}

func FillPie(graphics *Graphics, brush *Brush, x float32, y float32, width float32, height float32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillPie.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func FillPieI(graphics *Graphics, brush *Brush, x int32, y int32, width int32, height int32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillPieI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func FillPath(graphics *Graphics, brush *Brush, path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipFillPath.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func FillClosedCurve(graphics *Graphics, brush *Brush, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillClosedCurve.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func FillClosedCurveI(graphics *Graphics, brush *Brush, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipFillClosedCurveI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func FillClosedCurve2(graphics *Graphics, brush *Brush, points *PointF, count int32, tension float32, fillMode FillMode) Status {
	ret, _, _ := syscall.SyscallN(gdipFillClosedCurve2.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)),
		uintptr(fillMode))
	return Status(ret)
}

func FillClosedCurve2I(graphics *Graphics, brush *Brush, points *Point, count int32, tension float32, fillMode FillMode) Status {
	ret, _, _ := syscall.SyscallN(gdipFillClosedCurve2I.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)),
		uintptr(fillMode))
	return Status(ret)
}

func FillRegion(graphics *Graphics, brush *Brush, region *Region) Status {
	ret, _, _ := syscall.SyscallN(gdipFillRegion.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func DrawImage(graphics *Graphics, image *Image, x float32, y float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImage.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)))
	return Status(ret)
}

func DrawImageI(graphics *Graphics, image *Image, x int32, y int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImageI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(x),
		uintptr(y))
	return Status(ret)
}

func DrawImageRect(graphics *Graphics, image *Image, x float32, y float32, width float32, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImageRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func DrawImageRectI(graphics *Graphics, image *Image, x int32, y int32, width int32, height int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImageRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return Status(ret)
}

func DrawImagePoints(graphics *Graphics, image *Image, dstpoints *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImagePoints.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(dstpoints)),
		uintptr(count))
	return Status(ret)
}

func DrawImagePointsI(graphics *Graphics, image *Image, dstpoints *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImagePointsI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(dstpoints)),
		uintptr(count))
	return Status(ret)
}

func DrawImagePointRect(graphics *Graphics, image *Image, x float32, y float32, srcx float32, srcy float32, srcwidth float32, srcheight float32, srcUnit Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImagePointRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(srcx)),
		uintptr(math.Float32bits(srcy)),
		uintptr(math.Float32bits(srcwidth)),
		uintptr(math.Float32bits(srcheight)),
		uintptr(srcUnit))
	return Status(ret)
}

func DrawImagePointRectI(graphics *Graphics, image *Image, x int32, y int32, srcx int32, srcy int32, srcwidth int32, srcheight int32, srcUnit Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImagePointRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(x),
		uintptr(y),
		uintptr(srcx),
		uintptr(srcy),
		uintptr(srcwidth),
		uintptr(srcheight),
		uintptr(srcUnit))
	return Status(ret)
}

func DrawImageRectRect(graphics *Graphics, image *Image, dstx float32, dsty float32, dstwidth float32, dstheight float32, srcx float32, srcy float32, srcwidth float32, srcheight float32, srcUnit Unit, imageAttributes *ImageAttributes, callback uintptr, callbackData *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImageRectRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(math.Float32bits(dstx)),
		uintptr(math.Float32bits(dsty)),
		uintptr(math.Float32bits(dstwidth)),
		uintptr(math.Float32bits(dstheight)),
		uintptr(math.Float32bits(srcx)),
		uintptr(math.Float32bits(srcy)),
		uintptr(math.Float32bits(srcwidth)),
		uintptr(math.Float32bits(srcheight)),
		uintptr(srcUnit),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)))
	return Status(ret)
}

func DrawImageRectRectI(graphics *Graphics, image *Image, dstx int32, dsty int32, dstwidth int32, dstheight int32, srcx int32, srcy int32, srcwidth int32, srcheight int32, srcUnit Unit, imageAttributes *ImageAttributes, callback uintptr, callbackData *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImageRectRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(dstx),
		uintptr(dsty),
		uintptr(dstwidth),
		uintptr(dstheight),
		uintptr(srcx),
		uintptr(srcy),
		uintptr(srcwidth),
		uintptr(srcheight),
		uintptr(srcUnit),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)))
	return Status(ret)
}

func DrawImagePointsRect(graphics *Graphics, image *Image, points *PointF, count int32, srcx float32, srcy float32, srcwidth float32, srcheight float32, srcUnit Unit, imageAttributes *ImageAttributes, callback uintptr, callbackData *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImagePointsRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(srcx)),
		uintptr(math.Float32bits(srcy)),
		uintptr(math.Float32bits(srcwidth)),
		uintptr(math.Float32bits(srcheight)),
		uintptr(srcUnit),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)))
	return Status(ret)
}

func DrawImagePointsRectI(graphics *Graphics, image *Image, points *Point, count int32, srcx int32, srcy int32, srcwidth int32, srcheight int32, srcUnit Unit, imageAttributes *ImageAttributes, callback uintptr, callbackData *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImagePointsRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(srcx),
		uintptr(srcy),
		uintptr(srcwidth),
		uintptr(srcheight),
		uintptr(srcUnit),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)))
	return Status(ret)
}

func DrawImageFX(graphics *Graphics, image *Image, source *RectF, xForm *Matrix, effect *CGpEffect, imageAttributes *ImageAttributes, srcUnit Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawImageFX.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(xForm)),
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(srcUnit))
	return Status(ret)
}

func EnumerateMetafileDestPoint(graphics *Graphics, metafile *Metafile, destPoint *PointF, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileDestPoint.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoint)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileDestPointI(graphics *Graphics, metafile *Metafile, destPoint *Point, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileDestPointI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoint)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileDestRect(graphics *Graphics, metafile *Metafile, destRect *RectF, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileDestRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destRect)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileDestRectI(graphics *Graphics, metafile *Metafile, destRect *Rect, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileDestRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destRect)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileDestPoints(graphics *Graphics, metafile *Metafile, destPoints *PointF, count int32, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileDestPoints.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoints)),
		uintptr(count),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileDestPointsI(graphics *Graphics, metafile *Metafile, destPoints *Point, count int32, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileDestPointsI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoints)),
		uintptr(count),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileSrcRectDestPoint(graphics *Graphics, metafile *Metafile, destPoint *PointF, srcRect *RectF, srcUnit Unit, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileSrcRectDestPoint.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoint)),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(srcUnit),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileSrcRectDestPointI(graphics *Graphics, metafile *Metafile, destPoint *Point, srcRect *Rect, srcUnit Unit, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileSrcRectDestPointI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoint)),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(srcUnit),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileSrcRectDestRect(graphics *Graphics, metafile *Metafile, destRect *RectF, srcRect *RectF, srcUnit Unit, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileSrcRectDestRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destRect)),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(srcUnit),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileSrcRectDestRectI(graphics *Graphics, metafile *Metafile, destRect *Rect, srcRect *Rect, srcUnit Unit, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileSrcRectDestRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destRect)),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(srcUnit),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileSrcRectDestPoints(graphics *Graphics, metafile *Metafile, destPoints *PointF, count int32, srcRect *RectF, srcUnit Unit, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileSrcRectDestPoints.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoints)),
		uintptr(count),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(srcUnit),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func EnumerateMetafileSrcRectDestPointsI(graphics *Graphics, metafile *Metafile, destPoints *Point, count int32, srcRect *Rect, srcUnit Unit, callback uintptr, callbackData *byte, imageAttributes *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipEnumerateMetafileSrcRectDestPointsI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(destPoints)),
		uintptr(count),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(srcUnit),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)),
		uintptr(unsafe.Pointer(imageAttributes)))
	return Status(ret)
}

func PlayMetafileRecord(metafile *Metafile, recordType EmfPlusRecordType, flags uint32, dataSize uint32, data *win32.BYTE) Status {
	ret, _, _ := syscall.SyscallN(gdipPlayMetafileRecord.Addr(),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(recordType),
		uintptr(flags),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(data)))
	return Status(ret)
}

func SetClipGraphics(graphics *Graphics, srcgraphics *Graphics, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetClipGraphics.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(srcgraphics)),
		uintptr(combineMode))
	return Status(ret)
}

func SetClipRect(graphics *Graphics, x float32, y float32, width float32, height float32, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetClipRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(combineMode))
	return Status(ret)
}

func SetClipRectI(graphics *Graphics, x int32, y int32, width int32, height int32, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetClipRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(combineMode))
	return Status(ret)
}

func SetClipPath(graphics *Graphics, path *Path, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetClipPath.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(path)),
		uintptr(combineMode))
	return Status(ret)
}

func SetClipRegion(graphics *Graphics, region *Region, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetClipRegion.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(region)),
		uintptr(combineMode))
	return Status(ret)
}

func SetClipHrgn(graphics *Graphics, hRgn win32.HRGN, combineMode CombineMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetClipHrgn.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(hRgn),
		uintptr(combineMode))
	return Status(ret)
}

func ResetClip(graphics *Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipResetClip.Addr(),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func TranslateClip(graphics *Graphics, dx float32, dy float32) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateClip.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)))
	return Status(ret)
}

func TranslateClipI(graphics *Graphics, dx int32, dy int32) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateClipI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(dx),
		uintptr(dy))
	return Status(ret)
}

func GetClip(graphics *Graphics, region *Region) Status {
	ret, _, _ := syscall.SyscallN(gdipGetClip.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(region)))
	return Status(ret)
}

func GetClipBounds(graphics *Graphics, rect *RectF) Status {
	ret, _, _ := syscall.SyscallN(gdipGetClipBounds.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func GetClipBoundsI(graphics *Graphics, rect *Rect) Status {
	ret, _, _ := syscall.SyscallN(gdipGetClipBoundsI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func IsClipEmpty(graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsClipEmpty.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func GetVisibleClipBounds(graphics *Graphics, rect *RectF) Status {
	ret, _, _ := syscall.SyscallN(gdipGetVisibleClipBounds.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func GetVisibleClipBoundsI(graphics *Graphics, rect *Rect) Status {
	ret, _, _ := syscall.SyscallN(gdipGetVisibleClipBoundsI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func IsVisibleClipEmpty(graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisibleClipEmpty.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisiblePoint(graphics *Graphics, x float32, y float32, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisiblePoint.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisiblePointI(graphics *Graphics, x int32, y int32, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisiblePointI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisibleRect(graphics *Graphics, x float32, y float32, width float32, height float32, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisibleRect.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisibleRectI(graphics *Graphics, x int32, y int32, width int32, height int32, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisibleRectI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func SaveGraphics(graphics *Graphics, state *GraphicsState) Status {
	ret, _, _ := syscall.SyscallN(gdipSaveGraphics.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(state)))
	return Status(ret)
}

func RestoreGraphics(graphics *Graphics, state GraphicsState) Status {
	ret, _, _ := syscall.SyscallN(gdipRestoreGraphics.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(state))
	return Status(ret)
}

func BeginContainer(graphics *Graphics, dstrect *RectF, srcrect *RectF, unit Unit, state *GraphicsContainer) Status {
	ret, _, _ := syscall.SyscallN(gdipBeginContainer.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(dstrect)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unit),
		uintptr(unsafe.Pointer(state)))
	return Status(ret)
}

func BeginContainerI(graphics *Graphics, dstrect *Rect, srcrect *Rect, unit Unit, state *GraphicsContainer) Status {
	ret, _, _ := syscall.SyscallN(gdipBeginContainerI.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(dstrect)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unit),
		uintptr(unsafe.Pointer(state)))
	return Status(ret)
}

func BeginContainer2(graphics *Graphics, state *GraphicsContainer) Status {
	ret, _, _ := syscall.SyscallN(gdipBeginContainer2.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(state)))
	return Status(ret)
}

func EndContainer(graphics *Graphics, state GraphicsContainer) Status {
	ret, _, _ := syscall.SyscallN(gdipEndContainer.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(state))
	return Status(ret)
}

func GetMetafileHeaderFromEmf(hEmf win32.HANDLE, header *MetafileHeader) Status {
	ret, _, _ := syscall.SyscallN(gdipGetMetafileHeaderFromEmf.Addr(),
		uintptr(hEmf),
		uintptr(unsafe.Pointer(header)))
	return Status(ret)
}

func GetMetafileHeaderFromFile(filename *uint16, header *MetafileHeader) Status {
	ret, _, _ := syscall.SyscallN(gdipGetMetafileHeaderFromFile.Addr(),
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(header)))
	return Status(ret)
}

func GetMetafileHeaderFromStream(stream *win32.IStream, header *MetafileHeader) Status {
	ret, _, _ := syscall.SyscallN(gdipGetMetafileHeaderFromStream.Addr(),
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(header)))
	return Status(ret)
}

func GetMetafileHeaderFromMetafile(metafile *Metafile, header *MetafileHeader) Status {
	ret, _, _ := syscall.SyscallN(gdipGetMetafileHeaderFromMetafile.Addr(),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(header)))
	return Status(ret)
}

func GetHemfFromMetafile(metafile *Metafile, hEmf *win32.HANDLE) Status {
	ret, _, _ := syscall.SyscallN(gdipGetHemfFromMetafile.Addr(),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(hEmf)))
	return Status(ret)
}

func CreateStreamOnFile(filename *uint16, access uint32, stream **win32.IStream) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateStreamOnFile.Addr(),
		uintptr(unsafe.Pointer(filename)),
		uintptr(access),
		uintptr(unsafe.Pointer(stream)))
	return Status(ret)
}

func CreateMetafileFromWmf(hWmf win32.HANDLE, deleteWmf win32.BOOL, wmfPlaceableFileHeader *WmfPlaceableFileHeader, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMetafileFromWmf.Addr(),
		uintptr(hWmf),
		uintptr(deleteWmf),
		uintptr(unsafe.Pointer(wmfPlaceableFileHeader)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func CreateMetafileFromEmf(hEmf win32.HANDLE, deleteEmf win32.BOOL, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMetafileFromEmf.Addr(),
		uintptr(hEmf),
		uintptr(deleteEmf),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func CreateMetafileFromFile(file *uint16, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMetafileFromFile.Addr(),
		uintptr(unsafe.Pointer(file)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func CreateMetafileFromWmfFile(file *uint16, wmfPlaceableFileHeader *WmfPlaceableFileHeader, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMetafileFromWmfFile.Addr(),
		uintptr(unsafe.Pointer(file)),
		uintptr(unsafe.Pointer(wmfPlaceableFileHeader)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func CreateMetafileFromStream(stream *win32.IStream, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMetafileFromStream.Addr(),
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func RecordMetafile(referenceHdc win32.HDC, typ EmfType, frameRect *RectF, frameUnit MetafileFrameUnit, description *uint16, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipRecordMetafile.Addr(),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func RecordMetafileI(referenceHdc win32.HDC, typ EmfType, frameRect *Rect, frameUnit MetafileFrameUnit, description *uint16, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipRecordMetafileI.Addr(),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func RecordMetafileFileName(fileName *uint16, referenceHdc win32.HDC, typ EmfType, frameRect *RectF, frameUnit MetafileFrameUnit, description *uint16, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipRecordMetafileFileName.Addr(),
		uintptr(unsafe.Pointer(fileName)),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func RecordMetafileFileNameI(fileName *uint16, referenceHdc win32.HDC, typ EmfType, frameRect *Rect, frameUnit MetafileFrameUnit, description *uint16, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipRecordMetafileFileNameI.Addr(),
		uintptr(unsafe.Pointer(fileName)),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func RecordMetafileStream(stream *win32.IStream, referenceHdc win32.HDC, typ EmfType, frameRect *RectF, frameUnit MetafileFrameUnit, description *uint16, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipRecordMetafileStream.Addr(),
		uintptr(unsafe.Pointer(stream)),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func RecordMetafileStreamI(stream *win32.IStream, referenceHdc win32.HDC, typ EmfType, frameRect *Rect, frameUnit MetafileFrameUnit, description *uint16, metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipRecordMetafileStreamI.Addr(),
		uintptr(unsafe.Pointer(stream)),
		uintptr(referenceHdc),
		uintptr(typ),
		uintptr(unsafe.Pointer(frameRect)),
		uintptr(frameUnit),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(metafile)))
	return Status(ret)
}

func SetMetafileDownLevelRasterizationLimit(metafile *Metafile, metafileRasterizationLimitDpi uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetMetafileDownLevelRasterizationLimit.Addr(),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(metafileRasterizationLimitDpi))
	return Status(ret)
}

func GetMetafileDownLevelRasterizationLimit(metafile *Metafile, metafileRasterizationLimitDpi *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetMetafileDownLevelRasterizationLimit.Addr(),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(metafileRasterizationLimitDpi)))
	return Status(ret)
}

func GetImageDecodersSize(numDecoders *uint32, size *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageDecodersSize.Addr(),
		uintptr(unsafe.Pointer(numDecoders)),
		uintptr(unsafe.Pointer(size)))
	return Status(ret)
}

func GetImageDecoders(numDecoders uint32, size uint32, decoders *ImageCodecInfo) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageDecoders.Addr(),
		uintptr(numDecoders),
		uintptr(size),
		uintptr(unsafe.Pointer(decoders)))
	return Status(ret)
}

func GetImageEncodersSize(numEncoders *uint32, size *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageEncodersSize.Addr(),
		uintptr(unsafe.Pointer(numEncoders)),
		uintptr(unsafe.Pointer(size)))
	return Status(ret)
}

func GetImageEncoders(numEncoders uint32, size uint32, encoders *ImageCodecInfo) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageEncoders.Addr(),
		uintptr(numEncoders),
		uintptr(size),
		uintptr(unsafe.Pointer(encoders)))
	return Status(ret)
}

func Comment(graphics *Graphics, sizeData uint32, data *win32.BYTE) Status {
	ret, _, _ := syscall.SyscallN(gdipComment.Addr(),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(sizeData),
		uintptr(unsafe.Pointer(data)))
	return Status(ret)
}

