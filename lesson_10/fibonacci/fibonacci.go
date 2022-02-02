package fibonacci

func FibRecurse() func(uint, ...bool) uint64 {
	var cache = make(map[uint]uint64)

	var fib func(uint, ...bool) uint64
	fib = func(n uint, withCache ...bool) (result uint64) {
		if len(withCache) > 0 && withCache[0] {
			if result, ok := cache[n]; ok {
				return result
			}
			defer func() {
				cache[n] = result
			}()
		}
		if n <= 1 {
			return uint64(n)
		}
		return fib(n-1, withCache...) + fib(n-2, withCache...)
	}
	return fib
}

func FibArray(n int) uint64 {
	fib1 := 0
	fib2 := 1

	n -= 1

	for n > 0 {
		fib1, fib2 = fib2, fib1+fib2
		n -= 1
	}
	return uint64(fib2)
}
