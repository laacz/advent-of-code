use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/13.txt").expect("Failed to read file");
    let program = parse(&content);
    println!("part1: {}", part1(&program));
    println!("part2: {}", part2(&program));
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
            0 => self.get_memory(raw as usize),
            1 => raw,
            2 => self.get_memory((self.relative_base + raw) as usize),
            _ => panic!("Unknown parameter mode: {}", mode),
        }
    }

    fn get_write_addr(&self, offset: usize, mode: i64) -> usize {
        let raw = self.get_memory(self.pos + offset);
        match mode {
            0 => raw as usize,
            2 => (self.relative_base + raw) as usize,
            _ => panic!("Invalid write mode: {}", mode),
        }
    }

    fn add_input(&mut self, value: i64) {
        self.inputs.push(value);
    }

    fn run_until_output(&mut self) -> Option<i64> {
        while !self.halted {
            let instruction = self.get_memory(self.pos);
            let opcode = instruction % 100;

            let mode1 = (instruction / 100) % 10;
            let mode2 = (instruction / 1000) % 10;
            let mode3 = (instruction / 10000) % 10;

            match opcode {
                1 => {
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, a + b);
                    self.pos += 4;
                }
                2 => {
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, a * b);
                    self.pos += 4;
                }
                3 => {
                    if self.input_idx >= self.inputs.len() {
                        return None; // Need input, pause execution
                    }
                    let addr = self.get_write_addr(1, mode1);
                    self.set_memory(addr, self.inputs[self.input_idx]);
                    self.input_idx += 1;
                    self.pos += 2;
                }
                4 => {
                    let val = self.get_param(1, mode1);
                    self.pos += 2;
                    return Some(val);
                }
                5 => {
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    if a != 0 {
                        self.pos = b as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                6 => {
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    if a == 0 {
                        self.pos = b as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                7 => {
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, if a < b { 1 } else { 0 });
                    self.pos += 4;
                }
                8 => {
                    let a = self.get_param(1, mode1);
                    let b = self.get_param(2, mode2);
                    let addr = self.get_write_addr(3, mode3);
                    self.set_memory(addr, if a == b { 1 } else { 0 });
                    self.pos += 4;
                }
                9 => {
                    let a = self.get_param(1, mode1);
                    self.relative_base += a;
                    self.pos += 2;
                }
                99 => {
                    self.halted = true;
                    return None;
                }
                _ => {
                    panic!("Invalid opcode {} at position {}", opcode, self.pos);
                }
            }
        }
        None
    }
}

fn part1(program: &Vec<i64>) -> usize {
    let mut computer = IntcodeComputer::new(program);
    let mut blocks = 0;

    loop {
        let _x = match computer.run_until_output() {
            Some(v) => v,
            None => break,
        };
        let _y = match computer.run_until_output() {
            Some(v) => v,
            None => break,
        };
        let tile = match computer.run_until_output() {
            Some(v) => v,
            None => break,
        };
        if tile == 2 {
            blocks += 1;
        }
    }
    blocks
}

fn part2(program: &Vec<i64>) -> i64 {
    let mut computer = IntcodeComputer::new(program);
    computer.set_memory(0, 2); // Play for free

    let mut score = 0i64;
    let mut ball_x = 0i64;
    let mut paddle_x = 0i64;

    loop {
        let x = match computer.run_until_output() {
            Some(v) => v,
            None if computer.halted => break,
            None => {
                // Needs input
                let joystick = (ball_x - paddle_x).signum();
                computer.add_input(joystick);
                continue;
            }
        };
        let y = match computer.run_until_output() {
            Some(v) => v,
            None if computer.halted => break,
            None => {
                let joystick = (ball_x - paddle_x).signum();
                computer.add_input(joystick);
                continue;
            }
        };
        let tile = match computer.run_until_output() {
            Some(v) => v,
            None if computer.halted => break,
            None => {
                let joystick = (ball_x - paddle_x).signum();
                computer.add_input(joystick);
                continue;
            }
        };

        if x == -1 && y == 0 {
            score = tile;
        } else {
            match tile {
                3 => paddle_x = x,
                4 => ball_x = x,
                _ => {}
            }
        }
    }
    score
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_intcode_basic() {
        let program = parse("104,42,99");
        let mut computer = IntcodeComputer::new(&program);
        assert_eq!(computer.run_until_output(), Some(42));
        assert_eq!(computer.run_until_output(), None);
        assert!(computer.halted);
    }
}
