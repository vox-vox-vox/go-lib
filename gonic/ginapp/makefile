start:
	go run gin-pprof.go

benchcpu:
	go-wrk  -d=50 -c=50  http://localhost:4500/cpu/5

profcpu:
	go tool pprof http://127.0.0.1:4500/debug/pprof/profile
