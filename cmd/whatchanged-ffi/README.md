### macOS

```bash
go build -buildmode=c-shared -o dist/whatchanged.dylib github.com/release-lab/whatchanged/cmd/whatchanged-ffi
```

### windows

```bash
go build -buildmode=c-shared -o dist/whatchanged.dll github.com/release-lab/whatchanged/cmd/whatchanged-ffi
```

### linux

```bash
go build -buildmode=c-shared -o dist/whatchanged.so github.com/release-lab/whatchanged/cmd/whatchanged-ffi
```