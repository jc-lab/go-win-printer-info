package main

import (
	"fmt"
	win_printer_info "github.com/jc-lab/go-win-printer-info"
	"log"
	"strings"
)

func main() {
	helper := win_printer_info.NewPrinterHelper()

	fmt.Println("프린터 정보를 수집 중...")

	err := helper.GetPrinters()
	if err != nil {
		log.Fatalf("프린터 정보 수집 실패: %v", err)
	}

	fmt.Printf("\n발견된 프린터 수: %d\n", len(helper.PrinterList))
	fmt.Println(strings.Repeat("=", 80))

	for _, printer := range helper.PrinterList {
		fmt.Printf("프린터 #%d:\n", printer.Index)
		fmt.Printf("  이름: %s\n", printer.PrinterName)
		fmt.Printf("  기본 프린터: %v\n", printer.DefaultPrinter)
		if printer.ShareName != "" {
			fmt.Printf("  공유 이름: %s\n", printer.ShareName)
		}
		fmt.Printf("  포트: %s\n", printer.PortName)
		if printer.PrinterIP != "" {
			fmt.Printf("  IP 주소: %s\n", printer.PrinterIP)
		}
		if printer.Port != 0 {
			fmt.Printf("  포트 번호: %d\n", printer.Port)
		}
		fmt.Printf("  드라이버: %s\n", printer.DriverName)
		if printer.Location != "" {
			fmt.Printf("  위치: %s\n", printer.Location)
		}
		if printer.PrinterUUID != "" {
			fmt.Printf("  UUID: %s\n", printer.PrinterUUID)
		}
		if printer.ModelName != "" {
			fmt.Printf("  모델명: %s\n", printer.ModelName)
		}
		fmt.Printf("  상태: %d\n", printer.Status)
		fmt.Printf("  SNMP 활성화: %v\n", printer.SNMPEnable)
		fmt.Println(strings.Repeat("-", 40))
	}

	// 기본 프린터 정보 출력
	defaultPrinter, err := helper.GetDefaultPrinter()
	if err == nil && defaultPrinter != "" {
		fmt.Printf("\n기본 프린터: %s\n", defaultPrinter)
	}
}

// EXAMPLE:
// 발견된 프린터 수: 6
//================================================================================
//프린터 #1:
//  이름: IP - Epson ESC/P-R V4 Class Driver
//  기본 프린터: false
//  포트: 192.168.3.135_1
//  IP 주소: 192.168.3.135
//  포트 번호: 9100
//  드라이버: Epson ESC/P-R V4 Class Driver
//  상태: 0
//  SNMP 활성화: true
//----------------------------------------
//프린터 #2:
//  이름: WSD - EPSON3DEF75 (L6290 Series)
//  기본 프린터: false
//  포트: WSD-c4dc4c24-9590-4320-bdb7-f499aa9786a3
//  IP 주소: 192.168.1.100
//  드라이버: Epson ESC/P-R V4 Class Driver
//  위치: http://192.168.3.135:80/WSD/DEVICE
//  UUID: cfe92100-67c4-11d4-a45f-e0bb9e3def75
//  모델명: Epson WorkForce/L6290
//  상태: 0
//  SNMP 활성화: false
//----------------------------------------
//프린터 #3:
//  이름: OneNote for Windows 10
//  기본 프린터: false
//  포트: Microsoft.Office.OneNote_16001.14326.22348.0_x64__8wekyb3d8bbwe_microsoft.onenoteim_S-1-5-21-...
//  드라이버: Microsoft Software Printer Driver
//  상태: 0
//  SNMP 활성화: false
//----------------------------------------
//프린터 #4:
//  이름: Microsoft XPS Document Writer
//  기본 프린터: false
//  포트: PORTPROMPT:
//  드라이버: Microsoft XPS Document Writer v4
//  상태: 0
//  SNMP 활성화: false
//----------------------------------------
//프린터 #5:
//  이름: Microsoft Print to PDF
//  기본 프린터: false
//  포트: PORTPROMPT:
//  드라이버: Microsoft Print To PDF
//  상태: 0
//  SNMP 활성화: false
//----------------------------------------
//프린터 #6:
//  이름: Fax
//  기본 프린터: false
//  포트: SHRFAX:
//  IP 주소: SHRFAX
//  포트 번호: 9100
//  드라이버: Microsoft Shared Fax Driver
//  상태: 0
//  SNMP 활성화: false
//----------------------------------------
