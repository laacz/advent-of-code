fn main() {
    let content = std::fs::read_to_string("data/07.txt").expect("Failed to read file");
    let r = parse(&content);
    println!("part1: {}", part1(&r));
    println!("part2: {}", part2(&r));
}

fn parse(content: &str) -> Vec<i32> {
    content
        .lines()
        .next()
        .unwrap()
        .split(",")
        .map(|s| s.parse().unwrap())
        .collect()
}

fn get_param(program: &Vec<i32>, p: usize, mode: i32) -> i32 {
    if mode == 1 {
        return program[p] as i32;
    }
    program[program[p] as usize] as i32
}

fn run_program(ints: Vec<i32>, inputs: Vec<i32>) -> i32 {
    let mut program = ints.to_owned();
    let mut pos = 0;
    let mut input_idx = 0;
    let mut output = 0;

    while pos < program.len() {
        let opcode = program[pos] % 100;

        let param_modes = vec![
            program[pos] / 100 % 10,
            program[pos] / 1000 % 10,
            program[pos] / 10000 % 10,
        ];

        match opcode {
            1 => {
                let a = get_param(&program, pos + 1, param_modes[0]);
                let b = get_param(&program, pos + 2, param_modes[1]);
                let c = program[pos + 3];
                program[c as usize] = a + b;
                pos += 4;
            }
            2 => {
                let a = get_param(&program, pos + 1, param_modes[0]);
                let b = get_param(&program, pos + 2, param_modes[1]);
                let c = program[pos + 3];
                program[c as usize] = a * b;
                pos += 4;
            }
            3 => {
                let address = program[pos + 1] as usize;
                program[address] = inputs[input_idx];
                input_idx += 1;
                pos += 2;
            }
            4 => {
                output = get_param(&program, pos + 1, param_modes[0]);
                pos += 2;
            }
            5 => {
                let a = get_param(&program, pos + 1, param_modes[0]);
                let b = get_param(&program, pos + 2, param_modes[1]);
                if a != 0 {
                    pos = b as usize;
                } else {
                    pos += 3;
                }
            }
            6 => {
                let a = get_param(&program, pos + 1, param_modes[0]);
                let b = get_param(&program, pos + 2, param_modes[1]);
                if a == 0 {
                    pos = b as usize;
                } else {
                    pos += 3;
                }
            }
            7 => {
                let a = get_param(&program, pos + 1, param_modes[0]);
                let b = get_param(&program, pos + 2, param_modes[1]);
                let c = program[pos + 3] as usize;
                program[c] = if a < b { 1 } else { 0 };
                pos += 4;
            }
            8 => {
                let a = get_param(&program, pos + 1, param_modes[0]);
                let b = get_param(&program, pos + 2, param_modes[1]);
                let c = program[pos + 3] as usize;
                program[c] = if a == b { 1 } else { 0 };
                pos += 4;
            }
            99 => break,
            _ => {
                panic!("Invalid opcode {}", opcode);
            }
        }
    }

    output
}

fn permutations(items: Vec<i32>) -> Vec<Vec<i32>> {
    if items.len() == 1 {
        return vec![items];
    }

    let mut result = Vec::new();
    for i in 0..items.len() {
        let mut rest = items.clone();
        let first = rest.remove(i);
        for mut perm in permutations(rest) {
            perm.insert(0, first);
            result.push(perm);
        }
    }
    result
}

fn run_amplifiers(program: &Vec<i32>, phases: &Vec<i32>) -> i32 {
    let mut signal = 0;
    for phase in phases {
        signal = run_program(program.clone(), vec![*phase, signal]);
    }
    signal
}

fn part1(input: &Vec<i32>) -> i32 {
    permutations(vec![0, 1, 2, 3, 4])
        .iter()
        .map(|phases| run_amplifiers(input, phases))
        .max()
        .unwrap()
}

struct Amplifier {
    program: Vec<i32>,
    pos: usize,
    inputs: Vec<i32>,
    halted: bool,
}

impl Amplifier {
    fn new(program: Vec<i32>, phase: i32) -> Self {
        Self {
            program,
            pos: 0,
            inputs: vec![phase],
            halted: false,
        }
    }

    fn send_input(&mut self, value: i32) {
        self.inputs.push(value);
    }

