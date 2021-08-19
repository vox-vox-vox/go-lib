start:
	arun	go run main.go

benchcpu:
	go-wrk  -d=50 -c=50  http://localhost:4500/cpu/5

profcpu:
	go tool pprof http://127.0.0.1:4500/debug/pprof/profile

docker:
	docker build -t ginapp .
	docker run -p 4502:4501 --rm -it  ginapp ./main -p 4501
