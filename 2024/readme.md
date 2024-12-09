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