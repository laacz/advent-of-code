# Advent of Code 2022

These are solutions for the [Advent of Code 2022](https://adventofcode.com/2022) challenges.

This year it's Go time. Or to put it punly, let's give it a Go. Or just let's Go.

```bash
# Runs the tests first. If they pass, runs solutions for the given day.
run.sh day01
```

## Day 1

Easy peasy.

## Day 2

Lemon squeezy.

## Day 3

Nothing to see here.

## Day 4

Same here.

## Day 5

Straightforward, until I got to the second part. In certain cases Go modifies target slice, instead of copying. So I had
to resort to moving crates one-by one.

## Day 6

Shortest solution to date. Guessed the second part's task, so both parts use the same implementation.

## Day 7

Easy, but had to think about the right approach longer than I'd like.

## Day 8

Got stuck on failing tests, because missed that first lines do not count.

## Day 9

This took some while, because I messed up X and Y coordinates in one place. This was a relatively easy one nonetheless.

## Day 10

A few off by one one errors, but this was sweet.

## Day 11

Integers have a nasty habit of overflowing. After a half an hour or so sorted that out.

## Day 12

Graphs. Not a fan. Took some while.

## Day 13

Oh, this one would be so much easier using a loosely typed language. Fuuuuuuu... I don't have words. Took two days,
three
rewrites. On the bright side, I learned a lot about `interface{}`, `reflect` package and myself.

## Day 14

What a relief and fun after previous day. Well, day 13 was hard only because of Go typing and me not being used to it.

## Day 15

First part was easy enough.

Second part was a bit different. First - we generate a list of all lines at manhattan distance + 1. Then we check if any
of these lines intersect (well, we don't but still resulting list of points is not huge). If they do, we add the point
to the list of possible candidates for a place to have the beacon at. Then we just iterate over all these points and
find the first one which is not covered by any sensor.

## Day 16

On the fourth day I gave up. Will have to revisit this one and theory behind it. Graphs are my nemesis.

## Day 17

Part 1 was very sweet and I enjoyed that a lot. Part 2, however, was not fun at all. Although it was fairly obvious that
what we're looking for is a cycle, actual approach and with that its solution took me several attempts.

## Day 18

After previous few days this was pure joy. For the first part I created map of all cubes, and then counted sides,
subtracting adjoining ones. For the second part implemented simple BFS-based flood-fill.

## Day 19

This took some while to figure out. BFS might not be the fastest solution. Initially part 1 testcase took around a
minute. After pruning branches which get first geode later than others, got it down to .12s. For the second part this
didn't work well in terms of performance, so I added one more pruning step - check if maximum amount of bots per
blueprint are not exceeded.

## Day 20

Those days, when tests pass, but actual input doens't generate an acceptable result. The culprit was a newline at the
end of input data. Second part was a breeze after that. Took my time to refactor into smaller chunks to encapsulate
stuff and avoid func name collisions with builtins.

## Day 21

First part was easily by evaluating all yelled at monkeys in a loop, until there is no more unknown numbers. Second part
was tricky, since approach had to be completely different. When parsing, we mark a path from `humn` to `root`. Then we
solve for that, using existing solution from the part 1. I was afraid about the stack (it's recursive), but it turned
out that solution path is not that long. Longest was to figure out the right approach for the revese solver.

## Day 22

This is Spartaaaaaa! Part one was easy. Par two might have been the finneckiest AoC day so far. This implementation
is the fourth one after trying to adapt first part's solution, then solving generic folding and wrap-around solution,
then the graph approach. As a result I just cut two cubes (test and actual input), and hard-coded all the wrapping
rules. The hardest part, actually, was the debugging process. There were typos, there were just plain wrong definitions,
and there was the case of Down->Right wrapping, which wrapped differently for the test and for the actual input.

So here we go.

## Day 23

Straightforward until the point when tests pass, but actual input doesn't. This happened with part 2. The reason for
that was stupid logic error. I was annoyed by myself to the point where I did not do any performance refactoring, as
opposed to previous days.

## Day 24

Nice and simple flood-fill with a catch. Catch was in that a blizzard must clear the parth, so you can move there
again. Second part required no major modifications - just the ability to define start and end points.

## Day 25

And we're done. Except for day16, where I did not find a solution. This all actually gives me some relief. Looks like
this un somehow was much more stressful than I anticipated it would be. Oh, anf for the dec to SNAFU - went for a 
math'y solution without any fanciness all around.

So long, and thanks for all the fish. Well, not really. See you next year.