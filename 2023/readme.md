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

![Some debugging going on](https://raw.githubusercontent.com/laacz/advent-of-code/main/2023/day10/debug.png?raw=true)

## Day 11

Easy. First part got solved via the straightforward approach of inserting doubled rows and columns. Then rewrote to store only cols and rows with their corresponding expansion factor.

![Some more debugging going on](https://raw.githubusercontent.com/laacz/advent-of-code/main/2023/day11/debug.png?raw=true)

## Day 12

Struggled with this one almost to a point where I'd just peek at other solutions. Part two required only some memoization.

## Day 13

Both parts were really nice ones. I overly optimized, as was almost sure that the second part would be something huge. Well. It was not.

## Day 14

Part one was deceptively easy. Second one was simply a matter of detecting offset, cycles and then calculating the remaining cycles to take.

## Day 15

Part one was chill, par two was an unneccessary twist. What bothered me, that solution works not 100% of the time. Oh, well.

## Day 16

Bruted my way out of this.

## Day 17

Fuuuuuuuu... it took so long. I'm writing this entry on day 22. That's how long the first part took. Apparently Dijkstra is my actual nemesis.

## Day 18

First tried to compute area with a ray method. Didn't want the hassle with the edge cases, when ray occupies the same space as the edge. Thought about flood-fill, but a quick google search showed me Shoelace's algorithm for computing area of a polygon. After that - quick and easy.

## Day 19

Part one was a nightmare to parse, but easy to solve. Par two was some recursive range splitting fu I was not ready for. Fighted copy by reference vs copy by value for a while.

## Day 20

## Day 21

Part one was easy.

Part two was OMG. My solution has nothing to do with the test data.

This would not be possible without looking at the output of the actual input for the part 1. It's actually a diamond shape. Always. 

To validate that this propogates to other four adjoining squares, I tested the first part by modifying input so the starting point is in the middle of any of the four sides, as well as from all four corners.

So it all boiled down to calculating the covered tiles and then the hardest part - partially covered ones.

Each step expands our diamond by one tile in each direction. With the step 66 we overflow into horizontally and vertically adjoining tiles (65 steps to take from the center of the 131x131 tile).

That means that after taking 26501365 steps we've moved 26501365 steps away from the center. It fully covers (steps / 131)^2 tiles. 

Horizontally and vertically we need to take care of at most 2 tiles in each direction. Diagonally it's not that simple. 

However.

I decided to peek at the solutions and found that quadratic formula approach would be much easier. It does not work with test cases for some reason, but the actual input's result was accepted nontheless.

## Day 22

When all test cases pass, but input does not... 

Took another solution, within my input found first sequence of 10 blocks, for which my solution gives wrong output. 

![Debugging via console](https://raw.githubusercontent.com/laacz/advent-of-code/main/2023/day22/text-only-debug.png?raw=true)

My spatial extrapolation is not that good, so I went and wrote my first three.js script to visualize blocks ([debug.js](https://raw.githubusercontent.com/laacz/advent-of-code/main/2023/day22/debug.js)).

![Debugging via three.js rendering](https://raw.githubusercontent.com/laacz/advent-of-code/main/2023/day22/threejs.png?raw=true)
![Debugging via three.js editor](https://raw.githubusercontent.com/laacz/advent-of-code/main/2023/day22/threejs-v1.png?raw=true)

It immediately became clear that Blocks "F" and "B" are overlapping. Fixed algo, and all got good. 

When came the second part, I refactored everything to accomodate new requirements. Wanted to go recursive way, but finally used original `Drop` function to test if any other blocks would fall, when one has been removed.

Overall this took some time and was not easy at all. Good thing - refactored approach was much faster than theoriginal :)



