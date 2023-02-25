package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// brush
	gdipCloneBrush	*windows.LazyProc
	gdipDeleteBrush	*windows.LazyProc
	gdipGetBrushType	*windows.LazyProc

	// hatchbrush
	gdipCreateHatchBrush	*windows.LazyProc
	gdipGetHatchStyle	*windows.LazyProc
	gdipGetHatchForegroundColor	*windows.LazyProc
	gdipGetHatchBackgroundColor	*windows.LazyProc

	// lineargradientbrush
	gdipCreateLineBrush	*windows.LazyProc
	gdipCreateLineBrushI	*windows.LazyProc
	gdipCreateLineBrushFromRect	*windows.LazyProc
	gdipCreateLineBrushFromRectI	*windows.LazyProc
	gdipCreateLineBrushFromRectWithAngle	*windows.LazyProc
	gdipCreateLineBrushFromRectWithAngleI	*windows.LazyProc
	gdipSetLineColors	*windows.LazyProc
	gdipGetLineColors	*windows.LazyProc
	gdipGetLineRect	*windows.LazyProc
	gdipGetLineRectI	*windows.LazyProc
	gdipSetLineGammaCorrection	*windows.LazyProc
	gdipGetLineGammaCorrection	*windows.LazyProc
	gdipGetLineBlendCount	*windows.LazyProc
	gdipGetLineBlend	*windows.LazyProc
	gdipSetLineBlend	*windows.LazyProc
	gdipGetLinePresetBlendCount	*windows.LazyProc
	gdipGetLinePresetBlend	*windows.LazyProc
	gdipSetLinePresetBlend	*windows.LazyProc
	gdipSetLineSigmaBlend	*windows.LazyProc
	gdipSetLineLinearBlend	*windows.LazyProc
	gdipSetLineWrapMode	*windows.LazyProc
	gdipGetLineWrapMode	*windows.LazyProc
	gdipGetLineTransform	*windows.LazyProc
	gdipSetLineTransform	*windows.LazyProc
	gdipResetLineTransform	*windows.LazyProc
	gdipMultiplyLineTransform	*windows.LazyProc
	gdipTranslateLineTransform	*windows.LazyProc
	gdipScaleLineTransform	*windows.LazyProc
	gdipRotateLineTransform	*windows.LazyProc

	// pathgradientbrush
	gdipCreatePathGradient	*windows.LazyProc
	gdipCreatePathGradientI	*windows.LazyProc
	gdipCreatePathGradientFromPath	*windows.LazyProc
	gdipGetPathGradientCenterColor	*windows.LazyProc
	gdipSetPathGradientCenterColor	*windows.LazyProc
	gdipGetPathGradientSurroundColorsWithCount	*windows.LazyProc
	gdipSetPathGradientSurroundColorsWithCount	*windows.LazyProc
	gdipGetPathGradientPath	*windows.LazyProc
	gdipSetPathGradientPath	*windows.LazyProc
	gdipGetPathGradientCenterPoint	*windows.LazyProc
	gdipGetPathGradientCenterPointI	*windows.LazyProc
	gdipSetPathGradientCenterPoint	*windows.LazyProc
	gdipSetPathGradientCenterPointI	*windows.LazyProc
	gdipGetPathGradientRect	*windows.LazyProc
	gdipGetPathGradientRectI	*windows.LazyProc
	gdipGetPathGradientPointCount	*windows.LazyProc
	gdipGetPathGradientSurroundColorCount	*windows.LazyProc
	gdipSetPathGradientGammaCorrection	*windows.LazyProc
	gdipGetPathGradientGammaCorrection	*windows.LazyProc
	gdipGetPathGradientBlendCount	*windows.LazyProc
	gdipGetPathGradientBlend	*windows.LazyProc
	gdipSetPathGradientBlend	*windows.LazyProc
	gdipGetPathGradientPresetBlendCount	*windows.LazyProc
	gdipGetPathGradientPresetBlend	*windows.LazyProc
	gdipSetPathGradientPresetBlend	*windows.LazyProc
	gdipSetPathGradientSigmaBlend	*windows.LazyProc
	gdipSetPathGradientLinearBlend	*windows.LazyProc
	gdipGetPathGradientWrapMode	*windows.LazyProc
	gdipSetPathGradientWrapMode	*windows.LazyProc
	gdipGetPathGradientTransform	*windows.LazyProc
	gdipSetPathGradientTransform	*windows.LazyProc
	gdipResetPathGradientTransform	*windows.LazyProc
	gdipMultiplyPathGradientTransform	*windows.LazyProc
	gdipTranslatePathGradientTransform	*windows.LazyProc
	gdipScalePathGradientTransform	*windows.LazyProc
	gdipRotatePathGradientTransform	*windows.LazyProc
	gdipGetPathGradientFocusScales	*windows.LazyProc
	gdipSetPathGradientFocusScales	*windows.LazyProc

	// solidbrush
	gdipCreateSolidFill	*windows.LazyProc
	gdipSetSolidFillColor	*windows.LazyProc
	gdipGetSolidFillColor	*windows.LazyProc

	// texturebrush
	gdipCreateTexture	*windows.LazyProc
	gdipCreateTexture2	*windows.LazyProc
	gdipCreateTextureIA	*windows.LazyProc
	gdipCreateTexture2I	*windows.LazyProc
	gdipCreateTextureIAI	*windows.LazyProc
	gdipGetTextureTransform	*windows.LazyProc
	gdipSetTextureTransform	*windows.LazyProc
	gdipResetTextureTransform	*windows.LazyProc
	gdipMultiplyTextureTransform	*windows.LazyProc
	gdipTranslateTextureTransform	*windows.LazyProc
	gdipScaleTextureTransform	*windows.LazyProc
	gdipRotateTextureTransform	*windows.LazyProc
	gdipSetTextureWrapMode	*windows.LazyProc
	gdipGetTextureWrapMode	*windows.LazyProc
	gdipGetTextureImage	*windows.LazyProc
)

