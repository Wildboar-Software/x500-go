# X.500 Directory ASN.1 Data Structures in Go

This is a library containing Go equivalents of most of the data structures
described in the International Telecommuncations Union's X.500 series of
specifications for use in X.500 directories. They are decoded and encoded
using Golang's standard library's `encoding/asn1` module.

This library provides a few other conveniences you'll probably want on top
of that.

You can create attribute certificates as described in ITU-T Recommendation
X.509 by using this library. I might provide an example later on, but it's
pretty straight-forward.

## X.500 Directory Implementation

If you are interested in working with X.500 directories, consider checking
out [Meerkat DSA](https://wildboar-software.github.io/directory/), which,
to my knowledge, is one of the two free and open source X.500 directory
implementations out there, written by yours truly.

## Marshaling and Unmarshaling

This library provides `Marshal()` and `Unmarshal()` functions for converting
native structs into X.500 Attributes. This is so you can convert X.500
directory-specific data structures into those useful to your applications
directly.

For example, you can define a `Person` struct as such:

```go
type Person struct {
    CommonName []string `x500:"oid:2.5.4.3"`
    Surname    []string `x500:"oid:2.5.4.4"`
}
```

Then you can serialize an instance of it to X.500 attributes like so:

```go
p := Person{
    CommonName: []string{"Spongebob"},
    Surname:    []string{"Squarepants"},
}
attrs, err := x500.Marshal(p)
if err != nil {
    return err
}
// attrs now contains commonName and surname attributes
```

Unmarshaling is the reverse:

```go
p := Person{}
err := x500.Unmarshal(attrs, &p)
if err != nil {
    return err
}
// p.CommonName[0] == "Spongebob"
// p.Surname[0] == "Squarepants"
```

By default, all values are serialized into attributes if they are annotated
with `x500` tags and have an `oid:` tag to give them an OID, except in a
few cases: empty strings, zero-length object identifiers, and nil pointers
do not produce any attributes. In some cases, if the `must` tag is present
for these, attempting to serialize them with these falsy values will produce
an error instead of silently producing no attributes.

On the other hand, the `omitempty` tag can be used to explicitly request the
elision of attributes when the struct member tagged as such contains a
zero-value for its type. In other words. If a boolean-valued struct field
contains `false`, it will be encoded as an attribute with a `FALSE` `BOOLEAN`
value, but if you tag that field with `omitempty`, this field will only
produce an attribute if its value is `true`. This is in alignment with how
Golang's standard library's `json` module uses its `omitempty` tag, and
indeed where the name for this tag originates. An example:

```go
type Employee struct {
    Name        string `x500:"oid:2.5.4.3"`
    IsStinky    bool   `x500:"oid:1.2.3.6,omitempty"`
}
e := Employee{
    Name: "",
    IsStinky: false
}
attrs, err := x500.Marshal(e)
// len(attrs) == 0
// Name produced no attributes because it is an empty string.
```

Non-slice types are treated as single-valued attributes; if attributes
returned from the directory contain multiple values are unmarshaled into a
field that is not a slice, only the first value will be used. Slice types
properly handle multiple values as you would expect them to. If a value
itself is a SEQUENCE OF or SET OF, you can use the `list` tag to treat the
slice as a single value. This is primarily useful for representing values
having `PostalAddress` syntax. For example:

```go
type PostalThing struct {
    PostalAddress []string `x500:"oid:2.5.4.43,list"`
}
p := PostalThing{
    PostalAddress: []string{
        "P. Sherman",
        "42 Wallaby Way",
        "Sydney, AU",
    },
}
attrs, err := x500.Marshal(p)
// attrs[0] contains only one value, not three!
```

One huge limitation of the marshal/unmarshal API is no support for contexts.
There is one exception: you can marshal values to include language contexts
by using the `lang` and `uselang` tags. Example:

```go
type Person struct {
    CommonName  []string `x500:"oid:2.5.4.3"`
    Description []string `x500:"oid:2.5.4.4,uselang"`
}
p := Person{
    CommonName:     []string{"Jonathan Wilbur"},
    Description:    []string{"Always keepin it crispy"},
}
attrs, err := MarshalWithParams(p, "en")
// Now the description attribute has a value in valuesWithContext, which
// has a single languageContext, which has a context value of "en" to
// indicate that it is English.
```

