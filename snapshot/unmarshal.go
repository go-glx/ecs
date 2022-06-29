package snapshot

import (
	"encoding/xml"
	"fmt"
)

func unmarshalXML(src []byte) (StaticWorld, error) {
	sw := StaticWorld{}
	err := xml.Unmarshal(src, &sw)
	if err != nil {
		return StaticWorld{}, fmt.Errorf("failed unmarshal XML to staticWorld: %w", err)
	}

	return sw, nil
}
