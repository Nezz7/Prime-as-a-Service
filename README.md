Prime as a Service
=========
![Build](https://github.com/Nezz7/Prime-as-a-Service/workflows/Build/badge.svg)


Prime as a Service is a REST API implemented in [Go](http://golang.org) that provides some of the basic functionalities using prime numbers.


# Algorithm
Implemented algorithms
* [Sieve of Eratosthenes](https://cp-algorithms.com/algebra/sieve-of-eratosthenes.html) 
* [Prime factorization](https://cp-algorithms.com/algebra/factorization.html) 
* [Number of divisors](https://cp-algorithms.com/algebra/divisors.html) 

## Endpoints

| Path     |  Type        | Description        | 
|:-------------:|:-------------:|:-------------:|
| /api/primes/{lower:[0-9]+},{upper:[0-9]+} | GET | Return a list of prime numbers between lower and upper inclusive.|
| /api/primeFactors/{n:[0-9]+}              | GET | Return a list of prime facotrs of n and their power.  | 
| /api/numberOfDivisors/{n:[0-9]+}          | GET | Return the number of divisors of n.  | 
| /api/metrics/                             | GET | Return the metrics provided by promethus.  | 


# Performance
Performance depends on the size of max number. But as an example, it needs about 0.4 ms to produce the first 1,000,000 prime numbers.

```bash
$ go test -bench .  
    goos: linux
    goarch: amd64
    BenchmarkPrimesInRange-2   	    2359	    491736 ns/op
    PASS
    ok  	github.com/Nezz7/Prime-as-a-Service	    1.240s
```





