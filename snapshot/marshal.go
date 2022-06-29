package snapshot

import "encoding/xml"

func marshalXML(w StaticWorld) ([]byte, error) {
	return xml.MarshalIndent(w, "", "  ")
}
