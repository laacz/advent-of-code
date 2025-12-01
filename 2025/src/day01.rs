fn main() {
    let str = std::fs::read_to_string("data/01.txt").expect("Failed to read input file");
    let input = parse(str);
    println!("part1: {}", part1(&input));
    println!("part2: {}", part2(&input));
}

fn parse(str: String) -> Vec<(i32, i32)> {
    str.lines()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let dir = if line.starts_with('L') { -1 } else { 1 };
            let num = line[1..].parse::<i32>().unwrap();
            (dir, num)
        })
        .collect()
}

fn part1(input: &[(i32, i32)]) -> i32 {
    let mut pos = 50;
    let mut ret = 0;
    input.iter().for_each(|(dir, num)| {
        pos = (pos + dir * num + 100) % 100;
        if pos == 0 {
            ret += 1
        }
    });

    ret
}

fn part2(input: &[(i32, i32)]) -> i32 {
    let mut pos = 50;
    let mut ret = 0;
    input.iter().for_each(|(dir, num)| {
        for _ in 0..*num {
            pos += *dir;
            pos %= 100;
            if pos == 0 {
                ret += 1;
            } else {
            }
        }        
    });

    ret
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = parse("L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82".to_string());
        assert_eq!(part1(&input), 3);
    }

    #[test]
    fn test_part2() {
        let input = parse("L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82".to_string());
        assert_eq!(part2(&input), 6);
    }

    #[test]
    fn test_part2_zeroes() {
        let input = parse("L50\nL100".to_string());
        assert_eq!(part2(&input), 2);
       
    }
}
