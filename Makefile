build:
	go build -o daylang-server .

docker:
	docker run -p 5173:5173 daylang:local

image:
	docker build -t daylang:local .

run: build
	./daylang-server
