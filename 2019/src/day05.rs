fn main() {
    let content = std::fs::read_to_string("data/05.txt").expect("Failed to read file");
    let r: Vec<i32> = parse(&content);
    println!("part1: {}", part1(r.clone()));
    println!("part2: {}", part2(r.clone()));
}

fn parse(content: &str) -> Vec<i32> {
    content
        .lines()
        .next()
        .unwrap()
        .split(',')
        .map(|s| s.parse().unwrap())
        .collect()
}

fn get_param(program: &Vec<i32>, p: usize, mode: i32) -> i32 {
    if mode == 1 {
        return program[p] as i32;
    }
    program[program[p] as usize] as i32
}

fn run_program(ints: Vec<i32>) -> i32 {
    let mut program = ints.to_owned();
    let mut pos = 0;
    let mut output = String::from("0");

    while pos < program.len() {
        let opcode = program[pos] % 100;

        let param_modes = vec![
            program[pos] / 100 % 10,
            program[pos] / 1000 % 10,
            program[pos] / 10000 % 10,
        ];

        match opcode {
            1 => {
                let a = get_param(&program, pos+1, param_modes[0]);
                let b = get_param(&program, pos+2, param_modes[1]);
                let c = program[pos+3];

                program[c as usize] = a + b;
                pos += 4;
            }
            2 => {
                let a = get_param(&program, pos+1, param_modes[0]);
                let b = get_param(&program, pos+2, param_modes[1]);
                let c = program[pos+3];

                program[c as usize] = a * b;
                pos += 4;
            }
            3 => {
                let input = 1;
                let address = program[pos+1] as usize;

                program[address] = input;
                pos += 2;
            }
            4 => {
                let address = program[pos+1] as usize;

                output = program[address].to_string();
                pos += 2;
            }
            99 => break,
            _ => {
                panic!("Invalid opcode {}", opcode);
            }
        }
    }

    return output.parse().unwrap();
}

fn part1(input: Vec<i32>) -> i32 {
    run_program(input)
}

fn part2(input: Vec<i32>) -> i32 {
    run_program(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_output() {
        assert_eq!(run_program(parse("3,0,4,0,99")), 1);
    }
}
