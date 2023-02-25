package gdip

type Status int32
const (
	Ok                        Status = 0
	GenericError              Status = 1
	InvalidParameter          Status = 2
	OutOfMemory               Status = 3
	ObjectBusy                Status = 4
	InsufficientBuffer        Status = 5
	NotImplemented            Status = 6
	Win32Error                Status = 7
	WrongState                Status = 8
	Aborted                   Status = 9
	FileNotFound              Status = 10
	ValueOverflow             Status = 11
	AccessDenied              Status = 12
	UnknownImageFormat        Status = 13
	FontFamilyNotFound        Status = 14
	FontStyleNotFound         Status = 15
	NotTrueTypeFont           Status = 16
	UnsupportedGdiplusVersion Status = 17
	GdiplusNotInitialized     Status = 18
	PropertyNotFound          Status = 19
	PropertyNotSupported      Status = 20
	ProfileNotFound           Status = 21
)

type FillMode int32
const (
	FillModeAlternate FillMode = iota // 0
	FillModeWinding                   // 1
)

type WrapMode int32
const (
	WrapModeTile WrapMode = iota
	WrapModeTileFlipX
	WrapModeTileFlipY
	WrapModeTileFlipXY
	WrapModeClamp
)

type Unit int32
const (
	UnitWorld      Unit = iota // 0 -- World coordinate (non-physical unit)
	UnitDisplay                  // 1 -- Variable -- for PageTransform only
	UnitPixel                    // 2 -- Each unit is one device pixel.
	UnitPoint                    // 3 -- Each unit is a printer's point, or 1/72 inch.
	UnitInch                     // 4 -- Each unit is 1 inch.
	UnitDocument                 // 5 -- Each unit is 1/300 inch.
	UnitMillimeter               // 6 -- Each unit is 1 millimeter.
)


type QualityMode int32
const (
	QualityModeInvalid QualityMode = iota - 1
	QualityModeDefault
	QualityModeLow  // Best performance
	QualityModeHigh // Best rendering quality
)

type CompositingMode int32
const (
	CompositingModeSourceOver CompositingMode = iota // 0
	CompositingModeSourceCopy        // 1
)

type CompositingQuality int32
const (
	CompositingQualityInvalid CompositingQuality = iota + CompositingQuality(QualityModeInvalid)
	CompositingQualityDefault
	CompositingQualityHighSpeed
	CompositingQualityHighQuality
	CompositingQualityGammaCorrected
	CompositingQualityAssumeLinear
)

type InterpolationMode int32
const (
	InterpolationModeInvalid InterpolationMode = iota + InterpolationMode(QualityModeInvalid)
	InterpolationModeDefault
	InterpolationModeLowQuality
	InterpolationModeHighQuality
	InterpolationModeBilinear
	InterpolationModeBicubic
	InterpolationModeNearestNeighbor
	InterpolationModeHighQualityBilinear
	InterpolationModeHighQualityBicubic
)

type SmoothingMode int32
const (
	SmoothingModeInvalid     SmoothingMode = -1
	SmoothingModeDefault     SmoothingMode = 0
	SmoothingModeHighSpeed   SmoothingMode = 1
	SmoothingModeHighQuality SmoothingMode = 2
	SmoothingModeNone        SmoothingMode = 3
	SmoothingModeAntiAlias   SmoothingMode = 4

	SmoothingModeAntiAlias8x4 SmoothingMode = 4
	SmoothingModeAntiAlias8x8 SmoothingMode = 5
)

type PixelOffsetMode int32
const (
	PixelOffsetModeInvalid PixelOffsetMode = iota + PixelOffsetMode(QualityModeInvalid)
	PixelOffsetModeDefault
	PixelOffsetModeHighSpeed
	PixelOffsetModeHighQuality
	PixelOffsetModeNone // No pixel offset
	PixelOffsetModeHalf // Offset by -0.5, -0.5 for fast anti-alias perf
)

type TextRenderingHint int32
const (
	TextRenderingHintSystemDefault TextRenderingHint            = iota // Glyph with system default rendering hint
	TextRenderingHintSingleBitPerPixelGridFit        // Glyph bitmap with hinting
	TextRenderingHintSingleBitPerPixel               // Glyph bitmap without hinting
	TextRenderingHintAntiAliasGridFit                // Glyph anti-alias bitmap with hinting
	TextRenderingHintAntiAlias                       // Glyph anti-alias bitmap without hinting
	TextRenderingHintClearTypeGridFit                // Glyph CT bitmap with hinting
)


type BrushType int32
const (
	BrushTypeSolidColor BrushType = iota
	BrushTypeHatchFill
	BrushTypeTextureFill
	BrushTypePathGradient
	BrushTypeLinearGradient
)

type LineCap int32
const (
	LineCapFlat          LineCap = 0
	LineCapSquare        LineCap = 1
	LineCapRound         LineCap = 2
	LineCapTriangle      LineCap = 3
	LineCapNoAnchor      LineCap = 0x10
	LineCapSquareAnchor  LineCap = 0x11
	LineCapRoundAnchor   LineCap = 0x12
	LineCapDiamondAnchor LineCap = 0x13
	LineCapArrowAnchor   LineCap = 0x14
	LineCapCustom        LineCap = 0xff
	LineCapAnchorMask    LineCap = 0xf0
)

