package gdip

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/zzl/go-win32api/v2/win32"
	"golang.org/x/sys/windows"
)

var (

	// image
	gdipLoadImageFromStream	*windows.LazyProc
	gdipLoadImageFromFile	*windows.LazyProc
	gdipLoadImageFromStreamICM	*windows.LazyProc
	gdipLoadImageFromFileICM	*windows.LazyProc
	gdipCloneImage	*windows.LazyProc
	gdipDisposeImage	*windows.LazyProc
	gdipSaveImageToFile	*windows.LazyProc
	gdipSaveImageToStream	*windows.LazyProc
	gdipSaveAdd	*windows.LazyProc
	gdipSaveAddImage	*windows.LazyProc
	gdipGetImageGraphicsContext	*windows.LazyProc
	gdipGetImageBounds	*windows.LazyProc
	gdipGetImageDimension	*windows.LazyProc
	gdipGetImageType	*windows.LazyProc
	gdipGetImageWidth	*windows.LazyProc
	gdipGetImageHeight	*windows.LazyProc
	gdipGetImageHorizontalResolution	*windows.LazyProc
	gdipGetImageVerticalResolution	*windows.LazyProc
	gdipGetImageFlags	*windows.LazyProc
	gdipGetImageRawFormat	*windows.LazyProc
	gdipGetImagePixelFormat	*windows.LazyProc
	gdipGetImageThumbnail	*windows.LazyProc
	gdipGetEncoderParameterListSize	*windows.LazyProc
	gdipGetEncoderParameterList	*windows.LazyProc
	gdipImageGetFrameDimensionsCount	*windows.LazyProc
	gdipImageGetFrameDimensionsList	*windows.LazyProc
	gdipImageGetFrameCount	*windows.LazyProc
	gdipImageSelectActiveFrame	*windows.LazyProc
	gdipImageRotateFlip	*windows.LazyProc
	gdipGetImagePalette	*windows.LazyProc
	gdipSetImagePalette	*windows.LazyProc
	gdipGetImagePaletteSize	*windows.LazyProc
	gdipGetPropertyCount	*windows.LazyProc
	gdipGetPropertyIdList	*windows.LazyProc
	gdipGetPropertyItemSize	*windows.LazyProc
	gdipGetPropertyItem	*windows.LazyProc
	gdipGetPropertySize	*windows.LazyProc
	gdipGetAllPropertyItems	*windows.LazyProc
	gdipRemovePropertyItem	*windows.LazyProc
	gdipSetPropertyItem	*windows.LazyProc
	gdipFindFirstImageItem	*windows.LazyProc
	gdipFindNextImageItem	*windows.LazyProc
	gdipGetImageItemData	*windows.LazyProc
	gdipImageSetAbort	*windows.LazyProc
	gdipConvertToEmfPlus	*windows.LazyProc
	gdipConvertToEmfPlusToFile	*windows.LazyProc
	gdipConvertToEmfPlusToStream	*windows.LazyProc
	gdipImageForceValidation	*windows.LazyProc

	// imageattributes
	gdipCreateImageAttributes	*windows.LazyProc
	gdipCloneImageAttributes	*windows.LazyProc
	gdipDisposeImageAttributes	*windows.LazyProc
	gdipSetImageAttributesToIdentity	*windows.LazyProc
	gdipResetImageAttributes	*windows.LazyProc
	gdipSetImageAttributesColorMatrix	*windows.LazyProc
	gdipSetImageAttributesThreshold	*windows.LazyProc
	gdipSetImageAttributesGamma	*windows.LazyProc
	gdipSetImageAttributesNoOp	*windows.LazyProc
	gdipSetImageAttributesColorKeys	*windows.LazyProc
	gdipSetImageAttributesOutputChannel	*windows.LazyProc
	gdipSetImageAttributesOutputChannelColorProfile	*windows.LazyProc
	gdipSetImageAttributesRemapTable	*windows.LazyProc
	gdipSetImageAttributesWrapMode	*windows.LazyProc
	gdipSetImageAttributesICMMode	*windows.LazyProc
	gdipGetImageAttributesAdjustedPalette	*windows.LazyProc
	gdipSetImageAttributesCachedBackground	*windows.LazyProc
)

