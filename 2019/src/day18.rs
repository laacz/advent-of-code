use std::collections::{HashMap, HashSet, VecDeque};

fn main() {
    let content = std::fs::read_to_string("data/18.txt").expect("Failed to read file");
    let grid = parse(&content);
    println!("part1: {}", part1(&grid));
    println!("part2: {}", part2(&grid));
}

fn parse(content: &str) -> Vec<Vec<char>> {
    content.lines().map(|l| l.chars().collect()).collect()
}

fn add_key(keys: &str, key: char) -> String {
    if keys.contains(key) {
        return keys.to_string();
    }
    let mut chars: Vec<char> = keys.chars().collect();
    chars.push(key);
    chars.sort();
    chars.into_iter().collect()
}

fn part1(grid: &Vec<Vec<char>>) -> i32 {
    let dirs: [(i32, i32); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];

    let mut start = (0, 0);
    let mut total_keys = 0;
    for (y, row) in grid.iter().enumerate() {
        for (x, &c) in row.iter().enumerate() {
            if c == '@' {
                start = (x as i32, y as i32);
            }
            if c.is_ascii_lowercase() {
                total_keys += 1;
            }
        }
    }

    let mut q: VecDeque<((i32, i32), String, i32)> = VecDeque::new();
    let mut visited: HashSet<((i32, i32), String)> = HashSet::new();

    q.push_back((start, String::new(), 0));
    visited.insert((start, String::new()));

    while let Some(((x, y), keys, steps)) = q.pop_front() {
        if keys.len() == total_keys {
            return steps;
        }

        for (dx, dy) in &dirs {
            let nx = x + dx;
            let ny = y + dy;

            if ny < 0 || ny >= grid.len() as i32 || nx < 0 || nx >= grid[0].len() as i32 {
                continue;
            }

            let c = grid[ny as usize][nx as usize];

            if c == '#' {
                continue;
            }

            if c.is_ascii_uppercase() && !keys.contains(c.to_ascii_lowercase()) {
                continue;
            }

            let new_keys = if c.is_ascii_lowercase() {
                add_key(&keys, c)
            } else {
                keys.clone()
            };

            let state = ((nx, ny), new_keys.clone());
            if !visited.contains(&state) {
                visited.insert(state);
                q.push_back(((nx, ny), new_keys, steps + 1));
            }
        }
    }

    -1
}

fn bfs_from(grid: &Vec<Vec<char>>, start: (i32, i32)) -> HashMap<char, (i32, String)> {
    let dirs: [(i32, i32); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];
    let mut result: HashMap<char, (i32, String)> = HashMap::new();
    let mut visited: HashSet<(i32, i32)> = HashSet::new();
    let mut q: VecDeque<((i32, i32), i32, String)> = VecDeque::new();

    q.push_back((start, 0, String::new()));
    visited.insert(start);

    while let Some(((x, y), dist, doors)) = q.pop_front() {
        for (dx, dy) in &dirs {
            let nx = x + dx;
            let ny = y + dy;

            if ny < 0 || ny >= grid.len() as i32 || nx < 0 || nx >= grid[0].len() as i32 {
                continue;
            }
            if visited.contains(&(nx, ny)) {
                continue;
            }

            let c = grid[ny as usize][nx as usize];
            if c == '#' {
                continue;
            }

            visited.insert((nx, ny));

            let mut new_doors = doors.clone();
            if c.is_ascii_uppercase() {
                new_doors = add_key(&new_doors, c.to_ascii_lowercase());
            }

            if c.is_ascii_lowercase() {
                result.insert(c, (dist + 1, new_doors.clone()));
            }

            q.push_back(((nx, ny), dist + 1, new_doors));
        }
    }

    result
}

fn part2(grid: &Vec<Vec<char>>) -> i32 {
    let mut center = (0, 0);
    let mut total_keys = 0;
    for (y, row) in grid.iter().enumerate() {
        for (x, &c) in row.iter().enumerate() {
            if c == '@' {
                center = (x as i32, y as i32);
            }
            if c.is_ascii_lowercase() {
                total_keys += 1;
            }
        }
    }

    let mut grid = grid.clone();
    let (cx, cy) = (center.0 as usize, center.1 as usize);
    grid[cy][cx] = '#';
    grid[cy - 1][cx] = '#';
    grid[cy + 1][cx] = '#';
    grid[cy][cx - 1] = '#';
    grid[cy][cx + 1] = '#';
    grid[cy - 1][cx - 1] = '@';
    grid[cy - 1][cx + 1] = '@';
    grid[cy + 1][cx - 1] = '@';
    grid[cy + 1][cx + 1] = '@';

    let starts = [
        (cx as i32 - 1, cy as i32 - 1),
        (cx as i32 + 1, cy as i32 - 1),
        (cx as i32 - 1, cy as i32 + 1),
        (cx as i32 + 1, cy as i32 + 1),
    ];

    let mut distances: HashMap<char, HashMap<char, (i32, String)>> = HashMap::new();

    for (i, &start) in starts.iter().enumerate() {
        distances.insert((b'0' + i as u8) as char, bfs_from(&grid, start));
    }

    for (y, row) in grid.iter().enumerate() {
        for (x, &c) in row.iter().enumerate() {
            if c.is_ascii_lowercase() {
                distances.insert(c, bfs_from(&grid, (x as i32, y as i32)));
            }
        }
    }

    let start_positions: [char; 4] = ['0', '1', '2', '3'];

    let mut q: VecDeque<([char; 4], String, i32)> = VecDeque::new();
    let mut visited: HashMap<([char; 4], String), i32> = HashMap::new();

    q.push_back((start_positions, String::new(), 0));
    visited.insert((start_positions, String::new()), 0);

    let mut best = i32::MAX;

    while let Some((positions, keys, steps)) = q.pop_front() {
        if keys.len() == total_keys {
            best = best.min(steps);
            continue;
        }

        if steps >= best {
            continue;
        }

        for robot in 0..4 {
            let pos = positions[robot];
            if let Some(reachable) = distances.get(&pos) {
                for (&target_key, &(dist, ref doors_needed)) in reachable {
                    if keys.contains(target_key) {
                        continue;
                    }

                    if !doors_needed.chars().all(|d| keys.contains(d)) {
                        continue;
                    }

                    let new_keys = add_key(&keys, target_key);
                    let mut new_positions = positions;
                    new_positions[robot] = target_key;

                    let new_steps = steps + dist;
                    let state = (new_positions, new_keys.clone());

                    if let Some(&prev) = visited.get(&state) {
                        if new_steps >= prev {
                            continue;
                        }
                    }

                    visited.insert(state, new_steps);
                    q.push_back((new_positions, new_keys, new_steps));
                }
            }
        }
    }

    best
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example1() {
        let input = "#########
#b.A.@.a#
#########";
        assert_eq!(part1(&parse(input)), 8);
    }

    #[test]
    fn test_example2() {
        let input = "########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################";
        assert_eq!(part1(&parse(input)), 86);
    }

    #[test]
    fn test_example3() {
        let input = "########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################";
        assert_eq!(part1(&parse(input)), 132);
    }

    #[test]
    fn test_example4() {
        let input = "#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################";
        assert_eq!(part1(&parse(input)), 136);
    }

    #[test]
    fn test_example5() {
        let input = "########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################";
        assert_eq!(part1(&parse(input)), 81);
    }

}
