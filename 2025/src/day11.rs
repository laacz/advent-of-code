use std::collections::{HashMap, HashSet};

fn main() {
    let input = std::fs::read_to_string("data/11.txt").expect("Failed to read input file");
    let input = parse(&input);
    println!("part1: {}", part1(&input));
    println!("part2: {}", part2(&input));
}

type Graph = HashMap<String, Vec<String>>;

fn parse(input: &str) -> Graph {
    input
        .lines()
        .map(|l| l.trim())
        .filter(|l| !l.is_empty())
        .map(|line| {
            let mut parts = line.split(": ");
            let src = parts.next().unwrap().to_string();
            let dsts = parts
                .next()
                .unwrap()
                .split_whitespace()
                .map(|s| s.to_string())
                .collect();
            (src, dsts)
        })
        .collect()
}

fn all_paths(graph: &Graph, start: &str, target: &str) -> Vec<Vec<String>> {
    fn dfs(
        graph: &Graph,
        current: &str,
        target: &str,
        visited: &mut HashSet<String>,
        path: &mut Vec<String>,
        paths: &mut Vec<Vec<String>>,
    ) {
        path.push(current.to_string());

        if current == target {
            paths.push(path.clone());
        } else {
            visited.insert(current.to_string());
            if let Some(neighbors) = graph.get(current) {
                for neighbor in neighbors {
                    if !visited.contains(neighbor) {
                        dfs(graph, neighbor, target, visited, path, paths);
                    }
                }
            }
            visited.remove(current);
        }

        path.pop();
    }

    let mut visited = HashSet::new();
    let mut path = Vec::new();
    let mut paths = Vec::new();
    dfs(graph, start, target, &mut visited, &mut path, &mut paths);
    paths
}

fn count_paths_through(
    graph: &Graph,
    start: &str,
    target: &str,
    dac: &str,
    fft: &str,
) -> usize {
    fn dfs(
        graph: &Graph,
        current: &str,
        target: &str,
        dac: &str,
        fft: &str,
        seen_dac: bool,
        seen_fft: bool,
        visited: &mut HashSet<String>,
        memo: &mut HashMap<(String, bool, bool), usize>,
    ) -> usize {
        let seen_dac = seen_dac || current == dac;
        let seen_fft = seen_fft || current == fft;

        if current == target {
            return if seen_dac && seen_fft { 1 } else { 0 };
        }

        let key = (current.to_string(), seen_dac, seen_fft);
        if let Some(&cached) = memo.get(&key) {
            if !visited.contains(current) {
                return cached;
            }
        }

        visited.insert(current.to_string());
        let mut count = 0;

        if let Some(neighbors) = graph.get(current) {
            for neighbor in neighbors {
                if !visited.contains(neighbor) {
                    count += dfs(graph, neighbor, target, dac, fft, seen_dac, seen_fft, visited, memo);
                }
            }
        }

        visited.remove(current);
        memo.insert(key, count);
        count
    }

    let mut visited = HashSet::new();
    let mut memo = HashMap::new();
    dfs(graph, start, target, dac, fft, false, false, &mut visited, &mut memo)
}

fn part1(graph: &Graph) -> usize {
    all_paths(graph, "you", "out").len()
}

fn part2(graph: &Graph) -> usize {
    count_paths_through(graph, "svr", "out", "dac", "fft")
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = parse("\
        aaa: you hhh\n\
        you: bbb ccc\n\
        bbb: ddd eee\n\
        ccc: ddd eee fff\n\
        ddd: ggg\n\
        eee: out\n\
        fff: out\n\
        ggg: out\n\
        hhh: ccc fff iii\n\
        iii: out\
        ");
        assert_eq!(part1(&input), 5); 
    }

    #[test]
    fn test_part2() {
        let input = parse("\
        svr: aaa bbb\n\
        aaa: fft\n\
        fft: ccc\n\
        bbb: tty\n\
        tty: ccc\n\
        ccc: ddd eee\n\
        ddd: hub\n\
        hub: fff\n\
        eee: dac\n\
        dac: fff\n\
        fff: ggg hhh\n\
        ggg: out\n\
        hhh: out\
        ");
        assert_eq!(part2(&input), 2);
    }
}
