import time


def main(max_num):
    primes = []
    for x in range (2, max_num):
        prime = True
        for y in range(2, int(x**0.5) + 1):
            if x % y == 0:
                prime = False
                break
        if prime:
            primes.append(x)
    return primes

if __name__ == "__main__":

    startTime = time.perf_counter()
    max_num = 10000000
    primes = main(max_num)
    endTime = time.perf_counter()

    elapsed = endTime - startTime
    print(f"Found {len(primes)} primes to {str(max_num)} in {elapsed} seconds")
