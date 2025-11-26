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

fn run_program(ints: Vec<i32>, input_val: i32) -> i32 {
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
                let address = program[pos+1] as usize;

                program[address] = input_val;
                pos += 2;
            }
            4 => {
                let val = get_param(&program, pos+1, param_modes[0]);
                output = val.to_string();
                pos += 2;
            }
            5 => {
                let a = get_param(&program, pos+1, param_modes[0]);
                let b = get_param(&program, pos+2, param_modes[1]);
                if a != 0 {
                    pos = b as usize;
                } else {
                    pos += 3;
                }
            }
            6 => {
                let a = get_param(&program, pos+1, param_modes[0]);
                let b = get_param(&program, pos+2, param_modes[1]);
                if a == 0 {
                    pos = b as usize;
                } else {
                    pos += 3;
                }
            }
            7 => {
                let a = get_param(&program, pos+1, param_modes[0]);
                let b = get_param(&program, pos+2, param_modes[1]);
                let c = program[pos+3] as usize;
                program[c] = if a < b { 1 } else { 0 };
                pos += 4;
            }
            8 => {
                let a = get_param(&program, pos+1, param_modes[0]);
                let b = get_param(&program, pos+2, param_modes[1]);
                let c = program[pos+3] as usize;
                program[c] = if a == b { 1 } else { 0 };
                pos += 4;
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
    run_program(input, 1)
}

fn part2(input: Vec<i32>) -> i32 {
    run_program(input, 5)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_output() {
        assert_eq!(run_program(parse("3,0,4,0,99"), 1), 1);
    }

    #[test]
    fn test_eq_pos() {
        let prog = parse("3,9,8,9,10,9,4,9,99,-1,8");
        assert_eq!(run_program(prog.clone(), 8), 1);
        assert_eq!(run_program(prog.clone(), 7), 0);
    }

    #[test]
    fn test_lt_pos() {
        let prog = parse("3,9,7,9,10,9,4,9,99,-1,8");
        assert_eq!(run_program(prog.clone(), 7), 1);
        assert_eq!(run_program(prog.clone(), 8), 0);
    }

    #[test]
    fn test_eq_imm() {
        let prog = parse("3,3,1108,-1,8,3,4,3,99");
        assert_eq!(run_program(prog.clone(), 8), 1);
        assert_eq!(run_program(prog.clone(), 7), 0);
    }

    #[test]
    fn test_lt_imm() {
        let prog = parse("3,3,1107,-1,8,3,4,3,99");
        assert_eq!(run_program(prog.clone(), 7), 1);
        assert_eq!(run_program(prog.clone(), 8), 0);
    }

    #[test]
    fn test_jump_pos() {
        let prog = parse("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9");
        assert_eq!(run_program(prog.clone(), 0), 0);
        assert_eq!(run_program(prog.clone(), 1), 1);
    }

    #[test]
    fn test_jump_imm() {
        let prog = parse("3,3,1105,-1,9,1101,0,0,12,4,12,99,1");
        assert_eq!(run_program(prog.clone(), 0), 0);
        assert_eq!(run_program(prog.clone(), 1), 1);
    }

    #[test]
    fn test_large_example() {
        let prog = parse("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99");
        assert_eq!(run_program(prog.clone(), 7), 999);
        assert_eq!(run_program(prog.clone(), 8), 1000);
        assert_eq!(run_program(prog.clone(), 9), 1001);
    }
}
