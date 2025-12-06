use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/17.txt").expect("Failed to read file");
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

fn get_map(prog: &Vec<i64>) -> Vec<Vec<char>> {
    let mut cpu = Cpu::new(prog);
    let mut output = String::new();

    while let Some(v) = cpu.run() {
        output.push(v as u8 as char);
    }

    output
        .lines()
        .filter(|line| !line.is_empty())
        .map(|line| line.chars().collect())
        .collect()
}

fn print_map(map: &Vec<Vec<char>>) {
    let rows = map.len();
    let cols = if rows > 0 { map[0].len() } else { 0 };

    for y in 0..rows {
        for x in 0..cols {
            let c = map[y][x];
            let is_intersection = c == '#'
                && y > 0 && y < rows - 1 && x > 0 && x < cols - 1
                && map[y - 1][x] == '#'
                && map[y + 1][x] == '#'
                && map[y][x - 1] == '#'
                && map[y][x + 1] == '#';

            if is_intersection {
                print!("\x1b[1;96mO\x1b[0m");
            } else {
                match c {
                    '#' => print!("\x1b[90m#\x1b[0m"),
                    '.' => print!(" "),
                    '^' | 'v' | '<' | '>' => print!("\x1b[1;92m{}\x1b[0m", c),
                    _ => print!("{}", c),
                }
            }
        }
        println!();
    }
}

fn part1(prog: &Vec<i64>) -> i32 {
    let map = get_map(prog);
    print_map(&map);
    let rows = map.len();
    let cols = if rows > 0 { map[0].len() } else { 0 };

    let mut sum = 0;

    for y in 1..rows - 1 {
        for x in 1..cols - 1 {
            if map[y][x] == '#'
                && map[y - 1][x] == '#'
                && map[y + 1][x] == '#'
                && map[y][x - 1] == '#'
                && map[y][x + 1] == '#'
            {
                sum += (x as i32) * (y as i32);
            }
        }
    }

    sum
}

fn find_path(map: &Vec<Vec<char>>) -> Vec<String> {
    let rows = map.len() as i32;
    let cols = map[0].len() as i32;

    let mut pos = (0i32, 0i32);
    let mut dir = 0;

    for y in 0..rows {
        for x in 0..cols {
            let c = map[y as usize][x as usize];
            if c == '^' || c == 'v' || c == '<' || c == '>' {
                pos = (x, y);
                dir = match c {
                    '^' => 0,
                    '>' => 1,
                    'v' => 2,
                    '<' => 3,
                    _ => 0,
                };
            }
        }
    }

    let dirs: [(i32, i32); 4] = [(0, -1), (1, 0), (0, 1), (-1, 0)];
    let mut path = Vec::new();

    loop {
        let left = (dir + 3) % 4;
        let right = (dir + 1) % 4;

        let (ldx, ldy) = dirs[left];
        let (rdx, rdy) = dirs[right];

        let lpos = (pos.0 + ldx, pos.1 + ldy);
        let rpos = (pos.0 + rdx, pos.1 + rdy);

        let can_left = lpos.0 >= 0 && lpos.0 < cols && lpos.1 >= 0 && lpos.1 < rows
            && map[lpos.1 as usize][lpos.0 as usize] == '#';
        let can_right = rpos.0 >= 0 && rpos.0 < cols && rpos.1 >= 0 && rpos.1 < rows
            && map[rpos.1 as usize][rpos.0 as usize] == '#';

        if can_left {
            dir = left;
            path.push("L".to_string());
        } else if can_right {
            dir = right;
            path.push("R".to_string());
        } else {
            break;
        }

        let (dx, dy) = dirs[dir];
        let mut steps = 0;
        loop {
            let next = (pos.0 + dx, pos.1 + dy);
            if next.0 >= 0 && next.0 < cols && next.1 >= 0 && next.1 < rows
                && map[next.1 as usize][next.0 as usize] == '#'
            {
                pos = next;
                steps += 1;
            } else {
                break;
            }
        }
        path.push(steps.to_string());
    }

    path
}

fn try_match(path: &[String], a: &[String], b: &[String], c: &[String]) -> Option<Vec<String>> {
    let mut remaining = path;
    let mut main = Vec::new();

    while !remaining.is_empty() {
        if remaining.len() >= a.len() && &remaining[0..a.len()] == a {
            main.push("A".to_string());
            remaining = &remaining[a.len()..];
        } else if remaining.len() >= b.len() && &remaining[0..b.len()] == b {
            main.push("B".to_string());
            remaining = &remaining[b.len()..];
        } else if remaining.len() >= c.len() && &remaining[0..c.len()] == c {
            main.push("C".to_string());
            remaining = &remaining[c.len()..];
        } else {
            return None;
        }
    }

    let main_str = main.join(",");
    if main_str.len() <= 20 {
        Some(main)
    } else {
        None
    }
}

fn compress(path: &[String]) -> Option<(String, String, String, String)> {
    for a_len in 1..=10 {
        if a_len > path.len() {
            break;
        }
        let a: Vec<String> = path[0..a_len].to_vec();
        if a.join(",").len() > 20 {
            continue;
        }

        for b_start in a_len..path.len() {
            let mut idx = 0;
            let mut cur_start = 0;
            while idx < path.len() && idx < b_start {
                if path[idx..].starts_with(&a) {
                    idx += a_len;
                    cur_start = idx;
                } else {
                    break;
                }
            }
            if cur_start != b_start {
                continue;
            }

            for b_len in 1..=10 {
                if b_start + b_len > path.len() {
                    break;
                }
                let b: Vec<String> = path[b_start..b_start + b_len].to_vec();
                if b.join(",").len() > 20 {
                    continue;
                }

                let mut c_start = None;
                let mut idx = 0;
                while idx < path.len() {
                    if path[idx..].starts_with(&a) {
                        idx += a_len;
                    } else if path[idx..].starts_with(&b) {
                        idx += b_len;
                    } else {
                        c_start = Some(idx);
                        break;
                    }
                }

                if c_start.is_none() {
                    if let Some(main) = try_match(path, &a, &b, &[]) {
                        return Some((
                            main.join(","),
                            a.join(","),
                            b.join(","),
                            "".to_string(),
                        ));
                    }
                    continue;
                }

                let c_start = c_start.unwrap();

                for c_len in 1..=10 {
                    if c_start + c_len > path.len() {
                        break;
                    }
                    let c: Vec<String> = path[c_start..c_start + c_len].to_vec();
                    if c.join(",").len() > 20 {
                        continue;
                    }

                    if let Some(main) = try_match(path, &a, &b, &c) {
                        return Some((
                            main.join(","),
                            a.join(","),
                            b.join(","),
                            c.join(","),
                        ));
                    }
                }
            }
        }
    }
    None
}

fn part2(prog: &Vec<i64>) -> i64 {
    let map = get_map(prog);
    let path = find_path(&map);

    let (main, a, b, c) = compress(&path).expect("Could not compress path");

    let mut prog = prog.clone();
    prog[0] = 2;

    let mut cpu = Cpu::new(&prog);

    let input = format!("{}\n{}\n{}\n{}\nn\n", main, a, b, c);

    for ch in input.chars() {
        cpu.input(ch as i64);
    }

    let mut result = 0;
    while let Some(v) = cpu.run() {
        result = v;
    }

    result
}
