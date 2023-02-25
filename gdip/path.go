package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// graphicspath
	gdipCreatePath	*windows.LazyProc
	gdipCreatePath2	*windows.LazyProc
	gdipCreatePath2I	*windows.LazyProc
	gdipClonePath	*windows.LazyProc
	gdipDeletePath	*windows.LazyProc
	gdipResetPath	*windows.LazyProc
	gdipGetPointCount	*windows.LazyProc
	gdipGetPathTypes	*windows.LazyProc
	gdipGetPathPoints	*windows.LazyProc
	gdipGetPathPointsI	*windows.LazyProc
	gdipGetPathFillMode	*windows.LazyProc
	gdipSetPathFillMode	*windows.LazyProc
	gdipGetPathData	*windows.LazyProc
	gdipStartPathFigure	*windows.LazyProc
	gdipClosePathFigure	*windows.LazyProc
	gdipClosePathFigures	*windows.LazyProc
	gdipSetPathMarker	*windows.LazyProc
	gdipClearPathMarkers	*windows.LazyProc
	gdipReversePath	*windows.LazyProc
	gdipGetPathLastPoint	*windows.LazyProc
	gdipAddPathLine	*windows.LazyProc
	gdipAddPathLine2	*windows.LazyProc
	gdipAddPathArc	*windows.LazyProc
	gdipAddPathBezier	*windows.LazyProc
	gdipAddPathBeziers	*windows.LazyProc
	gdipAddPathCurve	*windows.LazyProc
	gdipAddPathCurve2	*windows.LazyProc
	gdipAddPathCurve3	*windows.LazyProc
	gdipAddPathClosedCurve	*windows.LazyProc
	gdipAddPathClosedCurve2	*windows.LazyProc
	gdipAddPathRectangle	*windows.LazyProc
	gdipAddPathRectangles	*windows.LazyProc
	gdipAddPathEllipse	*windows.LazyProc
	gdipAddPathPie	*windows.LazyProc
	gdipAddPathPolygon	*windows.LazyProc
	gdipAddPathPath	*windows.LazyProc
	gdipAddPathString	*windows.LazyProc
	gdipAddPathStringI	*windows.LazyProc
	gdipAddPathLineI	*windows.LazyProc
	gdipAddPathLine2I	*windows.LazyProc
	gdipAddPathArcI	*windows.LazyProc
	gdipAddPathBezierI	*windows.LazyProc
	gdipAddPathBeziersI	*windows.LazyProc
	gdipAddPathCurveI	*windows.LazyProc
	gdipAddPathCurve2I	*windows.LazyProc
	gdipAddPathCurve3I	*windows.LazyProc
	gdipAddPathClosedCurveI	*windows.LazyProc
	gdipAddPathClosedCurve2I	*windows.LazyProc
	gdipAddPathRectangleI	*windows.LazyProc
	gdipAddPathRectanglesI	*windows.LazyProc
	gdipAddPathEllipseI	*windows.LazyProc
	gdipAddPathPieI	*windows.LazyProc
	gdipAddPathPolygonI	*windows.LazyProc
	gdipFlattenPath	*windows.LazyProc
	gdipWindingModeOutline	*windows.LazyProc
	gdipWidenPath	*windows.LazyProc
	gdipWarpPath	*windows.LazyProc
	gdipTransformPath	*windows.LazyProc
	gdipGetPathWorldBounds	*windows.LazyProc
	gdipGetPathWorldBoundsI	*windows.LazyProc
	gdipIsVisiblePathPoint	*windows.LazyProc
	gdipIsVisiblePathPointI	*windows.LazyProc
	gdipIsOutlineVisiblePathPoint	*windows.LazyProc
	gdipIsOutlineVisiblePathPointI	*windows.LazyProc

	// pathiterator
	gdipCreatePathIter	*windows.LazyProc
	gdipDeletePathIter	*windows.LazyProc
	gdipPathIterNextSubpath	*windows.LazyProc
	gdipPathIterNextSubpathPath	*windows.LazyProc
	gdipPathIterNextPathType	*windows.LazyProc
	gdipPathIterNextMarker	*windows.LazyProc
	gdipPathIterNextMarkerPath	*windows.LazyProc
	gdipPathIterGetCount	*windows.LazyProc
	gdipPathIterGetSubpathCount	*windows.LazyProc
	gdipPathIterIsValid	*windows.LazyProc
	gdipPathIterHasCurve	*windows.LazyProc
	gdipPathIterRewind	*windows.LazyProc
	gdipPathIterEnumerate	*windows.LazyProc
	gdipPathIterCopyData	*windows.LazyProc
)

