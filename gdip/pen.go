package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// adjustablearrowcap
	gdipCreateAdjustableArrowCap	*windows.LazyProc
	gdipSetAdjustableArrowCapHeight	*windows.LazyProc
	gdipGetAdjustableArrowCapHeight	*windows.LazyProc
	gdipSetAdjustableArrowCapWidth	*windows.LazyProc
	gdipGetAdjustableArrowCapWidth	*windows.LazyProc
	gdipSetAdjustableArrowCapMiddleInset	*windows.LazyProc
	gdipGetAdjustableArrowCapMiddleInset	*windows.LazyProc
	gdipSetAdjustableArrowCapFillState	*windows.LazyProc
	gdipGetAdjustableArrowCapFillState	*windows.LazyProc

	// customlinecap
	gdipCreateCustomLineCap	*windows.LazyProc
	gdipDeleteCustomLineCap	*windows.LazyProc
	gdipCloneCustomLineCap	*windows.LazyProc
	gdipGetCustomLineCapType	*windows.LazyProc
	gdipSetCustomLineCapStrokeCaps	*windows.LazyProc
	gdipGetCustomLineCapStrokeCaps	*windows.LazyProc
	gdipSetCustomLineCapStrokeJoin	*windows.LazyProc
	gdipGetCustomLineCapStrokeJoin	*windows.LazyProc
	gdipSetCustomLineCapBaseCap	*windows.LazyProc
	gdipGetCustomLineCapBaseCap	*windows.LazyProc
	gdipSetCustomLineCapBaseInset	*windows.LazyProc
	gdipGetCustomLineCapBaseInset	*windows.LazyProc
	gdipSetCustomLineCapWidthScale	*windows.LazyProc
	gdipGetCustomLineCapWidthScale	*windows.LazyProc

	// pen
	gdipCreatePen1	*windows.LazyProc
	gdipCreatePen2	*windows.LazyProc
	gdipClonePen	*windows.LazyProc
	gdipDeletePen	*windows.LazyProc
	gdipSetPenWidth	*windows.LazyProc
	gdipGetPenWidth	*windows.LazyProc
	gdipSetPenUnit	*windows.LazyProc
	gdipGetPenUnit	*windows.LazyProc
	gdipSetPenLineCap197819	*windows.LazyProc
	gdipSetPenStartCap	*windows.LazyProc
	gdipSetPenEndCap	*windows.LazyProc
	gdipSetPenDashCap197819	*windows.LazyProc
	gdipGetPenStartCap	*windows.LazyProc
	gdipGetPenEndCap	*windows.LazyProc
	gdipGetPenDashCap197819	*windows.LazyProc
	gdipSetPenLineJoin	*windows.LazyProc
	gdipGetPenLineJoin	*windows.LazyProc
	gdipSetPenCustomStartCap	*windows.LazyProc
	gdipGetPenCustomStartCap	*windows.LazyProc
	gdipSetPenCustomEndCap	*windows.LazyProc
	gdipGetPenCustomEndCap	*windows.LazyProc
	gdipSetPenMiterLimit	*windows.LazyProc
	gdipGetPenMiterLimit	*windows.LazyProc
	gdipSetPenMode	*windows.LazyProc
	gdipGetPenMode	*windows.LazyProc
	gdipSetPenTransform	*windows.LazyProc
	gdipGetPenTransform	*windows.LazyProc
	gdipResetPenTransform	*windows.LazyProc
	gdipMultiplyPenTransform	*windows.LazyProc
	gdipTranslatePenTransform	*windows.LazyProc
	gdipScalePenTransform	*windows.LazyProc
	gdipRotatePenTransform	*windows.LazyProc
	gdipSetPenColor	*windows.LazyProc
	gdipGetPenColor	*windows.LazyProc
	gdipSetPenBrushFill	*windows.LazyProc
	gdipGetPenBrushFill	*windows.LazyProc
	gdipGetPenFillType	*windows.LazyProc
	gdipGetPenDashStyle	*windows.LazyProc
	gdipSetPenDashStyle	*windows.LazyProc
	gdipGetPenDashOffset	*windows.LazyProc
	gdipSetPenDashOffset	*windows.LazyProc
	gdipGetPenDashCount	*windows.LazyProc
	gdipSetPenDashArray	*windows.LazyProc
	gdipGetPenDashArray	*windows.LazyProc
	gdipGetPenCompoundCount	*windows.LazyProc
	gdipSetPenCompoundArray	*windows.LazyProc
	gdipGetPenCompoundArray	*windows.LazyProc
)

