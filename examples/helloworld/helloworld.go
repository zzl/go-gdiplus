package main

import (
	"github.com/zzl/go-gdiplus/gdip"
	"github.com/zzl/go-win32api/v2/win32"
	"runtime"
	"syscall"
)

/*
		//equivalent c++ code from msdn

		VOID OnPaint(HDC hdc)
		{
		   Graphics    graphics(hdc);

		   Pen      pen(Color(255, 0, 0, 255));
		   graphics.DrawLine(&pen, 32, 64, 200, 64);

		   SolidBrush  brush(Color(255, 0, 0, 255));
		   FontFamily  fontFamily(L"Times New Roman");
		   Font        font(&fontFamily, 24, FontStyleRegular, UnitPixel);
		   PointF      pointF(30.0f, 30.0f);

		   graphics.DrawString(L"Hello World!", -1, &font, pointF, &brush);
		}
*/
func OnPaint(hdc win32.HDC) {
	var graphics *gdip.Graphics

	gdip.CreateFromHDC(hdc, &graphics)
	defer gdip.DeleteGraphics(graphics)

	var pen *gdip.Pen
	gdip.CreatePen1(gdip.Color(0, 0, 255), 1, gdip.UnitWorld, &pen)
	defer gdip.DeletePen(pen)

	gdip.DrawLine(graphics, pen, 32, 100, 200, 100)

	//
	var brush *gdip.SolidFill
	gdip.CreateSolidFill(gdip.Color(255, 0, 0), &brush)
	defer gdip.DeleteBrush(&brush.Brush)

	var fontFamily *gdip.FontFamily
	gdip.CreateFontFamilyFromName(win32.StrToPwstr("Times New Roman"), nil, &fontFamily)
	defer gdip.DeleteFontFamily(fontFamily)

	var font *gdip.Font
	gdip.CreateFont(fontFamily, 24, gdip.FontStyleRegular, gdip.UnitPixel, &font)
	defer gdip.DeleteFont(font)

	gdip.DrawString(graphics, win32.StrToPwstr("Hello World!"), -1, font,
		&gdip.RectF{32, 64, 0, 0}, nil, &brush.Brush)
}

func main() {

	runtime.LockOSThread()

	gdip.GdiplusStartup()
	defer gdip.GdiplusShutdown()

	hInstance, _ := win32.GetModuleHandle(nil)
	CLASS_NAME := win32.StrToPwstr("GettingStarted")

	//
	var wc win32.WNDCLASS
	wc.LpfnWndProc = syscall.NewCallback(WindowProc)
	wc.HCursor, _ = win32.LoadCursor(win32.NULL, win32.IDC_ARROW)
	wc.HbrBackground = win32.GetStockObject(win32.WHITE_BRUSH)
	wc.LpszClassName = CLASS_NAME

	win32.RegisterClass(&wc)

	//
	hWnd, _ := win32.CreateWindowEx(
		0,
		CLASS_NAME,
		win32.StrToPwstr("Getting Started"),
		win32.WS_OVERLAPPEDWINDOW,
		win32.CW_USEDEFAULT,
		win32.CW_USEDEFAULT,
		400,
		300,
		win32.NULL,
		win32.NULL,
		hInstance,
		nil,
	)

	win32.ShowWindow(hWnd, win32.SW_SHOWDEFAULT)
	win32.UpdateWindow(hWnd)

	//
	var msg win32.MSG
	for {
		ok, _ := win32.GetMessage(&msg, win32.NULL, 0, 0)
		if ok == win32.FALSE {
			break
		}
		win32.TranslateMessage(&msg)
		win32.DispatchMessage(&msg)
	}
}

func WindowProc(hwnd win32.HWND, uMsg win32.UINT, wParam win32.WPARAM, lParam win32.LPARAM) win32.LRESULT {
	switch uMsg {
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		return 0
	case win32.WM_PAINT:
		var ps win32.PAINTSTRUCT
		hdc := win32.BeginPaint(hwnd, &ps)
		OnPaint(hdc)
		win32.EndPaint(hwnd, &ps)
		return 0
	}
	return win32.DefWindowProc(hwnd, uMsg, wParam, lParam)
}
