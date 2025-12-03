use std::collections::HashMap;

fn main() {
    let content = std::fs::read_to_string("data/14.txt").expect("Failed to read file");
    let r = parse(&content);
    println!("part1: {}", part1(&r));
    println!("part2: {}", part2(&r));
}

fn parse(content: &str) -> HashMap<String, (Vec<(String, i64)>, i64)> {
    let mut r = HashMap::new();
    for line in content.lines() {
        let line = line.trim();
        if line.is_empty() {
            continue;
        }
        let parts: Vec<&str> = line.split(" => ").collect();
        let mut inputs = Vec::new();
        for item in parts[0].split(", ") {
            let p: Vec<&str> = item.trim().split(' ').collect();
            let amt: i64 = p[0].parse().unwrap();
            let name = p[1].to_string();
            inputs.push((name, amt));
        }
        let out: Vec<&str> = parts[1].trim().split(' ').collect();
        let out_amt: i64 = out[0].parse().unwrap();
        let out_name = out[1].to_string();
        r.insert(out_name, (inputs, out_amt));
    }
    r
}

fn calc(r: &HashMap<String, (Vec<(String, i64)>, i64)>, fuel: i64) -> i64 {
    let mut need: HashMap<String, i64> = HashMap::new();
    let mut left: HashMap<String, i64> = HashMap::new();
    need.insert("FUEL".to_string(), fuel);

    loop {
        let chem = match need.iter().find(|(k, v)| *k != "ORE" && **v > 0) {
            Some((k, _)) => k.clone(),
            None => break,
        };

        let want = need[&chem];
        let have = *left.get(&chem).unwrap_or(&0);

        if have >= want {
            left.insert(chem.clone(), have - want);
            need.insert(chem, 0);
            continue;
        }

        let still = want - have;
        left.insert(chem.clone(), 0);

        let (inputs, batch) = &r[&chem];
        let times = (still + batch - 1) / batch;
        let made = times * batch;
        *left.entry(chem.clone()).or_insert(0) += made - still;
        need.insert(chem, 0);

        for (name, amt) in inputs {
            *need.entry(name.clone()).or_insert(0) += amt * times;
        }
    }

    *need.get("ORE").unwrap_or(&0)
}

fn part1(r: &HashMap<String, (Vec<(String, i64)>, i64)>) -> i64 {
    calc(r, 1)
}

fn part2(r: &HashMap<String, (Vec<(String, i64)>, i64)>) -> i64 {
    let total: i64 = 1_000_000_000_000;
    let per_fuel = calc(r, 1);
    let mut lo = total / per_fuel;
    let mut hi = lo * 2;

    while calc(r, hi) <= total {
        hi *= 2;
    }

    while lo < hi {
        let mid = (lo + hi + 1) / 2;
        if calc(r, mid) <= total {
            lo = mid;
        } else {
            hi = mid - 1;
        }
    }

    lo
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example1() {
        let input = "\
10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL";
        assert_eq!(part1(&parse(input)), 31);
    }

    #[test]
    fn test_example2() {
        let input = "\
9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL";
        assert_eq!(part1(&parse(input)), 165);
    }

    #[test]
    fn test_example3() {
        let input = "\
157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT";
        assert_eq!(part1(&parse(input)), 13312);
    }

    #[test]
    fn test_example4() {
        let input = "\
2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF";
        assert_eq!(part1(&parse(input)), 180697);
    }

    #[test]
    fn test_example5() {
        let input = "\
171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX";
        assert_eq!(part1(&parse(input)), 2210736);
    }
}
