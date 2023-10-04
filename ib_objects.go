package main

/*****************************************************************************\
  Originally I used the infoblox-go-client package.  While it worked quite well
  with Add, Update and Delete, Get was quite limited - both a name and a data
  value must always be provided.  I found no way to get all address records with,
  for instance, the IPv4 address xxxxx.  Perhaps I was missing something.
  Additionally, infoblox-go-client does not have a RecordAlias type, and I
  could not figure out how to add boot and PXE options.  Eventually, I backed
  off and went straight to WAPI calls.  Most of these records, though, are
  grabbed straight from infoblox-go-client.
\*****************************************************************************/

// A single string in a TXT record cannot exceed 255 bytes; longer strings will
// need to be broken up into sub-strings.
const maxDataStringSize = 250

type EA map[string]interface{}
type EASearch map[string]interface{}
type IBBase struct{}

type Network struct {
	IBBase
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Ea          EA     `json:"extattrs"`
	Comment     string `json:"comment"`
}

type FixedAddress struct {
	IBBase      `json:"-"`
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Comment     string `json:"comment"`
	IPv4Address string `json:"ipv4addr,omitempty"`
	IPv6Address string `json:"ipv6addr,omitempty"`
	Duid        string `json:"duid,omitempty"`
	Mac         string `json:"mac,omitempty"`
	Name        string `json:"name,omitempty"`
	MatchClient string `json:"match_client,omitempty"`
	Ea          EA     `json:"extattrs"`
	BootFile    string `json:"bootfile,omitempty"`
	BootServer  string `json:"bootserver,omitempty"`
	NextServer  string `json:"nextserver,omitempty"`
	Disable     bool   `json:"disable"`
}

type RecordA struct {
	IBBase   `json:"-"`
	Ref      string `json:"_ref,omitempty"`
	Ipv4Addr string `json:"ipv4addr,omitempty"`
	Name     string `json:"name,omitempty"`
	View     string `json:"view,omitempty"`
	Zone     string `json:"zone,omitempty"`
	UseTtl   bool   `json:"use_ttl"`
	Ttl      uint32 `json:"ttl"`
	Comment  string `json:"comment"`
	Ea       EA     `json:"extattrs"`
	Disable  bool   `json:"disable"`
}

type RecordPTR struct {
	IBBase   `json:"-"`
	Ref      string `json:"_ref,omitempty"`
	Ipv4Addr string `json:"ipv4addr,omitempty"`
	Ipv6Addr string `json:"ipv6addr,omitempty"`
	Name     string `json:"name,omitempty"`
	PtrdName string `json:"ptrdname,omitempty"`
	View     string `json:"view,omitempty"`
	Zone     string `json:"zone,omitempty"`
	Ea       EA     `json:"extattrs"`
	UseTtl   bool   `json:"use_ttl"`
	Ttl      uint32 `json:"ttl"`
	Comment  string `json:"comment"`
	Disable  bool   `json:"disable"`
}

type RecordCNAME struct {
	IBBase    `json:"-"`
	Ref       string `json:"_ref,omitempty"`
	Canonical string `json:"canonical,omitempty"`
	Name      string `json:"name,omitempty"`
	View      string `json:"view,omitempty"`
	Zone      string `json:"zone,omitempty"`
	Ea        EA     `json:"extattrs"`
	Comment   string `json:"comment"`
	UseTtl    bool   `json:"use_ttl"`
	Ttl       uint32 `json:"ttl"`
	Disable   bool   `json:"disable"`
}

type HostRecordIpv4Addr struct {
	IBBase     `json:"-"`
	Ipv4Addr   string `json:"ipv4addr,omitempty"`
	Ref        string `json:"_ref,omitempty"`
	Mac        string `json:"mac,omitempty"`
	Cidr       string `json:"network,omitempty"`
	EnableDhcp bool   `json:"configure_for_dhcp,omitempty"`
	BootFile   string `json:"bootfile,omitempty"`
	BootServer string `json:"bootserver,omitempty"`
	NextServer string `json:"nextserver,omitempty"`
}

type HostRecordIpv6Addr struct {
	IBBase     `json:"-"`
	Ipv6Addr   string `json:"ipv6addr,omitempty"`
	Ref        string `json:"_ref,omitempty"`
	Duid       string `json:"duid"`
	View       string `json:"view,omitempty"`
	Cidr       string `json:"network,omitempty"`
	EnableDhcp bool   `json:"configure_for_dhcp"`
}

type RecordHost struct {
	IBBase      `json:"-"`
	Ref         string               `json:"_ref,omitempty"`
	Ipv4Addr    string               `json:"ipv4addr,omitempty"`
	Ipv4Addrs   []HostRecordIpv4Addr `json:"ipv4addrs"`
	Ipv6Addr    string               `json:"ipv6addr,omitempty"`
	Ipv6Addrs   []HostRecordIpv6Addr `json:"ipv6addrs"`
	Name        string               `json:"name,omitempty"`
	View        string               `json:"view,omitempty"`
	Zone        string               `json:"zone,omitempty"`
	EnableDns   bool                 `json:"configure_for_dns"`
	NetworkView string               `json:"network_view,omitempty"`
	Comment     string               `json:"comment"`
	Ea          EA                   `json:"extattrs"`
	UseTtl      bool                 `json:"use_ttl"`
	Ttl         uint32               `json:"ttl"`
	Aliases     []string             `json:"aliases,omitempty"`
	Disable     bool                 `json:"disable"`
}

type RecordTXT struct {
	IBBase  `json:"-"`
	View    string `json:"view,omitempty"`
	Zone    string `json:"zone,omitempty"`
	Ref     string `json:"_ref,omitempty"`
	Name    string `json:"name,omitempty"`
	Text    string `json:"text,omitempty"`
	Ttl     uint32 `json:"ttl"`
	UseTtl  bool   `json:"use_ttl"`
	Comment string `json:"comment"`
	Ea      EA     `json:"extattrs"`
	Disable bool   `json:"disable"`
}

type RecordMX struct {
	IBBase     `json:"-"`
	Ref        string `json:"_ref,omitempty"`
	View       string `json:"view,omitempty"`
	Fqdn       string `json:"name,omitempty"`
	MX         string `json:"mail_exchanger,omitempty"`
	Preference uint32 `json:"preference"`
	Zone       string `json:"zone,omitempty"`
	Ttl        uint32 `json:"ttl"`
	UseTtl     bool   `json:"use_ttl"`
	Comment    string `json:"comment"`
	Ea         EA     `json:"extattrs"`
	Disable    bool   `json:"disable"`
}

// Define a record type for Aliases; such is missing from infoblox-go-client.
type RecordAlias struct {
	Ref        string `json:"_ref,omitempty"`
	Name       string `json:"name,omitempty"`
	Target     string `json:"target_name,omitempty"`
	TargetType string `json:"target_type,omitempty"`
	View       string `json:"view,omitempty"`
	Comment    string `json:"comment,omitempty"`
	UseTtl     bool   `json:"use_ttl"`
	Ttl        uint32 `json:"ttl"`
	Zone       string `json:"zone,omitempty"`
	Disable    bool   `json:"disable"`
}
