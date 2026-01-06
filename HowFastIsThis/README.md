#Wow, compiled to machine code is infact faster that compiled on runtime.

Started with lazyPrime.go and lazyPrime.py, with the understanding of the more fundemental diffs going into this.  However, it is fun to see how the tremendous speed delta between the two implementations fo the algo.

A more fair comparison would be to use NumPy in python, which would require a venv and an install -- significantly more overhead.

Same with using a threadpool/futures to comp the routined version.

Anyways:


lazyPrime.go (Not concurrent)
identified 664579 primes to 10000000 in 16.475891024s

routined_lazyPrime/lazyPrime.go (Concurrent at 64 routines)
identified 664579 primes to 10000000 in 813.660285ms secondsi

lazyPrime.py



Also would be slick to loop num routines to identifiy which would make the most sense for a system, providing a best time.  This could likely be inferred/calc'd from core count of machine, could image there being some variance due to schedulers to spec of proc.
