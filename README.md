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

## Settings file

Because payday can come at different times during the month, this script needs to be versatile.
The settings file, `payday.yml`, sees to that.
Because this program has a small userbase (me), the settings file is simple.
There are only the following settings:

`first_payday`: the first day of the month when pay checks are sent out.

`next_payday`: the next time a check goes out. This currently has two values:

* `TWO_WEEKS`: the nearest weekday two weeks from the first payday.
* `LAST_DAY`: the last weekday of the current month.

## To Do

* Create a test for the `main` function to test output and hour calculations.

## Update Log

**2024-03-31:** Modified code to use settings file.
Added tests for TWO_WEEK, but need to add tests for LAST_DAY.

**2024-03-30:** Updated NextPayday logic.

**2024-03-24:** Initial commit of the working script and tests of most methods.