func init() {

	// adjustablearrowcap
	gdipCreateAdjustableArrowCap = dll.NewProc("GdipCreateAdjustableArrowCap")
	gdipSetAdjustableArrowCapHeight = dll.NewProc("GdipSetAdjustableArrowCapHeight")
	gdipGetAdjustableArrowCapHeight = dll.NewProc("GdipGetAdjustableArrowCapHeight")
	gdipSetAdjustableArrowCapWidth = dll.NewProc("GdipSetAdjustableArrowCapWidth")
	gdipGetAdjustableArrowCapWidth = dll.NewProc("GdipGetAdjustableArrowCapWidth")
	gdipSetAdjustableArrowCapMiddleInset = dll.NewProc("GdipSetAdjustableArrowCapMiddleInset")
	gdipGetAdjustableArrowCapMiddleInset = dll.NewProc("GdipGetAdjustableArrowCapMiddleInset")
	gdipSetAdjustableArrowCapFillState = dll.NewProc("GdipSetAdjustableArrowCapFillState")
	gdipGetAdjustableArrowCapFillState = dll.NewProc("GdipGetAdjustableArrowCapFillState")

	// customlinecap
	gdipCreateCustomLineCap = dll.NewProc("GdipCreateCustomLineCap")
	gdipDeleteCustomLineCap = dll.NewProc("GdipDeleteCustomLineCap")
	gdipCloneCustomLineCap = dll.NewProc("GdipCloneCustomLineCap")
	gdipGetCustomLineCapType = dll.NewProc("GdipGetCustomLineCapType")
	gdipSetCustomLineCapStrokeCaps = dll.NewProc("GdipSetCustomLineCapStrokeCaps")
	gdipGetCustomLineCapStrokeCaps = dll.NewProc("GdipGetCustomLineCapStrokeCaps")
	gdipSetCustomLineCapStrokeJoin = dll.NewProc("GdipSetCustomLineCapStrokeJoin")
	gdipGetCustomLineCapStrokeJoin = dll.NewProc("GdipGetCustomLineCapStrokeJoin")
	gdipSetCustomLineCapBaseCap = dll.NewProc("GdipSetCustomLineCapBaseCap")
	gdipGetCustomLineCapBaseCap = dll.NewProc("GdipGetCustomLineCapBaseCap")
	gdipSetCustomLineCapBaseInset = dll.NewProc("GdipSetCustomLineCapBaseInset")
	gdipGetCustomLineCapBaseInset = dll.NewProc("GdipGetCustomLineCapBaseInset")
	gdipSetCustomLineCapWidthScale = dll.NewProc("GdipSetCustomLineCapWidthScale")
	gdipGetCustomLineCapWidthScale = dll.NewProc("GdipGetCustomLineCapWidthScale")

	// pen
	gdipCreatePen1 = dll.NewProc("GdipCreatePen1")
	gdipCreatePen2 = dll.NewProc("GdipCreatePen2")
	gdipClonePen = dll.NewProc("GdipClonePen")
	gdipDeletePen = dll.NewProc("GdipDeletePen")
	gdipSetPenWidth = dll.NewProc("GdipSetPenWidth")
	gdipGetPenWidth = dll.NewProc("GdipGetPenWidth")
	gdipSetPenUnit = dll.NewProc("GdipSetPenUnit")
	gdipGetPenUnit = dll.NewProc("GdipGetPenUnit")
	gdipSetPenLineCap197819 = dll.NewProc("GdipSetPenLineCap197819")
	gdipSetPenStartCap = dll.NewProc("GdipSetPenStartCap")
	gdipSetPenEndCap = dll.NewProc("GdipSetPenEndCap")
	gdipSetPenDashCap197819 = dll.NewProc("GdipSetPenDashCap197819")
	gdipGetPenStartCap = dll.NewProc("GdipGetPenStartCap")
	gdipGetPenEndCap = dll.NewProc("GdipGetPenEndCap")
	gdipGetPenDashCap197819 = dll.NewProc("GdipGetPenDashCap197819")
	gdipSetPenLineJoin = dll.NewProc("GdipSetPenLineJoin")
	gdipGetPenLineJoin = dll.NewProc("GdipGetPenLineJoin")
	gdipSetPenCustomStartCap = dll.NewProc("GdipSetPenCustomStartCap")
	gdipGetPenCustomStartCap = dll.NewProc("GdipGetPenCustomStartCap")
	gdipSetPenCustomEndCap = dll.NewProc("GdipSetPenCustomEndCap")
	gdipGetPenCustomEndCap = dll.NewProc("GdipGetPenCustomEndCap")
	gdipSetPenMiterLimit = dll.NewProc("GdipSetPenMiterLimit")
	gdipGetPenMiterLimit = dll.NewProc("GdipGetPenMiterLimit")
	gdipSetPenMode = dll.NewProc("GdipSetPenMode")
	gdipGetPenMode = dll.NewProc("GdipGetPenMode")
	gdipSetPenTransform = dll.NewProc("GdipSetPenTransform")
	gdipGetPenTransform = dll.NewProc("GdipGetPenTransform")
	gdipResetPenTransform = dll.NewProc("GdipResetPenTransform")
	gdipMultiplyPenTransform = dll.NewProc("GdipMultiplyPenTransform")
	gdipTranslatePenTransform = dll.NewProc("GdipTranslatePenTransform")
	gdipScalePenTransform = dll.NewProc("GdipScalePenTransform")
	gdipRotatePenTransform = dll.NewProc("GdipRotatePenTransform")
	gdipSetPenColor = dll.NewProc("GdipSetPenColor")
	gdipGetPenColor = dll.NewProc("GdipGetPenColor")
	gdipSetPenBrushFill = dll.NewProc("GdipSetPenBrushFill")
	gdipGetPenBrushFill = dll.NewProc("GdipGetPenBrushFill")
	gdipGetPenFillType = dll.NewProc("GdipGetPenFillType")
	gdipGetPenDashStyle = dll.NewProc("GdipGetPenDashStyle")
	gdipSetPenDashStyle = dll.NewProc("GdipSetPenDashStyle")
	gdipGetPenDashOffset = dll.NewProc("GdipGetPenDashOffset")
	gdipSetPenDashOffset = dll.NewProc("GdipSetPenDashOffset")
	gdipGetPenDashCount = dll.NewProc("GdipGetPenDashCount")
	gdipSetPenDashArray = dll.NewProc("GdipSetPenDashArray")
	gdipGetPenDashArray = dll.NewProc("GdipGetPenDashArray")
	gdipGetPenCompoundCount = dll.NewProc("GdipGetPenCompoundCount")
	gdipSetPenCompoundArray = dll.NewProc("GdipSetPenCompoundArray")
	gdipGetPenCompoundArray = dll.NewProc("GdipGetPenCompoundArray")
}

	// adjustablearrowcap
