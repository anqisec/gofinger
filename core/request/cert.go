package request

import (
	"crypto/tls"
	"net/url"
	"strings"
)

func getCertContent(uri string) string {
	uri = getAddr(uri)
	if len(uri) == 0 {
		return ""
	}
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", uri, tlsConfig)
	if err != nil {
		return ""
	}
	defer conn.Close()
	state := conn.ConnectionState()
	var builder strings.Builder
	for _, cert := range state.PeerCertificates {
		builder.WriteString("Issuer: ")
		builder.WriteString(cert.Issuer.String())
		builder.WriteString("Subject: ")
		builder.WriteString(cert.Subject.String())
		builder.WriteString("CommonName: ")
		builder.WriteString(cert.Issuer.CommonName)
		if len(cert.Issuer.OrganizationalUnit) != 0 {
			builder.WriteString("Organizational Unit: ")
			builder.WriteString(cert.Issuer.OrganizationalUnit[0])
		}
		if len(cert.Issuer.Organization) != 0 {
			builder.WriteString("Organization: ")
			builder.WriteString(cert.Issuer.Organization[0])
		}
		if len(cert.Issuer.Locality) != 0 {
			builder.WriteString("Locality: ")
			builder.WriteString(cert.Issuer.Locality[0])
		}
		builder.WriteString("CommonName: ")
		builder.WriteString(cert.Subject.CommonName)
		if len(cert.Subject.OrganizationalUnit) != 0 {
			builder.WriteString("Organizational Unit: ")
			builder.WriteString(cert.Subject.OrganizationalUnit[0])
		}
		if len(cert.Subject.Organization) != 0 {
			builder.WriteString("Organization: ")
			builder.WriteString(cert.Subject.Organization[0])
		}
		if len(cert.Subject.Organization) != 0 {
			builder.WriteString("Organization: ")
			builder.WriteString(cert.Subject.Organization[0])
		}
		if len(cert.Subject.Locality) != 0 {
			builder.WriteString("Locality: ")
			builder.WriteString(cert.Subject.Locality[0])
		}
	}
	return builder.String()
}
func getAddr(uri string) string {
	var builder strings.Builder
	parse, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	Scheme := parse.Scheme
	if Scheme != "https" {
		return ""
	}
	port := parse.Port()
	builder.WriteString(parse.Hostname())
	builder.WriteString(":")
	if len(port) == 0 {
		builder.WriteString("443")
	} else {
		builder.WriteString(port)
	}
	return builder.String()
}
