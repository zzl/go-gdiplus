package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// bitmap
	gdipCreateBitmapFromStream	*windows.LazyProc
	gdipCreateBitmapFromFile	*windows.LazyProc
	gdipCreateBitmapFromStreamICM	*windows.LazyProc
	gdipCreateBitmapFromFileICM	*windows.LazyProc
	gdipCreateBitmapFromScan0	*windows.LazyProc
	gdipCreateBitmapFromGraphics	*windows.LazyProc
	gdipCreateBitmapFromDirectDrawSurface	*windows.LazyProc
	gdipCreateBitmapFromGdiDib	*windows.LazyProc
	gdipCreateBitmapFromHBITMAP	*windows.LazyProc
	gdipCreateHBITMAPFromBitmap	*windows.LazyProc
	gdipCreateBitmapFromHICON	*windows.LazyProc
	gdipCreateHICONFromBitmap	*windows.LazyProc
	gdipCreateBitmapFromResource	*windows.LazyProc
	gdipCloneBitmapArea	*windows.LazyProc
	gdipCloneBitmapAreaI	*windows.LazyProc
	gdipBitmapLockBits	*windows.LazyProc
	gdipBitmapUnlockBits	*windows.LazyProc
	gdipBitmapGetPixel	*windows.LazyProc
	gdipBitmapSetPixel	*windows.LazyProc
	gdipBitmapSetResolution	*windows.LazyProc
	gdipBitmapConvertFormat	*windows.LazyProc
	gdipInitializePalette	*windows.LazyProc
	gdipBitmapApplyEffect	*windows.LazyProc
	gdipBitmapCreateApplyEffect	*windows.LazyProc
	gdipBitmapGetHistogram	*windows.LazyProc
	gdipBitmapGetHistogramSize	*windows.LazyProc
	gdipCreateEffect	*windows.LazyProc
	gdipDeleteEffect	*windows.LazyProc
	gdipGetEffectParameterSize	*windows.LazyProc
	gdipSetEffectParameters	*windows.LazyProc
	gdipGetEffectParameters	*windows.LazyProc
	gdipTestControl	*windows.LazyProc

	// cachedbitmap
	gdipCreateCachedBitmap	*windows.LazyProc
	gdipDeleteCachedBitmap	*windows.LazyProc
	gdipDrawCachedBitmap	*windows.LazyProc
	gdipEmfToWmfBits	*windows.LazyProc
)

func init() {

	// bitmap
	gdipCreateBitmapFromStream = dll.NewProc("GdipCreateBitmapFromStream")
	gdipCreateBitmapFromFile = dll.NewProc("GdipCreateBitmapFromFile")
	gdipCreateBitmapFromStreamICM = dll.NewProc("GdipCreateBitmapFromStreamICM")
	gdipCreateBitmapFromFileICM = dll.NewProc("GdipCreateBitmapFromFileICM")
	gdipCreateBitmapFromScan0 = dll.NewProc("GdipCreateBitmapFromScan0")
	gdipCreateBitmapFromGraphics = dll.NewProc("GdipCreateBitmapFromGraphics")
	gdipCreateBitmapFromDirectDrawSurface = dll.NewProc("GdipCreateBitmapFromDirectDrawSurface")
	gdipCreateBitmapFromGdiDib = dll.NewProc("GdipCreateBitmapFromGdiDib")
	gdipCreateBitmapFromHBITMAP = dll.NewProc("GdipCreateBitmapFromHBITMAP")
	gdipCreateHBITMAPFromBitmap = dll.NewProc("GdipCreateHBITMAPFromBitmap")
	gdipCreateBitmapFromHICON = dll.NewProc("GdipCreateBitmapFromHICON")
	gdipCreateHICONFromBitmap = dll.NewProc("GdipCreateHICONFromBitmap")
	gdipCreateBitmapFromResource = dll.NewProc("GdipCreateBitmapFromResource")
	gdipCloneBitmapArea = dll.NewProc("GdipCloneBitmapArea")
	gdipCloneBitmapAreaI = dll.NewProc("GdipCloneBitmapAreaI")
	gdipBitmapLockBits = dll.NewProc("GdipBitmapLockBits")
	gdipBitmapUnlockBits = dll.NewProc("GdipBitmapUnlockBits")
	gdipBitmapGetPixel = dll.NewProc("GdipBitmapGetPixel")
	gdipBitmapSetPixel = dll.NewProc("GdipBitmapSetPixel")
	gdipBitmapSetResolution = dll.NewProc("GdipBitmapSetResolution")
	gdipBitmapConvertFormat = dll.NewProc("GdipBitmapConvertFormat")
	gdipInitializePalette = dll.NewProc("GdipInitializePalette")
	gdipBitmapApplyEffect = dll.NewProc("GdipBitmapApplyEffect")
	gdipBitmapCreateApplyEffect = dll.NewProc("GdipBitmapCreateApplyEffect")
	gdipBitmapGetHistogram = dll.NewProc("GdipBitmapGetHistogram")
	gdipBitmapGetHistogramSize = dll.NewProc("GdipBitmapGetHistogramSize")
	gdipCreateEffect = dll.NewProc("GdipCreateEffect")
	gdipDeleteEffect = dll.NewProc("GdipDeleteEffect")
	gdipGetEffectParameterSize = dll.NewProc("GdipGetEffectParameterSize")
	gdipSetEffectParameters = dll.NewProc("GdipSetEffectParameters")
	gdipGetEffectParameters = dll.NewProc("GdipGetEffectParameters")
	gdipTestControl = dll.NewProc("GdipTestControl")

	// cachedbitmap
	gdipCreateCachedBitmap = dll.NewProc("GdipCreateCachedBitmap")
	gdipDeleteCachedBitmap = dll.NewProc("GdipDeleteCachedBitmap")
	gdipDrawCachedBitmap = dll.NewProc("GdipDrawCachedBitmap")
	gdipEmfToWmfBits = dll.NewProc("GdipEmfToWmfBits")
}

	// bitmap
