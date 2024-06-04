package delivery

import "encoding/xml"

type SoapRequest struct {
	XMLName   xml.Name    `xml:"soapenv:Envelope"`
	XMLNsSoap string      `xml:"xmlns:soapenv,attr"`
	XMLNsUrn  string      `xml:"xmlns:urn,attr"`
	Header    interface{} `xml:"soapenv:Header"`
	Body      interface{} `xml:"soapenv:Body"`
}

type RequestBody struct {
	RequestRecovery RequestRecovery `xml:"urn:ZP2PRFC_GUIA_REMISION_ENTREGA"`
}

type RequestRecovery struct {
	PciMotivo  string `xml:"PCI_MOTIVO"`
	PciTipoRef string `xml:"PCI_TIPO_REF"`
	PciVbeln   string `xml:"PCI_VBELN,omitempty"`
	PciPosnr   string `xml:"PCI_POSNR,omitempty"`
	PciVstel   string `xml:"PCI_VSTEL"`
	PciVgbel   string `xml:"PCI_VGBEL,omitempty"`
	PciErdat   string `xml:"PCI_ERDAT"`
	PciErdat1  string `xml:"PCI_ERDAT1"`
	PciMaktx   string `xml:"PCI_MAKTX,omitempty"`
	PciMatnr   string `xml:"PCI_MATNR,omitempty"`
	PciWerks   string `xml:"PCI_WERKS,omitempty"`
}
