fn main() {
    let input = std::fs::read_to_string("data/02.txt").expect("Failed to read input file");
    let ranges = parse(&input);
    println!("part1: {}", part1(&ranges));
    println!("part2: {}", part2(&ranges));
}

fn parse(input: &str) -> Vec<(u64, u64)> {
    input
        .trim()
        .split(',')
        .filter(|s| !s.is_empty())
        .map(|range| {
            let mut parts = range.split('-');
            let start: u64 = parts.next().unwrap().parse().unwrap();
            let end: u64 = parts.next().unwrap().parse().unwrap();
            (start, end)
        })
        .collect()
}

// any odd number of digits has 0 invalid ids
// 2 digits - 1-9 - 9 invalid ids
// 4 digits - 10-99 - 90 invalid ids
// 6 digits - 100-999 - 900 invalid ids
// 8 digits - 1000-9999 - 9000 invalid ids
// so it's 9 * 10 ^ (n/2)
// 
// 3859_3856-3859_3862 - subset of 9000 invalid ids.
// 990_244-1_009_337 - subset of 900 invalid ids
fn sum_invalids(start: u64, end: u64, n: u32) -> u64 {
    let multiplier = 10u64.pow(n) + 1;

    let min_base = if n == 1 { 1 } else { 10u64.pow(n - 1) };
    let max_base = 10u64.pow(n) - 1;

    let base_lo = (start + multiplier - 1) / multiplier;
    let base_hi = end / multiplier;

    let lo = base_lo.max(min_base);
    let hi = base_hi.min(max_base);

    if lo > hi {
        return 0;
    }

    let count = hi - lo + 1;
    let sum_of_bases = count * (lo + hi) / 2;

    multiplier * sum_of_bases
}

fn part1(ranges: &[(u64, u64)]) -> u64 {
    let mut total = 0;

    for &(start, end) in ranges {
        for n in 1..=10 {
            total += sum_invalids(start, end, n);
        }
    }

    total
}

fn is_repeated(n: u64) -> bool {
    let s = n.to_string();
    let len = s.len();

    for pat_len in 1..=len / 2 {
        if len % pat_len == 0 {
            let pattern = &s[..pat_len];
            if pattern.repeat(len / pat_len) == s {
                return true;
            }
        }
    }
    false
}

fn part2(ranges: &[(u64, u64)]) -> u64 {
    let mut total = 0;
    for &(start, end) in ranges {
        for n in start..=end {
            if is_repeated(n) {
                total += n;
            }
        }
    }
    total
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
        let ranges = parse(input);
        assert_eq!(part1(&ranges), 1227775554);
    }

    #[test]
    fn test_part2() {
        let input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
        let ranges = parse(input);
        assert_eq!(part2(&ranges), 4174379265);
    }
}
