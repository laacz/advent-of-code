fn main() {
    let input = std::fs::read_to_string("data/06.txt").expect("Failed to read input file");
    let cols = parse(&input);
    println!("part1: {}", part1(&cols));
    println!("part2: {}", part2(&cols));
}

fn parse(input: &str) -> Vec<Vec<String>> {
    let lines: Vec<&str> = input.trim().lines().collect();
    let mut ret = vec![];

    let op_positions: Vec<usize> = lines
        .iter()
        .last()
        .unwrap()
        .char_indices()
        .filter(|(_, c)| *c == '*' || *c == '+')
        .map(|(i, _)| i)
        .collect();

    let num_pos = op_positions.len() - 1;
    for i in 0..num_pos + 1 {
        let mut rows = vec![];
        let pos = op_positions[i];
        for row in lines.iter() {
            let end_pos = if i < num_pos {
                op_positions[i + 1] - 1
            } else {
                row.len()
            };
            rows.push(row[pos..end_pos].to_string());
        }
        ret.push(rows);
    }

    ret
}

fn part1(nums: &[Vec<String>]) -> usize {
    let mut ret = 0 as usize;

    for col in 0..nums.len() {
        let rows = &nums[col];
        let mut result = rows.first().unwrap().trim().parse::<usize>().unwrap();
        let op = rows.last().unwrap().trim();
        for r in 1..rows.len() - 1 {
            let s = &rows[r];
            let num = s.trim().parse::<usize>().unwrap();

            if op == "*" {
                result *= num;
            } else {
                result += num;
            }
        }
        ret += result;
    }

    ret
}

fn part2(nums: &[Vec<String>]) -> usize {
    let mut ret = 0 as usize;

    let row_count = nums.first().unwrap().len();
    for group in nums.iter() {
        let op = group.last().unwrap().trim();
        let mut interim = if op == "*" { 1 } else { 0 };

        for pos in 0..group.first().unwrap().len() {
            let mut num = String::new();
            for row in 0..row_count - 1 {
                num.push(group[row].chars().nth(pos).unwrap());
            }

            let num = num.trim().parse::<usize>().unwrap();

            if op == "*" {
                interim *= num;
            } else {
                interim += num;
            }
        }
        ret += interim;
    }

    ret
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "\
123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   + ";

    #[test]
    fn test_part1() {
        let input = parse(INPUT);
        assert_eq!(part1(&input), 4277556);
    }

    #[test]
    fn test_part2() {
        let input = parse(INPUT);
        assert_eq!(part2(&input), 3263827);
    }
}