func init() {

	// image
	gdipLoadImageFromStream = dll.NewProc("GdipLoadImageFromStream")
	gdipLoadImageFromFile = dll.NewProc("GdipLoadImageFromFile")
	gdipLoadImageFromStreamICM = dll.NewProc("GdipLoadImageFromStreamICM")
	gdipLoadImageFromFileICM = dll.NewProc("GdipLoadImageFromFileICM")
	gdipCloneImage = dll.NewProc("GdipCloneImage")
	gdipDisposeImage = dll.NewProc("GdipDisposeImage")
	gdipSaveImageToFile = dll.NewProc("GdipSaveImageToFile")
	gdipSaveImageToStream = dll.NewProc("GdipSaveImageToStream")
	gdipSaveAdd = dll.NewProc("GdipSaveAdd")
	gdipSaveAddImage = dll.NewProc("GdipSaveAddImage")
	gdipGetImageGraphicsContext = dll.NewProc("GdipGetImageGraphicsContext")
	gdipGetImageBounds = dll.NewProc("GdipGetImageBounds")
	gdipGetImageDimension = dll.NewProc("GdipGetImageDimension")
	gdipGetImageType = dll.NewProc("GdipGetImageType")
	gdipGetImageWidth = dll.NewProc("GdipGetImageWidth")
	gdipGetImageHeight = dll.NewProc("GdipGetImageHeight")
	gdipGetImageHorizontalResolution = dll.NewProc("GdipGetImageHorizontalResolution")
	gdipGetImageVerticalResolution = dll.NewProc("GdipGetImageVerticalResolution")
	gdipGetImageFlags = dll.NewProc("GdipGetImageFlags")
	gdipGetImageRawFormat = dll.NewProc("GdipGetImageRawFormat")
	gdipGetImagePixelFormat = dll.NewProc("GdipGetImagePixelFormat")
	gdipGetImageThumbnail = dll.NewProc("GdipGetImageThumbnail")
	gdipGetEncoderParameterListSize = dll.NewProc("GdipGetEncoderParameterListSize")
	gdipGetEncoderParameterList = dll.NewProc("GdipGetEncoderParameterList")
	gdipImageGetFrameDimensionsCount = dll.NewProc("GdipImageGetFrameDimensionsCount")
	gdipImageGetFrameDimensionsList = dll.NewProc("GdipImageGetFrameDimensionsList")
	gdipImageGetFrameCount = dll.NewProc("GdipImageGetFrameCount")
	gdipImageSelectActiveFrame = dll.NewProc("GdipImageSelectActiveFrame")
	gdipImageRotateFlip = dll.NewProc("GdipImageRotateFlip")
	gdipGetImagePalette = dll.NewProc("GdipGetImagePalette")
	gdipSetImagePalette = dll.NewProc("GdipSetImagePalette")
	gdipGetImagePaletteSize = dll.NewProc("GdipGetImagePaletteSize")
	gdipGetPropertyCount = dll.NewProc("GdipGetPropertyCount")
	gdipGetPropertyIdList = dll.NewProc("GdipGetPropertyIdList")
	gdipGetPropertyItemSize = dll.NewProc("GdipGetPropertyItemSize")
	gdipGetPropertyItem = dll.NewProc("GdipGetPropertyItem")
	gdipGetPropertySize = dll.NewProc("GdipGetPropertySize")
	gdipGetAllPropertyItems = dll.NewProc("GdipGetAllPropertyItems")
	gdipRemovePropertyItem = dll.NewProc("GdipRemovePropertyItem")
	gdipSetPropertyItem = dll.NewProc("GdipSetPropertyItem")
	gdipFindFirstImageItem = dll.NewProc("GdipFindFirstImageItem")
	gdipFindNextImageItem = dll.NewProc("GdipFindNextImageItem")
	gdipGetImageItemData = dll.NewProc("GdipGetImageItemData")
	gdipImageSetAbort = dll.NewProc("GdipImageSetAbort")
	gdipConvertToEmfPlus = dll.NewProc("GdipConvertToEmfPlus")
	gdipConvertToEmfPlusToFile = dll.NewProc("GdipConvertToEmfPlusToFile")
	gdipConvertToEmfPlusToStream = dll.NewProc("GdipConvertToEmfPlusToStream")
	gdipImageForceValidation = dll.NewProc("GdipImageForceValidation")

	// imageattributes
	gdipCreateImageAttributes = dll.NewProc("GdipCreateImageAttributes")
	gdipCloneImageAttributes = dll.NewProc("GdipCloneImageAttributes")
	gdipDisposeImageAttributes = dll.NewProc("GdipDisposeImageAttributes")
	gdipSetImageAttributesToIdentity = dll.NewProc("GdipSetImageAttributesToIdentity")
	gdipResetImageAttributes = dll.NewProc("GdipResetImageAttributes")
	gdipSetImageAttributesColorMatrix = dll.NewProc("GdipSetImageAttributesColorMatrix")
	gdipSetImageAttributesThreshold = dll.NewProc("GdipSetImageAttributesThreshold")
	gdipSetImageAttributesGamma = dll.NewProc("GdipSetImageAttributesGamma")
	gdipSetImageAttributesNoOp = dll.NewProc("GdipSetImageAttributesNoOp")
	gdipSetImageAttributesColorKeys = dll.NewProc("GdipSetImageAttributesColorKeys")
	gdipSetImageAttributesOutputChannel = dll.NewProc("GdipSetImageAttributesOutputChannel")
	gdipSetImageAttributesOutputChannelColorProfile = dll.NewProc("GdipSetImageAttributesOutputChannelColorProfile")
	gdipSetImageAttributesRemapTable = dll.NewProc("GdipSetImageAttributesRemapTable")
	gdipSetImageAttributesWrapMode = dll.NewProc("GdipSetImageAttributesWrapMode")
	gdipSetImageAttributesICMMode = dll.NewProc("GdipSetImageAttributesICMMode")
	gdipGetImageAttributesAdjustedPalette = dll.NewProc("GdipGetImageAttributesAdjustedPalette")
	gdipSetImageAttributesCachedBackground = dll.NewProc("GdipSetImageAttributesCachedBackground")
}

	// image
