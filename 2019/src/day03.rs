use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/03.txt").expect("Failed to read file");
    let w: Vec<Vec<&str>> = parse(&content);
    println!("part1: {}", part1(&w[0], &w[1]));
    println!("part2: {}", part2(&w[0], &w[1]));
}

fn parse(content: &str) -> Vec<Vec<&str>> {
    content
        .lines()
        .map(|line| line.split(',').collect())
        .collect()
}

fn trace_wire(wire: &Vec<&str>) -> HashMap<(i32, i32), i32> {
    let mut map = HashMap::new();
    let mut x = 0;
    let mut y = 0;
    let mut steps = 0;

    for mv in wire {
        let dir = mv.chars().next().unwrap();
        let len: i32 = mv[1..].parse().unwrap();
        let (dx, dy) = match dir {
            'R' => (1, 0),
            'L' => (-1, 0),
            'U' => (0, 1),
            'D' => (0, -1),
            _ => panic!("Invalid direction"),
        };
        for _ in 0..len {
            (x, y) = (x + dx, y + dy);
            steps += 1;
            map.entry((x, y)).or_insert(steps);
        }
    }

    map
}

fn part1(w1: &Vec<&str>, w2: &Vec<&str>) -> i32 {
    let result = trace_wire(w1);
    let result2 = trace_wire(w2);
    let mut min = i32::MAX;
    for (k, _v) in result {
        if result2.contains_key(&k) {
            min = min.min(k.0.abs() + k.1.abs())
        }
    }

    min
}

fn part2(w2: &Vec<&str>, w1: &Vec<&str>) -> i32 {
    let result = trace_wire(w1);
    let result2 = trace_wire(w2);
    let mut min = i32::MAX;
    for (k, v) in result {
        if result2.contains_key(&k) {
            min = min.min(v + result2[&k]);
        }
    }
    min
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let c = parse(
            "R75,D30,R83,U83,L12,D49,R71,U7,L72\n\
            U62,R66,U55,R34,D71,R55,D58,R83",
        );
        assert_eq!(part1(&c[0], &c[1]), 159);
        assert_eq!(part2(&c[0], &c[1]), 610);
        let c = parse(
            "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\n\
            U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
        );
        assert_eq!(part1(&c[0], &c[1]), 135);
        assert_eq!(part2(&c[0], &c[1]), 410);
    }
}
