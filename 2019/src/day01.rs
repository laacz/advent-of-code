fn main() {
    let masses: Vec<i32> = read_file_ints("data/01.txt");
    println!("part1: {}", part1(&masses));
    println!("part2: {}", part2(&masses));
}

fn read_file_ints(filename: &str) -> Vec<i32> {
    std::fs::read_to_string(filename)
        .expect("Failed to read file")
        .lines()
        .map(|line| line.parse().expect("Failed to parse line"))
        .collect()
}

fn part1(masses: &[i32]) -> i32 {
    masses.iter().map(|mass| mass / 3 - 2).sum()
}

fn part2(masses: &[i32]) -> i32 {
    masses.iter().map(|mass| fuel_for_mass(*mass)).sum()
}

fn fuel_for_mass(mass: i32) -> i32 {
    let mut fuel = mass;
    let mut total = 0;
    while fuel > 0 {
        fuel = fuel / 3 - 2;
        if fuel > 0 {
            total += fuel;
        }
    }
    total
}
