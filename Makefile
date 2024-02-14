default: windows linux
	@echo "Built for all"

flags := -ldflags "-s -w"

windows:
	@echo "Building for Windows"	
	@GOOS=windows GOARCH=amd64 go build -o pkt2xml.exe $(flags) *.go

linux:
	@echo "Building for Linux"
	@GOOS=linux GOARCH=amd64 go build -o pkt2xml $(flags) *.go

clean:
	rm -f pkt2xml pkt2xml.exe