# Payday

## About

Small script to get familiar with [Go](https://go.dev/) and writing Go tests.
It is used to calculate the next payday for a bi-monthly pay period.
This script is based on a [Python version](https://gist.github.com/andreburto/66fb46e2a7ae63cb777eb0023deae5bb) written a while pack.

## Usage

### Build

```
go build
```

### Test

```
go test -v
```

### Run

```
./payday
```

## To Do

* Create a test for the `main` function to test output and hour calculations.
* Turn script into microservice. (LOL jk)

## Update Log

**2024-03-24:** Initial commit of the working script and tests of most methods.
