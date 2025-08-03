windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc garble -literals -tiny build -ldflags="-s -w"  -o amsi.exe ; upx amsi.exe

clean:
	rm -f loader_linux amsi.exe
