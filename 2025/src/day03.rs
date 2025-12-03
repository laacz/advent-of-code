fn main() {
    let input = std::fs::read_to_string("data/03.txt").expect("Failed to read input file");
    let ranges = parse(&input);
    println!("part1: {}", part1(&ranges));
    println!("part2: {}", part2(&ranges));
}

fn parse(input: &str) -> Vec<Vec<u64>> {
    input
        .trim()
        .lines()
        .map(|bank| {
            bank.chars().map(|c| c.to_digit(10).unwrap() as u64).collect::<Vec<u64>>()
        })
        .collect()
}

fn part1(batteries: &[Vec<u64>]) -> u64 {
    let mut ret  = 0;

    for bank in batteries {
        let mut max_joltage = 0;
        for i in 0..bank.len()-1 {
            let val = bank[i];
            max_joltage = max_joltage.max(val as u64 * 10 + bank[i+1..].iter().max().unwrap());
        }

        ret += max_joltage;
    }

    ret
}

fn part2(batteries: &[Vec<u64>]) -> u64 {
    let mut ret = 0;

    for bank in batteries {
        let mut max_joltage: u64 = 0;
        let mut pos = 0;
        for i in 0..12 {
            let val = bank[pos..bank.len()-11+i].iter().max().unwrap();
            pos += bank[pos..bank.len()-11+i].iter().position(|x| x == val).unwrap()+1;

            max_joltage += val * 10u64.pow(11 - i as u32);
        }

        ret += max_joltage;
    }

    ret
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "\
            987654321111111\n\
            811111111111119\n\
            234234234234278\n\
            818181911112111\
        ";
        let p = parse(input);
        assert_eq!(part1(&p), 357);
    }

    #[test]
    fn test_part2() {
        let input = "\
            987654321111111\n\
            811111111111119\n\
            234234234234278\n\
            818181911112111\
        ";
        let p = parse(input);
        assert_eq!(part2(&p), 3121910778619);
    }
}
