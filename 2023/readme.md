# Advent of Code 2023

These are solutions for the [Advent of Code 2023](https://adventofcode.com/2023) challenges.

Continuing in Go. However, doing this without the use of any code assistant. Will be weird at first.

Note to self. I actually do not have much time this year to dedicate to this. That's the reason I'm going with a familiar language. Nevertheless - I suspect this will go on long after the actual event is over.

```bash
for folder in $(find ./ -name 'day*' -type d); do cd $folder && go test && go run . && cd ..; done
```

## Day 01

Nothing out of the ordinary for the first day. Just a pesky itsy-bitsy off-by one :)

## Day 02

Easy peasy.

## Day 03

Part 2 caught me a bit off-guard. After short deliberations with a person as non-programmer as me (that being myself) decided to go ad-hoc for the secodn part.

## Day 04

First part was trivial, second - easy enough once you read all the requirements. I feel that there was some fancy algorithm intended to be used here, but I'll leave it to those who feel them.

## Day 05

Bruteforced my way out of it. Obviously could be made better by working with building e2e map, then ranges and their overlaps/intersections, rather than iterating through all the seed numbers.

## Day 06

Oddly - parsing was the one that took the most time to write. But bruting is fine and quck.

## Day 07

Easy enough. For the first part its just a matter of a sort. Second part was just a matter of iterating through all the non jocker cards in the hand and replacing jockers with each one of them, writing down which hand would be stronger.

## Day 08

First part - easy. Second part - tried to brute force first, didn't work. And it came to me unexpectedly easy - for each of the starting point there is a finite number of steps to be taken to get to the finish. 

When I was close to finishing, it came to me once again. The realisation that it won't work, as after each 'last' step the next cycle is of different length. I finished it anyway, and to my absolute surprise and astonishment - answer was correct.

So in effect, there were no clues in the problem that paths are always cycling at same length.

## Day 09

First time it did not work with the actual input, though tests passed. Checked output for each line - there it is, negative numbers. Replaced check for reduced slice sum being `>0` with `!=0` and it still did not work. After twenty minutes or so there it was - an edge case when all numbers summed up to zero. Fixed that and we're golden. 

Second part was the same, except with heads instead of tails.

## Day 10

With off and on took almost whole day. No actional struggle with the first part where the hardest part was to build BFS'able tree. 

As for the second part - at first I was baffled and could now think of anything. But, once I plotted it on screen, solution became obvious - find out if the given point is inside the polygon. Winding numbers and that's it.

![alt text](https://github.com/laacz/aoc-2023/blob/main/day10/debug.png?raw=true)

## Day 11

Easy. First part got solved via the straightforward approach of inserting doubled rows and columns. Then rewrote to store only cols and rows with their corresponding expansion factor.

![alt text](https://github.com/laacz/aoc-2023/blob/main/day11/debug.png?raw=true)

## Day 12

Struggled with this one almost to a point where I'd just peek at other solutions. Part two required only some memoization.

## Day 13

Both parts were really nice ones. I overly optimized, as was almost sure that the second part would be something huge. Well. It was not.
