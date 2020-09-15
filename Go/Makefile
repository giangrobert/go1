#install go get github.com/cespare/reflex
build:
	go build -o server main.go

run: build
	./server

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.' make run

#root@79924527312a:/go# make watch