func LoadImageFromStream(stream *win32.IStream, image **Image) Status {
	ret, _, _ := syscall.SyscallN(gdipLoadImageFromStream.Addr(), 
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(image)))
	return Status(ret)
}

func LoadImageFromFile(filename *uint16, image **Image) Status {
	ret, _, _ := syscall.SyscallN(gdipLoadImageFromFile.Addr(), 
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(image)))
	return Status(ret)
}

func LoadImageFromStreamICM(stream *win32.IStream, image **Image) Status {
	ret, _, _ := syscall.SyscallN(gdipLoadImageFromStreamICM.Addr(), 
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(image)))
	return Status(ret)
}

func LoadImageFromFileICM(filename *uint16, image **Image) Status {
	ret, _, _ := syscall.SyscallN(gdipLoadImageFromFileICM.Addr(), 
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(image)))
	return Status(ret)
}

func CloneImage(image *Image, cloneImage **Image) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneImage.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(cloneImage)))
	return Status(ret)
}

func DisposeImage(image *Image) Status {
	ret, _, _ := syscall.SyscallN(gdipDisposeImage.Addr(), 
		uintptr(unsafe.Pointer(image)))
	return Status(ret)
}

func SaveImageToFile(image *Image, filename *uint16, clsidEncoder *win32.CLSID, encoderParams *EncoderParameters) Status {
	ret, _, _ := syscall.SyscallN(gdipSaveImageToFile.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(filename)),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(encoderParams)))
	return Status(ret)
}

func SaveImageToStream(image *Image, stream *win32.IStream, clsidEncoder *win32.CLSID, encoderParams *EncoderParameters) Status {
	ret, _, _ := syscall.SyscallN(gdipSaveImageToStream.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(encoderParams)))
	return Status(ret)
}

func SaveAdd(image *Image, encoderParams *EncoderParameters) Status {
	ret, _, _ := syscall.SyscallN(gdipSaveAdd.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(encoderParams)))
	return Status(ret)
}

