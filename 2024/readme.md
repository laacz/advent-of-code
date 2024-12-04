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
