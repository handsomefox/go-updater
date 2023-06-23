clean:
	rm -rf ./build
	rm -f archive.tar.gz

build: clean
	git submodule update
	mkdir ./build
	tar czf archive.tar.gz ./vendor/update-golang/
	go build -o ./build/goupdate --ldflags "-s -w" main.go

run: build
	./build/goupdate

install: build
	sudo cp ./build/goupdate /usr/bin/
