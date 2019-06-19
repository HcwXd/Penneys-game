# Penneys-game

## What is Penney's Game
> [Wiki](https://en.wikipedia.org/wiki/Penney%27s_game)

Penney's Game requires two players to play.

First, Player A selects a sequence of heads and tails (of length 3 or larger), and shows this sequence to player B. 

Player B then selects another sequence of heads and tails of the same length. 

Subsequently, a fair coin is tossed until either player A's or player B's sequence appears as a consecutive subsequence of the coin toss outcomes. The player whose sequence appears first wins.

## Web GUI Test Tool
[Link](https://hcwxd.github.io/Penneys-game/index.html)

## Generator

```
├── main.go
├── basic
│   ├── main.go
│   └── result
├── question1
│   ├── main
│   ├── main.go
│   ├── result
│   └── run.sh
└── question2
    ├── main
    ├── main.go
    └── result
```
### General generator

`main.go`
```bash
go build main.go
./main [sequence_len repeat_time head_probability]

# Output: For all permutation of A, we calculate B's best strategy and winning rate.
```
### Basic Question

Assuming A choose first and B second. Can you propose a strategy for A and for B to optimize their chance of winning? 

What is the winning rate of B when n varies, given both A and B would choose their best strategy?


`/basic/main.go`
```bash
go build main.go
./main [sequence_len]

# Output: For all permutation of A, we calculate B's best strategy and winning rate. We can then find out best strategy for A
```
### Advanced Question 1

What is the best strategy Assuming the coin is unfair, what shall be the strategy for A and B under this circumstance?

`/question1/main.go`
```bash
go build main.go

./main [sequence_len repeat_time head_probability]
# Or
./run.sh

# Output: As basic question, while we adjust sequence_len and head_probability
```

### Advanced Question 2

Best strategy if A can change one single choice in his sequence after B making his decision? What’s the winning rate of B now? 



`/question2/main.go`
```bash
go build main.go
./main [sequence_len]

# Output: For all permutation of A, we first calculate B's best strategy and winning rate. Then we try to change A's one single choice and find out which change is the optimal choice.
```