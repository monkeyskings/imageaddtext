#GOENV SET
GOCMD=go
GOBUILD=${GOCMD} build
#make
default:
	mkdir -p ./editimage
	$(GOBUILD)  -o imagetools main.go
	mv imagetools ./editimage
	cp -r config ./editimage
	cp -r font ./editimage
	cp -r output ./editimage
	cp -r data ./editimage

#make clean
clean:
	rm -rf ./editimage