func init() {

	// graphicspath
	gdipCreatePath = dll.NewProc("GdipCreatePath")
	gdipCreatePath2 = dll.NewProc("GdipCreatePath2")
	gdipCreatePath2I = dll.NewProc("GdipCreatePath2I")
	gdipClonePath = dll.NewProc("GdipClonePath")
	gdipDeletePath = dll.NewProc("GdipDeletePath")
	gdipResetPath = dll.NewProc("GdipResetPath")
	gdipGetPointCount = dll.NewProc("GdipGetPointCount")
	gdipGetPathTypes = dll.NewProc("GdipGetPathTypes")
	gdipGetPathPoints = dll.NewProc("GdipGetPathPoints")
	gdipGetPathPointsI = dll.NewProc("GdipGetPathPointsI")
	gdipGetPathFillMode = dll.NewProc("GdipGetPathFillMode")
	gdipSetPathFillMode = dll.NewProc("GdipSetPathFillMode")
	gdipGetPathData = dll.NewProc("GdipGetPathData")
	gdipStartPathFigure = dll.NewProc("GdipStartPathFigure")
	gdipClosePathFigure = dll.NewProc("GdipClosePathFigure")
	gdipClosePathFigures = dll.NewProc("GdipClosePathFigures")
	gdipSetPathMarker = dll.NewProc("GdipSetPathMarker")
	gdipClearPathMarkers = dll.NewProc("GdipClearPathMarkers")
	gdipReversePath = dll.NewProc("GdipReversePath")
	gdipGetPathLastPoint = dll.NewProc("GdipGetPathLastPoint")
	gdipAddPathLine = dll.NewProc("GdipAddPathLine")
	gdipAddPathLine2 = dll.NewProc("GdipAddPathLine2")
	gdipAddPathArc = dll.NewProc("GdipAddPathArc")
	gdipAddPathBezier = dll.NewProc("GdipAddPathBezier")
	gdipAddPathBeziers = dll.NewProc("GdipAddPathBeziers")
	gdipAddPathCurve = dll.NewProc("GdipAddPathCurve")
	gdipAddPathCurve2 = dll.NewProc("GdipAddPathCurve2")
	gdipAddPathCurve3 = dll.NewProc("GdipAddPathCurve3")
	gdipAddPathClosedCurve = dll.NewProc("GdipAddPathClosedCurve")
	gdipAddPathClosedCurve2 = dll.NewProc("GdipAddPathClosedCurve2")
	gdipAddPathRectangle = dll.NewProc("GdipAddPathRectangle")
	gdipAddPathRectangles = dll.NewProc("GdipAddPathRectangles")
	gdipAddPathEllipse = dll.NewProc("GdipAddPathEllipse")
	gdipAddPathPie = dll.NewProc("GdipAddPathPie")
	gdipAddPathPolygon = dll.NewProc("GdipAddPathPolygon")
	gdipAddPathPath = dll.NewProc("GdipAddPathPath")
	gdipAddPathString = dll.NewProc("GdipAddPathString")
	gdipAddPathStringI = dll.NewProc("GdipAddPathStringI")
	gdipAddPathLineI = dll.NewProc("GdipAddPathLineI")
	gdipAddPathLine2I = dll.NewProc("GdipAddPathLine2I")
	gdipAddPathArcI = dll.NewProc("GdipAddPathArcI")
	gdipAddPathBezierI = dll.NewProc("GdipAddPathBezierI")
	gdipAddPathBeziersI = dll.NewProc("GdipAddPathBeziersI")
	gdipAddPathCurveI = dll.NewProc("GdipAddPathCurveI")
	gdipAddPathCurve2I = dll.NewProc("GdipAddPathCurve2I")
	gdipAddPathCurve3I = dll.NewProc("GdipAddPathCurve3I")
	gdipAddPathClosedCurveI = dll.NewProc("GdipAddPathClosedCurveI")
	gdipAddPathClosedCurve2I = dll.NewProc("GdipAddPathClosedCurve2I")
	gdipAddPathRectangleI = dll.NewProc("GdipAddPathRectangleI")
	gdipAddPathRectanglesI = dll.NewProc("GdipAddPathRectanglesI")
	gdipAddPathEllipseI = dll.NewProc("GdipAddPathEllipseI")
	gdipAddPathPieI = dll.NewProc("GdipAddPathPieI")
	gdipAddPathPolygonI = dll.NewProc("GdipAddPathPolygonI")
	gdipFlattenPath = dll.NewProc("GdipFlattenPath")
	gdipWindingModeOutline = dll.NewProc("GdipWindingModeOutline")
	gdipWidenPath = dll.NewProc("GdipWidenPath")
	gdipWarpPath = dll.NewProc("GdipWarpPath")
	gdipTransformPath = dll.NewProc("GdipTransformPath")
	gdipGetPathWorldBounds = dll.NewProc("GdipGetPathWorldBounds")
	gdipGetPathWorldBoundsI = dll.NewProc("GdipGetPathWorldBoundsI")
	gdipIsVisiblePathPoint = dll.NewProc("GdipIsVisiblePathPoint")
	gdipIsVisiblePathPointI = dll.NewProc("GdipIsVisiblePathPointI")
	gdipIsOutlineVisiblePathPoint = dll.NewProc("GdipIsOutlineVisiblePathPoint")
	gdipIsOutlineVisiblePathPointI = dll.NewProc("GdipIsOutlineVisiblePathPointI")

	// pathiterator
	gdipCreatePathIter = dll.NewProc("GdipCreatePathIter")
	gdipDeletePathIter = dll.NewProc("GdipDeletePathIter")
	gdipPathIterNextSubpath = dll.NewProc("GdipPathIterNextSubpath")
	gdipPathIterNextSubpathPath = dll.NewProc("GdipPathIterNextSubpathPath")
	gdipPathIterNextPathType = dll.NewProc("GdipPathIterNextPathType")
	gdipPathIterNextMarker = dll.NewProc("GdipPathIterNextMarker")
	gdipPathIterNextMarkerPath = dll.NewProc("GdipPathIterNextMarkerPath")
	gdipPathIterGetCount = dll.NewProc("GdipPathIterGetCount")
	gdipPathIterGetSubpathCount = dll.NewProc("GdipPathIterGetSubpathCount")
	gdipPathIterIsValid = dll.NewProc("GdipPathIterIsValid")
	gdipPathIterHasCurve = dll.NewProc("GdipPathIterHasCurve")
	gdipPathIterRewind = dll.NewProc("GdipPathIterRewind")
	gdipPathIterEnumerate = dll.NewProc("GdipPathIterEnumerate")
	gdipPathIterCopyData = dll.NewProc("GdipPathIterCopyData")
}

	// graphicspath
