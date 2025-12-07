fn main() {
    let input = std::fs::read_to_string("data/07.txt").expect("Failed to read input file");
    let input = parse(&input);
    println!("part1: {}", part1(&input));
    println!("part2: {}", part2(&input));
}

fn parse(input: &str) -> Vec<Vec<u8>> {
    input
        .trim()
        .lines()
        .map(|line| line.bytes().collect())
        .collect()
}

fn print_map(map: &[Vec<u8>]) {
    for row in map {
        for &cell in row {
            print!("{}", cell as char);
        }
        println!();
    }
}

fn part1(input: &[Vec<u8>]) -> usize {
    let mut ret = 0;
    let mut map = input.to_owned();

    let start = map[0].iter().position(|&c| c == b'S').unwrap();
    map[1][start] = b'|';

    println!("{}", start);
    for row in 2..map.len() {
        for col in 0..map[row].len() {
            if map[row - 1][col] != b'|' {
                continue;
            }

            match map[row][col] {
                b'.' => map[row][col] = b'|',
                b'^' => {
                    ret += 1;
                    if col > 0 && map[row][col - 1] == b'.' {
                        map[row][col - 1] = b'|';
                    }
                    if col < map[row].len() - 1 && map[row][col + 1] == b'.' {
                        map[row][col + 1] = b'|';
                    }
                }
                _ => {}
            }
        }
    }

    print_map(&map);
    ret
}

fn part2(input: &[Vec<u8>]) -> usize {
    let map = input.to_owned();
    let len = map[0].len();

    let mut timelines: Vec<usize> = vec![0; len];

    let start = map[0].iter().position(|&c| c == b'S').unwrap();
    timelines[start] = 1;

    for row in 1..map.len() {
        let mut next_timelines: Vec<usize> = vec![0; len];

        for col in 0..len {
            if timelines[col] == 0 {
                continue;
            }

            match map[row][col] {
                b'.' | b'S' => {
                    next_timelines[col] += timelines[col];
                }
                b'^' => {
                    if col > 0 {
                        next_timelines[col - 1] += timelines[col];
                    }
                    if col < len - 1 {
                        next_timelines[col + 1] += timelines[col];
                    }
                }
                _ => {}
            }
        }


        timelines = next_timelines;
        // println!("{:?}:", timelines);
    }

    timelines.iter().sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "\
    .......S.......\n\
    ...............\n\
    .......^.......\n\
    ...............\n\
    ......^.^......\n\
    ...............\n\
    .....^.^.^.....\n\
    ...............\n\
    ....^.^...^....\n\
    ...............\n\
    ...^.^...^.^...\n\
    ...............\n\
    ..^...^.....^..\n\
    ...............\n\
    .^.^.^.^.^...^.\n\
    ...............\
    ";

    #[test]
    fn test_part1() {
        let input = parse(INPUT);
        assert_eq!(part1(&input), 21);
    }

    #[test]
    fn test_part2() {
        let input = parse(INPUT);
        assert_eq!(part2(&input), 40);
    }
}
