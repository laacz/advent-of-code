use std::collections::{HashMap, HashSet, VecDeque};

fn main() {
    let content = std::fs::read_to_string("data/15.txt").expect("Failed to read file");
    let prog = parse(&content);
    println!("part1: {}", part1(&prog));
    println!("part2: {}", part2(&prog));
}

fn parse(content: &str) -> Vec<i64> {
    content.lines().next().unwrap().split(',').map(|s| s.parse().unwrap()).collect()
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
        Self { mem, pos: 0, base: 0, inp: Vec::new(), inp_idx: 0, halted: false }
    }

    fn clone(&self) -> Self {
        Self {
            mem: self.mem.clone(),
            pos: self.pos,
            base: self.base,
            inp: self.inp.clone(),
            inp_idx: self.inp_idx,
            halted: self.halted,
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
                    self.set(a, if self.param(1, m1) < self.param(2, m2) { 1 } else { 0 });
                    self.pos += 4;
                }
                8 => {
                    let a = self.addr(3, m3);
                    self.set(a, if self.param(1, m1) == self.param(2, m2) { 1 } else { 0 });
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

fn explore(prog: &Vec<i64>) -> (HashMap<(i32, i32), i32>, (i32, i32)) {
    let dirs: [(i32, i32); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];
    let mut map: HashMap<(i32, i32), i32> = HashMap::new();
    let mut q: VecDeque<((i32, i32), Cpu)> = VecDeque::new();
    let mut oxy = (0, 0);

    q.push_back(((0, 0), Cpu::new(prog)));
    map.insert((0, 0), 0);

    while let Some(((x, y), cpu)) = q.pop_front() {
        for (cmd, (dx, dy)) in dirs.iter().enumerate() {
            let nx = x + dx;
            let ny = y + dy;
            if map.contains_key(&(nx, ny)) {
                continue;
            }

            let mut c = cpu.clone();
            c.input((cmd + 1) as i64);
            let status = c.run().unwrap();

            match status {
                0 => {
                    map.insert((nx, ny), -1);
                }
                1 | 2 => {
                    map.insert((nx, ny), map[&(x, y)] + 1);
                    q.push_back(((nx, ny), c));
                    if status == 2 {
                        oxy = (nx, ny);
                    }
                }
                _ => {}
            }
        }
    }

    (map, oxy)
}

fn part1(prog: &Vec<i64>) -> i32 {
    let (map, oxy) = explore(prog);
    print_map(&map, oxy);
    map[&oxy]
}

fn part2(prog: &Vec<i64>) -> i32 {
    let dirs: [(i32, i32); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];
    let (map, oxy) = explore(prog);

    let mut seen: HashSet<(i32, i32)> = HashSet::new();
    let mut q: VecDeque<((i32, i32), i32)> = VecDeque::new();
    q.push_back((oxy, 0));
    seen.insert(oxy);
    let mut max_t = 0;

    while let Some(((x, y), t)) = q.pop_front() {
        max_t = max_t.max(t);
        for (dx, dy) in &dirs {
            let nx = x + dx;
            let ny = y + dy;
            if seen.contains(&(nx, ny)) {
                continue;
            }
            if let Some(&v) = map.get(&(nx, ny)) {
                if v >= 0 {
                    seen.insert((nx, ny));
                    q.push_back(((nx, ny), t + 1));
                }
            }
        }
    }

    max_t
}

fn print_map(map: &HashMap<(i32, i32), i32>, oxy: (i32, i32)) {
    let min_x = map.keys().map(|(x, _)| *x).min().unwrap();
    let max_x = map.keys().map(|(x, _)| *x).max().unwrap();
    let min_y = map.keys().map(|(_, y)| *y).min().unwrap();
    let max_y = map.keys().map(|(_, y)| *y).max().unwrap();

    for y in min_y..=max_y {
        for x in min_x..=max_x {
            if (x, y) == (0, 0) {
                print!("\x1b[1;92mS\x1b[0m");
            } else if (x, y) == oxy {
                print!("\x1b[1;96mO\x1b[0m");
            } else if let Some(&v) = map.get(&(x, y)) {
                if v < 0 {
                    print!("\x1b[90m█\x1b[0m");
                } else {
                    print!("\x1b[90m·\x1b[0m");
                }
            } else {
                print!(" ");
            }
        }
        println!();
    }
}
