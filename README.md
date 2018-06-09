# Effective Concurrency with Go
A tutorial on concurrent program design and how to use the context package for
cancellation, timeouts and passing "request scoped data" across goroutines.

## If running in a docker container
1. docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
1. setup .bashrc to have ~/go/bin in PATH or export PATH=~/go/bin:$PATH
1. present -http 0.0.0.0:3999 -orighost 192.168.99.100
