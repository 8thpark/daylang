build:
	go build -o daylang-server .

docker:
	docker run -p 5174:5174 daylang:local

image:
	docker build -t daylang:local .

run: build
	./daylang-server
