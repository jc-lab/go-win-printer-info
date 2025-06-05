//go:build windows
// +build windows

package winprinter

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sys/windows"
	"unsafe"
)

// Windows DLL 로드
var (
	winspool              = windows.NewLazyDLL("winspool.drv")
	procEnumPrinters      = winspool.NewProc("EnumPrintersW")
	procGetDefaultPrinter = winspool.NewProc("GetDefaultPrinterW")
	procOpenPrinter       = winspool.NewProc("OpenPrinterW")
	procClosePrinter      = winspool.NewProc("ClosePrinter")
	procXcvData           = winspool.NewProc("XcvDataW")
)

// CallOpenPrinter - OpenPrinter API 호출
func CallOpenPrinter(printerName *uint16, hPrinter *windows.Handle, defaults *PRINTER_DEFAULTS) error {
	ret, _, err := procOpenPrinter.Call(
		uintptr(unsafe.Pointer(printerName)),
		uintptr(unsafe.Pointer(hPrinter)),
		uintptr(unsafe.Pointer(defaults)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// CallClosePrinter - ClosePrinter API 호출
func CallClosePrinter(hPrinter windows.Handle) error {
	ret, _, err := procClosePrinter.Call(uintptr(hPrinter))
	if ret == 0 {
		return err
	}
	return nil
}

// CallXcvData - XcvData API 호출
func CallXcvData(
	hXcv windows.Handle,
	dataName *uint16,
	inputData uintptr,
	inputSize uintptr,
	outputData uintptr,
	outputSize uintptr,
	neededSize *uint32,
	status *uint32,
) error {
	ret, _, err := procXcvData.Call(
		uintptr(hXcv),
		uintptr(unsafe.Pointer(dataName)),
		inputData,
		inputSize,
		outputData,
		outputSize,
		uintptr(unsafe.Pointer(neededSize)),
		uintptr(unsafe.Pointer(status)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

func CallGetDefaultPrinter(buffer *uint16, size *uint32) error {
	ret, _, err := procGetDefaultPrinter.Call(
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(size)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// GetDefaultPrinter 함수
func GetDefaultPrinter() (string, error) {
	var size uint32

	if err := CallGetDefaultPrinter(nil, &size); err != nil {
		return "", errors.Wrap(err, "GetDefaultPrinter failed")
	}

	if size == 0 {
		return "", nil
	}

	if int(size) < 0 || int(size) > 1024 {
		return "", fmt.Errorf("GetDefaultPrinter returned invalid size (%d)", size)
	}

	buffer := make([]uint16, size)
	if err := CallGetDefaultPrinter(&buffer[0], &size); err != nil {
		return "", errors.Wrap(err, "GetDefaultPrinter failed")
	}

	if int(size) < 0 || int(size) > len(buffer) {
		return "", fmt.Errorf("GetDefaultPrinter returned invalid size (%d)", size)
	}

	return windows.UTF16ToString(buffer[:size]), nil
}

func CallEnumPrinters(flags uint32, name uintptr, level uint32, pPrinterEnum uintptr, cbBuf uint32, pcbNeeded, pcReturned *uint32) error {
	ret, _, err := procEnumPrinters.Call(
		uintptr(flags),
		name,
		uintptr(level),
		pPrinterEnum,
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(pcbNeeded)),
		uintptr(unsafe.Pointer(pcReturned)),
	)
	if ret == 0 {
		return err
	}
	return nil
}
