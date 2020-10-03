build:
	docker build . -t go-dock

run:
	docker run -v /home/north/projects/golangpublic/public:/dist/public -p 3000:3000 go-dock 

.PHONY: build run