type LineJoin int32
const (
	LineJoinMiter LineJoin = iota
	LineJoinBevel
	LineJoinRound
	LineJoinMiterClipped
)

type DashCap int32
const (
	DashCapFlat DashCap = iota
	DashCapRound
	DashCapTriangle
)

type DashStyle int32
const (
	DashStyleSolid DashStyle = iota
	DashStyleDash
	DashStyleDot
	DashStyleDashDot
	DashStyleDashDotDot
	DashStyleCustom
)

type PenAlignment int32
const (
	PenAlignmentCenter PenAlignment = iota
	PenAlignmentInset
)

type MatrixOrder int32
const (
	MatrixOrderPrepend MatrixOrder = iota
	MatrixOrderAppend
)

type PenType int32
const (
	PenTypeSolidColor PenType = iota
	PenTypeHatchFill
	PenTypeTextureFill
	PenTypePathGradient
	PenTypeLinearGradient
	PenTypeUnknown PenType = -1
)

type FlushIntention int32
const (
	FlushIntentionFlush FlushIntention = iota
	FlushIntentionSync
)

type CoordinateSpace int32
const (
	CoordinateSpaceWorld CoordinateSpace = iota
	CoordinateSpacePage
	CoordinateSpaceDevice
)

type CombineMode int32
const (
	CombineModeReplace CombineMode = iota
	CombineModeIntersect
	CombineModeUnion
	CombineModeXor
	CombineModeExclude
	CombineModeComplement
)

type ColorAdjustType int32
const (
	ColorAdjustTypeDefault ColorAdjustType = iota
	ColorAdjustTypeBitmap
	ColorAdjustTypeBrush
	ColorAdjustTypePen
	ColorAdjustTypeText
	ColorAdjustTypeCount
	ColorAdjustTypeAny
)

type ImageLockMode int32
const (
	ImageLockModeRead         ImageLockMode = 0x0001
	ImageLockModeWrite        ImageLockMode = 0x0002
	ImageLockModeUserInputBuf ImageLockMode = 0x0004
)

type HatchStyle int32
const (
	HatchStyleHorizontal HatchStyle = iota
	HatchStyleVertical
	HatchStyleForwardDiagonal
	HatchStyleBackwardDiagonal
	HatchStyleCross
	HatchStyleDiagonalCross
	HatchStyle05Percent
	HatchStyle10Percent
	HatchStyle20Percent
	HatchStyle25Percent
	HatchStyle30Percent
	HatchStyle40Percent
	HatchStyle50Percent
	HatchStyle60Percent
	HatchStyle70Percent
	HatchStyle75Percent
	HatchStyle80Percent
	HatchStyle90Percent
	HatchStyleLightDownwardDiagonal
	HatchStyleLightUpwardDiagonal
	HatchStyleDarkDownwardDiagonal
	HatchStyleDarkUpwardDiagonal
	HatchStyleWideDownwardDiagonal
	HatchStyleWideUpwardDiagonal
	HatchStyleLightVertical
	HatchStyleLightHorizontal
	HatchStyleNarrowVertical
	HatchStyleNarrowHorizontal
	HatchStyleDarkVertical
	HatchStyleDarkHorizontal
	HatchStyleDashedDownwardDiagonal
	HatchStyleDashedUpwardDiagonal
	HatchStyleDashedHorizontal
	HatchStyleDashedVertical
	HatchStyleSmallConfetti
	HatchStyleLargeConfetti
	HatchStyleZigZag
	HatchStyleWave
	HatchStyleDiagonalBrick
	HatchStyleHorizontalBrick
	HatchStyleWeave
	HatchStylePlaid
	HatchStyleDivot
	HatchStyleDottedGrid
	HatchStyleDottedDiamond
	HatchStyleShingle
	HatchStyleTrellis
	HatchStyleSphere
	HatchStyleSmallGrid
	HatchStyleSmallCheckerBoard
	HatchStyleLargeCheckerBoard
	HatchStyleOutlinedDiamond
	HatchStyleSolidDiamond
	HatchStyleTotal
	HatchStyleLargeGrid HatchStyle = HatchStyleCross
	HatchStyleMin       HatchStyle = HatchStyleHorizontal
	HatchStyleMax       HatchStyle = HatchStyleTotal - 1
)

type WarpMode int32
const (
	WarpModePerspective WarpMode = iota
	WarpModeBilinear
)