func init() {

	// brush
	gdipCloneBrush = dll.NewProc("GdipCloneBrush")
	gdipDeleteBrush = dll.NewProc("GdipDeleteBrush")
	gdipGetBrushType = dll.NewProc("GdipGetBrushType")

	// hatchbrush
	gdipCreateHatchBrush = dll.NewProc("GdipCreateHatchBrush")
	gdipGetHatchStyle = dll.NewProc("GdipGetHatchStyle")
	gdipGetHatchForegroundColor = dll.NewProc("GdipGetHatchForegroundColor")
	gdipGetHatchBackgroundColor = dll.NewProc("GdipGetHatchBackgroundColor")

	// lineargradientbrush
	gdipCreateLineBrush = dll.NewProc("GdipCreateLineBrush")
	gdipCreateLineBrushI = dll.NewProc("GdipCreateLineBrushI")
	gdipCreateLineBrushFromRect = dll.NewProc("GdipCreateLineBrushFromRect")
	gdipCreateLineBrushFromRectI = dll.NewProc("GdipCreateLineBrushFromRectI")
	gdipCreateLineBrushFromRectWithAngle = dll.NewProc("GdipCreateLineBrushFromRectWithAngle")
	gdipCreateLineBrushFromRectWithAngleI = dll.NewProc("GdipCreateLineBrushFromRectWithAngleI")
	gdipSetLineColors = dll.NewProc("GdipSetLineColors")
	gdipGetLineColors = dll.NewProc("GdipGetLineColors")
	gdipGetLineRect = dll.NewProc("GdipGetLineRect")
	gdipGetLineRectI = dll.NewProc("GdipGetLineRectI")
	gdipSetLineGammaCorrection = dll.NewProc("GdipSetLineGammaCorrection")
	gdipGetLineGammaCorrection = dll.NewProc("GdipGetLineGammaCorrection")
	gdipGetLineBlendCount = dll.NewProc("GdipGetLineBlendCount")
	gdipGetLineBlend = dll.NewProc("GdipGetLineBlend")
	gdipSetLineBlend = dll.NewProc("GdipSetLineBlend")
	gdipGetLinePresetBlendCount = dll.NewProc("GdipGetLinePresetBlendCount")
	gdipGetLinePresetBlend = dll.NewProc("GdipGetLinePresetBlend")
	gdipSetLinePresetBlend = dll.NewProc("GdipSetLinePresetBlend")
	gdipSetLineSigmaBlend = dll.NewProc("GdipSetLineSigmaBlend")
	gdipSetLineLinearBlend = dll.NewProc("GdipSetLineLinearBlend")
	gdipSetLineWrapMode = dll.NewProc("GdipSetLineWrapMode")
	gdipGetLineWrapMode = dll.NewProc("GdipGetLineWrapMode")
	gdipGetLineTransform = dll.NewProc("GdipGetLineTransform")
	gdipSetLineTransform = dll.NewProc("GdipSetLineTransform")
	gdipResetLineTransform = dll.NewProc("GdipResetLineTransform")
	gdipMultiplyLineTransform = dll.NewProc("GdipMultiplyLineTransform")
	gdipTranslateLineTransform = dll.NewProc("GdipTranslateLineTransform")
	gdipScaleLineTransform = dll.NewProc("GdipScaleLineTransform")
	gdipRotateLineTransform = dll.NewProc("GdipRotateLineTransform")

	// pathgradientbrush
	gdipCreatePathGradient = dll.NewProc("GdipCreatePathGradient")
	gdipCreatePathGradientI = dll.NewProc("GdipCreatePathGradientI")
	gdipCreatePathGradientFromPath = dll.NewProc("GdipCreatePathGradientFromPath")
	gdipGetPathGradientCenterColor = dll.NewProc("GdipGetPathGradientCenterColor")
	gdipSetPathGradientCenterColor = dll.NewProc("GdipSetPathGradientCenterColor")
	gdipGetPathGradientSurroundColorsWithCount = dll.NewProc("GdipGetPathGradientSurroundColorsWithCount")
	gdipSetPathGradientSurroundColorsWithCount = dll.NewProc("GdipSetPathGradientSurroundColorsWithCount")
	gdipGetPathGradientPath = dll.NewProc("GdipGetPathGradientPath")
	gdipSetPathGradientPath = dll.NewProc("GdipSetPathGradientPath")
	gdipGetPathGradientCenterPoint = dll.NewProc("GdipGetPathGradientCenterPoint")
	gdipGetPathGradientCenterPointI = dll.NewProc("GdipGetPathGradientCenterPointI")
	gdipSetPathGradientCenterPoint = dll.NewProc("GdipSetPathGradientCenterPoint")
	gdipSetPathGradientCenterPointI = dll.NewProc("GdipSetPathGradientCenterPointI")
	gdipGetPathGradientRect = dll.NewProc("GdipGetPathGradientRect")
	gdipGetPathGradientRectI = dll.NewProc("GdipGetPathGradientRectI")
	gdipGetPathGradientPointCount = dll.NewProc("GdipGetPathGradientPointCount")
	gdipGetPathGradientSurroundColorCount = dll.NewProc("GdipGetPathGradientSurroundColorCount")
	gdipSetPathGradientGammaCorrection = dll.NewProc("GdipSetPathGradientGammaCorrection")
	gdipGetPathGradientGammaCorrection = dll.NewProc("GdipGetPathGradientGammaCorrection")
	gdipGetPathGradientBlendCount = dll.NewProc("GdipGetPathGradientBlendCount")
	gdipGetPathGradientBlend = dll.NewProc("GdipGetPathGradientBlend")
	gdipSetPathGradientBlend = dll.NewProc("GdipSetPathGradientBlend")
	gdipGetPathGradientPresetBlendCount = dll.NewProc("GdipGetPathGradientPresetBlendCount")
	gdipGetPathGradientPresetBlend = dll.NewProc("GdipGetPathGradientPresetBlend")
	gdipSetPathGradientPresetBlend = dll.NewProc("GdipSetPathGradientPresetBlend")
	gdipSetPathGradientSigmaBlend = dll.NewProc("GdipSetPathGradientSigmaBlend")
	gdipSetPathGradientLinearBlend = dll.NewProc("GdipSetPathGradientLinearBlend")
	gdipGetPathGradientWrapMode = dll.NewProc("GdipGetPathGradientWrapMode")
	gdipSetPathGradientWrapMode = dll.NewProc("GdipSetPathGradientWrapMode")
	gdipGetPathGradientTransform = dll.NewProc("GdipGetPathGradientTransform")
	gdipSetPathGradientTransform = dll.NewProc("GdipSetPathGradientTransform")
	gdipResetPathGradientTransform = dll.NewProc("GdipResetPathGradientTransform")
	gdipMultiplyPathGradientTransform = dll.NewProc("GdipMultiplyPathGradientTransform")
	gdipTranslatePathGradientTransform = dll.NewProc("GdipTranslatePathGradientTransform")
	gdipScalePathGradientTransform = dll.NewProc("GdipScalePathGradientTransform")
	gdipRotatePathGradientTransform = dll.NewProc("GdipRotatePathGradientTransform")
	gdipGetPathGradientFocusScales = dll.NewProc("GdipGetPathGradientFocusScales")
	gdipSetPathGradientFocusScales = dll.NewProc("GdipSetPathGradientFocusScales")

	// solidbrush
	gdipCreateSolidFill = dll.NewProc("GdipCreateSolidFill")
	gdipSetSolidFillColor = dll.NewProc("GdipSetSolidFillColor")
	gdipGetSolidFillColor = dll.NewProc("GdipGetSolidFillColor")

	// texturebrush
	gdipCreateTexture = dll.NewProc("GdipCreateTexture")
	gdipCreateTexture2 = dll.NewProc("GdipCreateTexture2")
	gdipCreateTextureIA = dll.NewProc("GdipCreateTextureIA")
	gdipCreateTexture2I = dll.NewProc("GdipCreateTexture2I")
	gdipCreateTextureIAI = dll.NewProc("GdipCreateTextureIAI")
	gdipGetTextureTransform = dll.NewProc("GdipGetTextureTransform")
	gdipSetTextureTransform = dll.NewProc("GdipSetTextureTransform")
	gdipResetTextureTransform = dll.NewProc("GdipResetTextureTransform")
	gdipMultiplyTextureTransform = dll.NewProc("GdipMultiplyTextureTransform")
	gdipTranslateTextureTransform = dll.NewProc("GdipTranslateTextureTransform")
	gdipScaleTextureTransform = dll.NewProc("GdipScaleTextureTransform")
	gdipRotateTextureTransform = dll.NewProc("GdipRotateTextureTransform")
	gdipSetTextureWrapMode = dll.NewProc("GdipSetTextureWrapMode")
	gdipGetTextureWrapMode = dll.NewProc("GdipGetTextureWrapMode")
	gdipGetTextureImage = dll.NewProc("GdipGetTextureImage")
}

	// brush
