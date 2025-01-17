package x500

import (
	"encoding/asn1"
	"testing"
)

func TestDirectoryString (t *testing.T) {
  ds1 := asn1.RawValue{
    Tag: asn1.TagT61String,
    Bytes: []byte("Hello"),
  }
  ds2 := asn1.RawValue{
    Tag: asn1.TagPrintableString,
    Bytes: []byte("Hello"),
  }
  ds3 := asn1.RawValue{
    Tag: asn1.TagBMPString,
    Bytes: []byte{0, 'H', 0, 'e', 0, 'l', 0, 'l', 0, 'o'},
  }
  ds4 := asn1.RawValue{
    Tag: 28,
    Bytes: []byte{0, 0, 0, 'H', 0, 0, 0, 'e', 0, 0, 0, 'l', 0, 0, 0, 'l', 0, 0, 0, 'o'},
  }
  ds5 := asn1.RawValue{
    Tag: asn1.TagUTF8String,
    Bytes: []byte("Hello"),
  }

  ds := []asn1.RawValue{ds1, ds2, ds3, ds4, ds5}

  for i, v := range ds {
    str, err := DirectoryStringToString(v)
    if err != nil {
      t.Error(err)
      return
    }
    if str != "Hello" {
      t.Errorf("actual decoded string was %s (len %d, iteration %d)", str, len(str), i)
      return
    }
  }

}
