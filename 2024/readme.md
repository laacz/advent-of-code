# Advent of Code 2024

These are my solutions for the [Advent of Code 2024](https://adventofcode.com/2024) challenges.

I feel this frustration with fights I'll be having. Waiting for day 18 or so. Wanted to go with Zig, but settled with Go, as I do not want to add to the AoC anxiety by learning a new language.

I do not know why each year I keep doing this. I'm not a competitive person. I do not like to be under pressure. But that last test which succeeds along with the correct answer to your input sometime in early 2025 gives a rush.

This year rules are three:

- No code assistant, no LLMs.
- No peeking unless I'm on a single challenge for third attempt, which typically is third day plus resource availability.
- Not sharing the actual input. Its against rules which I missed last year.

```bash
# Running a single day
# put your input into input.txt
cd day01 && go test && go run .

# Running all
for folder in $(find ./ -name 'day*' -type d)
do
    cd $folder
    go test && go run .
    cd ..
done
```

## Day 1

Nothing of note.

## Day 2

Part 2 took a bit longer than I expected. Same off-by-one errors.

## Day 3

Parsers and state machines on day3. That's anew. Went non-regex, as I wanted it to be clean and avoid full rewrite for part 2. It appears that regexes would have worked.

## Day 4

Didn't bother searching for the best algorithm. Just brute forced it. Easy peasy. Might look at possible optimizations later.

_Update_. Looked at possible optimizations later that day. Just abstracted stuff away, wrote some tests. Now I'm a bit more content with myself. Bruteforce still...

## Day 5

Easy. Just a bit of sorting.

## Day 6

Part one was a simple robot implementation. However, part to took some thinking regarding an optimal obstacle placement and loop detection. For placement - on a path. For loop detection - same wall hit in same direction was optimal enough.

Bruteforce works (7s on my i7, 6s on m2). One probable optimization might be working only with turns.

_Update_. Just routinely threw it all into goroutines. Twice as fast.

## Day 7

Got lucky in part 1, as only two operations wew suspiciously similar to base 2 counting system which otherwise is known as binary. And, as expected, second part was a ternary one. Which got me thinking and debugging for good 30 minutes. 2 seconds are good enough for me.

## Day 8

It was an easy basic geometry/vectors, but the accompanying text for part2 almost mislead me. Nothing especially hard about today.

![Testcase visualization](https://raw.githubusercontent.com/laacz/advent-of-code/main/2024/day08/tests.png?raw=true)

## Day 9

Without a coffee cup or two, I could not understand the task at first. When I did, defragmentation it was.

On the first attempt, I was trying to be smart with intervals, etc. Overengineered a bit to a point where I had no clue what I was doing. Then tried to work with strings. It was fine, but, of course, nobody realized that there are larger indices than 255 :) Rewriting to `[]int` was a breeze.

Second part took a lot more time than needed, as I am well too keen to introduce off-by-one errors.

Later returned for some microoptimizations and removal of dead code parts.

## Day 10

This one was easy. Reused BFS code from tjhe previous year. First one is a list of distinct targets, second - a list of distinct paths. Then tried to polish it and forked everything up beyond the length of the undo buffer.

## Day 11

First was a brute force approach. Then had to rewrite everything. Mentioned strict order can be ignored. And we should deal with counts, not individual stones. Part2 implements hashing and that's all.

## Day 12

This year hard problems start on day 12. Floodfill was an obvious one. Calculating perimeter was easy as well - just increment when you see an outside.

Counting edges was as simple as counting corners. Just take a look at 2x2 squares. If an odd number, it's a corner. There was an edge case I spent a lot of time on. It's when the 2x2 catches two corners of the same letter.

In the following example there are four plots. This threw me off for a while.

```
AB
BA
```

A bit later optimized so it does not look at the whole map for each corner detection.

## Day 13

What is it with programming tasks and algebra? Year over year I really forget more and more of the math stuff, as I do not need it in my daily life. So this was challenging.

First part, of course, was bruteforce. Second got me thinking. It looked like a trivial case of solving two linear equation system for two unknowns. A bit of googling around and refreshing my knowledge regarding solving these by using determinants, and voila - first part passes all the tests.

Second part did not come together, tho. Cause was that it had solutions where solutions are not positive integers. All good after accounting for that.

## Day 14

Easy. For the second part I made some assumptions and was patient with bruteforce. And voila. Assumption was that the tree would have some straight horizontal lines, so it was just a matter of pattern matching. If that pattern would not consist of all ones, I'd be screwed.

Afterwards made a huge optimization of string representation of the robots' state. So it's quick enough now.

## Day 15

Lanterfish sokoban! :) This was a true pleasure. However, it took a while.

I started off with a trivial solver for part 1. Worked like a charm.

For second part I took a thinking pause. Came up with an algorithm. Rushed to implement it. It was not the right one. Going up/down and increasing range to check for obstacles. Everything worked, except for the test case's 311th step. Had to come up with another approach - identify all boxes and then just move them. If at least one box is not movable, we're blocked.

![Testcase visualization](https://raw.githubusercontent.com/laacz/advent-of-code/main/2024/day15/sokoban.png?raw=true)

## Day 16

Dijkstra. Composite graph, where turns should be part of the state. I. Do. Not. Like. Graphs. I got all the paths in the part1, so part2 was already solved. Oddly late in the game for the first time this year. Got a bit fancy with directions tho, but got stuck on path reconstruction, as I often do.

## Day 17

Dang output. I lost half an hour generating test cases. All tests pass, submission is rejected. Then submitted a string, where output was joined with commas, as stated in the task...

For the second part I had to abort bruteforcing and look for hints. `<<3` was enough. These types of tasks are not my forte. Or even my normale.

## Day 18

Same old Dijkstra. For the second part just same pathfinding until there is no path. Didn't write a single test line tho. No need.

## Day 19

First part was simple recursion until we got to one where memoization would be needed (search was taking too long). But overall - easy peasy.