func CreatePath(brushMode FillMode, path **Path) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePath.Addr(), 
		uintptr(brushMode),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func CreatePath2(points *PointF, types *win32.BYTE, count int32, fillMode FillMode, path **Path) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePath2.Addr(), 
		uintptr(unsafe.Pointer(points)),
		uintptr(unsafe.Pointer(types)),
		uintptr(count),
		uintptr(fillMode),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func CreatePath2I(points *Point, types *win32.BYTE, count int32, fillMode FillMode, path **Path) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePath2I.Addr(), 
		uintptr(unsafe.Pointer(points)),
		uintptr(unsafe.Pointer(types)),
		uintptr(count),
		uintptr(fillMode),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func ClonePath(path *Path, clonePath **Path) Status {
	ret, _, _ := syscall.SyscallN(gdipClonePath.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(clonePath)))
	return Status(ret)
}

func DeletePath(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipDeletePath.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func ResetPath(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipResetPath.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func GetPointCount(path *Path, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPointCount.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func GetPathTypes(path *Path, types *win32.BYTE, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathTypes.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(types)),
		uintptr(count))
	return Status(ret)
}

func GetPathPoints(p0 *Path, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathPoints.Addr(), 
		uintptr(unsafe.Pointer(p0)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func GetPathPointsI(p0 *Path, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathPointsI.Addr(), 
		uintptr(unsafe.Pointer(p0)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func GetPathFillMode(path *Path, fillmode *FillMode) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathFillMode.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(fillmode)))
	return Status(ret)
}

func SetPathFillMode(path *Path, fillmode FillMode) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathFillMode.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(fillmode))
	return Status(ret)
}

func GetPathData(path *Path, pathData *PathData) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathData.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(pathData)))
	return Status(ret)
}

func StartPathFigure(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipStartPathFigure.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func ClosePathFigure(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipClosePathFigure.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func ClosePathFigures(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipClosePathFigures.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func SetPathMarker(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPathMarker.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func ClearPathMarkers(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipClearPathMarkers.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func ReversePath(path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipReversePath.Addr(), 
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func GetPathLastPoint(path *Path, lastPoint *PointF) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathLastPoint.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(lastPoint)))
	return Status(ret)
}

func AddPathLine(path *Path, x1 float32, y1 float32, x2 float32, y2 float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathLine.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x1)),
		uintptr(math.Float32bits(y1)),
		uintptr(math.Float32bits(x2)),
		uintptr(math.Float32bits(y2)))
	return Status(ret)
}

