package navervod

import (
	"encoding/xml"
	"fmt"
	"io"
)

func startElement(dec *xml.Decoder) (start xml.StartElement, err error) {
	var ok bool
	for {
		token, err := dec.Token()
		if err != nil {
			return xml.StartElement{}, err
		}
		start, ok = token.(xml.StartElement)
		if ok {
			return start, nil
		}
	}
}

func DecodeXML(r io.Reader) (interface{}, error) {
	dec := xml.NewDecoder(r)
	start, err := startElement(dec)
	if err != nil {
		return nil, err
	}
	var v interface{}
	switch start.Name.Local {
	case "error":
		v = new(ErrorTag)
	case "videoSource":
		v = new(VideoSourceTag)
	default:
		return nil, fmt.Errorf("cannot interpret <%s> xml document",
			start.Name.Local)
	}
	return v, dec.DecodeElement(v, &start)
}