func CreateBitmapFromStream(stream *win32.IStream, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromStream.Addr(), 
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromFile(filename *uint16, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromFile.Addr(), 
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromStreamICM(stream *win32.IStream, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromStreamICM.Addr(), 
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromFileICM(filename *uint16, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromFileICM.Addr(), 
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromScan0(width int32, height int32, stride int32, format PixelFormat, scan0 *win32.BYTE, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromScan0.Addr(), 
		uintptr(width),
		uintptr(height),
		uintptr(stride),
		uintptr(format),
		uintptr(unsafe.Pointer(scan0)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromGraphics(width int32, height int32, target *Graphics, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromGraphics.Addr(), 
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(target)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromDirectDrawSurface(surface *win32.IUnknown, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromDirectDrawSurface.Addr(), 
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromGdiDib(gdiBitmapInfo *win32.BITMAPINFO, gdiBitmapData *byte, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromGdiDib.Addr(), 
		uintptr(unsafe.Pointer(gdiBitmapInfo)),
		uintptr(unsafe.Pointer(gdiBitmapData)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateBitmapFromHBITMAP(hbm win32.HBITMAP, hpal win32.HPALETTE, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromHBITMAP.Addr(), 
		uintptr(hbm),
		uintptr(hpal),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateHBITMAPFromBitmap(bitmap *Bitmap, hbmReturn *win32.HBITMAP, background ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateHBITMAPFromBitmap.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(hbmReturn)),
		uintptr(background))
	return Status(ret)
}

func CreateBitmapFromHICON(hicon win32.HICON, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromHICON.Addr(), 
		uintptr(hicon),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CreateHICONFromBitmap(bitmap *Bitmap, hbmReturn *win32.HICON) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateHICONFromBitmap.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(hbmReturn)))
	return Status(ret)
}

func CreateBitmapFromResource(hInstance win32.HINSTANCE, lpBitmapName *uint16, bitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateBitmapFromResource.Addr(), 
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpBitmapName)),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func CloneBitmapArea(x float32, y float32, width float32, height float32, format PixelFormat, srcBitmap *Bitmap, dstBitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneBitmapArea.Addr(), 
		uintptr(math.Float32bits(x)),
		uintptr(math.Float32bits(y)),
		uintptr(math.Float32bits(width)),
		uintptr(math.Float32bits(height)),
		uintptr(format),
		uintptr(unsafe.Pointer(srcBitmap)),
		uintptr(unsafe.Pointer(dstBitmap)))
	return Status(ret)
}

func CloneBitmapAreaI(x int32, y int32, width int32, height int32, format PixelFormat, srcBitmap *Bitmap, dstBitmap **Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneBitmapAreaI.Addr(), 
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(unsafe.Pointer(srcBitmap)),
		uintptr(unsafe.Pointer(dstBitmap)))
	return Status(ret)
}

func BitmapLockBits(bitmap *Bitmap, rect *Rect, flags uint32, format PixelFormat, lockedBitmapData *BitmapData) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapLockBits.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(flags),
		uintptr(format),
		uintptr(unsafe.Pointer(lockedBitmapData)))
	return Status(ret)
}

func BitmapUnlockBits(bitmap *Bitmap, lockedBitmapData *BitmapData) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapUnlockBits.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(lockedBitmapData)))
	return Status(ret)
}

func BitmapGetPixel(bitmap *Bitmap, x int32, y int32, color *ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapGetPixel.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(color)))
	return Status(ret)
}

func BitmapSetPixel(bitmap *Bitmap, x int32, y int32, color ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapSetPixel.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(x),
		uintptr(y),
		uintptr(color))
	return Status(ret)
}

func BitmapSetResolution(bitmap *Bitmap, xdpi float32, ydpi float32) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapSetResolution.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(math.Float32bits(xdpi)),
		uintptr(math.Float32bits(ydpi)))
	return Status(ret)
}

func BitmapConvertFormat(pInputBitmap *Bitmap, format PixelFormat, dithertype DitherType, palettetype PaletteType, palette *ColorPalette, alphaThresholdPercent float32) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapConvertFormat.Addr(), 
		uintptr(unsafe.Pointer(pInputBitmap)),
		uintptr(format),
		uintptr(dithertype),
		uintptr(palettetype),
		uintptr(unsafe.Pointer(palette)),
		uintptr(math.Float32bits(alphaThresholdPercent)))
	return Status(ret)
}

