package data

import (
	"encoding/xml"
)

type NmapRun struct {
	XMLName          xml.Name `xml:"nmaprun"`
	Scanner          string   `xml:"scanner,attr"`
	Args             string   `xml:"args,attr"`
	Version          string   `xml:"version,attr"`
	XMLOutputVersion string   `xml:"xmloutputversion,attr"`
	Hosts            []Host   `xml:"host"`
	RunStats         RunStats `xml:"runstats"`
}

type Host struct {
	StartTime int64     `xml:"starttime,attr"`
	EndTime   int64     `xml:"endtime,attr"`
	Status    Status    `xml:"status"`
	Address   Address   `xml:"address"`
	Hostnames Hostnames `xml:"hostnames"`
	Ports     Ports     `xml:"ports"`
}

type Status struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL string `xml:"reason_ttl,attr"`
}

type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type Hostnames struct {
	Hostname []Hostname `xml:"hostname"`
}

type Hostname struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type Ports struct {
	PortList []Port `xml:"port"`
}

type Port struct {
	Protocol string    `xml:"protocol,attr"`
	PortID   int       `xml:"portid,attr"`
	State    PortState `xml:"state"`
	Service  *Service  `xml:"service,omitempty"`
}

type PortState struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL string `xml:"reason_ttl,attr"`
}

type Service struct {
	Name   string `xml:"name,attr"`
	Method string `xml:"method,attr"`
	Conf   int    `xml:"conf,attr"`
}

type RunStats struct {
	Finished Finished `xml:"finished"`
	Hosts    Hosts    `xml:"hosts"`
}

type Finished struct {
	Time    int64  `xml:"time,attr"`
	Elapsed string `xml:"elapsed,attr"`
	Summary string `xml:"summary,attr"`
	Exit    string `xml:"exit,attr"`
}

type Hosts struct {
	Up    int `xml:"up,attr"`
	Down  int `xml:"down,attr"`
	Total int `xml:"total,attr"`
}
