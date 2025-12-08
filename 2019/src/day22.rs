use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/22.txt").expect("Failed to read file");
    let input = parse(&content);
    println!("part1: {}", part1(&input));
    println!("part2: {}", part2(&input));
}

fn parse(content: &str) -> Vec<(usize, isize)> {
    content
        .trim()
        .lines()
        .map(|line| {
            let parts: Vec<&str> = line.split_whitespace().collect();
            match parts[0] {
                "deal" if parts[1] == "into" => (0, 0),
                "deal" if parts[1] == "with" => (2, parts[3].parse::<isize>().unwrap()),
                "cut" => (1, parts[1].parse::<isize>().unwrap()), 
                _ => panic!("Unknown instruction {}", line),
            }
        })
        .collect()
}

struct Stack {
    cards: Vec<usize>,
}

impl Stack {
    fn new(n: usize) -> Self {
        Stack {
            cards: (0..n).collect(),
        }
    }

    fn deal_into_new_stack(&mut self) {
        self.cards.reverse();
    }

    fn cut(&mut self, n: isize) {
        let len = self.cards.len() as isize;
        let n = if n < 0 { len + n } else { n } as usize;
        let mut new_cards = Vec::with_capacity(self.cards.len());
        new_cards.extend_from_slice(&self.cards[n..]);
        new_cards.extend_from_slice(&self.cards[..n]);
        self.cards = new_cards;
    }

    fn deal_with_increment(&mut self, n: usize) {
        let len = self.cards.len();
        let mut new_cards = vec![0; len];
        for (i, &card) in self.cards.iter().enumerate() {
            new_cards[(i * n) % len] = card;
        }
        self.cards = new_cards;
    }
}

fn part1(prog: &Vec<(usize, isize)>) -> i64 {
    let mut stack = Stack::new(10007);

    for &(instr, arg) in prog {
        match instr {
            0 => stack.deal_into_new_stack(),
            1 => stack.cut(arg),
            2 => stack.deal_with_increment(arg as usize),
            _ => panic!("Unknown instruction"),
        }
    }

    stack.cards.iter().position(|&x| x == 2019).unwrap() as i64
}

fn modd(x: i128, n: i128) -> i128 { ((x % n) + n) % n }

fn modpow(mut base: i128, mut exp: i128, n: i128) -> i128 {
    let mut result = 1i128;
    base = modd(base, n);
    while exp > 0 {
        if exp % 2 == 1 { result = modd(result * base, n); }
        exp /= 2;
        base = modd(base * base, n);
    }
    result
}

fn moddiv(a: i128, b: i128, n: i128) -> i128 { modd(a * modpow(b, n - 2, n), n) }

fn part2(prog: &Vec<(usize, isize)>) -> i128 {
    const DECK_SIZE: i128 = 119315717514047;
    const SHUFFLE_COUNT: i128 = 101741582076661;

    // new_position = num_pos * old_position + shift
    // compute num_pos and shift
    let mut num_pos: i128 = 1;
    let mut shift: i128 = 0;
    
    for &(instr, arg) in prog {
        match instr {
            0 => { // deal into new stack: reverse the deck
                num_pos = modd(-num_pos, DECK_SIZE);
                shift = modd(-shift - 1, DECK_SIZE);
            }
            1 => { // cut N: move top N cards to bottom
                shift = modd(shift - arg as i128, DECK_SIZE);
            }
            2 => { // deal with increment N: spread cards out
                num_pos = modd(num_pos * arg as i128, DECK_SIZE);
                shift = modd(shift * arg as i128, DECK_SIZE);
            }
            _ => panic!("Unknown instruction"),
        }
    }

    // simple math old_pos = (new_pos - shift) / num_pos
    let reverse_num_pos = moddiv(1, num_pos, DECK_SIZE);
    let reverse_shift = modd(-shift * reverse_num_pos, DECK_SIZE);

    // this i had to look up
    // - final_num_pos = reverse_num_pos^k
    // - final_shift = reverse_shift * (reverse_num_pos^k - 1) / (reverse_num_pos - 1)
    let final_num_pos = modpow(reverse_num_pos, SHUFFLE_COUNT, DECK_SIZE);
    let final_shift = modd(
        reverse_shift * moddiv(final_num_pos - 1, reverse_num_pos - 1, DECK_SIZE),
        DECK_SIZE
    );

    modd(final_num_pos * 2020 + final_shift, DECK_SIZE)
}

