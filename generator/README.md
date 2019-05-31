# Simple Generator

> This generator
> 1. Generate all possible sequences of player A
> 2. For each sequence, it find the best sequence of player B that counters player A
> 3. Output all sequences found in step 2 and sort them. Thus the one with highest winning rate is player A's optimal strategy

## Usage

1. Compile
```
go build main.go
```

2. Execution (Default sequence len = 5)

```
./main [sequence len]
```