func AddPathLine2(path *Path, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathLine2.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathArc(path *Path, x float32, y float32, width float32, height float32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathArc.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func AddPathBezier(path *Path, x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32, x4 float32, y4 float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathBezier.Addr(), 
		uintptr(unsafe.Pointer(path)),
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

func AddPathBeziers(path *Path, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathBeziers.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathCurve(path *Path, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathCurve.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathCurve2(path *Path, points *PointF, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathCurve2.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func AddPathCurve3(path *Path, points *PointF, count int32, offset int32, numberOfSegments int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathCurve3.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(offset),
		uintptr(numberOfSegments),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func AddPathClosedCurve(path *Path, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathClosedCurve.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathClosedCurve2(path *Path, points *PointF, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathClosedCurve2.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func AddPathRectangle(path *Path, x float32, y float32, width float32, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathRectangle.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func AddPathRectangles(path *Path, rects *RectF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathRectangles.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count))
	return Status(ret)
}

func AddPathEllipse(path *Path, x float32, y float32, width float32, height float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathEllipse.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)))
	return Status(ret)
}

func AddPathPie(path *Path, x float32, y float32, width float32, height float32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathPie.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func AddPathPolygon(path *Path, points *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathPolygon.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathPath(path *Path, addingPath *Path, connect win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathPath.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(addingPath)),
		uintptr(connect))
	return Status(ret)
}

func AddPathString(path *Path, string *uint16, length int32, family *FontFamily, style FontStyle, emSize float32, layoutRect *RectF, format *StringFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathString.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(string)),
		uintptr(length),
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(math.Float32bits(emSize)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func AddPathStringI(path *Path, string *uint16, length int32, family *FontFamily, style FontStyle, emSize float32, layoutRect *Rect, format *StringFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathStringI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(string)),
		uintptr(length),
		uintptr(unsafe.Pointer(family)),
		uintptr(style),
		uintptr(math.Float32bits(emSize)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func AddPathLineI(path *Path, x1 int32, y1 int32, x2 int32, y2 int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathLineI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
	return Status(ret)
}

func AddPathLine2I(path *Path, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathLine2I.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathArcI(path *Path, x int32, y int32, width int32, height int32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathArcI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func AddPathBezierI(path *Path, x1 int32, y1 int32, x2 int32, y2 int32, x3 int32, y3 int32, x4 int32, y4 int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathBezierI.Addr(), 
		uintptr(unsafe.Pointer(path)),
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

func AddPathBeziersI(path *Path, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathBeziersI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathCurveI(path *Path, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathCurveI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathCurve2I(path *Path, points *Point, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathCurve2I.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func AddPathCurve3I(path *Path, points *Point, count int32, offset int32, numberOfSegments int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathCurve3I.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(offset),
		uintptr(numberOfSegments),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func AddPathClosedCurveI(path *Path, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathClosedCurveI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func AddPathClosedCurve2I(path *Path, points *Point, count int32, tension float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathClosedCurve2I.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(tension)))
	return Status(ret)
}

func AddPathRectangleI(path *Path, x int32, y int32, width int32, height int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathRectangleI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return Status(ret)
}

func AddPathRectanglesI(path *Path, rects *Rect, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathRectanglesI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count))
	return Status(ret)
}

func AddPathEllipseI(path *Path, x int32, y int32, width int32, height int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathEllipseI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
	return Status(ret)
}

func AddPathPieI(path *Path, x int32, y int32, width int32, height int32, startAngle float32, sweepAngle float32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathPieI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(math.Float32bits(startAngle)),
		uintptr(math.Float32bits(sweepAngle)))
	return Status(ret)
}

func AddPathPolygonI(path *Path, points *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipAddPathPolygonI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count))
	return Status(ret)
}

func FlattenPath(path *Path, matrix *Matrix, flatness float32) Status {
	ret, _, _ := syscall.SyscallN(gdipFlattenPath.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(flatness)))
	return Status(ret)
}

func WindingModeOutline(path *Path, matrix *Matrix, flatness float32) Status {
	ret, _, _ := syscall.SyscallN(gdipWindingModeOutline.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(flatness)))
	return Status(ret)
}

func WidenPath(nativePath *Path, pen *Pen, matrix *Matrix, flatness float32) Status {
	ret, _, _ := syscall.SyscallN(gdipWidenPath.Addr(), 
		uintptr(unsafe.Pointer(nativePath)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(flatness)))
	return Status(ret)
}

func WarpPath(path *Path, matrix *Matrix, points *PointF, count int32, srcx float32, srcy float32, srcwidth float32, srcheight float32, warpMode WarpMode, flatness float32) Status {
	ret, _, _ := syscall.SyscallN(gdipWarpPath.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(math.Float32bits(srcx)),
		uintptr(math.Float32bits(srcy)),
		uintptr(math.Float32bits(srcwidth)),
		uintptr(math.Float32bits(srcheight)),
		uintptr(warpMode),
		uintptr(math.Float32bits(flatness)))
	return Status(ret)
}

func TransformPath(path *Path, matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipTransformPath.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func GetPathWorldBounds(path *Path, bounds *RectF, matrix *Matrix, pen *Pen) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathWorldBounds.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(bounds)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pen)))
	return Status(ret)
}

func GetPathWorldBoundsI(path *Path, bounds *Rect, matrix *Matrix, pen *Pen) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPathWorldBoundsI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(bounds)),
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pen)))
	return Status(ret)
}

