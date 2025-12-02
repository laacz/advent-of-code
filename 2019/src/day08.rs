fn main() {
    let content = std::fs::read_to_string("data/08.txt").expect("Failed to read file");
    let r = parse(&content);
    println!("part1: {}", part1(&r));
    println!("part2:\n{}", part2(&r));
}

fn parse(content: &str) -> Vec<u32> {
    content
        .lines()
        .next()
        .unwrap()
        .trim()
        .chars()
        .map(|c| c.to_digit(10).unwrap())
        .collect()
}

fn part1(input: &Vec<u32>) -> i32 {
    let mut layers: Vec<Vec<u32>> = Vec::new();

    for chunk in input.chunks(25 * 6) {
        layers.push(chunk.to_vec());
    }

    let mut min_zeroes = usize::MAX;
    let mut res = 0;

    for layer in layers.iter() {
        let zeroes = layer.iter().filter(|&&x| x == 0).count();
        if zeroes < min_zeroes {
            min_zeroes = zeroes;
            res = layer.iter().filter(|&&x| x == 1).count()
                * layer.iter().filter(|&&x| x == 2).count();
        }
    }

    res as i32
}

fn part2(input: &Vec<u32>) -> String {
    let mut ret = String::new();
    let mut layers: Vec<Vec<u32>> = Vec::new();

    for chunk in input.chunks(25 * 6) {
        layers.push(chunk.to_vec());
    }

    let mut image = vec![2; 25 * 6];
    for y in 0..6 {
        for x in 0..25 {
            for layer in layers.iter() {
                if image[x + y * 25] == 2 {
                    image[x + y * 25] = layer[x + y * 25];
                }
            }
            if image[x + y * 25] == 1 {
                ret.push('#');
            } else {
                ret.push(' ');
            }
        }
        ret.push('\n');
    }

    ret
}