func CloneBrush(brush *Brush, cloneBrush **Brush) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneBrush.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(cloneBrush)))
	return Status(ret)
}

func DeleteBrush(brush *Brush) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteBrush.Addr(), 
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func GetBrushType(brush *Brush, typ *BrushType) Status {
	ret, _, _ := syscall.SyscallN(gdipGetBrushType.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(typ)))
	return Status(ret)
}


	// hatchbrush
func CreateHatchBrush(hatchstyle HatchStyle, forecol ARGB, backcol ARGB, brush **Hatch) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateHatchBrush.Addr(), 
		uintptr(hatchstyle),
		uintptr(forecol),
		uintptr(backcol),
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func GetHatchStyle(brush *Hatch, hatchstyle *HatchStyle) Status {
	ret, _, _ := syscall.SyscallN(gdipGetHatchStyle.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(hatchstyle)))
	return Status(ret)
}

func GetHatchForegroundColor(brush *Hatch, forecol *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGetHatchForegroundColor.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(forecol)))
	return Status(ret)
}

func GetHatchBackgroundColor(brush *Hatch, backcol *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGetHatchBackgroundColor.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(backcol)))
	return Status(ret)
}


	// lineargradientbrush
func CreateLineBrush(point1 *PointF, point2 *PointF, color1 ARGB, color2 ARGB, wrapMode WrapMode, lineGradient **LineGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateLineBrush.Addr(), 
		uintptr(unsafe.Pointer(point1)),
		uintptr(unsafe.Pointer(point2)),
		uintptr(color1),
		uintptr(color2),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
	return Status(ret)
}

