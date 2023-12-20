package main

import (
	"errors"
	"fmt"
	"github.com/nitram509/gofitz/pkg/scpd"
	"strings"
	"unicode"
)

const spaces = "                                                                                                     "

func string2SnakeCase(str string) string {
	str = strings.ReplaceAll(str, "X_AVM-DE_", "Avm")
	str = strings.ReplaceAll(str, "WebDAV", "Webdav")
	str = strings.ReplaceAll(str, "FRITZ", "Fritz")
	str = strings.ReplaceAll(str, "VoIP", "Voip")
	str = strings.ReplaceAll(str, "WLAN", "Wlan")
	str = strings.ReplaceAll(str, "UUID", "Uuid")
	str = strings.ReplaceAll(str, "SSID", "Ssid")
	str = strings.ReplaceAll(str, "DECT", "Dect")
	str = strings.ReplaceAll(str, "UPnP", "Upnp")
	str = strings.ReplaceAll(str, "DNS", "Dns")
	str = strings.ReplaceAll(str, "FTP", "Ftp")
	str = strings.ReplaceAll(str, "NTP", "Ntp")
	str = strings.ReplaceAll(str, "SMB", "Smb")
	str = strings.ReplaceAll(str, "SID", "Sid")
	str = strings.ReplaceAll(str, "UID", "Uid")
	str = strings.ReplaceAll(str, "URL", "Url")
	str = strings.ReplaceAll(str, "WAN", "Wan")
	str = strings.ReplaceAll(str, "VPN", "Vpn")
	str = strings.ReplaceAll(str, "PFS", "Pfs")
	str = strings.ReplaceAll(str, "DSL", "Dsl")
	str = strings.ReplaceAll(str, "LAN", "Lan")
	str = strings.ReplaceAll(str, "USP", "Usp")
	str = strings.ReplaceAll(str, "TAM", "Tam")
	str = strings.ReplaceAll(str, "MAC", "Mac")
	str = strings.ReplaceAll(str, "CGI", "CGI")
	str = strings.ReplaceAll(str, "IP", "Ip")
	str = strings.ReplaceAll(str, "ID", "Id")
	str = strings.ReplaceAll(str, "TR", "Tr")
	str = strings.ReplaceAll(str, "TV", "Tv")
	str = strings.ReplaceAll(str, "X_", "X")
	sb := strings.Builder{}
	for i := 0; i < len(str); i++ {
		s := str[i]
		if unicode.IsUpper(rune(s)) && (sb.Len() > 0) {
			sb.WriteString("_")
		}
		sb.WriteString(strings.ToLower(string(s)))
	}
	return sb.String()
}

func determineTypeName(spec scpd.ServiceControlledProtocolDescriptions, relatedStateVariable string) string {
	for _, serviceState := range spec.ServiceStateTable {
		if relatedStateVariable == serviceState.Name {
			switch strings.ToLower(serviceState.DataType) {
			// return type with trailing spaces for formatting
			case "string":
				return "string"
			case "datetime":
				return "string"
			case "uuid":
				return "string"
			case "ui4":
				return "int   "
			case "i4":
				return "int   "
			case "ui2":
				return "int   "
			case "ui1":
				return "int   "
			case "boolean":
				return "bool  "
			default:
				panic(errors.New(fmt.Sprintf("relatedStateVariable '%s' unknown type '%s'",
					relatedStateVariable,
					strings.ToLower(serviceState.DataType))))
			}
		}
	}
	panic(errors.New(fmt.Sprintf("relatedStateVariable '%s' not found", relatedStateVariable)))
}

func string2CamelCase(str string) string {
	sb := strings.Builder{}
	// turn into CamelCase
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		if sb.Len() == 0 {
			sb.WriteString(strings.ToUpper(s))
		} else if s == "_" {
			i = i + 1
			s := string(str[i])
			sb.WriteString(strings.ToUpper(s))
		} else {
			sb.WriteString(s)
		}
	}
	return sb.String()
}
