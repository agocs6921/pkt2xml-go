# pkt2xml

A tool for converting PKT/PKA Packet Tracer files to XML and back.

## Usage

### Encryption
*(XML to PKT)*
```sh
pkt2xml encrypt <file>.xml
```

### Decryption
*(PKT to XML)*
```sh
pkt2xml decrypt <file>.pkt
```

## Build

To build for your platform
```sh
go build *.go
```

To build for all platforms supported
```sh
make
```

## Known issues

- Decrypting then re-encrypting the output doesn't result in the same file due to incorrect file size bytes
- Doesn't check if the XML file is a valid Packet Tracer XML file