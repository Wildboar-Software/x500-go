# X.500 Directory Access Protocol (DAP) Client in Go

This is an implementation an X.500 Directory Access Protocol (DAP) client in
Go as described in ITU-T Recommendation X.519. It currently only supports the
use of the Internet Directly-Mapped (IDM) protocol also described in the same
standard, but this library was implemented in such a way as to facilitate easy
implementation of OSI transport later on, if desired.

This library was developed and tested against
[Meerkat DSA](https://wildboar-software.github.io/directory/), which, to my
knowledge, is one of two free and open source X.500 directories. If you are
interested in working with X.500 directories, consider checking it out!

## Notes for Users

You will need to set `priority` in the service controls. Golang defaults enums
to 0, even if you use the `default:1` tag on a struct member.

The `timeLimit` field of the service controls is populated by the timeout
specified with the `Context` object. If there is no such timeout specified
either in the service controls, or in the context object, the request will not
timeout unless configured to do so at the TCP or TLS layers. The `sizeLimit` and
`attributeSizeLimit` fields get populated automatically with sensible defaults
unless you supply your own values.

This library supports requesting an attribute certificate from the DSA per the
[private extension used by Meerkat DSA](https://wildboar-software.github.io/directory/docs/attr-cert).

Known issues: https://github.com/golang/go/issues/27426
Any `SEQUENCE OF SET` type will fail to be unmarshalled.
