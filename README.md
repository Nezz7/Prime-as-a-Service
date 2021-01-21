Prime as a Service
=========
![Build](https://github.com/Nezz7/Prime-as-a-Service/workflows/Build/badge.svg)


Prime as a Service is a REST API implemented in [Go](http://golang.org) that provides some of the basic functionalities using prime numbers.
This service was deployed using a Docker Swarm cluster and monitored with Grafana and Prometheus.


# Algorithm
Implemented algorithms :
* [Sieve of Eratosthenes](https://cp-algorithms.com/algebra/sieve-of-eratosthenes.html) 
* [Prime factorization](https://cp-algorithms.com/algebra/factorization.html) 
* [Number of divisors](https://cp-algorithms.com/algebra/divisors.html) 

## Endpoints

| Path                                  | Type  | Description                                                      | 
|:--------------------------------------| :---: |:-----------------------------------------------------------------|
| /primes/lower&upper               |  GET  | Return a list of prime numbers between lower and upper inclusive.|
| /prime-factors/n                  |  GET  | Return a list of prime facotrs of n and their power.             | 
| /number-of-divisors/n             |  GET  | Return the number of divisors of n.                              | 
| /metrics                          |  GET  | Return the metrics provided by Prometheus.                        |      

## Deploy

1. Clone this repository <br>
`$ git clone https://github.com/Nezz7/prime-as-service.git`
2. Change the current working directory <br>
`$ cd /prime-as-a-service`<br>
3. Create the docker image : server-image  <br>
`$ docker build -t server-image .`<br>
4. Initialize the swarm <br>
`$ docker swarm init`<br>
Don't forget to save the token.<br>
5. Deploy the stack to the swarm<br>
`$ docker stack deploy --compose-file docker-compose.yml  $SERVICE_NAME`<br>
6. Check that it’s running <br>
`$ docker stack services $SERVICE_NAME`<br>
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





