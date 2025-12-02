use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/09.txt").expect("Failed to read file");
    let r = parse(&content);
    println!("part1: {}", part1(&r));
    println!("part2: {}", part2(&r));
}

fn parse(content: &str) -> Vec<i64> {
    content
        .lines()
        .next()
        .unwrap()
        .split(',')
        .map(|s| s.parse().unwrap())
        .collect()
}

struct IntcodeComputer {
    memory: HashMap<usize, i64>,
    pos: usize,
    relative_base: i64,
    inputs: Vec<i64>,
    input_idx: usize,
    outputs: Vec<i64>,
    halted: bool,
}

impl IntcodeComputer {
    fn new(program: &Vec<i64>) -> Self {
        let mut memory = HashMap::new();
        for (i, &val) in program.iter().enumerate() {
            memory.insert(i, val);
        }
        Self {
            memory,
            pos: 0,
            relative_base: 0,
            inputs: Vec::new(),
            input_idx: 0,
            outputs: Vec::new(),
            halted: false,
        }
    }

    fn get_memory(&self, addr: usize) -> i64 {
        *self.memory.get(&addr).unwrap_or(&0)
    }

    fn set_memory(&mut self, addr: usize, value: i64) {
        self.memory.insert(addr, value);
    }

    fn get_param(&self, offset: usize, mode: i64) -> i64 {
        let raw = self.get_memory(self.pos + offset);
        match mode {
            0 => self.get_memory(raw as usize), // position mode
            1 => raw,                            // immediate mode
            2 => self.get_memory((self.relative_base + raw) as usize), // relative mode
            _ => panic!("Unknown parameter mode: {}", mode),
        }
    }

    fn get_write_addr(&self, offset: usize, mode: i64) -> usize {
        let raw = self.get_memory(self.pos + offset);
        match mode {
            0 => raw as usize,                              // position mode
            2 => (self.relative_base + raw) as usize,       // relative mode
            _ => panic!("Invalid write mode: {}", mode),
        }
    }

    fn add_input(&mut self, value: i64) {
        self.inputs.push(value);
    }

    fn run(&mut self) {
        while !self.halted {
            let instruction = self.get_memory(self.pos);
            let opcode = instruction % 100;

            let mode1 = (instruction / 100) % 10;
            let mode2 = (instruction / 1000) % 10;
            let mode3 = (instruction / 10000) % 10;

            match opcode {
                1 => {
                    // Add
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, a + b);
                    self.pos += 4;
                }
                2 => {
                    // Multiply
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, a * b);
                    self.pos += 4;
                }
                3 => {
                    // Input
                    if self.input_idx >= self.inputs.len() {
                        panic!("No input available");
                    }
                    let addr = self.get_write_addr(1, mode1);
                    self.set_memory(addr, self.inputs[self.input_idx]);
                    self.input_idx += 1;
                    self.pos += 2;
                }
                4 => {
                    // Output
                    let val = self.get_param(1, mode1);
                    self.outputs.push(val);
                    self.pos += 2;
                }
                5 => {
                    // Jump-if-true
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    if a != 0 {
                        self.pos = b as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                6 => {
                    // Jump-if-false
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    if a == 0 {
                        self.pos = b as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                7 => {
                    // Less than
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, if a < b { 1 } else { 0 });
                    self.pos += 4;
                }
                8 => {
                    // Equals
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, if a == b { 1 } else { 0 });
                    self.pos += 4;
                }
                9 => {
                    // Adjust relative base
                    let a = self.get_param(1, mode1);
                    self.relative_base += a;
                    self.pos += 2;
                }
                99 => {
                    self.halted = true;
                }
                _ => {
                    panic!("Invalid opcode {} at position {}", opcode, self.pos);
                }
            }
        }
    }
}

fn run_program(program: &Vec<i64>, input: i64) -> Vec<i64> {
    let mut computer = IntcodeComputer::new(program);
    computer.add_input(input);
    computer.run();
    computer.outputs
}

fn part1(input: &Vec<i64>) -> i64 {
    let outputs = run_program(input, 1);
    *outputs.last().unwrap()
}

fn part2(input: &Vec<i64>) -> i64 {
    let outputs = run_program(input, 2);
    *outputs.last().unwrap()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_quine() {
        let program = parse("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99");
        let outputs = run_program(&program, 0);
        assert_eq!(outputs, program);
    }

    #[test]
    fn test_16_digit() {
        let program = parse("1102,34915192,34915192,7,4,7,99,0");
        let outputs = run_program(&program, 0);
        assert_eq!(outputs.len(), 1);
        assert_eq!(outputs[0].to_string().len(), 16);
    }

    #[test]
    fn test_large_number() {
        let program = parse("104,1125899906842624,99");
        let outputs = run_program(&program, 0);
        assert_eq!(outputs, vec![1125899906842624]);
    }
}
