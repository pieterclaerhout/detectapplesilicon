PACKAGE := github.com/pieterclaerhout/detectapplesilicon
OUTFILE := detectapplesilicon

detectapplesilicon: main.go
	GOOS=darwin GOARCH=amd64 gotip build -o $(OUTFILE)-amd64 $(PACKAGE)
	GOOS=darwin GOARCH=arm64 gotip build -o $(OUTFILE)-arm64 $(PACKAGE)
	lipo -create -output $(OUTFILE) $(OUTFILE)-amd64 $(OUTFILE)-arm64
	codesign -s - $(OUTFILE)
	rm -f $(OUTFILE)-amd64 $(OUTFILE)-arm64

run: detectapplesilicon
	@./$(OUTFILE)
