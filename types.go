package main

import "time"

type DeviceResp struct {
	Status []Device `json:"status"`
}
type Arguments struct {
	Name      string `json:"Name"`
	Type      string `json:"Type"`
	Mandatory bool   `json:"Mandatory"`
}
type Actions struct {
	Function  string      `json:"Function"`
	Name      string      `json:"Name"`
	Arguments []Arguments `json:"Arguments"`
}
type Names struct {
	Name   string `json:"Name"`
	Source string `json:"Source"`
	Suffix string `json:"Suffix"`
	ID     string `json:"Id"`
}
type DeviceTypes struct {
	Type   string `json:"Type"`
	Source string `json:"Source"`
	ID     string `json:"Id"`
}
type Bdd struct {
	CloudVersion        string `json:"CloudVersion"`
	BDDRequestsSent     int    `json:"BDDRequestsSent"`
	BDDRequestsAnswered int    `json:"BDDRequestsAnswered"`
	BDDRequestsFailed   int    `json:"BDDRequestsFailed"`
	DeviceName          string `json:"DeviceName"`
	DeviceType          string `json:"DeviceType"`
	ModelName           string `json:"ModelName"`
	OperatingSystem     string `json:"OperatingSystem"`
	SoftwareVersion     string `json:"SoftwareVersion"`
	Manufacturer        string `json:"Manufacturer"`
	MACVendor           string `json:"MACVendor"`
	DeviceCategory      string `json:"DeviceCategory"`
}
type IPv4Address struct {
	Address       string `json:"Address"`
	Status        string `json:"Status"`
	Scope         string `json:"Scope"`
	AddressSource string `json:"AddressSource"`
	Reserved      bool   `json:"Reserved"`
	ID            string `json:"Id"`
}
type IPv6Address struct {
	Address       string `json:"Address"`
	Status        string `json:"Status"`
	Scope         string `json:"Scope"`
	AddressSource string `json:"AddressSource"`
	ID            string `json:"Id"`
}
type Priority struct {
	Configuration string `json:"Configuration"`
	Type          string `json:"Type"`
}
type Security struct {
	Algorithm string `json:"Algorithm"`
	Score     int    `json:"Score"`
	Scores    []any  `json:"Scores"`
}
type WANAccess struct {
	BlockedReasons string `json:"BlockedReasons"`
}
type MDNSService struct {
	Name        string `json:"Name"`
	ServiceName string `json:"ServiceName"`
	Domain      string `json:"Domain"`
	Port        string `json:"Port"`
	Text        string `json:"Text"`
	ID          string `json:"Id"`
}
type SSWSta struct {
	SupportedStandards           string `json:"SupportedStandards"`
	Supports24GHz                bool   `json:"Supports24GHz"`
	Supports5GHz                 bool   `json:"Supports5GHz"`
	Supports6GHz                 bool   `json:"Supports6GHz"`
	ReconnectClass               string `json:"ReconnectClass"`
	FailedSteerCount             int    `json:"FailedSteerCount"`
	SuccessSteerCount            int    `json:"SuccessSteerCount"`
	AvgSteeringTime              int    `json:"AvgSteeringTime"`
	SupportedUNIIBands           string `json:"SupportedUNIIBands"`
	VendorSpecificElementOUIList string `json:"VendorSpecificElementOUIList"`
}
type Ssw struct {
	Capabilities string `json:"Capabilities"`
	CurrentMode  string `json:"CurrentMode"`
}
type Device struct {
	Key                          string        `json:"Key"`
	DiscoverySource              string        `json:"DiscoverySource"`
	Name                         string        `json:"Name"`
	DeviceType                   string        `json:"DeviceType"`
	Active                       bool          `json:"Active"`
	Tags                         string        `json:"Tags"`
	FirstSeen                    time.Time     `json:"FirstSeen"`
	LastConnection               time.Time     `json:"LastConnection"`
	LastChanged                  time.Time     `json:"LastChanged"`
	Master                       string        `json:"Master"`
	DeviceCategory               string        `json:"DeviceCategory,omitempty"`
	VendorClassID                string        `json:"VendorClassID,omitempty"`
	UserClassID                  string        `json:"UserClassID,omitempty"`
	ClientID                     string        `json:"ClientID,omitempty"`
	SerialNumber                 string        `json:"SerialNumber"`
	ProductClass                 string        `json:"ProductClass,omitempty"`
	Oui                          string        `json:"OUI,omitempty"`
	DHCPOption55                 string        `json:"DHCPOption55,omitempty"`
	ModelName                    string        `json:"ModelName,omitempty"`
	OperatingSystem              string        `json:"OperatingSystem,omitempty"`
	SoftwareVersion              string        `json:"SoftwareVersion,omitempty"`
	Manufacturer                 string        `json:"Manufacturer,omitempty"`
	IPAddress                    string        `json:"IPAddress,omitempty"`
	IPAddressSource              string        `json:"IPAddressSource,omitempty"`
	Location                     string        `json:"Location,omitempty"`
	PhysAddress                  string        `json:"PhysAddress,omitempty"`
	Layer2Interface              string        `json:"Layer2Interface,omitempty"`
	InterfaceName                string        `json:"InterfaceName,omitempty"`
	MACVendor                    string        `json:"MACVendor,omitempty"`
	Owner                        string        `json:"Owner,omitempty"`
	UniqueID                     string        `json:"UniqueID,omitempty"`
	Index                        string        `json:"Index"`
	Actions                      []Actions     `json:"Actions,omitempty"`
	Names                        []Names       `json:"Names"`
	DeviceTypes                  []DeviceTypes `json:"DeviceTypes"`
	Bdd                          Bdd           `json:"BDD,omitempty"`
	ModelNames                   []any         `json:"ModelNames,omitempty"`
	OperatingSystems             []any         `json:"OperatingSystems,omitempty"`
	SoftwareVersions             []any         `json:"SoftwareVersions,omitempty"`
	Manufacturers                []any         `json:"Manufacturers,omitempty"`
	IPv4Address                  []IPv4Address `json:"IPv4Address,omitempty"`
	IPv6Address                  []IPv6Address `json:"IPv6Address,omitempty"`
	Locations                    []any         `json:"Locations,omitempty"`
	Groups                       []any         `json:"Groups,omitempty"`
	Priority                     Priority      `json:"Priority,omitempty"`
	Security                     Security      `json:"Security,omitempty"`
	UserAgents                   []any         `json:"UserAgents,omitempty"`
	WANAccess                    WANAccess     `json:"WANAccess,omitempty"`
	MDNSService                  []MDNSService `json:"mDNSService,omitempty"`
	MDNSRecord                   []any         `json:"mDNSRecord,omitempty"`
	SignalStrength               int           `json:"SignalStrength,omitempty"`
	SignalNoiseRatio             int           `json:"SignalNoiseRatio,omitempty"`
	LastDataDownlinkRate         int           `json:"LastDataDownlinkRate,omitempty"`
	LastDataUplinkRate           int           `json:"LastDataUplinkRate,omitempty"`
	EncryptionMode               string        `json:"EncryptionMode,omitempty"`
	LinkBandwidth                string        `json:"LinkBandwidth,omitempty"`
	SecurityModeEnabled          string        `json:"SecurityModeEnabled,omitempty"`
	HtCapabilities               string        `json:"HtCapabilities,omitempty"`
	VhtCapabilities              string        `json:"VhtCapabilities,omitempty"`
	HeCapabilities               string        `json:"HeCapabilities,omitempty"`
	SupportedMCS                 string        `json:"SupportedMCS,omitempty"`
	AuthenticationState          bool          `json:"AuthenticationState,omitempty"`
	OperatingStandard            string        `json:"OperatingStandard,omitempty"`
	OperatingFrequencyBand       string        `json:"OperatingFrequencyBand,omitempty"`
	AvgSignalStrengthByChain     int           `json:"AvgSignalStrengthByChain,omitempty"`
	MaxBandwidthSupported        string        `json:"MaxBandwidthSupported,omitempty"`
	MaxDownlinkRateSupported     int           `json:"MaxDownlinkRateSupported,omitempty"`
	MaxDownlinkRateReached       int           `json:"MaxDownlinkRateReached,omitempty"`
	DownlinkMCS                  int           `json:"DownlinkMCS,omitempty"`
	DownlinkBandwidth            int           `json:"DownlinkBandwidth,omitempty"`
	DownlinkShortGuard           bool          `json:"DownlinkShortGuard,omitempty"`
	UplinkMCS                    int           `json:"UplinkMCS,omitempty"`
	UplinkBandwidth              int           `json:"UplinkBandwidth,omitempty"`
	UplinkShortGuard             bool          `json:"UplinkShortGuard,omitempty"`
	MaxUplinkRateSupported       int           `json:"MaxUplinkRateSupported,omitempty"`
	MaxUplinkRateReached         int           `json:"MaxUplinkRateReached,omitempty"`
	MaxTxSpatialStreamsSupported int           `json:"MaxTxSpatialStreamsSupported,omitempty"`
	MaxRxSpatialStreamsSupported int           `json:"MaxRxSpatialStreamsSupported,omitempty"`
	SSWSta                       SSWSta        `json:"SSWSta,omitempty"`
	Description                  string        `json:"Description,omitempty"`
	HardwareVersion              string        `json:"HardwareVersion,omitempty"`
	BootLoaderVersion            string        `json:"BootLoaderVersion,omitempty"`
	FirewallLevel                string        `json:"FirewallLevel,omitempty"`
	LinkType                     string        `json:"LinkType,omitempty"`
	LinkState                    string        `json:"LinkState,omitempty"`
	ConnectionProtocol           string        `json:"ConnectionProtocol,omitempty"`
	ConnectionState              string        `json:"ConnectionState,omitempty"`
	LastConnectionError          string        `json:"LastConnectionError,omitempty"`
	ConnectionIPv4Address        string        `json:"ConnectionIPv4Address,omitempty"`
	ConnectionIPv6Address        string        `json:"ConnectionIPv6Address,omitempty"`
	RemoteGateway                string        `json:"RemoteGateway,omitempty"`
	DNSServers                   string        `json:"DNSServers,omitempty"`
	Internet                     bool          `json:"Internet,omitempty"`
	Iptv                         bool          `json:"IPTV,omitempty"`
	Telephony                    bool          `json:"Telephony,omitempty"`
	Alternative                  []string      `json:"Alternative,omitempty"`
	Ssw                          Ssw           `json:"SSW,omitempty"`
	Type                         string        `json:"Type,omitempty"`
	ManufacturerURL              string        `json:"ManufacturerURL,omitempty"`
	ModelDescription             string        `json:"ModelDescription,omitempty"`
	ModelNumber                  string        `json:"ModelNumber,omitempty"`
	ModelURL                     string        `json:"ModelURL,omitempty"`
	Udn                          string        `json:"UDN,omitempty"`
	Upc                          string        `json:"UPC,omitempty"`
	PresentationURL              string        `json:"PresentationURL,omitempty"`
	Server                       string        `json:"Server,omitempty"`
	Service                      []any         `json:"Service,omitempty"`
}
