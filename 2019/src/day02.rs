fn main() {
    let mut ints: Vec<i32> = read_file_ints_comma("data/02.txt");
    ints[1] = 12;
    ints[2] = 2;
    println!("part1: {}", part1(&ints));
    println!("part2: {}", part2(&ints));
}

fn read_file_ints_comma(filename: &str) -> Vec<i32> {
    std::fs::read_to_string(filename)
        .expect("Failed to read file")
        .lines()
        .next()
        .unwrap()
        .split(',')
        .map(|line| line.parse().expect("Failed to parse line"))
        .collect()
}

fn run_program(ints: &[i32]) -> i32 {
    let mut program = ints.to_owned();
    let mut pos = 0;

    while pos < program.len() {
        let opcode = program[pos];
        match opcode {
            1 => {
                let a = program[pos + 1] as usize;
                let b = program[pos + 2] as usize;
                let c = program[pos + 3] as usize;
                program[c] = program[a] + program[b];
            }
            2 => {
                let a = program[pos + 1] as usize;
                let b = program[pos + 2] as usize;
                let c = program[pos + 3] as usize;
                program[c] = program[a] * program[b];
            }
            99 => break,
            _ => panic!("Invalid opcode"),
        }
        pos += 4;
    }
    program[0]
}

fn part1(ints: &[i32]) -> i32 {
    run_program(ints)
}

fn part2(ints: &[i32]) -> i32 {
    for noun in 0..100 {
        for verb in 0..100 {
            let mut program = ints.to_owned();
            program[1] = noun;
            program[2] = verb;
            if run_program(&program) == 19690720 {
                return 100 * noun + verb;
            }
        }
    }
    -1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(&vec![1, 0, 0, 0, 99]), 2);
        assert_eq!(part1(&vec![2, 3, 0, 3, 99]), 2);
        assert_eq!(part1(&vec![2, 4, 4, 5, 99, 0]), 2);
        assert_eq!(part1(&vec![1, 1, 1, 4, 99, 5, 6, 0, 99]), 30);
    }
}
