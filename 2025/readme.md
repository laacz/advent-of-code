# Advent of Code 2025

https://adventofcode.com/2025/

Rust. Let's cry alone.

## Day 1

Could not figure out math for part 2 with left rotations. Ended up brute forcing. Feeling ashamed already. More anxiety to come!

## Day 2

Already a challenge. Part 1 looks clean, but because I cleaned up some 50 lines of spaghetti or so. For part 2 did brute force, as input did not give extreme ranges.

## Day 3

This was so neat. Part 1 was simple. In part 2 we had to think about algorithmic approach of finding max in a given window, noting its position, doing that 12 times.

## Day 4

Nice and simple.

![Final test case state](https://raw.githubusercontent.com/laacz/advent-of-code/main/2025/day04.png?raw=true)

## Day 5

First part - meh. Second part came easy after I figured out that intervals need to be sorted first.

## Day 6

There's always one where parsing is harder than computing. For second part parser had to be rewritten, part1 solution got adjusted. Yes, it's suboptimal. I might return and implement better logic (parse once, use more idiomatic rust, etc).

## Day 7

Nice one. For part two just multiply by two each time a split occurs, but count universes, not beams (two beams can occupy one spot and hit one splitter).

## Day 8

Today I learned about Unions.

## Day 9

This was a sexy task. For part 2 I first went for naive approach. It did not work out :) After looking at input cardinality, just compressed it by sorting individual coords, and using their indices. After that was a simple flood fill and then brute.\

![My grid](https://raw.githubusercontent.com/laacz/advent-of-code/main/2025/day09.png?raw=true)

## Day 10

Could not do part2 without hints.

Part 1 was easy and worked with bruteforce.

For part 2 bruteforce was not an option. I was getting nowhere with DFS as well - it was taking too long for the actual input. No optimizations of pruning strategy worked.

In the end gave up and looked at The Megathread. Chose the Gaussian elimination. Then almost died implementing it in *Rust*. Solution is still slow, but there is not enough of pride left to do anything about it.


## Day 11

Part 1 - easy. Simple DFS with backtracking. I misread instructions at first (find paths to `bbb`) and spent some time searching for imaginary bugs in my implementation.

Part 2 had too many paths, eh? Memoization to the rescue.