func SaveAddImage(image *Image, newImage *Image, encoderParams *EncoderParameters) Status {
	ret, _, _ := syscall.SyscallN(gdipSaveAddImage.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(newImage)),
		uintptr(unsafe.Pointer(encoderParams)))
	return Status(ret)
}

func GetImageGraphicsContext(image *Image, graphics **Graphics) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageGraphicsContext.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(graphics)))
	return Status(ret)
}

func GetImageBounds(image *Image, srcRect *RectF, srcUnit *Unit) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageBounds.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(srcRect)),
		uintptr(unsafe.Pointer(srcUnit)))
	return Status(ret)
}

func GetImageDimension(image *Image, width *float32, height *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageDimension.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(width)),
		uintptr(unsafe.Pointer(height)))
	return Status(ret)
}

func GetImageType(image *Image, typ *ImageType) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageType.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(typ)))
	return Status(ret)
}

func GetImageWidth(image *Image, width *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageWidth.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(width)))
	return Status(ret)
}

func GetImageHeight(image *Image, height *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageHeight.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(height)))
	return Status(ret)
}

func GetImageHorizontalResolution(image *Image, resolution *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageHorizontalResolution.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(resolution)))
	return Status(ret)
}

func GetImageVerticalResolution(image *Image, resolution *float32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageVerticalResolution.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(resolution)))
	return Status(ret)
}

func GetImageFlags(image *Image, flags *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageFlags.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(flags)))
	return Status(ret)
}

func GetImageRawFormat(image *Image, format *syscall.GUID) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageRawFormat.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func GetImagePixelFormat(image *Image, format *PixelFormat) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImagePixelFormat.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(format)))
	return Status(ret)
}

func GetImageThumbnail(image *Image, thumbWidth uint32, thumbHeight uint32, thumbImage **Image, callback uintptr, callbackData *byte) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageThumbnail.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(thumbWidth),
		uintptr(thumbHeight),
		uintptr(unsafe.Pointer(thumbImage)),
		uintptr(callback),
		uintptr(unsafe.Pointer(callbackData)))
	return Status(ret)
}

func GetEncoderParameterListSize(image *Image, clsidEncoder *win32.CLSID, size *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetEncoderParameterListSize.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(size)))
	return Status(ret)
}

func GetEncoderParameterList(image *Image, clsidEncoder *win32.CLSID, size uint32, buffer *EncoderParameters) Status {
	ret, _, _ := syscall.SyscallN(gdipGetEncoderParameterList.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(size),
		uintptr(unsafe.Pointer(buffer)))
	return Status(ret)
}

func ImageGetFrameDimensionsCount(image *Image, count *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipImageGetFrameDimensionsCount.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func ImageGetFrameDimensionsList(image *Image, dimensionIDs *syscall.GUID, count uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipImageGetFrameDimensionsList.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(dimensionIDs)),
		uintptr(count))
	return Status(ret)
}

func ImageGetFrameCount(image *Image, dimensionID *syscall.GUID, count *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipImageGetFrameCount.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(dimensionID)),
		uintptr(unsafe.Pointer(count)))
	return Status(ret)
}

func ImageSelectActiveFrame(image *Image, dimensionID *syscall.GUID, frameIndex uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipImageSelectActiveFrame.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(dimensionID)),
		uintptr(frameIndex))
	return Status(ret)
}

func ImageRotateFlip(image *Image, rfType RotateFlipType) Status {
	ret, _, _ := syscall.SyscallN(gdipImageRotateFlip.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(rfType))
	return Status(ret)
}

func GetImagePalette(image *Image, palette *ColorPalette, size int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImagePalette.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(palette)),
		uintptr(size))
	return Status(ret)
}

func SetImagePalette(image *Image, palette *ColorPalette) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImagePalette.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(palette)))
	return Status(ret)
}

func GetImagePaletteSize(image *Image, size *int32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImagePaletteSize.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(size)))
	return Status(ret)
}

func GetPropertyCount(image *Image, numOfProperty *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPropertyCount.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(numOfProperty)))
	return Status(ret)
}

