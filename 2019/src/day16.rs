fn main() {
    let content = std::fs::read_to_string("data/16.txt").expect("Failed to read file");
    let input = parse(&content);
    println!("part1: {}", part1(&input));
    println!("part2: {}", part2(&input));
}

fn parse(content: &str) -> Vec<i32> {
    content
        .trim()
        .lines()
        .next()
        .unwrap()
        .chars()
        .map(|s| (s as i32) - ('0' as i32))
        .collect()
}

fn pattern_val(output_pos: usize, input_pos: usize) -> i32 {
    const BASE_PATTERN: [i32; 4] = [0, 1, 0, -1];
    let index = (input_pos + 1) / (output_pos + 1) % 4;
    BASE_PATTERN[index]
}

fn fft_phase(input: &[i32]) -> Vec<i32> {
    let n = input.len();
    let mut output = vec![0; n];

    for i in 0..n {
        let mut sum: i32 = 0;
        for j in 0..n {
            sum += input[j] * pattern_val(i, j);
        }
        output[i] = (sum.abs() % 10) as i32;
    }

    output
}

fn part1(input: &[i32]) -> String {
    let mut signal = input.to_vec();

    for _ in 0..100 {
        signal = fft_phase(&signal);
    }

    signal[0..8]
        .iter()
        .map(|d| d.to_string())
        .collect::<String>()
}

fn part2(input: &[i32]) -> String {
    let offset: usize = input[0..7]
        .iter()
        .fold(0, |acc, &d| acc * 10 + d as usize);

    let full_len = input.len() * 10000;

    let mut signal: Vec<i32> = (offset..full_len)
        .map(|i| input[i % input.len()])
        .collect();

    for _ in 0..100 {
        let mut suffix_sum = 0;
        for i in (0..signal.len()).rev() {
            suffix_sum = (suffix_sum + signal[i]) % 10;
            signal[i] = suffix_sum;
        }
    }

    signal[0..8]
        .iter()
        .map(|d| d.to_string())
        .collect::<String>()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_fft_phases() {
        let input = parse("12345678");
        let after1 = fft_phase(&input);
        assert_eq!(after1, vec![4, 8, 2, 2, 6, 1, 5, 8]);

        let after2 = fft_phase(&after1);
        assert_eq!(after2, vec![3, 4, 0, 4, 0, 4, 3, 8]);

        let after3 = fft_phase(&after2);
        assert_eq!(after3, vec![0, 3, 4, 1, 5, 5, 1, 8]);

        let after4 = fft_phase(&after3);
        assert_eq!(after4, vec![0, 1, 0, 2, 9, 4, 9, 8]);
    }

    #[test]
    fn test_part1_example1() {
        assert_eq!(part1(&parse("80871224585914546619083218645595")), "24176176");
    }

    #[test]
    fn test_part1_example2() {
        assert_eq!(part1(&parse("19617804207202209144916044189917")), "73745418");
    }

    #[test]
    fn test_part1_example3() {
        assert_eq!(part1(&parse("69317163492948606335995924319873")), "52432133");
    }

    #[test]
    fn test_part2_example1() {
        assert_eq!(part2(&parse("03036732577212944063491565474664")), "84462026");
    }

    #[test]
    fn test_part2_example2() {
        assert_eq!(part2(&parse("02935109699940807407585447034323")), "78725270");
    }

    #[test]
    fn test_part2_example3() {
        assert_eq!(part2(&parse("03081770884921959731165446850517")), "53553731");
    }
}
