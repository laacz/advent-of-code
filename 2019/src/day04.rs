fn main() {
    let content = std::fs::read_to_string("data/04.txt").expect("Failed to read file");
    let r: Vec<i32> = parse(&content);
    println!("part1: {}", part1(r[0], r[1]));
    println!("part2: {}", part2(r[0], r[1]));
}

fn parse(content: &str) -> Vec<i32> {
    content
        .lines()
        .next()
        .unwrap()
        .split('-')
        .map(|s| s.parse().unwrap())
        .collect()
}

fn is_valid(password: &str, v2: bool) -> bool {
    let digits = password
        .chars()
        .map(|c| c.to_digit(10).unwrap())
        .collect::<Vec<u32>>();
    let mut has_double = false;
    for i in 1..digits.len() {
        if digits[i] < digits[i - 1] {
            return false;
        }
        if digits[i] == digits[i - 1] {
            has_double = true;
        }
    }

    if v2 {
        let mut idx = 0;
        while idx < digits.len() {
            let current = digits[idx];
            let mut count = 1;
            while idx + count < digits.len() && digits[idx + count] == current {
                count += 1;
            }
            if count == 2 {
                return true;
            }
            idx += count;
        }
        false
    } else {
        has_double
    }
}

fn part1(start: i32, end: i32) -> i32 {
    let mut count = 0;
    for i in start..=end {
        if is_valid(&i.to_string(), false) {
            count += 1;
        }
    }
    count
}

fn part2(start: i32, end: i32) -> i32 {
    let mut count = 0;
    for i in start..=end {
        if is_valid(&i.to_string(), true) {
            count += 1;
        }
    }
    count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(is_valid("111111", false), true);
        assert_eq!(is_valid("223450", false), false);
        assert_eq!(is_valid("123789", false), false);
        assert_eq!(is_valid("112233", true), true);
        assert_eq!(is_valid("123444", true), false);
        assert_eq!(is_valid("111122", true), true);
    }
}
