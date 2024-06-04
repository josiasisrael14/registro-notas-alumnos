package model

import "encoding/xml"

type SoapRequest struct {
	XMLName   xml.Name    `xml:"soapenv:Envelope"`
	XMLNsSoap string      `xml:"xmlns:soapenv,attr"`
	XMLNsUrn  string      `xml:"xmlns:urn,attr"`
	Header    interface{} `xml:"soapenv:Header"`
	Body      interface{} `xml:"soapenv:Body"`
}

type RequestBody struct {
	RequestRecovery RequestRecovery `xml:"urn:ZP2PRFC_GUIA_REMISION_PEDIDO"`
}

type RequestRecovery struct {
	PciAedat   string `xml:"PCI_AEDAT"`
	PciAedat1  string `xml:"PCI_AEDAT1"`
	PciBsart   string `xml:"PCI_BSART,omitempty"`
	PciBukrs   string `xml:"PCI_BUKRS,omitempty"`
	PciEbelen  string `xml:"PCI_EBELN,omitempty"`
	PciEbelp   string `xml:"PCI_EBELP,omitempty"`
	PciLifnr   string `xml:"PCI_LIFNR,omitempty"`
	PciMatnr   string `xml:"PCI_MATNR,omitempty"`
	PciMotivo  string `xml:"PCI_MOTIVO"`
	PciReswk   string `xml:"PCI_RESWK,omitempty"`
	PciTipoRef string `xml:"PCI_TIPO_REF"`
	PciWerks   string `xml:"PCI_WERKS,omitempty"`
}
