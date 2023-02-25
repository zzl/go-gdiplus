package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// matrix
	gdipCreateMatrix	*windows.LazyProc
	gdipCreateMatrix2	*windows.LazyProc
	gdipCreateMatrix3	*windows.LazyProc
	gdipCreateMatrix3I	*windows.LazyProc
	gdipCloneMatrix	*windows.LazyProc
	gdipDeleteMatrix	*windows.LazyProc
	gdipSetMatrixElements	*windows.LazyProc
	gdipMultiplyMatrix	*windows.LazyProc
	gdipTranslateMatrix	*windows.LazyProc
	gdipScaleMatrix	*windows.LazyProc
	gdipRotateMatrix	*windows.LazyProc
	gdipShearMatrix	*windows.LazyProc
	gdipInvertMatrix	*windows.LazyProc
	gdipTransformMatrixPoints	*windows.LazyProc
	gdipTransformMatrixPointsI	*windows.LazyProc
	gdipVectorTransformMatrixPoints	*windows.LazyProc
	gdipVectorTransformMatrixPointsI	*windows.LazyProc
	gdipGetMatrixElements	*windows.LazyProc
	gdipIsMatrixInvertible	*windows.LazyProc
	gdipIsMatrixIdentity	*windows.LazyProc
	gdipIsMatrixEqual	*windows.LazyProc
)

func init() {

	// matrix
	gdipCreateMatrix = dll.NewProc("GdipCreateMatrix")
	gdipCreateMatrix2 = dll.NewProc("GdipCreateMatrix2")
	gdipCreateMatrix3 = dll.NewProc("GdipCreateMatrix3")
	gdipCreateMatrix3I = dll.NewProc("GdipCreateMatrix3I")
	gdipCloneMatrix = dll.NewProc("GdipCloneMatrix")
	gdipDeleteMatrix = dll.NewProc("GdipDeleteMatrix")
	gdipSetMatrixElements = dll.NewProc("GdipSetMatrixElements")
	gdipMultiplyMatrix = dll.NewProc("GdipMultiplyMatrix")
	gdipTranslateMatrix = dll.NewProc("GdipTranslateMatrix")
	gdipScaleMatrix = dll.NewProc("GdipScaleMatrix")
	gdipRotateMatrix = dll.NewProc("GdipRotateMatrix")
	gdipShearMatrix = dll.NewProc("GdipShearMatrix")
	gdipInvertMatrix = dll.NewProc("GdipInvertMatrix")
	gdipTransformMatrixPoints = dll.NewProc("GdipTransformMatrixPoints")
	gdipTransformMatrixPointsI = dll.NewProc("GdipTransformMatrixPointsI")
	gdipVectorTransformMatrixPoints = dll.NewProc("GdipVectorTransformMatrixPoints")
	gdipVectorTransformMatrixPointsI = dll.NewProc("GdipVectorTransformMatrixPointsI")
	gdipGetMatrixElements = dll.NewProc("GdipGetMatrixElements")
	gdipIsMatrixInvertible = dll.NewProc("GdipIsMatrixInvertible")
	gdipIsMatrixIdentity = dll.NewProc("GdipIsMatrixIdentity")
	gdipIsMatrixEqual = dll.NewProc("GdipIsMatrixEqual")
}

	// matrix
func CreateMatrix(matrix **Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func CreateMatrix2(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32, matrix **Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMatrix2.Addr(), 
		uintptr(math.Float32bits(m11)),
		uintptr(math.Float32bits(m12)),
		uintptr(math.Float32bits(m21)),
		uintptr(math.Float32bits(m22)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func CreateMatrix3(rect *RectF, dstplg *PointF, matrix **Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMatrix3.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(dstplg)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func CreateMatrix3I(rect *Rect, dstplg *Point, matrix **Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateMatrix3I.Addr(), 
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(dstplg)),
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func CloneMatrix(matrix *Matrix, cloneMatrix **Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(cloneMatrix)))
	return Status(ret)
}

func DeleteMatrix(matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func SetMatrixElements(matrix *Matrix, m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetMatrixElements.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(m11)),
		uintptr(math.Float32bits(m12)),
		uintptr(math.Float32bits(m21)),
		uintptr(math.Float32bits(m22)),
		uintptr(math.Float32bits(dx)),
		uintptr(math.Float32bits(dy)))
	return Status(ret)
}

func MultiplyMatrix(matrix *Matrix, matrix2 *Matrix, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipMultiplyMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(matrix2)),
		uintptr(order))
	return Status(ret)
}

func TranslateMatrix(matrix *Matrix, offsetX float32, offsetY float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipTranslateMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(offsetX)),
		uintptr(math.Float32bits(offsetY)),
		uintptr(order))
	return Status(ret)
}

func ScaleMatrix(matrix *Matrix, scaleX float32, scaleY float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipScaleMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(scaleX)),
		uintptr(math.Float32bits(scaleY)),
		uintptr(order))
	return Status(ret)
}

func RotateMatrix(matrix *Matrix, angle float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipRotateMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(angle)),
		uintptr(order))
	return Status(ret)
}

func ShearMatrix(matrix *Matrix, shearX float32, shearY float32, order MatrixOrder) Status {
	ret, _, _ := syscall.SyscallN(gdipShearMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(math.Float32bits(shearX)),
		uintptr(math.Float32bits(shearY)),
		uintptr(order))
	return Status(ret)
}

func InvertMatrix(matrix *Matrix) Status {
	ret, _, _ := syscall.SyscallN(gdipInvertMatrix.Addr(), 
		uintptr(unsafe.Pointer(matrix)))
	return Status(ret)
}

func TransformMatrixPoints(matrix *Matrix, pts *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipTransformMatrixPoints.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pts)),
		uintptr(count))
	return Status(ret)
}

func TransformMatrixPointsI(matrix *Matrix, pts *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipTransformMatrixPointsI.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pts)),
		uintptr(count))
	return Status(ret)
}

func VectorTransformMatrixPoints(matrix *Matrix, pts *PointF, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipVectorTransformMatrixPoints.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pts)),
		uintptr(count))
	return Status(ret)
}

func VectorTransformMatrixPointsI(matrix *Matrix, pts *Point, count int32) Status {
	ret, _, _ := syscall.SyscallN(gdipVectorTransformMatrixPointsI.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(pts)),
		uintptr(count))
	return Status(ret)
}

func GetMatrixElements(matrix *Matrix, matrixOut *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetMatrixElements.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(matrixOut)))
	return Status(ret)
}

func IsMatrixInvertible(matrix *Matrix, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsMatrixInvertible.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsMatrixIdentity(matrix *Matrix, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsMatrixIdentity.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

func IsMatrixEqual(matrix *Matrix, matrix2 *Matrix, result *win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipIsMatrixEqual.Addr(), 
		uintptr(unsafe.Pointer(matrix)),
		uintptr(unsafe.Pointer(matrix2)),
		uintptr(unsafe.Pointer(result)))
	return Status(ret)
}

