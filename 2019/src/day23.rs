use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/23.txt").expect("Failed to read file");
    let prog = parse(&content);
    println!("part1: {}", part1(&prog));
    println!("part2: {}", part2(&prog));
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

struct Cpu {
    mem: HashMap<usize, i64>,
    pos: usize,
    base: i64,
    inp: Vec<i64>,
    inp_idx: usize,
    halted: bool,
}

impl Cpu {
    fn new(prog: &Vec<i64>) -> Self {
        let mut mem = HashMap::new();
        for (i, &v) in prog.iter().enumerate() {
            mem.insert(i, v);
        }
        Self {
            mem,
            pos: 0,
            base: 0,
            inp: Vec::new(),
            inp_idx: 0,
            halted: false,
        }
    }

    fn get(&self, a: usize) -> i64 {
        *self.mem.get(&a).unwrap_or(&0)
    }

    fn set(&mut self, a: usize, v: i64) {
        self.mem.insert(a, v);
    }

    fn param(&self, off: usize, mode: i64) -> i64 {
        let raw = self.get(self.pos + off);
        match mode {
            0 => self.get(raw as usize),
            1 => raw,
            2 => self.get((self.base + raw) as usize),
            _ => panic!(),
        }
    }

    fn addr(&self, off: usize, mode: i64) -> usize {
        let raw = self.get(self.pos + off);
        match mode {
            0 => raw as usize,
            2 => (self.base + raw) as usize,
            _ => panic!(),
        }
    }

    fn input(&mut self, v: i64) {
        self.inp.push(v);
    }

    fn run(&mut self) -> Option<i64> {
        while !self.halted {
            let op = self.get(self.pos);
            let (m1, m2, m3) = ((op / 100) % 10, (op / 1000) % 10, (op / 10000) % 10);
            match op % 100 {
                1 => {
                    let a = self.addr(3, m3);
                    self.set(a, self.param(1, m1) + self.param(2, m2));
                    self.pos += 4;
                }
                2 => {
                    let a = self.addr(3, m3);
                    self.set(a, self.param(1, m1) * self.param(2, m2));
                    self.pos += 4;
                }
                3 => {
                    if self.inp_idx >= self.inp.len() {
                        return None;
                    }
                    let a = self.addr(1, m1);
                    self.set(a, self.inp[self.inp_idx]);
                    self.inp_idx += 1;
                    self.pos += 2;
                }
                4 => {
                    let v = self.param(1, m1);
                    self.pos += 2;
                    return Some(v);
                }
                5 => {
                    if self.param(1, m1) != 0 {
                        self.pos = self.param(2, m2) as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                6 => {
                    if self.param(1, m1) == 0 {
                        self.pos = self.param(2, m2) as usize;
                    } else {
                        self.pos += 3;
                    }
                }
                7 => {
                    let a = self.addr(3, m3);
                    self.set(
                        a,
                        if self.param(1, m1) < self.param(2, m2) {
                            1
                        } else {
                            0
                        },
                    );
                    self.pos += 4;
                }
                8 => {
                    let a = self.addr(3, m3);
                    self.set(
                        a,
                        if self.param(1, m1) == self.param(2, m2) {
                            1
                        } else {
                            0
                        },
                    );
                    self.pos += 4;
                }
                9 => {
                    self.base += self.param(1, m1);
                    self.pos += 2;
                }
                99 => {
                    self.halted = true;
                    return None;
                }
                _ => panic!(),
            }
        }
        None
    }
}

fn part1(prog: &Vec<i64>) -> i64 {
    let mut cpus: Vec<Cpu> = (0..50).map(|_| Cpu::new(prog)).collect();
    let mut queues: Vec<Vec<(i64, i64)>> = vec![Vec::new(); 50];
    let mut outputs: Vec<Vec<i64>> = vec![Vec::new(); 50];

    for i in 0..50 {
        cpus[i].input(i as i64);
    }

    loop {
        for i in 0..50 {
            if let Some((x, y)) = queues[i].first().copied() {
                queues[i].remove(0);
                cpus[i].input(x);
                cpus[i].input(y);
            } else {
                cpus[i].input(-1);
            }

            while let Some(v) = cpus[i].run() {
                outputs[i].push(v);
                if outputs[i].len() == 3 {
                    let dest = outputs[i][0];
                    let x = outputs[i][1];
                    let y = outputs[i][2];
                    outputs[i].clear();

                    if dest == 255 {
                        return y;
                    }

                    if dest >= 0 && dest < 50 {
                        queues[dest as usize].push((x, y));
                    }
                }
            }
        }
    }
}

fn part2(prog: &Vec<i64>) -> i64 {
    let mut cpus: Vec<Cpu> = (0..50).map(|_| Cpu::new(prog)).collect();
    let mut queues: Vec<Vec<(i64, i64)>> = vec![Vec::new(); 50];
    let mut outputs: Vec<Vec<i64>> = vec![Vec::new(); 50];

    for i in 0..50 {
        cpus[i].input(i as i64);
    }

    let mut nat: Option<(i64, i64)> = None;
    let mut last_nat_y: Option<i64> = None;

    loop {
        let mut any_activity = false;

        for i in 0..50 {
            if let Some((x, y)) = queues[i].first().copied() {
                queues[i].remove(0);
                cpus[i].input(x);
                cpus[i].input(y);
                any_activity = true;
            } else {
                cpus[i].input(-1);
            }

            while let Some(v) = cpus[i].run() {
                outputs[i].push(v);
                if outputs[i].len() == 3 {
                    let dest = outputs[i][0];
                    let x = outputs[i][1];
                    let y = outputs[i][2];
                    outputs[i].clear();
                    any_activity = true;

                    if dest == 255 {
                        nat = Some((x, y));
                    } else if dest >= 0 && dest < 50 {
                        queues[dest as usize].push((x, y));
                    }
                }
            }
        }

        let all_queues_empty = queues.iter().all(|q| q.is_empty());
        if !any_activity && all_queues_empty {
            if let Some((x, y)) = nat {
                if last_nat_y == Some(y) {
                    return y;
                }
                last_nat_y = Some(y);
                queues[0].push((x, y));
            }
        }
    }
}
