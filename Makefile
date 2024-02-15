default: windows linux
	@echo "Built for all"

executable_name := pkt2xml
flags := -ldflags "-s -w"

windows:
	@echo "Building for Windows"	
	@GOOS=windows GOARCH=amd64 go build -o $(executable_name).exe $(flags) *.go

linux:
	@echo "Building for Linux"
	@GOOS=linux GOARCH=amd64 go build -o $(executable_name) $(flags) *.go

clean:
	rm -f $(executable_name) $(executable_name).exe