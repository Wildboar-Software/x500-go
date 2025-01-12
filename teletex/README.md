# Teletex (T.61) String Decoding in Go

This is a library for converting Teletex (T.61) strings to UTF-8. This is
mostly useful for X.500 directories, X.509 PKI and other cryptograhpic matters.
Basically, all you need from this library is `TeletexToUTF8()` which takes the
raw bytes of Teletex and returns a native Golang UTF-8 `string`.

There is no functionality for encoding a string as Teletex. Teletex is an
obsolete text encoding, so you shouldn't use it. Hence this library implements
a one-way ticket to Unicode.

Note that Teletex is basically the same as ASCII for all characters below and
including code point 0x7F (`ESC`). So when encoding a `DirectoryString`, as is
used in X.500 directories and X.509 PKI, `TeletexString` may actually be a
desirable encoding, since it is less strict than `PrintableString`, but still
overlaps with ASCII and is still one-byte-per-character (assuming ASCII only).

See [the Wikipedia page](https://en.wikipedia.org/wiki/ITU_T.61) for more
information about this string encoding.

## X.500 Directory Implementation

If you are interested in working with X.500 directories, consider checking
out [Meerkat DSA](https://wildboar-software.github.io/directory/), which,
to my knowledge, is one of the two free and open source X.500 directory
implementations out there, written by yours truly.