func IsVisiblePathPoint(path *Path, x float32, y float32, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisiblePathPoint.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsVisiblePathPointI(path *Path, x int32, y int32, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsVisiblePathPointI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsOutlineVisiblePathPoint(path *Path, x float32, y float32, pen *Pen, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsOutlineVisiblePathPoint.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsOutlineVisiblePathPointI(path *Path, x int32, y int32, pen *Pen, graphics *Graphics, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsOutlineVisiblePathPointI.Addr(), 
		uintptr(unsafe.Pointer(path)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(pen)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}


	// pathiterator
func CreatePathIter(iterator **PathIterator, path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipCreatePathIter.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func DeletePathIter(iterator *PathIterator) Status {
	ret, _, _ := syscall.SyscallN(gdipDeletePathIter.Addr(), 
		uintptr(unsafe.Pointer(iterator)))
	return Status(ret)
}

func PathIterNextSubpath(iterator *PathIterator, resultCount *int32, startIndex *int32, endIndex *int32, isClosed *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterNextSubpath.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(startIndex)),
		uintptr(unsafe.Pointer(endIndex)),
		uintptr(unsafe.Pointer(isClosed)))
	return Status(ret)
}

func PathIterNextSubpathPath(iterator *PathIterator, resultCount *int32, path *Path, isClosed *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterNextSubpathPath.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(isClosed)))
	return Status(ret)
}

func PathIterNextPathType(iterator *PathIterator, resultCount *int32, pathType *win32.BYTE, startIndex *int32, endIndex *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterNextPathType.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(pathType)),
		uintptr(unsafe.Pointer(startIndex)),
		uintptr(unsafe.Pointer(endIndex)))
	return Status(ret)
}

func PathIterNextMarker(iterator *PathIterator, resultCount *int32, startIndex *int32, endIndex *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterNextMarker.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(startIndex)),
		uintptr(unsafe.Pointer(endIndex)))
	return Status(ret)
}

func PathIterNextMarkerPath(iterator *PathIterator, resultCount *int32, path *Path) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterNextMarkerPath.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(path)))
	return Status(ret)
}

func PathIterGetCount(iterator *PathIterator, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterGetCount.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func PathIterGetSubpathCount(iterator *PathIterator, count *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterGetSubpathCount.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func PathIterIsValid(iterator *PathIterator, valid *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterIsValid.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(valid)))
	return Status(ret)
}

func PathIterHasCurve(iterator *PathIterator, hasCurve *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterHasCurve.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(hasCurve)))
	return Status(ret)
}

func PathIterRewind(iterator *PathIterator) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterRewind.Addr(), 
		uintptr(unsafe.Pointer(iterator)))
	return Status(ret)
}

func PathIterEnumerate(iterator *PathIterator, resultCount *int32, points *PointF, types *win32.BYTE, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterEnumerate.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(points)),
		uintptr(unsafe.Pointer(types)),
		uintptr(count))
	return Status(ret)
}

func PathIterCopyData(iterator *PathIterator, resultCount *int32, points *PointF, types *win32.BYTE, startIndex int32, endIndex int32) Status {
	ret, _, _ := syscall.SyscallN(gdipPathIterCopyData.Addr(), 
		uintptr(unsafe.Pointer(iterator)),
		uintptr(unsafe.Pointer(resultCount)),
		uintptr(unsafe.Pointer(points)),
		uintptr(unsafe.Pointer(types)),
		uintptr(startIndex),
		uintptr(endIndex))
	return Status(ret)
}

