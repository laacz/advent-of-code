use std::collections::{HashMap, HashSet};
use std::io::{self, BufRead, Write};

fn main() {
    let content = std::fs::read_to_string("data/25.txt").expect("Failed to read file");
    let prog = parse(&content);

    let args: Vec<String> = std::env::args().collect();
    if args.len() > 1 && args[1] == "-i" {
        play_interactive(&prog);
    } else {
        println!("part1: {}", part1(&prog));
    }
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

#[derive(Clone)]
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

    fn input_str(&mut self, s: &str) {
        for c in s.chars() {
            self.input(c as i64);
        }
        self.input(10);
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

    fn run_to_output(&mut self) -> String {
        let mut output = String::new();
        while let Some(v) = self.run() {
            if v < 128 {
                output.push(v as u8 as char);
            } else {
                output.push_str(&format!("[{}]", v));
            }
        }
        output
    }

    fn send_cmd(&mut self, cmd: &str) -> String {
        self.input_str(cmd);
        self.run_to_output()
    }
}

fn opposite(dir: &str) -> &str {
    match dir {
        "north" => "south",
        "south" => "north",
        "east" => "west",
        "west" => "east",
        _ => panic!("Unknown direction"),
    }
}

fn parse_room(output: &str) -> (String, Vec<String>, Vec<String>) {
    let mut name = String::new();
    let mut doors = Vec::new();
    let mut items = Vec::new();
    let mut in_doors = false;
    let mut in_items = false;

    for line in output.lines() {
        if line.starts_with("== ") && line.ends_with(" ==") {
            name = line[3..line.len() - 3].to_string();
            in_doors = false;
            in_items = false;
        } else if line == "Doors here lead:" {
            in_doors = true;
            in_items = false;
        } else if line == "Items here:" {
            in_doors = false;
            in_items = true;
        } else if line.starts_with("- ") {
            if in_doors {
                doors.push(line[2..].to_string());
            } else if in_items {
                items.push(line[2..].to_string());
            }
        } else if line.is_empty() {
            in_doors = false;
            in_items = false;
        }
    }
    (name, doors, items)
}

const DANGEROUS_ITEMS: &[&str] = &[
    "giant electromagnet",
    "infinite loop",
    "molten lava",
    "photons",
    "escape pod",
];

fn part1(prog: &Vec<i64>) -> String {
    let mut cpu = Cpu::new(prog);
    let output = cpu.run_to_output();

    let mut visited: HashSet<String> = HashSet::new();
    let mut collected_items: Vec<String> = Vec::new();
    let mut checkpoint_path: Vec<String> = Vec::new();
    let mut security_dir = String::new();

    fn explore(
        cpu: &mut Cpu,
        visited: &mut HashSet<String>,
        collected: &mut Vec<String>,
        checkpoint_path: &mut Vec<String>,
        security_dir: &mut String,
        path: &mut Vec<String>,
        output: &str,
    ) {
        let (room_name, doors, items) = parse_room(output);
        if room_name.is_empty() || visited.contains(&room_name) {
            return;
        }
        visited.insert(room_name.clone());

        for item in &items {
            if !DANGEROUS_ITEMS.contains(&item.as_str()) {
                cpu.send_cmd(&format!("take {}", item));
                collected.push(item.clone());
            }
        }

        if room_name == "Security Checkpoint" {
            *checkpoint_path = path.clone();
        }

        for dir in &doors {
            if room_name == "Security Checkpoint" {
                let came_from = path.last().map(|d| opposite(d));
                if came_from != Some(dir.as_str()) {
                    *security_dir = dir.clone();
                }
                continue;
            }

            path.push(dir.clone());
            let new_output = cpu.send_cmd(dir);
            let (new_room, _, _) = parse_room(&new_output);

            if !new_room.is_empty() && !visited.contains(&new_room) {
                explore(
                    cpu,
                    visited,
                    collected,
                    checkpoint_path,
                    security_dir,
                    path,
                    &new_output,
                );
            }

            cpu.send_cmd(opposite(dir));
            path.pop();
        }
    }

    let mut path = Vec::new();
    explore(
        &mut cpu,
        &mut visited,
        &mut collected_items,
        &mut checkpoint_path,
        &mut security_dir,
        &mut path,
        &output,
    );

    for dir in &checkpoint_path {
        cpu.send_cmd(dir);
    }

    let n = collected_items.len();
    for mask in 1..(1 << n) {
        for item in &collected_items {
            cpu.send_cmd(&format!("drop {}", item));
        }

        let mut holding = Vec::new();
        for (i, item) in collected_items.iter().enumerate() {
            if mask & (1 << i) != 0 {
                cpu.send_cmd(&format!("take {}", item));
                holding.push(item.as_str());
            }
        }

        let result = cpu.send_cmd(&security_dir);

        let failed = result.contains("ejected")
            || result.contains("lighter")
            || result.contains("heavier")
            || result.contains("Alert");

        if !failed {
            let mut digits = String::new();
            for c in result.chars() {
                if c.is_ascii_digit() {
                    digits.push(c);
                } else if !digits.is_empty() {
                    if digits.len() >= 6 {
                        return format!("{} (items: {})", digits, holding.join(", "));
                    }
                    digits.clear();
                }
            }
            if digits.len() >= 6 {
                return format!("{} (items: {})", digits, holding.join(", "));
            }
            for word in result.split_whitespace() {
                let clean: String = word.chars().filter(|c| c.is_ascii_digit()).collect();
                if clean.len() >= 6 {
                    return format!("{} (items: {})", clean, holding.join(", "));
                }
            }
            return format!(
                "Passed with items: {}\nFull output:\n{}",
                holding.join(", "),
                result
            );
        }
    }

    "No solution found".to_string()
}

fn parse_items(output: &str) -> Vec<String> {
    let mut items = Vec::new();
    let mut in_items = false;
    for line in output.lines() {
        if line == "Items here:" {
            in_items = true;
            continue;
        }
        if in_items {
            if line.starts_with("- ") {
                items.push(line[2..].to_string());
            } else if !line.is_empty() {
                break;
            }
        }
    }
    items
}

fn play_interactive(prog: &Vec<i64>) {
    let mut cpu = Cpu::new(prog);
    let mut last_safe_state: Option<Cpu> = None;

    let output = cpu.run_to_output();
    let mut room_items = parse_items(&output);

    print!("{}", output);
    io::stdout().flush().unwrap();

    let stdin = io::stdin();

    loop {
        let saved_state = cpu.clone();

        print!("> ");
        io::stdout().flush().unwrap();

        let mut line = String::new();
        if stdin.lock().read_line(&mut line).unwrap() == 0 {
            break;
        }
        let line = line.trim();

        if line == "quit" || line == "q" {
            break;
        }

        let cmd = match line {
            "n" => "north".to_string(),
            "s" => "south".to_string(),
            "e" => "east".to_string(),
            "w" => "west".to_string(),
            "i" => "inv".to_string(),
            _ if line.starts_with("t ") => {
                let prefix = &line[2..];
                if let Some(item) = room_items.iter().find(|i| i.starts_with(prefix)) {
                    format!("take {}", item)
                } else {
                    format!("take {}", prefix)
                }
            }
            _ if line.starts_with("d ") => format!("drop {}", &line[2..]),
            _ => line.to_string(),
        };

        cpu.input_str(&cmd);

        let output = cpu.run_to_output();

        let died = output.contains("you are ejected back to the checkpoint")
            || output.contains("You don't survive")
            || output.contains("you can't move")
            || output.contains("infinite loop")
            || output.contains("You're launched into space")
            || cpu.halted;

        if died && !cpu.halted {
            println!("\n*** EVAPORATED! Restoring to last safe state... ***\n");
            if let Some(ref state) = last_safe_state {
                cpu = state.clone();
                let check_cpu = cpu.clone();
                cpu.input_str("inv");
                let inv_output = cpu.run_to_output();
                cpu = check_cpu;
                println!("Current inventory:{}", inv_output);
            } else {
                cpu = Cpu::new(prog);
                let output = cpu.run_to_output();
                print!("{}", output);
            }
        } else {
            print!("{}", output);
            let new_items = parse_items(&output);
            if !new_items.is_empty() || output.contains("Items here:") {
                room_items = new_items;
            }
            if !cpu.halted {
                last_safe_state = Some(saved_state);
            }
        }

        io::stdout().flush().unwrap();

        if cpu.halted {
            println!("\n*** Game ended ***");
            break;
        }
    }
}
