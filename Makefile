build: build_server

all: air_build_url_shortener build

clean:
	rm -r tmp/*
	rm -r bin/*

build_server:
	go build -o bin/url-shortener cmd/server/url-shortener.go

copy_config:
	cp .env bin/

copy_data:
	mkdir -p bin/data
	cp data/urlshort.db bin/data/

execute:
	cd bin
	./otp-server

run: build copy_config copy_data execute

air_build_url_shortener:
	go build -o tmp/url-shortener-air cmd/server/url-shortener.go

air_copy_config:
	cp .env tmp/

air_copy_data:
	mkdir -p tmp/data
	cp data/urlshort.db tmp/data/

air: clean_air air_copy_config air_copy_data air_build_url_shortener

clean_air:
	rm -rf tmp/*