func GetPropertyIdList(image *Image, numOfProperty uint32, list *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPropertyIdList.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(numOfProperty),
		uintptr(unsafe.Pointer(list)))
	return Status(ret)
}

func GetPropertyItemSize(image *Image, propId uint32, size *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPropertyItemSize.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(propId),
		uintptr(unsafe.Pointer(size)))
	return Status(ret)
}

func GetPropertyItem(image *Image, propId uint32, propSize uint32, buffer *PropertyItem) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPropertyItem.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(propId),
		uintptr(propSize),
		uintptr(unsafe.Pointer(buffer)))
	return Status(ret)
}

func GetPropertySize(image *Image, totalBufferSize *uint32, numProperties *uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipGetPropertySize.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(totalBufferSize)),
		uintptr(unsafe.Pointer(numProperties)))
	return Status(ret)
}

func GetAllPropertyItems(image *Image, totalBufferSize uint32, numProperties uint32, allItems *PropertyItem) Status {
	ret, _, _ := syscall.SyscallN(gdipGetAllPropertyItems.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(totalBufferSize),
		uintptr(numProperties),
		uintptr(unsafe.Pointer(allItems)))
	return Status(ret)
}

func RemovePropertyItem(image *Image, propId uint32) Status {
	ret, _, _ := syscall.SyscallN(gdipRemovePropertyItem.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(propId))
	return Status(ret)
}

func SetPropertyItem(image *Image, item *PropertyItem) Status {
	ret, _, _ := syscall.SyscallN(gdipSetPropertyItem.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
	return Status(ret)
}

func FindFirstImageItem(image *Image, item *ImageItemData) Status {
	ret, _, _ := syscall.SyscallN(gdipFindFirstImageItem.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
	return Status(ret)
}

func FindNextImageItem(image *Image, item *ImageItemData) Status {
	ret, _, _ := syscall.SyscallN(gdipFindNextImageItem.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
	return Status(ret)
}

func GetImageItemData(image *Image, item *ImageItemData) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageItemData.Addr(), 
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(item)))
	return Status(ret)
}

func ImageSetAbort(pImage *Image, pIAbort *GdiplusAbort) Status {
	ret, _, _ := syscall.SyscallN(gdipImageSetAbort.Addr(), 
		uintptr(unsafe.Pointer(pImage)),
		uintptr(unsafe.Pointer(pIAbort)))
	return Status(ret)
}

func ConvertToEmfPlus(refGraphics *Graphics, metafile *Metafile, conversionSuccess *win32.BOOL, emfType EmfType, description *uint16, out_metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipConvertToEmfPlus.Addr(), 
		uintptr(unsafe.Pointer(refGraphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(conversionSuccess)),
		uintptr(emfType),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(out_metafile)))
	return Status(ret)
}

func ConvertToEmfPlusToFile(refGraphics *Graphics, metafile *Metafile, conversionSuccess *win32.BOOL, filename *uint16, emfType EmfType, description *uint16, out_metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipConvertToEmfPlusToFile.Addr(), 
		uintptr(unsafe.Pointer(refGraphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(conversionSuccess)),
		uintptr(unsafe.Pointer(filename)),
		uintptr(emfType),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(out_metafile)))
	return Status(ret)
}

func ConvertToEmfPlusToStream(refGraphics *Graphics, metafile *Metafile, conversionSuccess *win32.BOOL, stream *win32.IStream, emfType EmfType, description *uint16, out_metafile **Metafile) Status {
	ret, _, _ := syscall.SyscallN(gdipConvertToEmfPlusToStream.Addr(), 
		uintptr(unsafe.Pointer(refGraphics)),
		uintptr(unsafe.Pointer(metafile)),
		uintptr(unsafe.Pointer(conversionSuccess)),
		uintptr(unsafe.Pointer(stream)),
		uintptr(emfType),
		uintptr(unsafe.Pointer(description)),
		uintptr(unsafe.Pointer(out_metafile)))
	return Status(ret)
}

func ImageForceValidation(image *Image) Status {
	ret, _, _ := syscall.SyscallN(gdipImageForceValidation.Addr(), 
		uintptr(unsafe.Pointer(image)))
	return Status(ret)
}


	// imageattributes