func CreateLineBrushI(point1 *Point, point2 *Point, color1 ARGB, color2 ARGB, wrapMode WrapMode, lineGradient **LineGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateLineBrushI.Addr(), 
		uintptr(unsafe.Pointer(point1)),
		uintptr(unsafe.Pointer(point2)),
		uintptr(color1),
		uintptr(color2),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
	return Status(ret)
}

func CreateLineBrushFromRect(rect *RectF, color1 ARGB, color2 ARGB, mode LinearGradientMode, wrapMode WrapMode, lineGradient **LineGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateLineBrushFromRect.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1),
		uintptr(color2),
		uintptr(mode),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
	return Status(ret)
}

func CreateLineBrushFromRectI(rect *Rect, color1 ARGB, color2 ARGB, mode LinearGradientMode, wrapMode WrapMode, lineGradient **LineGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateLineBrushFromRectI.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1),
		uintptr(color2),
		uintptr(mode),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
	return Status(ret)
}

func CreateLineBrushFromRectWithAngle(rect *RectF, color1 ARGB, color2 ARGB, angle float32, isAngleScalable win32.BOOL, wrapMode WrapMode, lineGradient **LineGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateLineBrushFromRectWithAngle.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1),
		uintptr(color2),
		uintptr(math.Float32bits(angle)),
		uintptr(isAngleScalable),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
	return Status(ret)
}

