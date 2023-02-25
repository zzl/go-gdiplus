package gdip

import (
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// memory
	gdipAlloc	*windows.LazyProc
	gdipFree	*windows.LazyProc

	// notification
	gdiplusNotificationHook	*windows.LazyProc
	gdiplusNotificationUnhook	*windows.LazyProc
)

func init() {

	// memory
	gdipAlloc = dll.NewProc("GdipAlloc")
	gdipFree = dll.NewProc("GdipFree")

	// notification
	gdiplusNotificationHook = dll.NewProc("GdiplusNotificationHook")
	gdiplusNotificationUnhook = dll.NewProc("GdiplusNotificationUnhook")
}

	// memory
func Alloc(size uint) Status {
	ret, _, _ := syscall.SyscallN(gdipAlloc.Addr(), 
		uintptr(size))
	return Status(ret)
}

func Free(ptr *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipFree.Addr(), 
		uintptr(unsafe.Pointer(ptr)))
	return Status(ret)
}


	// notification
func lusNotificationHook(token *win32.ULONG_PTR) Status {
	ret, _, _ := syscall.SyscallN(gdiplusNotificationHook.Addr(), 
		uintptr(unsafe.Pointer(token)))
	return Status(ret)
}

func lusNotificationUnhook(token win32.ULONG_PTR) byte {
	ret, _, _ := syscall.SyscallN(gdiplusNotificationUnhook.Addr(), 
		uintptr(token))
	return byte(ret)
}

