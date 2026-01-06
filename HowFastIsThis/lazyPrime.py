import time


def main():
    max = 100000
    primes = []
    for x in range ( max):
        prime = True
        for y in range(2, x-1):
            if x % y == 0:
                prime = False
                break
        if prime:
            primes.append(x)
    return primes

if __name__ == "__main__":
    startTime = time.perf_counter()
    primes = main()
    endTime = time.perf_counter()

    elapsed = endTime - startTime
    print(f"Found {len(primes)} primes in {elapsed} seconds")