func CreateLineBrushFromRectWithAngleI(rect *Rect, color1 ARGB, color2 ARGB, angle float32, isAngleScalable win32.BOOL, wrapMode WrapMode, lineGradient **LineGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateLineBrushFromRectWithAngleI.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(color1),
		uintptr(color2),
		uintptr(math.Float32bits(angle)),
		uintptr(isAngleScalable),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))
	return Status(ret)
}

func SetLineColors(brush *LineGradient, color1 ARGB, color2 ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLineColors.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(color1),
		uintptr(color2))
	return Status(ret)
}

func GetLineColors(brush *LineGradient, colors *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineColors.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(colors)))
	return Status(ret)
}

func GetLineRect(brush *LineGradient, rect *RectF) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineRect.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func GetLineRectI(brush *LineGradient, rect *Rect) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineRectI.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func SetLineGammaCorrection(brush *LineGradient, useGammaCorrection win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLineGammaCorrection.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(useGammaCorrection))
	return Status(ret)
}

func GetLineGammaCorrection(brush *LineGradient, useGammaCorrection *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineGammaCorrection.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(useGammaCorrection)))
	return Status(ret)
}

func GetLineBlendCount(brush *LineGradient, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineBlendCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func GetLineBlend(brush *LineGradient, blend *float32, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func SetLineBlend(brush *LineGradient, blend *float32, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLineBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func GetLinePresetBlendCount(brush *LineGradient, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLinePresetBlendCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func GetLinePresetBlend(brush *LineGradient, blend *ARGB, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLinePresetBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func SetLinePresetBlend(brush *LineGradient, blend *ARGB, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLinePresetBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func SetLineSigmaBlend(brush *LineGradient, focus float32, scale float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLineSigmaBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(focus)),
		uintptr(math.Float32bits(scale)))
	return Status(ret)
}

func SetLineLinearBlend(brush *LineGradient, focus float32, scale float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLineLinearBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(focus)),
		uintptr(math.Float32bits(scale)))
	return Status(ret)
}

func SetLineWrapMode(brush *LineGradient, wrapmode WrapMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLineWrapMode.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(wrapmode))
	return Status(ret)
}

func GetLineWrapMode(brush *LineGradient, wrapmode *WrapMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineWrapMode.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(wrapmode)))
	return Status(ret)
}

func GetLineTransform(brush *LineGradient, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetLineTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func SetLineTransform(brush *LineGradient, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipSetLineTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func ResetLineTransform(brush *LineGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipResetLineTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func MultiplyLineTransform(brush *LineGradient, matrix *Matrix, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipMultiplyLineTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
	return Status(ret)
}

func TranslateLineTransform(brush *LineGradient, dx float32, dy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateLineTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)),
		uintptr(order))
	return Status(ret)
}

func ScaleLineTransform(brush *LineGradient, sx float32, sy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipScaleLineTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(sx)),
		uintptr(math.Float32bits(sy)),
		uintptr(order))
	return Status(ret)
}

