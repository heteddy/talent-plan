.PHONY: all

all: test bench

test:
	go test
clean:
	rm *.out mergesort.test
bench:
	#go test -bench Benchmark -run xx -count 5 -benchmem
	go test -bench Benchmark -run xx -count 5 -benchmem -cpuprofile cpu.out -memprofile mem.out


