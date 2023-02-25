package gdip

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var (
	gdiplusShutdown *windows.LazyProc
	gdiplusStartup  *windows.LazyProc
)

var (
	dll *windows.LazyDLL
)

var (
	token uintptr
)

func init() {
	dll = windows.NewLazySystemDLL("gdiplus.dll")

	gdiplusShutdown = dll.NewProc("GdiplusShutdown")
	gdiplusStartup = dll.NewProc("GdiplusStartup")
}

func GdiplusStartup() Status {

	var input GdiplusStartupInput
	var output GdiplusStartupOutput

	input.GdiplusVersion = 1

	ret, _, _ := syscall.SyscallN(gdiplusStartup.Addr(),
		uintptr(unsafe.Pointer(&token)),
		uintptr(unsafe.Pointer(&input)),
		uintptr(unsafe.Pointer(&output)))

	return Status(ret)
}

func GdiplusShutdown() {
	syscall.SyscallN(gdiplusShutdown.Addr(), token)
}
