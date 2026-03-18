# Query 
![Coverage](https://img.shields.io/badge/Coverage-45.4%25-yellow)

[![Go](https://github.com/macinnir/goquery/actions/workflows/go.yml/badge.svg)](https://github.com/macinnir/goquery/actions/workflows/go.yml)

# Performance 

```
# https://graphviz.org/download/#mac
brew install graphviz

go test -bench=. ./core/lib/utils/query/testgen/. -cpuprofile cpu.prof
go tool pprof -svg cpu.prof > cpu.svg

go test -bench=. -trace trace.out ./core/lib/utils/query/testgen/.
go tool trace trace.out

# Data Race: Two goroutines access the same variable concurrently 
# and at least one of the accesses is a write 
# https://golang.org/doc/articles/race_detector
go test -race
```
