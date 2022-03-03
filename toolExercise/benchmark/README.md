go test -bench=.
> 正则表达  

go test -bench=. -count 5

go test -bench=. -count 5 -run=^#

go test -bench=. -benchtime=10s

go test -bench=. -benchtime=100x

go test -bench=. -benchtime=10s -benchmem

## article
* [Benchmarking in Golang: Improving function performance - LogRocket Blog](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)

