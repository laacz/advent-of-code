# 2019

Refreshing my rusty Rust skills, since I've decided to fight The Borrow Checker once again. [Did not go well 4 years ago](https://github.com/laacz/aoc-2021-rust).

Will be doing this until it gets too hard.

## Day 01 .. 05

Easy enough.

## Day 06

First part was easy enough. Got lost with actual Rust implementation tho.

Struggled a bit more with the second part. Mainly due the unfamiliarity with Rust. Somehow it takes a longer while to translate my thinking into Rust than into Go, for example.

## Day 07

This. Just. Sucked.

## Day 08

A compensation for day 07 probably.

## Day 09

That damned Intcode computer again. To make my life even harder, decided to implement it as a type. I dislike these large programs as AoC solutions. Part 2 was a nice surprise for the late night.

## Day 10

Task was simple, implementation not. Still getting used to the ways of Rust. Regarding solution - the formulas are not living in my head, so looked them up. Without web search I wouldn't be able to code anything beyond simple straightforward stuff.

## Day 11

Once more that damned Intcode computer. Nevertheless - look how pretty! Yeah, procrastinated a bit :)

![HSV rainbow!](https://raw.githubusercontent.com/laacz/advent-of-code/main/2019/day11.png?raw=true)

## Day 12

So much to write. Small errors, off-by-ones, off-by-13322932's, moon-affects-itself features. This time around I learned and remembered few more things about Rust. Like `iter_mut()`, init and fill a vec,

## Day 13

Easy with Intcode. Except for part 2. Which was not easy. Used optionals outside return values for the first time.

## Day 14

Here be graphs (binary trees for now).

## Day 15

All of my nemesis. Nemesises. Nemesi. Wait, googling. ... Back with nemeses! Their names are BFS, DFS, Dijkstra, and more. Here they come. First part requires us to do BFS with the Intcode computer and its program. BFS guarantees that first found path is the shortest. For part 2 we to BFS on the actual generated map.

![The map!](https://raw.githubusercontent.com/laacz/advent-of-code/main/2019/day15.png?raw=true)

## Day 16

For the second part I went for help. Help was not very helpful tbh. Re-read the insight, asked an LLM to ELIH (explain like i'm hungover). Then off to the races.

## Day 17

This was brutal. No BFS on Dijkstra (yet to come, eh?), but still.

## Day 18

Simple BFS at first. State - position and keys collected. Well, not simple, as you can't go though doors you don't have keys for.

Second part took a while to figure out. Blind BFS was too huge, so converted to a distance graph instead of movement graph. This was hard. I doubt that my spahetti is optimal.
