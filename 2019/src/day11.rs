use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/11.txt").expect("Failed to read file");
    let program = parse(&content);
    println!("part1: {}", part1(&program));
    println!("part2:");
    part2(&program);
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
                        panic!("No input available");
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

#[derive(Clone, Copy)]
enum Direction {
    Up,
    Right,
    Down,
    Left,
}

impl Direction {
    fn turn_left(self) -> Direction {
        match self {
            Direction::Up => Direction::Left,
            Direction::Left => Direction::Down,
            Direction::Down => Direction::Right,
            Direction::Right => Direction::Up,
        }
    }

    fn turn_right(self) -> Direction {
        match self {
            Direction::Up => Direction::Right,
            Direction::Right => Direction::Down,
            Direction::Down => Direction::Left,
            Direction::Left => Direction::Up,
        }
    }

    fn delta(self) -> (i32, i32) {
        match self {
            Direction::Up => (0, -1),
            Direction::Right => (1, 0),
            Direction::Down => (0, 1),
            Direction::Left => (-1, 0),
        }
    }
}

fn run_robot(program: &Vec<i64>, start_color: i64) -> HashMap<(i32, i32), i64> {
    let mut computer = IntcodeComputer::new(program);
    let mut panels: HashMap<(i32, i32), i64> = HashMap::new();
    let mut pos = (0, 0);
    let mut direction = Direction::Up;

    if start_color != 0 {
        panels.insert(pos, start_color);
    }

    loop {
        let current_color = *panels.get(&pos).unwrap_or(&0);
        computer.add_input(current_color);

        let paint_color = match computer.run_until_output() {
            Some(c) => c,
            None => break,
        };

        let turn = match computer.run_until_output() {
            Some(t) => t,
            None => break,
        };

        panels.insert(pos, paint_color);

        direction = if turn == 0 {
            direction.turn_left()
        } else {
            direction.turn_right()
        };

        let (dx, dy) = direction.delta();
        pos = (pos.0 + dx, pos.1 + dy);
    }

    panels
}

fn part1(program: &Vec<i64>) -> usize {
    run_robot(program, 0).len()
}

fn rainbow(t: f64) -> (u8, u8, u8) {
    let t = t * 6.0;
    let (r, g, b) = match t as u32 {
        0 => (1.0, t, 0.0),           // red -> yellow
        1 => (2.0 - t, 1.0, 0.0),     // yellow -> green
        2 => (0.0, 1.0, t - 2.0),     // green -> cyan
        3 => (0.0, 4.0 - t, 1.0),     // cyan -> blue
        4 => (t - 4.0, 0.0, 1.0),     // blue -> magenta
        _ => (1.0, 0.0, 6.0 - t),     // magenta -> red
    };
    ((r * 255.0) as u8, (g * 255.0) as u8, (b * 255.0) as u8)
}

fn part2(program: &Vec<i64>) {
    let panels = run_robot(program, 1);

    let min_x = panels.keys().map(|&(x, _)| x).min().unwrap();
    let max_x = panels.keys().map(|&(x, _)| x).max().unwrap();
    let min_y = panels.keys().map(|&(_, y)| y).min().unwrap();
    let max_y = panels.keys().map(|&(_, y)| y).max().unwrap();

    for y in min_y..=max_y {
        for x in min_x..=max_x {
            let color = *panels.get(&(x, y)).unwrap_or(&0);
            if color == 1 {
                let t = (x - min_x) as f64 / (max_x - min_x) as f64;
                let (r, g, b) = rainbow(t);
                print!("\x1b[38;2;{};{};{}m██\x1b[0m", r, g, b);
            } else {
                print!("  ");
            }
        }
        println!();
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_direction() {
        let dir = Direction::Up;
        assert!(matches!(dir.turn_left(), Direction::Left));
        assert!(matches!(dir.turn_right(), Direction::Right));

        let dir = Direction::Left;
        assert!(matches!(dir.turn_left(), Direction::Down));
        assert!(matches!(dir.turn_right(), Direction::Up));
    }
}
