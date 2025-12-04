fn main() {
    let input = std::fs::read_to_string("data/04.txt").expect("Failed to read input file");
    let ranges = parse(&input);
    println!("part1: {}", part1(&ranges));
    println!("part2: {}", part2(&ranges));
}

fn parse(input: &str) -> Vec<Vec<usize>> {
    input
        .trim()
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| if c == '@' { 1 } else { 0 })
                .collect::<Vec<usize>>()
        })
        .collect()
}

fn print_map(input: &[Vec<usize>]) {
    for row in input {
        for &cell in row {
            if cell == 1 {
                print!("\x1b[92m@\x1b[0m"); // Bright green
            } else if cell == 2 {
                print!("\x1b[91mx\x1b[0m"); // Bright   
            } else {
                print!("\x1b[90m·\x1b[0m"); // Dark gray middle dot
            }
        }
        println!();
    }
}

fn num_rolls_around(map: &[Vec<usize>], x: usize, y: usize) -> usize {
    const DIRS: [(isize, isize); 8] = [
        (-1, -1),
        (0, -1),
        (1, -1),
        (-1, 0),
        (1, 0),
        (-1, 1),
        (0, 1),
        (1, 1),
    ];
    let mut ret = 0;
    for (dx, dy) in DIRS.iter() {
        let nx = x as isize + dx;
        let ny = y as isize + dy;
        if nx >= 0
            && ny >= 0
            && (ny as usize) < map.len()
            && (nx as usize) < map[0].len()
            && map[ny as usize][nx as usize] == 1
        {
            ret += 1;
        }
    }

    ret
}

fn part1(input: &[Vec<usize>]) -> usize {
    let mut ret: usize = 0;

    for y in 0..input.len() {
        for x in 0..input[0].len() {
            if input[y][x] == 1 && num_rolls_around(input, x, y) < 4 {
                ret += 1;
            }
        }
    }

    ret
}

fn part2(input: &[Vec<usize>]) -> usize {
    let mut ret: usize = 0;
    let mut map = input.to_vec();

    loop {
        let mut to_remove: Vec<(usize, usize)> = vec![];
        for y in 0..map.len() {
            for x in 0..map[0].len() {
                if map[y][x] == 1 && num_rolls_around(&map, x, y) < 4 {
                    to_remove.push((x, y));
                }
            }
        }
        for (x, y) in to_remove.iter() {
            map[*y][*x] = 2;
        }
        if to_remove.len() == 0 {
            break;  
        }
        ret += to_remove.len();
    }

    print_map(&map);

    ret
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "\
            ..@@.@@@@.\n\
            @@@.@.@.@@\n\
            @@@@@.@.@@\n\
            @.@@@@..@.\n\
            @@.@@@@.@@\n\
            .@@@@@@@.@\n\
            .@.@.@.@@@\n\
            @.@@@.@@@@\n\
            .@@@@@@@@.\n\
            @.@.@@@.@.\
        ";
        let p = parse(input);
        assert_eq!(part1(&p), 13);
    }

    #[test]
    fn test_part2() {
        let input = "\
            ..@@.@@@@.\n\
            @@@.@.@.@@\n\
            @@@@@.@.@@\n\
            @.@@@@..@.\n\
            @@.@@@@.@@\n\
            .@@@@@@@.@\n\
            .@.@.@.@@@\n\
            @.@@@.@@@@\n\
            .@@@@@@@@.\n\
            @.@.@@@.@.\
        ";
        let p = parse(input);
        assert_eq!(part2(&p), 423);
    }
}