type StringFormatFlags uint32
const (
	StringFormatFlagsDirectionRightToLeft  StringFormatFlags = 0x00000001
	StringFormatFlagsDirectionVertical     StringFormatFlags = 0x00000002
	StringFormatFlagsNoFitBlackBox         StringFormatFlags = 0x00000004
	StringFormatFlagsDisplayFormatControl  StringFormatFlags = 0x00000020
	StringFormatFlagsNoFontFallback        StringFormatFlags = 0x00000400
	StringFormatFlagsMeasureTrailingSpaces StringFormatFlags = 0x00000800
	StringFormatFlagsNoWrap                StringFormatFlags = 0x00001000
	StringFormatFlagsLineLimit             StringFormatFlags = 0x00002000

	StringFormatFlagsNoClip    StringFormatFlags = 0x00004000
	StringFormatFlagsBypassGDI StringFormatFlags = 0x80000000
)

type StringAlignment int32
const (
	StringAlignmentNear StringAlignment = iota
	StringAlignmentCenter
	StringAlignmentFar
)

type HotkeyPrefix int32
const (
	HotkeyPrefixNone HotkeyPrefix = iota
	HotkeyPrefixShow
	HotkeyPrefixHide
)

type StringTrimming int32
const (
	StringTrimmingNone StringTrimming = iota
	StringTrimmingCharacter
	StringTrimmingWord
	StringTrimmingEllipsisCharacter
	StringTrimmingEllipsisWord
	StringTrimmingEllipsisPath
)

type StringDigitSubstitute int32
const (
	StringDigitSubstituteUser StringDigitSubstitute = iota
	StringDigitSubstituteNone
	StringDigitSubstituteNational
	StringDigitSubstituteTraditional
)

type FontStyle int32
const (
	FontStyleRegular    FontStyle = 0
	FontStyleBold       FontStyle = 1
	FontStyleItalic     FontStyle = 2
	FontStyleBoldItalic FontStyle = 3
	FontStyleUnderline  FontStyle = 4
	FontStyleStrikeout  FontStyle = 8
)

type RotateFlipType int32
const (
	RotateNoneFlipNone RotateFlipType = iota
	Rotate90FlipNone
	Rotate180FlipNone
	Rotate270FlipNone
	RotateNoneFlipX
	Rotate90FlipX
	Rotate180FlipX
	Rotate270FlipX
	RotateNoneFlipY  RotateFlipType = Rotate180FlipX
	Rotate90FlipY    RotateFlipType = Rotate270FlipX
	Rotate180FlipY   RotateFlipType = RotateNoneFlipX
	Rotate270FlipY   RotateFlipType = Rotate90FlipX
	RotateNoneFlipXY RotateFlipType = Rotate180FlipNone
	Rotate90FlipXY   RotateFlipType = Rotate270FlipNone
	Rotate180FlipXY  RotateFlipType = RotateNoneFlipNone
	Rotate270FlipXY  RotateFlipType = Rotate90FlipNone
)

type DitherType int32
const (
	DitherTypeNone DitherType = iota
	DitherTypeSolid
	DitherTypeOrdered4x4
	DitherTypeOrdered8x8
	DitherTypeOrdered16x16
	DitherTypeSpiral4x4
	DitherTypeSpiral8x8
	DitherTypeDualSpiral4x4
	DitherTypeDualSpiral8x8
	DitherTypeErrorDiffusion
	DitherTypeMax
)

type PaletteType int32
const (
	PaletteTypeCustom PaletteType = iota
	PaletteTypeOptimal
	PaletteTypeFixedBW
	PaletteTypeFixedHalftone8
	PaletteTypeFixedHalftone27
	PaletteTypeFixedHalftone64
	PaletteTypeFixedHalftone125
	PaletteTypeFixedHalftone216
	PaletteTypeFixedHalftone252
	PaletteTypeFixedHalftone256
)

type HistogramFormat int32
const (
	HistogramFormatARGB HistogramFormat = iota
	HistogramFormatPARGB
	HistogramFormatRGB
	HistogramFormatGray
	HistogramFormatB
	HistogramFormatG
	HistogramFormatR
	HistogramFormatA
)

type GpTestControlEnum int32
const (
	TestControlForceBilinear GpTestControlEnum = iota
	TestControlNoICM
	TestControlGetBuildNumber
)

type LinearGradientMode int32
const (
	LinearGradientModeHorizontal LinearGradientMode = iota
	LinearGradientModeVertical
	LinearGradientModeForwardDiagonal
	LinearGradientModeBackwardDiagonal
)

type ImageType int32
const (
	ImageTypeUnknown ImageType = iota
	ImageTypeBitmap
	ImageTypeMetafile
)

type ColorMatrixFlags int32
const (
	ColorMatrixFlagsDefault ColorMatrixFlags = iota
	ColorMatrixFlagsSkipGrays
	ColorMatrixFlagsAltGray
)

type ColorChannelFlags int32
const (
	ColorChannelFlagsC ColorChannelFlags = iota
	ColorChannelFlagsM
	ColorChannelFlagsY
	ColorChannelFlagsK
	ColorChannelFlagsLast
)

type CustomLineCapType int32
const (
	CustomLineCapTypeDefault CustomLineCapType = iota
	CustomLineCapTypeAdjustableArrow
)

//
type EmfPlusRecordType int32
const (
	//...
)

type EmfType int32
const (
	//...
)

type MetafileFrameUnit int32
const (
	//...
)
