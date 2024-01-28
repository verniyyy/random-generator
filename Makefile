APPNAME:=rgen

build:
	@go build -o bin/$(APPNAME)

install: build
	@chmod 755 bin/$(APPNAME)
	@sudo cp -p bin/$(APPNAME) /usr/bin/