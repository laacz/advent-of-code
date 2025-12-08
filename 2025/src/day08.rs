use std::usize;

fn main() {
    let input = std::fs::read_to_string("data/08.txt").expect("Failed to read input file");
    let input = parse(&input);
    println!("part1: {}", part1(&input, 1000));
    println!("part2: {}", part2(&input));
}

#[derive(Debug)]
struct Vec3 {
    x: u32,
    y: u32,
    z: u32,
}

impl Vec3 {
    fn distance2(&self, other: &Vec3) -> u64 {
        (self.x.abs_diff(other.x) as u64).pow(2)
            + (self.y.abs_diff(other.y) as u64).pow(2)
            + (self.z.abs_diff(other.z) as u64).pow(2)
    }
}

fn parse(input: &str) -> Vec<Vec3> {
    input
        .trim()
        .lines()
        .map(|line| {
            let parts: Vec<u32> = line
                .split(',')
                .map(|s| s.parse().expect("Failed to parse number"))
                .collect();
            Vec3 {
                x: parts[0],
                y: parts[1],
                z: parts[2],
            }
        })
        .collect()
}

fn find(parent: &mut [usize], x: usize) -> usize {
    if parent[x] != x {
        parent[x] = find(parent, parent[x]);
    }
    parent[x]
}

fn union(parent: &mut [usize], size: &mut [usize], x: usize, y: usize) {
    let root_x = find(parent, x);
    let root_y = find(parent, y);

    if root_x == root_y {
        return;
    }

    if size[root_x] < size[root_y] {
        parent[root_x] = root_y;
        size[root_y] += size[root_x];
    } else {
        parent[root_y] = root_x;
        size[root_x] += size[root_y];
    }
}

fn part1(input: &[Vec3], limit: usize) -> usize {
    let mut distances: Vec<(u64, usize, usize)> = vec![];
    for (i, a) in input.iter().enumerate() {
        for (j, b) in input.iter().enumerate().skip(i + 1) {
            let dist2 = a.distance2(b);
            distances.push((dist2, i, j));
        }
    }
    distances.sort_by_key(|&(d, _, _)| d);

    let mut parents = (0..input.len()).collect::<Vec<usize>>();
    let mut size = vec![1; input.len()];

    for &(_, i, j) in distances.iter().take(limit) {
        union(&mut parents, &mut size, i, j);
    }

    let mut root_sizes = vec![];
    for i in 0..input.len() {
        if parents[i] == i {
            root_sizes.push(size[i]);
        }
    }
    root_sizes.sort_by(|a, b| b.cmp(a));
    root_sizes.iter().take(3).product()
}

fn part2(input: &[Vec3]) -> u64 {
    let mut distances: Vec<(u64, usize, usize)> = vec![];
    for (i, a) in input.iter().enumerate() {
        for (j, b) in input.iter().enumerate().skip(i + 1) {
            let dist2 = a.distance2(b);
            distances.push((dist2, i, j));
        }
    }
    distances.sort_by_key(|&(d, _, _)| d);

    let mut parents = (0..input.len()).collect::<Vec<usize>>();
    let mut size = vec![1; input.len()];

    let mut last_i = 0;
    let mut last_j = 0;

    for &(_, i, j) in &distances {
        if find(&mut parents, i) != find(&mut parents, j) {
            last_i = i;
            last_j = j;
            union(&mut parents, &mut size, i, j);
        }
    }

    input[last_i].x as u64 * input[last_j].x as u64
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "\
    162,817,812\n\
    57,618,57\n\
    906,360,560\n\
    592,479,940\n\
    352,342,300\n\
    466,668,158\n\
    542,29,236\n\
    431,825,988\n\
    739,650,466\n\
    52,470,668\n\
    216,146,977\n\
    819,987,18\n\
    117,168,530\n\
    805,96,715\n\
    346,949,466\n\
    970,615,88\n\
    941,993,340\n\
    862,61,35\n\
    984,92,344\n\
    425,690,689\
    ";

    #[test]
    fn test_part1() {
        let input = parse(INPUT);
        assert_eq!(part1(&input, 10), 40);
    }

    #[test]
    fn test_part2() {
        let input = parse(INPUT);
        assert_eq!(part2(&input), 25272);
    }
}