func RotateLineTransform(brush *LineGradient, angle float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipRotateLineTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(angle)),
		uintptr(order))
	return Status(ret)
}


	// pathgradientbrush
func CreatePathGradient(points *PointF, count int32, wrapMode WrapMode, polyGradient **PathGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePathGradient.Addr(), 
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(polyGradient)))
	return Status(ret)
}

func CreatePathGradientI(points *Point, count int32, wrapMode WrapMode, polyGradient **PathGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePathGradientI.Addr(), 
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(polyGradient)))
	return Status(ret)
}

func CreatePathGradientFromPath(path *Path, polyGradient **PathGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePathGradientFromPath.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(polyGradient)))
	return Status(ret)
}

func GetPathGradientCenterColor(brush *PathGradient, colors *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientCenterColor.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(colors)))
	return Status(ret)
}

func SetPathGradientCenterColor(brush *PathGradient, colors ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientCenterColor.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(colors))
	return Status(ret)
}

func GetPathGradientSurroundColorsWithCount(brush *PathGradient, color *ARGB, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientSurroundColorsWithCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(color)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func SetPathGradientSurroundColorsWithCount(brush *PathGradient, color *ARGB, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientSurroundColorsWithCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(color)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func GetPathGradientPath(brush *PathGradient, path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientPath.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func SetPathGradientPath(brush *PathGradient, path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientPath.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func GetPathGradientCenterPoint(brush *PathGradient, points *PointF) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientCenterPoint.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)))
	return Status(ret)
}

func GetPathGradientCenterPointI(brush *PathGradient, points *Point) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientCenterPointI.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)))
	return Status(ret)
}

func SetPathGradientCenterPoint(brush *PathGradient, points *PointF) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientCenterPoint.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)))
	return Status(ret)
}

func SetPathGradientCenterPointI(brush *PathGradient, points *Point) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientCenterPointI.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(points)))
	return Status(ret)
}

func GetPathGradientRect(brush *PathGradient, rect *RectF) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientRect.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func GetPathGradientRectI(brush *PathGradient, rect *Rect) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientRectI.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(rect)))
	return Status(ret)
}

func GetPathGradientPointCount(brush *PathGradient, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientPointCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func GetPathGradientSurroundColorCount(brush *PathGradient, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientSurroundColorCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func SetPathGradientGammaCorrection(brush *PathGradient, useGammaCorrection win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientGammaCorrection.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(useGammaCorrection))
	return Status(ret)
}

func GetPathGradientGammaCorrection(brush *PathGradient, useGammaCorrection *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientGammaCorrection.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(useGammaCorrection)))
	return Status(ret)
}

func GetPathGradientBlendCount(brush *PathGradient, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientBlendCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func GetPathGradientBlend(brush *PathGradient, blend *float32, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func SetPathGradientBlend(brush *PathGradient, blend *float32, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func GetPathGradientPresetBlendCount(brush *PathGradient, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientPresetBlendCount.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func GetPathGradientPresetBlend(brush *PathGradient, blend *ARGB, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientPresetBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func SetPathGradientPresetBlend(brush *PathGradient, blend *ARGB, positions *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientPresetBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(blend)),
		uintptr(unsafe.Pointer(positions)),
		uintptr(count))
	return Status(ret)
}

func SetPathGradientSigmaBlend(brush *PathGradient, focus float32, scale float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientSigmaBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(focus)),
		uintptr(math.Float32bits(scale)))
	return Status(ret)
}

func SetPathGradientLinearBlend(brush *PathGradient, focus float32, scale float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientLinearBlend.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(focus)),
		uintptr(math.Float32bits(scale)))
	return Status(ret)
}

func GetPathGradientWrapMode(brush *PathGradient, wrapmode *WrapMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientWrapMode.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(wrapmode)))
	return Status(ret)
}

func SetPathGradientWrapMode(brush *PathGradient, wrapmode WrapMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientWrapMode.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(wrapmode))
	return Status(ret)
}