func CreateAdjustableArrowCap(height float32, width float32, isFilled win32.BOOL, cap **AdjustableArrowCap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateAdjustableArrowCap.Addr(), 
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(width)),
		uintptr(isFilled),
		uintptr(unsafe.Pointer(cap)))
	return Status(ret)
}

func SetAdjustableArrowCapHeight(cap *AdjustableArrowCap, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetAdjustableArrowCapHeight.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func GetAdjustableArrowCapHeight(cap *AdjustableArrowCap, height *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetAdjustableArrowCapHeight.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(unsafe.Pointer(height)))
	return Status(ret)
}

func SetAdjustableArrowCapWidth(cap *AdjustableArrowCap, width float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetAdjustableArrowCapWidth.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(math.Float32bits(width)))
	return Status(ret)
}

func GetAdjustableArrowCapWidth(cap *AdjustableArrowCap, width *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetAdjustableArrowCapWidth.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(unsafe.Pointer(width)))
	return Status(ret)
}

func SetAdjustableArrowCapMiddleInset(cap *AdjustableArrowCap, middleInset float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetAdjustableArrowCapMiddleInset.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(math.Float32bits(middleInset)))
	return Status(ret)
}

func GetAdjustableArrowCapMiddleInset(cap *AdjustableArrowCap, middleInset *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetAdjustableArrowCapMiddleInset.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(unsafe.Pointer(middleInset)))
	return Status(ret)
}

func SetAdjustableArrowCapFillState(cap *AdjustableArrowCap, fillState win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipSetAdjustableArrowCapFillState.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(fillState))
	return Status(ret)
}

func GetAdjustableArrowCapFillState(cap *AdjustableArrowCap, fillState *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipGetAdjustableArrowCapFillState.Addr(), 
		uintptr(unsafe.Pointer(cap)),
		uintptr(unsafe.Pointer(fillState)))
	return Status(ret)
}


	// customlinecap
func CreateCustomLineCap(fillPath *Path, strokePath *Path, baseCap LineCap, baseInset float32, customCap **CustomLineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateCustomLineCap.Addr(), 
		uintptr(unsafe.Pointer(fillPath)),
		uintptr(unsafe.Pointer(strokePath)),
		uintptr(baseCap),
		uintptr(math.Float32bits(baseInset)),
		uintptr(unsafe.Pointer(customCap)))
	return Status(ret)
}

