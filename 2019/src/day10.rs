use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/10.txt").expect("Failed to read file");
    let r = parse(&content);
    println!("part1: {}", part1(&r));
    println!("part2: {}", part2(&r));
}

fn parse(content: &str) -> Vec<Vec<i64>> {
    content
        .trim()
        .lines()
        .map(|s| s.chars().map(|c| if c == '#' { 1 } else { 0 }).collect())
        .collect()
}

fn gcd(a: i64, b: i64) -> i64 {
    if b == 0 { a } else { gcd(b, a % b) }
}

fn visible_asteroids(field: &Vec<Vec<i64>>, x: usize, y: usize) -> usize {
    let mut angles: HashMap<(i64, i64), bool> = HashMap::new();
    let height = field.len() as i64;
    let width = field[0].len() as i64;

    for j in 0..height {
        for i in 0..width {
            if field[j as usize][i as usize] == 1 && !(i == x as i64 && j == y as i64) {
                let dy = j - y as i64;
                let dx = i - x as i64;
                let gcd = gcd(dx.abs(), dy.abs());
                let angle = (dx / gcd, dy / gcd);
                angles.insert(angle, true);
            }
        }
    }

    angles.len()
}

fn best_station(input: &Vec<Vec<i64>>) -> (usize, usize) {
    let mut best = (0, 0);
    let mut max = 0;
    for (y, row) in input.iter().enumerate() {
        for (x, cell) in row.iter().enumerate() {
            if *cell == 1 {
                let visible = visible_asteroids(input, x, y);
                if visible > max {
                    max = visible;
                    best = (x, y);
                }
            }
        }
    }
    best
}

fn part1(input: &Vec<Vec<i64>>) -> usize {
    let (x, y) = best_station(input);
    visible_asteroids(input, x, y)
}

fn part2(input: &Vec<Vec<i64>>) -> i64 {
    use std::collections::BTreeMap;
    
    let (sx, sy) = best_station(input);
    let (sx, sy) = (sx as i64, sy as i64);

    let mut by_angle: BTreeMap<i64, Vec<(i64, i64, i64)>> = BTreeMap::new();
    for (y, row) in input.iter().enumerate() {
        for (x, &c) in row.iter().enumerate() {
            if c == 1 && !(x as i64 == sx && y as i64 == sy) {
                let (dx, dy) = (x as i64 - sx, y as i64 - sy);
                let mut angle = (dx as f64).atan2(-dy as f64);
                if angle < 0.0 { angle += 2.0 * std::f64::consts::PI; }
                let angle_key = (angle * 1e9) as i64;
                let dist = dx * dx + dy * dy;
                by_angle.entry(angle_key).or_default().push((dist, x as i64, y as i64));
            }
        }
    }

    for group in by_angle.values_mut() {
        group.sort();
    }

    let mut count = 0;
    loop {
        for group in by_angle.values_mut() {
            if !group.is_empty() {
                let (_, x, y) = group.remove(0);
                count += 1;
                if count == 200 {
                    return x * 100 + y;
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        for (input, expected) in vec![
            (
                ".#..#\n\
                 .....\n\
                 #####\n\
                 ....#\n\
                 ...##",
                8,
            ),
            (
                "......#.#.\n\
                 #..#.#....\n\
                 ..#######.\n\
                 .#.#.###..\n\
                 .#..#.....\n\
                 ..#....#.#\n\
                 #..#....#.\n\
                 .##.#..###\n\
                 ##...#..#.\n\
                 .#....####",
                33,
            ),
            (
                "#.#...#.#.\n\
                 .###....#.\n\
                 .#....#...\n\
                 ##.#.#.#.#\n\
                 ....#.#.#.\n\
                 .##..###.#\n\
                 ..#...##..\n\
                 ..##....##\n\
                 ......#...\n\
                 .####.###.",
                35,
            ),
            (
                ".#..#..###\n\
                 ####.###.#\n\
                 ....###.#.\n\
                 ..###.##.#\n\
                 ##.##.#.#.\n\
                 ....###..#\n\
                 ..#.#..#.#\n\
                 #..#.#.###\n\
                 .##...##.#\n\
                 .....#.#..",
                41,
            ),
            (
                ".#..##.###...#######\n\
                 ##.############..##.\n\
                 .#.######.########.#\n\
                 .###.#######.####.#.\n\
                 #####.##.#.##.###.##\n\
                 ..#####..#.#########\n\
                 ####################\n\
                 #.####....###.#.#.##\n\
                 ##.#################\n\
                 #####.##.###..####..\n\
                 ..######..##.#######\n\
                 ####.##.####...##..#\n\
                 .#####..#.######.###\n\
                 ##...#.##########...\n\
                 #.##########.#######\n\
                 .####.#.###.###.#.##\n\
                 ....##.##.###..#####\n\
                 .#.#.###########.###\n\
                 #.#.#.#####.####.###\n\
                 ###.##.####.##.#..##",
                210,
            )
        ] {
            let parsed = parse(input);
            assert_eq!(part1(&parsed), expected);
        }
    }

    #[test]
    fn test_part2() {
        let input = ".#..##.###...#######\n\
                           ##.############..##.\n\
                           .#.######.########.#\n\
                           .###.#######.####.#.\n\
                           #####.##.#.##.###.##\n\
                           ..#####..#.#########\n\
                           ####################\n\
                           #.####....###.#.#.##\n\
                           ##.#################\n\
                           #####.##.###..####..\n\
                           ..######..##.#######\n\
                           ####.##.####...##..#\n\
                           .#####..#.######.###\n\
                           ##...#.##########...\n\
                           #.##########.#######\n\
                           .####.#.###.###.#.##\n\
                           ....##.##.###..#####\n\
                           .#.#.###########.###\n\
                           #.#.#.#####.####.###\n\
                           ###.##.####.##.#..##";
        let parsed = parse(input);
        assert_eq!(part2(&parsed), 802);
    }

}