    fn run_until_output(&mut self) -> Option<i32> {
        let mut input_idx = 0;

        while self.pos < self.program.len() {
            let opcode = self.program[self.pos] % 100;

            let param_modes = vec![
                self.program[self.pos] / 100 % 10,
                self.program[self.pos] / 1000 % 10,
                self.program[self.pos] / 10000 % 10,
            ];

            let get_param = |prog: &Vec<i32>, p: usize, mode: i32| -> i32 {
                if mode == 1 {
                    prog[p] as i32
                } else {
                    prog[prog[p] as usize] as i32
                }
            };

            match opcode {
                1 => {
                    let a = get_param(&self.program, self.pos + 1, param_modes[0]);
                    let b = get_param(&self.program, self.pos + 2, param_modes[1]);
                    let c = self.program[self.pos + 3];
                    self.program[c as usize] = a + b;
                    self.pos += 4;
                }
                2 => {
                    let a = get_param(&self.program, self.pos + 1, param_modes[0]);
                    let b = get_param(&self.program, self.pos + 2, param_modes[1]);
                    let c = self.program[self.pos + 3];
                    self.program[c as usize] = a * b;
                    self.pos += 4;
                }
                3 => {
                    if input_idx >= self.inputs.len() {
                        return None; // Wait for input
                    }
                    let address = self.program[self.pos + 1] as usize;
                    self.program[address] = self.inputs[input_idx];
                    input_idx += 1;
                    self.pos += 2;
                }
                4 => {
                    let output = get_param(&self.program, self.pos + 1, param_modes[0]);
                    self.pos += 2;
                    self.inputs.drain(0..input_idx);
                    return Some(output);
                }
                5 => {
                    let a = get_param(&self.program, self.pos + 1, param_modes[0]);
                    let b = get_param(&self.program, self.pos + 2, param_modes[1]);
                    if a != 0 {
                        self.pos = b as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                6 => {
                    let a = get_param(&self.program, self.pos + 1, param_modes[0]);
                    let b = get_param(&self.program, self.pos + 2, param_modes[1]);
                    if a == 0 {
                        self.pos = b as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                7 => {
                    let a = get_param(&self.program, self.pos + 1, param_modes[0]);
                    let b = get_param(&self.program, self.pos + 2, param_modes[1]);
                    let c = self.program[self.pos + 3] as usize;
                    self.program[c] = if a < b { 1 } else { 0 };
                    self.pos += 4;
                }
                8 => {
                    let a = get_param(&self.program, self.pos + 1, param_modes[0]);
                    let b = get_param(&self.program, self.pos + 2, param_modes[1]);
                    let c = self.program[self.pos + 3] as usize;
                    self.program[c] = if a == b { 1 } else { 0 };
                    self.pos += 4;
                }
                99 => {
                    self.halted = true;
                    self.inputs.drain(0..input_idx);
                    return None;
                }
                _ => {
                    panic!("Invalid opcode {}", opcode);
                }
            }
        }

        self.halted = true;
        None
    }
}

fn run_feedback_loop(program: &Vec<i32>, phases: &Vec<i32>) -> i32 {
    let mut amps: Vec<Amplifier> = phases
        .iter()
        .map(|&phase| Amplifier::new(program.clone(), phase))
        .collect();

    let mut signal = 0;
    let mut last_output = 0;

    loop {
        let mut all_halted = true;

        for amp in &mut amps {
            if amp.halted {
                continue;
            }
            all_halted = false;
            amp.send_input(signal);
            if let Some(output) = amp.run_until_output() {
                signal = output;
                last_output = output;
            }
        }

        if all_halted {
            break;
        }
    }

    last_output
}

fn part2(input: &Vec<i32>) -> i32 {
    permutations(vec![5, 6, 7, 8, 9])
        .iter()
        .map(|phases| run_feedback_loop(input, phases))
        .max()
        .unwrap()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(&parse("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")),
            43210,
        );
        assert_eq!(
            part1(&parse(
                "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
            )),
            54321,
        );
        assert_eq!(
            part1(&parse(
                "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
            )),
            65210,
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(&parse(
                "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
            )),
            139629729,
        );
        assert_eq!(
            part2(&parse(
                "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
            )),
            18216,
        );
    }
}