func CreateImageAttributes(imageattr **ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipCreateImageAttributes.Addr(), 
		uintptr(unsafe.Pointer(imageattr)))
	return Status(ret)
}

func CloneImageAttributes(imageattr *ImageAttributes, cloneImageattr **ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipCloneImageAttributes.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(unsafe.Pointer(cloneImageattr)))
	return Status(ret)
}

func DisposeImageAttributes(imageattr *ImageAttributes) Status {
	ret, _, _ := syscall.SyscallN(gdipDisposeImageAttributes.Addr(), 
		uintptr(unsafe.Pointer(imageattr)))
	return Status(ret)
}

func SetImageAttributesToIdentity(imageattr *ImageAttributes, typ ColorAdjustType) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesToIdentity.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ))
	return Status(ret)
}

func ResetImageAttributes(imageattr *ImageAttributes, typ ColorAdjustType) Status {
	ret, _, _ := syscall.SyscallN(gdipResetImageAttributes.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ))
	return Status(ret)
}

func SetImageAttributesColorMatrix(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL, colorMatrix *ColorMatrix, grayMatrix *ColorMatrix, flags ColorMatrixFlags) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesColorMatrix.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag),
		uintptr(unsafe.Pointer(colorMatrix)),
		uintptr(unsafe.Pointer(grayMatrix)),
		uintptr(flags))
	return Status(ret)
}

func SetImageAttributesThreshold(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL, threshold float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesThreshold.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag),
		uintptr(math.Float32bits(threshold)))
	return Status(ret)
}

func SetImageAttributesGamma(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL, gamma float32) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesGamma.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag),
		uintptr(math.Float32bits(gamma)))
	return Status(ret)
}

func SetImageAttributesNoOp(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesNoOp.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag))
	return Status(ret)
}

func SetImageAttributesColorKeys(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL, colorLow ARGB, colorHigh ARGB) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesColorKeys.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag),
		uintptr(colorLow),
		uintptr(colorHigh))
	return Status(ret)
}

func SetImageAttributesOutputChannel(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL, channelFlags ColorChannelFlags) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesOutputChannel.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag),
		uintptr(channelFlags))
	return Status(ret)
}

func SetImageAttributesOutputChannelColorProfile(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL, colorProfileFilename *uint16) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesOutputChannelColorProfile.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag),
		uintptr(unsafe.Pointer(colorProfileFilename)))
	return Status(ret)
}

func SetImageAttributesRemapTable(imageattr *ImageAttributes, typ ColorAdjustType, enableFlag win32.BOOL, mapSize uint32, map4 *ColorMap) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesRemapTable.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(typ),
		uintptr(enableFlag),
		uintptr(mapSize),
		uintptr(unsafe.Pointer(map4)))
	return Status(ret)
}

func SetImageAttributesWrapMode(imageAttr *ImageAttributes, wrap WrapMode, argb ARGB, clamp win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesWrapMode.Addr(), 
		uintptr(unsafe.Pointer(imageAttr)),
		uintptr(wrap),
		uintptr(argb),
		uintptr(clamp))
	return Status(ret)
}

func SetImageAttributesICMMode(imageAttr *ImageAttributes, on win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesICMMode.Addr(), 
		uintptr(unsafe.Pointer(imageAttr)),
		uintptr(on))
	return Status(ret)
}

func GetImageAttributesAdjustedPalette(imageAttr *ImageAttributes, colorPalette *ColorPalette, colorAdjustType ColorAdjustType) Status {
	ret, _, _ := syscall.SyscallN(gdipGetImageAttributesAdjustedPalette.Addr(), 
		uintptr(unsafe.Pointer(imageAttr)),
		uintptr(unsafe.Pointer(colorPalette)),
		uintptr(colorAdjustType))
	return Status(ret)
}

func SetImageAttributesCachedBackground(imageattr *ImageAttributes, enableFlag win32.BOOL) Status {
	ret, _, _ := syscall.SyscallN(gdipSetImageAttributesCachedBackground.Addr(), 
		uintptr(unsafe.Pointer(imageattr)),
		uintptr(enableFlag))
	return Status(ret)
}

