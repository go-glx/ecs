package snapshot

import (
	"bytes"
	"encoding/xml"
)

var xmlOutputReplacePatterns = [][2][]byte{
	{[]byte("></systems>"), []byte("/>")},
	{[]byte("></prop>"), []byte("/>")},
}

func marshalXML(w StaticWorld) ([]byte, error) {
	data, err := xml.MarshalIndent(w, "", "  ")
	if err != nil {
		return nil, err
	}

	return beautify(data), nil
}

func beautify(src []byte) []byte {
	for _, pattern := range xmlOutputReplacePatterns {
		src = bytes.ReplaceAll(src, pattern[0], pattern[1])
	}

	return src
}
