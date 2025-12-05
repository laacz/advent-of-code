fn main() {
    let input = std::fs::read_to_string("data/05.txt").expect("Failed to read input file");
    let (ranges, ids) = parse(&input);
    println!("part1: {}", part1(&ranges, &ids));
    println!("part2: {}", part2(&ranges));
}

fn parse(input: &str) -> (Vec<Vec<usize>>, Vec<usize>) {
    let parts = input.split("\n\n").collect::<Vec<&str>>();
    let intervals = parts[0]
        .trim()
        .lines()
        .map(|line| {
            let bounds = line
                .split('-')
                .map(|x| x.parse::<usize>().unwrap())
                .collect::<Vec<usize>>();
            bounds
        })
        .collect::<Vec<Vec<usize>>>();
    let ids = parts[1]
        .trim()
        .lines()
        .map(|id| id.parse().unwrap())
        .collect::<Vec<usize>>();

    (intervals, ids)
}

fn part1(ranges: &[Vec<usize>], ids: &[usize]) -> usize {
    let mut ret = 0;

    for id in ids {
        for range in ranges {
            if id >= &range[0] && id <= &range[1] {
                ret += 1;
                break;
            }
        }
    }

    ret
}

fn part2(ranges: &[Vec<usize>]) -> usize {
    let mut ret = 0;

    let sorted_ranges = {
        let mut r = ranges.to_vec();
        r.sort_by(|a, b| a[0].cmp(&b[0]));
        r
    };

    let mut merged_ranges: Vec<Vec<usize>> = Vec::new();

    for range in &sorted_ranges {
        if merged_ranges.is_empty() {
            merged_ranges.push(range.clone());
            continue;
        }

        let last_range = merged_ranges.last_mut().unwrap();
        if range[0] <= last_range[1] {
            last_range[1] = last_range[1].max(range[1]);
        } else {
            merged_ranges.push(range.clone());
        }
    }

    merged_ranges.iter().for_each(|r| {
        ret += r[1] - r[0] + 1;
    });

    ret
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "\
            3-5\n\
            10-14\n\
            16-20\n\
            12-18\n\
            \n\
            1\n\
            5\n\
            8\n\
            11\n\
            17\n\
            32\
        ";
        let p = parse(input);
        assert_eq!(part1(&p.0, &p.1), 3);
    }

    #[test]
    fn test_part2() {
        let input = "\
            3-5\n\
            10-14\n\
            16-20\n\
            12-18\n\
            \n\
            1\n\
            5\n\
            8\n\
            11\n\
            17\n\
            32\
        ";
        let p = parse(input);
        assert_eq!(part2(&p.0), 14);
    }
}