func GetPathGradientTransform(brush *PathGradient, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func SetPathGradientTransform(brush *PathGradient, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func ResetPathGradientTransform(brush *PathGradient) Status {
	ret, _, _ := syscall.SyscallN(gdipResetPathGradientTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func MultiplyPathGradientTransform(brush *PathGradient, matrix *Matrix, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipMultiplyPathGradientTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
	return Status(ret)
}

func TranslatePathGradientTransform(brush *PathGradient, dx float32, dy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslatePathGradientTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)),
		uintptr(order))
	return Status(ret)
}

func ScalePathGradientTransform(brush *PathGradient, sx float32, sy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipScalePathGradientTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(sx)),
		uintptr(math.Float32bits(sy)),
		uintptr(order))
	return Status(ret)
}

func RotatePathGradientTransform(brush *PathGradient, angle float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipRotatePathGradientTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(angle)),
		uintptr(order))
	return Status(ret)
}

func GetPathGradientFocusScales(brush *PathGradient, xScale *float32, yScale *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathGradientFocusScales.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(xScale)),
		uintptr(unsafe.Pointer(yScale)))
	return Status(ret)
}

func SetPathGradientFocusScales(brush *PathGradient, xScale float32, yScale float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathGradientFocusScales.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(xScale)),
		uintptr(math.Float32bits(yScale)))
	return Status(ret)
}


	// solidbrush
func CreateSolidFill(color ARGB, brush **SolidFill) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateSolidFill.Addr(), 
		uintptr(color),
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func SetSolidFillColor(brush *SolidFill, color ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipSetSolidFillColor.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(color))
	return Status(ret)
}

func GetSolidFillColor(brush *SolidFill, color *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGetSolidFillColor.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(color)))
	return Status(ret)
}


	// texturebrush
func CreateTexture(image *Image, wrapmode WrapMode, texture **Texture) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateTexture.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(wrapmode),
		uintptr(unsafe.Pointer(texture)))
	return Status(ret)
}

func CreateTexture2(image *Image, wrapmode WrapMode, x float32, y float32, width float32, height float32, texture **Texture) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateTexture2.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(wrapmode),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(unsafe.Pointer(texture)))
	return Status(ret)
}

func CreateTextureIA(image *Image, imageAttributes *ImageAttributes, x float32, y float32, width float32, height float32, texture **Texture) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateTextureIA.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(unsafe.Pointer(texture)))
	return Status(ret)
}

func CreateTexture2I(image *Image, wrapmode WrapMode, x int32, y int32, width int32, height int32, texture **Texture) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateTexture2I.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(wrapmode),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(texture)))
	return Status(ret)
}

func CreateTextureIAI(image *Image, imageAttributes *ImageAttributes, x int32, y int32, width int32, height int32, texture **Texture) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateTextureIAI.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(imageAttributes)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(texture)))
	return Status(ret)
}

func GetTextureTransform(brush *Texture, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetTextureTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func SetTextureTransform(brush *Texture, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipSetTextureTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func ResetTextureTransform(brush *Texture) Status {
	ret, _, _ := syscall.SyscallN(gdipResetTextureTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func MultiplyTextureTransform(brush *Texture, matrix *Matrix, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipMultiplyTextureTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
	return Status(ret)
}

func TranslateTextureTransform(brush *Texture, dx float32, dy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateTextureTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)),
		uintptr(order))
	return Status(ret)
}

func ScaleTextureTransform(brush *Texture, sx float32, sy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipScaleTextureTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(sx)),
		uintptr(math.Float32bits(sy)),
		uintptr(order))
	return Status(ret)
}

func RotateTextureTransform(brush *Texture, angle float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipRotateTextureTransform.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(angle)),
		uintptr(order))
	return Status(ret)
}

func SetTextureWrapMode(brush *Texture, wrapmode WrapMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetTextureWrapMode.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(wrapmode))
	return Status(ret)
}

func GetTextureWrapMode(brush *Texture, wrapmode *WrapMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetTextureWrapMode.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(wrapmode)))
	return Status(ret)
}

func GetTextureImage(brush *Texture, image **Image) Status {
	ret, _, _ := syscall.SyscallN(gdipGetTextureImage.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(image)))
	return Status(ret)
}