func DeleteCustomLineCap(customCap *CustomLineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteCustomLineCap.Addr(), 
		uintptr(unsafe.Pointer(customCap)))
	return Status(ret)
}

func CloneCustomLineCap(customCap *CustomLineCap, clonedCap **CustomLineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneCustomLineCap.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(clonedCap)))
	return Status(ret)
}

func GetCustomLineCapType(customCap *CustomLineCap, capType *CustomLineCapType) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCustomLineCapType.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(capType)))
	return Status(ret)
}

func SetCustomLineCapStrokeCaps(customCap *CustomLineCap, startCap LineCap, endCap LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetCustomLineCapStrokeCaps.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(startCap),
		uintptr(endCap))
	return Status(ret)
}

func GetCustomLineCapStrokeCaps(customCap *CustomLineCap, startCap *LineCap, endCap *LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCustomLineCapStrokeCaps.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(startCap)),
		uintptr(unsafe.Pointer(endCap)))
	return Status(ret)
}

func SetCustomLineCapStrokeJoin(customCap *CustomLineCap, lineJoin LineJoin) Status {
	ret, _, _ := syscall.SyscallN(gdipSetCustomLineCapStrokeJoin.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(lineJoin))
	return Status(ret)
}

func GetCustomLineCapStrokeJoin(customCap *CustomLineCap, lineJoin *LineJoin) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCustomLineCapStrokeJoin.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(lineJoin)))
	return Status(ret)
}

func SetCustomLineCapBaseCap(customCap *CustomLineCap, baseCap LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetCustomLineCapBaseCap.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(baseCap))
	return Status(ret)
}

func GetCustomLineCapBaseCap(customCap *CustomLineCap, baseCap *LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCustomLineCapBaseCap.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(baseCap)))
	return Status(ret)
}

func SetCustomLineCapBaseInset(customCap *CustomLineCap, inset float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetCustomLineCapBaseInset.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(math.Float32bits(inset)))
	return Status(ret)
}

func GetCustomLineCapBaseInset(customCap *CustomLineCap, inset *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCustomLineCapBaseInset.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(inset)))
	return Status(ret)
}

func SetCustomLineCapWidthScale(customCap *CustomLineCap, widthScale float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetCustomLineCapWidthScale.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(math.Float32bits(widthScale)))
	return Status(ret)
}

func GetCustomLineCapWidthScale(customCap *CustomLineCap, widthScale *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetCustomLineCapWidthScale.Addr(), 
		uintptr(unsafe.Pointer(customCap)),
		uintptr(unsafe.Pointer(widthScale)))
	return Status(ret)
}


	// pen
func CreatePen1(color ARGB, width float32, unit Unit, pen **Pen) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePen1.Addr(), 
		uintptr(color),
		uintptr(math.Float32bits(width)),
		uintptr(unit),
		uintptr(unsafe.Pointer(pen)))
	return Status(ret)
}

func CreatePen2(brush *Brush, width float32, unit Unit, pen **Pen) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePen2.Addr(), 
		uintptr(unsafe.Pointer(brush)),
		uintptr(math.Float32bits(width)),
		uintptr(unit),
		uintptr(unsafe.Pointer(pen)))
	return Status(ret)
}

func ClonePen(pen *Pen, clonepen **Pen) Status {
	ret, _, _ := syscall.SyscallN(gdipClonePen.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(clonepen)))
	return Status(ret)
}

func DeletePen(pen *Pen) Status {
	ret, _, _ := syscall.SyscallN(gdipDeletePen.Addr(), 
		uintptr(unsafe.Pointer(pen)))
	return Status(ret)
}

func SetPenWidth(pen *Pen, width float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenWidth.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(width)))
	return Status(ret)
}

func GetPenWidth(pen *Pen, width *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenWidth.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(width)))
	return Status(ret)
}

func SetPenUnit(pen *Pen, unit Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenUnit.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unit))
	return Status(ret)
}

func GetPenUnit(pen *Pen, unit *Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenUnit.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(unit)))
	return Status(ret)
}

func SetPenLineCap197819(pen *Pen, startCap LineCap, endCap LineCap, dashCap DashCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenLineCap197819.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(startCap),
		uintptr(endCap),
		uintptr(dashCap))
	return Status(ret)
}

func SetPenStartCap(pen *Pen, startCap LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenStartCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(startCap))
	return Status(ret)
}

func SetPenEndCap(pen *Pen, endCap LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenEndCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(endCap))
	return Status(ret)
}