func InitializePalette(palette *ColorPalette, palettetype PaletteType, optimalColors int32, useTransparentColor win32.BOOL, bitmap *Bitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipInitializePalette.Addr(), 
		uintptr(unsafe.Pointer(palette)),
		uintptr(palettetype),
		uintptr(optimalColors),
		uintptr(useTransparentColor),
		uintptr(unsafe.Pointer(bitmap)))
	return Status(ret)
}

func BitmapApplyEffect(bitmap *Bitmap, effect *CGpEffect, roi *win32.RECT, useAuxData win32.BOOL, auxData **byte, auxDataSize *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapApplyEffect.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(roi)),
		uintptr(useAuxData),
		uintptr(unsafe.Pointer(auxData)),
		uintptr(unsafe.Pointer(auxDataSize)))
	return Status(ret)
}

func BitmapCreateApplyEffect(inputBitmaps **Bitmap, numInputs int32, effect *CGpEffect, roi *win32.RECT, outputRect *win32.RECT, outputBitmap **Bitmap, useAuxData win32.BOOL, auxData **byte, auxDataSize *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapCreateApplyEffect.Addr(), 
		uintptr(unsafe.Pointer(inputBitmaps)),
		uintptr(numInputs),
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(roi)),
		uintptr(unsafe.Pointer(outputRect)),
		uintptr(unsafe.Pointer(outputBitmap)),
		uintptr(useAuxData),
		uintptr(unsafe.Pointer(auxData)),
		uintptr(unsafe.Pointer(auxDataSize)))
	return Status(ret)
}

func BitmapGetHistogram(bitmap *Bitmap, format HistogramFormat, NumberOfEntries uint32, channel0 *uint32, channel1 *uint32, channel2 *uint32, channel3 *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapGetHistogram.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(format),
		uintptr(NumberOfEntries),
		uintptr(unsafe.Pointer(channel0)),
		uintptr(unsafe.Pointer(channel1)),
		uintptr(unsafe.Pointer(channel2)),
		uintptr(unsafe.Pointer(channel3)))
	return Status(ret)
}

func BitmapGetHistogramSize(format HistogramFormat, NumberOfEntries *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipBitmapGetHistogramSize.Addr(), 
		uintptr(format),
		uintptr(unsafe.Pointer(NumberOfEntries)))
	return Status(ret)
}

func CreateEffect(guid syscall.GUID, effect **CGpEffect) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateEffect.Addr(), 
		uintptr(unsafe.Pointer(&guid)),
		uintptr(unsafe.Pointer(effect)))
	return Status(ret)
}

func DeleteEffect(effect *CGpEffect) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteEffect.Addr(), 
		uintptr(unsafe.Pointer(effect)))
	return Status(ret)
}

func GetEffectParameterSize(effect *CGpEffect, size *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetEffectParameterSize.Addr(), 
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(size)))
	return Status(ret)
}

func SetEffectParameters(effect *CGpEffect, params *byte, size uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetEffectParameters.Addr(), 
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(params)),
		uintptr(size))
	return Status(ret)
}

func GetEffectParameters(effect *CGpEffect, size *uint32, params *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipGetEffectParameters.Addr(), 
		uintptr(unsafe.Pointer(effect)),
		uintptr(unsafe.Pointer(size)),
		uintptr(unsafe.Pointer(params)))
	return Status(ret)
}

func TestControl(control GpTestControlEnum, param *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipTestControl.Addr(), 
		uintptr(control),
		uintptr(unsafe.Pointer(param)))
	return Status(ret)
}


	// cachedbitmap
func CreateCachedBitmap(bitmap *Bitmap, graphics *Graphics, cachedBitmap **CachedBitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateCachedBitmap.Addr(), 
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(cachedBitmap)))
	return Status(ret)
}

func DeleteCachedBitmap(cachedBitmap *CachedBitmap) Status {
	ret, _, _ := syscall.SyscallN(gdipDeleteCachedBitmap.Addr(), 
		uintptr(unsafe.Pointer(cachedBitmap)))
	return Status(ret)
}

func DrawCachedBitmap(graphics *Graphics, cachedBitmap *CachedBitmap, x int32, y int32) Status {
	ret, _, _ := syscall.SyscallN(gdipDrawCachedBitmap.Addr(), 
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(cachedBitmap)),
		uintptr(x),
		uintptr(y))
	return Status(ret)
}

func EmfToWmfBits(hemf win32.HANDLE, cbData16 uint32, pData16 *byte, iMapMode int32, eFlags int32) uint32 {
	ret, _, _ := syscall.SyscallN(gdipEmfToWmfBits.Addr(), 
		uintptr(hemf),
		uintptr(cbData16),
		uintptr(unsafe.Pointer(pData16)),
		uintptr(iMapMode),
		uintptr(eFlags))
	return uint32(ret)
}

