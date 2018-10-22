# not used yet
test:
	go test ./comms

# build folie for several platforms and compress for release
builds: version.go
	@ rm -rf $@; mkdir $@
	@ echo Re-building binaries:
	@ echo "  MacOSX 64-bit"
	@ GOOS=darwin GOARCH=amd64 go build -a -o $@/folie-macos64
	@ echo "  Windows 32-bit"
	@ GOOS=windows GOARCH=386 go build -a -o $@/folie-windows.exe
	@ echo "  Windows 64-bit"
	@ GOOS=windows GOARCH=amd64 go build -a -o $@/folie-windows64.exe
	@ echo "  Linux 32-bit"
	@ GOOS=linux GOARCH=386 go build -a -o $@/folie-linux
	@ echo "  Linux 64-bit"
	@ GOOS=linux GOARCH=amd64 go build -a -o $@/folie-linux64
	@ echo "  ARMv6 32-bit"
	@ GOOS=linux GOARCH=arm GOARM=6 go build -a -o $@/folie-arm
#	@ echo "  ARMv8 64-bit"
#	@ GOOS=linux GOARCH=arm64 go build -a -o $@/folie-arm64
	@ echo "  FreeBSD 64-bit"
	@ GOOS=freebsd GOARCH=amd64 go build -a -o $@/folie-freebsd64
	@ gzip $@/folie-*
	@ echo; ls -l builds/*; echo

version.go:
	@ echo "VERSION = `git describe --tags`"
	@ sed -i '' -e "s/VERSION.*/VERSION = \"`git describe --tags`\"/" $@

clean:
	rm -rf folie folie.exe builds

.PHONY: test builds version.go clean
