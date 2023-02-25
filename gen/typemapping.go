package gen

type TypeMapping struct {
	CType    string
	GoType   string
	CallConv string
	GdipType bool
}

var typeMappings map[string]TypeMapping

func initTypeMappings() {
	typeMappings = map[string]TypeMapping{
		"REAL": {
			GoType:   "float32",
			CallConv: "math.Float32bits(@)",
		},
		"HICON": {
			GoType:   "win32.HICON",
			CallConv: "@",
		},
		"LOGFONTW": {
			GoType: "win32.LOGFONT",
		},
		"HANDLE": {
			GoType: "win32.HANDLE",
		},
		"BYTE": {
			GoType: "win32.BYTE",
		},
		"ULONG_PTR": {
			GoType: "win32.ULONG_PTR",
		},
		//ImageCodecInfo?
		"GUID": {
			GoType: "syscall.GUID",
		},
		"CompositingQuality": {
			GoType:   "int32",
			GdipType: true,
		},
		"INT": {
			GoType: "int32",
		},
		"RECT": {
			GoType: "win32.RECT",
		},
		"UINT": {
			GoType: "uint32",
		},
		"HDC": {
			GoType: "win32.HDC",
		},
		"ColorChannelFlags": {
			GoType:   "int32",
			GdipType: true,
		},
		"RotateFlipType": {
			GoType:   "int32",
			GdipType: true,
		},
		"HistogramFormat": {
			GoType:   "int32",
			GdipType: true,
		},
		"StringDigitSubstitute": {
			GoType:   "int32",
			GdipType: true,
		},
		"CompositingMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"GraphicsContainer": {
			GoType:   "uint32",
			GdipType: true,
		},
		"PixelFormat": {
			GoType:   "int32",
			GdipType: true,
		},
		"CustomLineCapType": {
			GoType:   "int32",
			GdipType: true,
		},
		"SmoothingMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"CLSID": {
			GoType: "win32.CLSID",
		},
		"WrapMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"PixelOffsetMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"WarpMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"ImageType": {
			GoType:   "int32",
			GdipType: true,
		},
		"PROPID": {
			GoType: "uint32",
		},
		"StringAlignment": {
			GoType:   "int32",
			GdipType: true,
		},
		"WCHAR": {
			GoType: "uint16",
		},
		"InterpolationMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"StringTrimming": {
			GoType:   "int32",
			GdipType: true,
		},
		"HPALETTE": {
			GoType: "win32.HPALETTE",
		},
		"HINSTANC": {
			GoType: "win32.HINSTANCE",
		},
		"GpTestControlEnum": {
			GoType:   "int32",
			GdipType: true,
		},
		"BOOL": {
			GoType: "win32.BOOL",
		},
		"ColorMatrixFlags": {
			GoType:   "int32",
			GdipType: true,
		},
		"size_t": {
			GoType: "uint",
		},
		"HBITMAP": {
			GoType: "win32.HBITMAP",
		},
		"GraphicsState": {
			GoType:   "uint32",
			GdipType: true,
		},
		"Unit": {
			GoType:   "int32",
			GdipType: true,
		},
		"TextRenderingHint": {
			GoType:   "int32",
			GdipType: true,
		},
		"PaletteType": {
			GoType:   "int32",
			GdipType: true,
		},
		"HRGN": {
			GoType: "win32.HRGN",
		},
		//ColorMatrix?
		//ColorPalette?
		"LPBYTE": {
			GoType: "*byte",
		},
		"LANGID": {
			GoType: "win32.LANGID",
		},
		"HWND": {
			GoType: "win32.HWND",
		},
		"GpHatchStyle": {
			GoType:   "int32",
			GdipType: true,
		},
		"GpWrapMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"DitherType": {
			GoType:   "int32",
			GdipType: true,
		},
		"ColorAdjustType": {
			GoType:   "int32",
			GdipType: true,
		},
		"GpBrushType": {
			GoType:   "int32",
			GdipType: true,
		},
		"UINT16": {
			GoType: "win32.UINT16",
		},
		"GpDashCap": {
			GoType:   "int32",
			GdipType: true,
		},
		"GpLineGradient": {
			GoType:   "int32",
			GdipType: true,
		},
		"GpLineJoin": {
			GoType:   "int32",
			GdipType: true,
		},
		"LinearGradientMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"GpPenType": {
			GoType:   "int32",
			GdipType: true,
		},
		"CombineMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"IStream": {
			GoType: "win32.IStream",
		},
		"IDirectDrawSurface7": {
			GoType: "win32.IUnknown",
		},
		"HINSTANCE": {
			GoType: "win32.HINSTANCE",
		},
		"GpRect": {
			GoType:   "Rect",
			GdipType: true,
		},
		"void": {
			GoType: "byte",
		},
		"VOID": {
			GoType: "byte",
		},
		"HENHMETAFILE": {
			GoType: "win32.HANDLE",
		},
		"BITMAPINFO": {
			GoType: "win32.BITMAPINFO",
		},
		"GpRectF": {
			GoType:   "RectF",
			GdipType: true,
		},
		"GpPointF": {
			GoType:   "PointF",
			GdipType: true,
		},
		"GpPoint": {
			GoType:   "Point",
			GdipType: true,
		},
		"GpPathGradient": {
			GoType:   "GpBrush",
			GdipType: true,
		},
		"GpTexture": {
			GoType:   "GpBrush",
			GdipType: true,
		},
		"GpCoordinateSpace": {
			GoType:   "int32",
			GdipType: true,
		},
		"GpFillMode": {
			GoType:   "int32",
			GdipType: true,
		},
		"EnumerateMetafileProc": {
			GoType: "uintptr",
		},
		"EmfPlusRecordType": {
			GoType:   "int32",
			GdipType: true,
		},
		"EmfType": {
			GoType:   "int32",
			GdipType: true,
		},
		"MetafileFrameUnit": {
			GoType:   "int32",
			GdipType: true,
		},
		"HMETAFILE": {
			GoType: "win32.HANDLE",
		},
		"DrawImageAbort": {
			GoType: "uintptr",
		},
		"GpFlushIntention": {
			GoType:   "int32",
			GdipType: true,
		},
		"GetThumbnailImageAbort": {
			GoType: "uintptr",
		},
	}
}
