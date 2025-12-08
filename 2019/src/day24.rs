use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/24.txt").expect("Failed to read file");
    let input = parse(&content);
    println!("part1: {}", part1(input));
    println!("part2: {}", part2(input));
}

fn parse(content: &str) -> u32 {
    content
        .lines()
        .flat_map(|line| line.chars())
        .enumerate()
        .fold(0, |acc, (i, c)| if c == '#' { acc | (1 << i) } else { acc })
}

fn step(state: u32) -> u32 {
    let mut new_state = 0;
    for i in 0..25 {
        let mut neighbors = 0;
        if i % 5 > 0 && (state & (1 << (i - 1))) != 0 {
            neighbors += 1;
        }
        if i % 5 < 4 && (state & (1 << (i + 1))) != 0 {
            neighbors += 1;
        }
        if i / 5 > 0 && (state & (1 << (i - 5))) != 0 {
            neighbors += 1;
        }
        if i / 5 < 4 && (state & (1 << (i + 5))) != 0 {
            neighbors += 1;
        }

        if (state & (1 << i)) != 0 {
            if neighbors == 1 {
                new_state |= 1 << i;
            }
        } else {
            if neighbors == 1 || neighbors == 2 {
                new_state |= 1 << i;
            }
        }
    }
    new_state
}

fn part1(input: u32) -> u32 {
    let mut seen = HashMap::new();
    let mut state = input;

    loop {
        if seen.contains_key(&state) {
            return state;
        }
        seen.insert(state, true);
        state = step(state);
    }
}

fn count_neighbors_recursive(levels: &HashMap<i32, u32>, level: i32, pos: usize) -> u32 {
    let state = *levels.get(&level).unwrap_or(&0);
    let outer = *levels.get(&(level - 1)).unwrap_or(&0);
    let inner = *levels.get(&(level + 1)).unwrap_or(&0);

    let mut neighbors = 0;

    if pos % 5 == 0 {
        if (outer & (1 << 11)) != 0 {
            neighbors += 1;
        }
    } else if pos == 13 {
        for p in [4, 9, 14, 19, 24] {
            if (inner & (1 << p)) != 0 {
                neighbors += 1;
            }
        }
    } else {
        if (state & (1 << (pos - 1))) != 0 {
            neighbors += 1;
        }
    }

    if pos % 5 == 4 {
        if (outer & (1 << 13)) != 0 {
            neighbors += 1;
        }
    } else if pos == 11 {
        for p in [0, 5, 10, 15, 20] {
            if (inner & (1 << p)) != 0 {
                neighbors += 1;
            }
        }
    } else {
        if (state & (1 << (pos + 1))) != 0 {
            neighbors += 1;
        }
    }

    if pos / 5 == 0 {
        if (outer & (1 << 7)) != 0 {
            neighbors += 1;
        }
    } else if pos == 17 {
        for p in [20, 21, 22, 23, 24] {
            if (inner & (1 << p)) != 0 {
                neighbors += 1;
            }
        }
    } else {
        if (state & (1 << (pos - 5))) != 0 {
            neighbors += 1;
        }
    }

    if pos / 5 == 4 {
        if (outer & (1 << 17)) != 0 {
            neighbors += 1;
        }
    } else if pos == 7 {
        for p in [0, 1, 2, 3, 4] {
            if (inner & (1 << p)) != 0 {
                neighbors += 1;
            }
        }
    } else {
        if (state & (1 << (pos + 5))) != 0 {
            neighbors += 1;
        }
    }

    neighbors
}

fn step_recursive(levels: &HashMap<i32, u32>) -> HashMap<i32, u32> {
    let min_level = *levels.keys().min().unwrap_or(&0) - 1;
    let max_level = *levels.keys().max().unwrap_or(&0) + 1;

    let mut new_levels = HashMap::new();

    for level in min_level..=max_level {
        let state = *levels.get(&level).unwrap_or(&0);
        let mut new_state = 0u32;

        for pos in 0..25 {
            if pos == 12 {
                continue;
            }
            let neighbors = count_neighbors_recursive(levels, level, pos);
            let is_bug = (state & (1 << pos)) != 0;

            if is_bug {
                if neighbors == 1 {
                    new_state |= 1 << pos;
                }
            } else {
                if neighbors == 1 || neighbors == 2 {
                    new_state |= 1 << pos;
                }
            }
        }

        if new_state != 0 {
            new_levels.insert(level, new_state);
        }
    }

    new_levels
}

fn part2(input: u32) -> u32 {
    let mut levels = HashMap::new();
    levels.insert(0, input);

    for _ in 0..200 {
        levels = step_recursive(&levels);
    }

    levels.values().map(|&state| state.count_ones()).sum()
}
