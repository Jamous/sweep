## Build linux binary
go build .
mv sweep bin/linux_x86/

## Build windows binary
GOOS=windows GOARCH=amd64 go build .
mv sweep.exe bin/windows_x86
