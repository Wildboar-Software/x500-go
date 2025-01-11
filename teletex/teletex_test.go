// SPDX-License-Identifier: MIT
package teletex

import (
	"testing"
)

func TestItReturnsPureASCIIUnchanged(t *testing.T) {
	input := "Jonathan M. Wilbur"
	if TeletexToUTF8([]byte(input)) != input {
		t.FailNow()
	}
}

func TestTranslatesUnequivalentCharacters(t *testing.T) {
	input := []byte{'B', 'i', 'g', 0xA4, 'M', 'o', 'n', 'e', 'y', 0xA4}
	output := TeletexToUTF8(input)
	if output != "Big$Money$" {
		t.FailNow()
	}
}

func TestTransposesAndTranslatesDiacritics(t *testing.T) {
	input := []byte{'B', 'i', 'g', 'B', 0xC4, 'o', 0xC5, 'i'}
	output := TeletexToUTF8(input)
	if output != "BigBo\u0303i\u0304" {
		t.FailNow()
	}
}

func TestEmptyString(t *testing.T) {
	input := []byte{}
	output := TeletexToUTF8(input)
	if output != "" {
		t.FailNow()
	}
}

//     fn it_decodes_an_empty_string() {
//         let input = "";
//         match teletex_to_utf8(input.as_bytes()) {
//             Cow::Borrowed(s) => assert_eq!(s, input),
//             _ => panic!(),
//         }
//     }

// }
