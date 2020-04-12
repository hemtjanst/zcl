# ZCL

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/hemtjan.st/zcl) [![CI badge](https://github.com/hemtjanst/zcl/workflows/CI/badge.svg)](https://github.com/hemtjanst/zcl/actions?query=workflow%3ACI)

This is a Zigbee™ Cluster Library. It's primarily aimed at Go users but also
contains a TypeScript library.

## Generated code

The packages under `cluster/` and `zdo/` are auto-generated from definitions in
[hemtjanst/zyaml](https://github.com/hemtjanst/zyaml) using the templates in
`template/`.

Regenerate code by running:

```sh
# Create go/typescript source code
go run hemtjan.st/zcl/cmd/zclgen -definition-path ../path-to-zyaml

# Test and build generated code
go test hemtjan.st/zcl/...
go build hemtjan.st/zcl/...

# Transpile typescript
tsc -d ./cluster/all.ts --outFile ../javascript-output/zigbee.js --sourceMap --module amd
```

## Attribution

Zigbee is a trademark of the Zigbee Alliance.
