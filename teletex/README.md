# Teletex (T.61) String Decoding in Go

Basically, all you need from this library is `TeletexToUTF8()` which takes the
raw bytes of Teletex and returns a native Golang UTF-8 `string`.

There is no functionality for encoding a string as Teletex. Teletex is an
obsolete text encoding, so you shouldn't use it. Hence this library implements
a one-way ticket to Unicode.

This is mostly useful for X.509 certificates and other cryptographic artifacts,
as well as X.500 directories.

Note that Teletex is basically the same as ASCII for all characters below and
including code point 0x7F (`ESC`). So when encoding a `DirectoryString`, as is
used in X.500 directories and X.509 PKI, `TeletexString` may actually be a
desirable encoding, since it is less strict than `PrintableString`, but still
overlaps with ASCII and is still one-byte-per-character (assuming ASCII only).

See [the Wikipedia page](https://en.wikipedia.org/wiki/ITU_T.61) for more
information about this string encoding.

