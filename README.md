## Generated code

The packages under `cluster/` and `zdo/` are auto-generated from definitions in [github.com/hemtjanst/zyaml](https://github.com/hemtjanst/zyaml) 
using the templates in `template/`  

Regenerate code by running:
```
# Create go/typescript source code
go run hemtjan.st/zcl/cmd/zclgen -definition-path ../path-to-zyaml

# Test and build generated code
go test hemtjan.st/zcl/...
go build hemtjan.st/zcl/...

# Transpile typescript
tsc -d ./cluster/all.ts --outFile ../javascript-output/zigbee.js --sourceMap --module amd
```


