use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/06.txt").expect("Failed to read file");
    let r: Vec<(&str, &str)> = parse(&content);
    println!("part1: {}", part1(&r));
    println!("part2: {}", part2(&r));
}

fn parse(content: &str) -> Vec<(&str, &str)> {
    content
        .lines()
        .filter_map(|line| {
            let parts: Vec<&str> = line.split(')').collect();
            Some((parts[1], parts[0]))
        })
        .collect()
}

fn part1(input: &Vec<(&str, &str)>) -> i32 {
    let mut parent_map: HashMap<&str, &str> = HashMap::new();
    for (child, parent) in input.iter() {
        parent_map.insert(child, parent);
    }

    let mut total = 0;
    for (child, _) in input.iter() {
        let mut current = *child;
        while let Some(&parent) = parent_map.get(current) {
            total += 1;
            current = parent;
        }
    }

    total
}

fn part2(input: &Vec<(&str, &str)>) -> i32 {
    let mut parent_map: HashMap<&str, &str> = HashMap::new();
    for (child, parent) in input.iter() {
        parent_map.insert(child, parent);
    }

    let mut you_ancestors: HashMap<&str, i32> = HashMap::new();
    let mut current: &str = *parent_map.get("YOU").unwrap();
    let mut distance = 0;
    while let Some(&parent) = parent_map.get(current) {
        you_ancestors.insert(current, distance);
        distance += 1;
        current = parent;
    }
    you_ancestors.insert(current, distance);

    current = *parent_map.get("SAN").unwrap();
    distance = 0;
    while !you_ancestors.contains_key(current) {
        distance += 1;
        current = *parent_map.get(current).unwrap();
    }

    distance + you_ancestors.get(current).unwrap()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            part1(&parse(
                "\
            COM)B\n\
            B)C\n\
            C)D\n\
            D)E\n\
            E)F\n\
            B)G\n\
            G)H\n\
            D)I\n\
            E)J\n\
            J)K\n\
            K)L"
            )),
            42
        );
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            part2(&parse(
                "\
            COM)B\n\
            B)C\n\
            C)D\n\
            D)E\n\
            E)F\n\
            B)G\n\
            G)H\n\
            D)I\n\
            E)J\n\
            J)K\n\
            K)L\n\
            K)YOU\n\
            I)SAN"
            )),
            4
        );
    }
}
