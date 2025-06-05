package winprinter

// Windows API constants
const (
	PRINTER_ENUM_LOCAL        = 0x00000002
	PRINTER_ENUM_CONNECTIONS  = 0x00000004
	ERROR_INSUFFICIENT_BUFFER = 122
	SERVER_ACCESS_ADMINISTER  = 0x01
)

// PRINTER_INFO_2 https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-info-2
type PRINTER_INFO_2 struct {
	ServerName         *uint16
	PrinterName        *uint16
	ShareName          *uint16
	PortName           *uint16
	DriverName         *uint16
	Comment            *uint16
	Location           *uint16
	DevMode            *byte
	SepFile            *uint16
	PrintProcessor     *uint16
	Datatype           *uint16
	Parameters         *uint16
	SecurityDescriptor *byte
	Attributes         uint32
	Priority           uint32
	DefaultPriority    uint32
	StartTime          uint32
	UntilTime          uint32
	Status             uint32
	CJobs              uint32
	AveragePPM         uint32
}

// PRINTER_DEFAULTS https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-defaults
type PRINTER_DEFAULTS struct {
	Datatype      *uint16
	DevMode       *byte
	DesiredAccess uint32
}

// CONFIG_INFO_DATA_1 https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/tcpxcv/ns-tcpxcv-_config_info_data_1
type CONFIG_INFO_DATA_1 struct {
	Reserved [128]byte
	Version  uint32
}

const MAX_PORTNAME_LEN = 64
const MAX_NETWORKNAME_LEN = 49
const MAX_SNMP_COMMUNITY_STR_LEN = 33
const MAX_QUEUENAME_LEN = 33
const MAX_IPADDR_STR_LEN = 16
const RESERVED_BYTE_ARRAY_SIZE = 540

// PORT_DATA_1 https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/tcpxcv/ns-tcpxcv-_port_data_1
type PORT_DATA_1 struct {
	PortName      [MAX_PORTNAME_LEN]uint16
	Version       uint32
	Protocol      uint32
	CbSize        uint32
	Reserved      uint32
	HostAddress   [MAX_NETWORKNAME_LEN]uint16
	SNMPCommunity [MAX_SNMP_COMMUNITY_STR_LEN]uint16
	DoubleSpool   uint32
	Queue         [MAX_QUEUENAME_LEN]uint16
	IPAddress     [MAX_IPADDR_STR_LEN]uint16
	Reserved2     [RESERVED_BYTE_ARRAY_SIZE]byte
	PortNumber    uint32
	SNMPEnabled   uint32
	SNMPDevIndex  uint32
}