func SetPenDashCap197819(pen *Pen, dashCap DashCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenDashCap197819.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(dashCap))
	return Status(ret)
}

func GetPenStartCap(pen *Pen, startCap *LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenStartCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(startCap)))
	return Status(ret)
}

func GetPenEndCap(pen *Pen, endCap *LineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenEndCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(endCap)))
	return Status(ret)
}

func GetPenDashCap197819(pen *Pen, dashCap *DashCap) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenDashCap197819.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dashCap)))
	return Status(ret)
}

func SetPenLineJoin(pen *Pen, lineJoin LineJoin) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenLineJoin.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(lineJoin))
	return Status(ret)
}

func GetPenLineJoin(pen *Pen, lineJoin *LineJoin) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenLineJoin.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(lineJoin)))
	return Status(ret)
}

func SetPenCustomStartCap(pen *Pen, customCap *CustomLineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenCustomStartCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return Status(ret)
}

func GetPenCustomStartCap(pen *Pen, customCap **CustomLineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenCustomStartCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return Status(ret)
}

func SetPenCustomEndCap(pen *Pen, customCap *CustomLineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenCustomEndCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return Status(ret)
}

func GetPenCustomEndCap(pen *Pen, customCap **CustomLineCap) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenCustomEndCap.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(customCap)))
	return Status(ret)
}

func SetPenMiterLimit(pen *Pen, miterLimit float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenMiterLimit.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(miterLimit)))
	return Status(ret)
}

func GetPenMiterLimit(pen *Pen, miterLimit *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenMiterLimit.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(miterLimit)))
	return Status(ret)
}

func SetPenMode(pen *Pen, penMode PenAlignment) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenMode.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(penMode))
	return Status(ret)
}

func GetPenMode(pen *Pen, penMode *PenAlignment) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenMode.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(penMode)))
	return Status(ret)
}

func SetPenTransform(pen *Pen, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenTransform.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func GetPenTransform(pen *Pen, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenTransform.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func ResetPenTransform(pen *Pen) Status {
	ret, _, _ := syscall.SyscallN(gdipResetPenTransform.Addr(), 
		uintptr(unsafe.Pointer(pen)))
	return Status(ret)
}

func MultiplyPenTransform(pen *Pen, matrix *Matrix, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipMultiplyPenTransform.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(order))
	return Status(ret)
}

func TranslatePenTransform(pen *Pen, dx float32, dy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslatePenTransform.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)),
		uintptr(order))
	return Status(ret)
}

func ScalePenTransform(pen *Pen, sx float32, sy float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipScalePenTransform.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(sx)),
		uintptr(math.Float32bits(sy)),
		uintptr(order))
	return Status(ret)
}

func RotatePenTransform(pen *Pen, angle float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipRotatePenTransform.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(angle)),
		uintptr(order))
	return Status(ret)
}

func SetPenColor(pen *Pen, argb ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenColor.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(argb))
	return Status(ret)
}

func GetPenColor(pen *Pen, argb *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenColor.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(argb)))
	return Status(ret)
}

func SetPenBrushFill(pen *Pen, brush *Brush) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenBrushFill.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func GetPenBrushFill(pen *Pen, brush **Brush) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenBrushFill.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(brush)))
	return Status(ret)
}

func GetPenFillType(pen *Pen, typ *PenType) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenFillType.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(typ)))
	return Status(ret)
}

func GetPenDashStyle(pen *Pen, dashstyle *DashStyle) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenDashStyle.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dashstyle)))
	return Status(ret)
}

func SetPenDashStyle(pen *Pen, dashstyle DashStyle) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenDashStyle.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(dashstyle))
	return Status(ret)
}

func GetPenDashOffset(pen *Pen, offset *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenDashOffset.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(offset)))
	return Status(ret)
}

func SetPenDashOffset(pen *Pen, offset float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenDashOffset.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(math.Float32bits(offset)))
	return Status(ret)
}

func GetPenDashCount(pen *Pen, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenDashCount.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func SetPenDashArray(pen *Pen, dash *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenDashArray.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return Status(ret)
}

func GetPenDashArray(pen *Pen, dash *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenDashArray.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return Status(ret)
}

func GetPenCompoundCount(pen *Pen, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenCompoundCount.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func SetPenCompoundArray(pen *Pen, dash *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPenCompoundArray.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return Status(ret)
}

func GetPenCompoundArray(pen *Pen, dash *float32, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPenCompoundArray.Addr(), 
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(dash)),
		uintptr(count))
	return Status(ret)
}

