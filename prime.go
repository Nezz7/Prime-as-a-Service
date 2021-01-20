package main

import (
	"fmt"
)
//MAXN represent the maximum number supported 
const MAXN = 1000000
var p [] int 
var primes [] int

//SieveOfEratosthenes an implementation of Sieve algorithm in O(Nlog(N))
// returns a slice of size n where p[i] equals to the greatest prime divisor of i
// and returns the primes numbers <= n
func SieveOfEratosthenes(n int) ([] int, [] int) {
	var p = make([] int, n + 1)
	var primes = make([] int, 0)
	p[0] = 1;
	for i := 1; i <= n; i++{
		p[i] = i;
	}
	for i := 2; i <= n; i++{
		if(p[i] == i){
			primes = append(primes, i)
			for j := 2 * i; j <= n; j+= i{
				p[j] = i;
			}
		}
	}

	return p, primes
}
func getP() [] int {
	if p == nil {
		p, primes = SieveOfEratosthenes(MAXN)
	}
	return p
}
func getPrimes() [] int {
	if primes == nil {
		p, primes = SieveOfEratosthenes(MAXN)
	}
	return primes
}


type pair struct{
	num int
	pow int
}

func getPrimeFact(n int) [] pair{
	var result = make([] pair, 0)
	p = getP();
	for n > 1  {
		var d = p[n]
		cnt := 0
		for n % d == 0 {
			cnt ++
			n /= d
		}
		result = append(result, pair{d, cnt})
	}
	return result
}

func numberOfDivisors(n int) int{
	numDiv := 1;
	pf := getPrimeFact(n)
	for i := 0; i < len(pf); i++{
		numDiv *= (pf[i].pow + 1)
	}
	return numDiv
}
func getPrimesInRange(low int, high int)[] int{
	primes = getPrimes();
	left  := 0
	right := len(primes) - 1
	start := 0
	for left <= right {
		mid := (left + right) / 2;
		if primes[mid] >= low {
			right = mid - 1
			start = mid - 1
		}else {
			left = mid + 1
		}
	}
	var result = make([] int, 0)
	for i := start; i < len(primes) && primes[i] <= high; i++{
		result = append(result, primes[i]);
	}
	return result;

}

func main() {
		fmt.Println("hello world")
		_, primes := SieveOfEratosthenes(10)
		for i := 0; i < len(primes); i++ {
			fmt.Print(primes[i], " ")
		}
		pf := getPrimeFact(60)
		for i := 0; i < len(pf); i++{
			fmt.Println(pf[i].num, " ", pf[i].pow)
		}
		r := getPrimesInRange(5, 100)
		for i := 0; i < len(r); i++ {
			fmt.Print(r[i], " ")
		}

}
