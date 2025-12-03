fn main() {
    let content = std::fs::read_to_string("data/12.txt").expect("Failed to read file");
    println!("part1: {}", part1(&parse(&content)));
    println!("part2: {}", part2(&parse(&content)));
}

#[derive(Clone, PartialEq)]
struct Vec3 {
    x: i32,
    y: i32,
    z: i32,
}

#[derive(Clone, PartialEq)]
struct Moon {
    pos: Vec3,
    vel: Vec3,
}

fn parse(content: &str) -> Vec<Moon> {
    content
        .lines()
        .filter(|line| !line.trim().is_empty())
        .map(|line| {
            let line = &line[1..line.len() - 1];
            let parts: Vec<&str> = line.split(',').collect();
            let x = parts[0].split('=').nth(1).unwrap().parse().unwrap();
            let y = parts[1].split('=').nth(1).unwrap().parse().unwrap();
            let z = parts[2].split('=').nth(1).unwrap().parse().unwrap();
            Moon {
                pos: Vec3 { x, y, z },
                vel: Vec3 { x: 0, y: 0, z: 0 },
            }
        })
        .collect()
}

fn step(moons: &mut Vec<Moon>) {
    for i in 0..moons.len() {
        for j in (i + 1)..moons.len() {
            if moons[i].pos.x < moons[j].pos.x {
                moons[i].vel.x += 1;
                moons[j].vel.x -= 1;
            } else if moons[i].pos.x > moons[j].pos.x {
                moons[i].vel.x -= 1;
                moons[j].vel.x += 1;
            }

            if moons[i].pos.y < moons[j].pos.y {
                moons[i].vel.y += 1;
                moons[j].vel.y -= 1;
            } else if moons[i].pos.y > moons[j].pos.y {
                moons[i].vel.y -= 1;
                moons[j].vel.y += 1;
            }

            if moons[i].pos.z < moons[j].pos.z {
                moons[i].vel.z += 1;
                moons[j].vel.z -= 1;
            } else if moons[i].pos.z > moons[j].pos.z {
                moons[i].vel.z -= 1;
                moons[j].vel.z += 1;
            }
        }
    }

    for moon in moons.iter_mut() {
        moon.pos.x += moon.vel.x;
        moon.pos.y += moon.vel.y;
        moon.pos.z += moon.vel.z;
    }
}

fn energy(moon: &Moon) -> i32 {
    let potential = moon.pos.x.abs() + moon.pos.y.abs() + moon.pos.z.abs();
    let kinetic = moon.vel.x.abs() + moon.vel.y.abs() + moon.vel.z.abs();
    potential * kinetic
}

fn simulate(moons: &Vec<Moon>, steps: usize) -> Vec<Moon> {
    let mut m = moons.clone();
    for _ in 0..steps {
        step(&mut m);
    }
    m
}

fn part1(moons: &Vec<Moon>) -> i32 {
    let m = simulate(moons, 1000);
    m.iter().map(|moon| energy(moon)).sum()
}

fn find_cycle(positions: &Vec<i32>) -> i64 {
    let initial_pos = positions.clone();
    let initial_vel = vec![0i32; positions.len()];

    let mut pos = positions.clone();
    let mut vel = vec![0i32; positions.len()];
    let mut steps: i64 = 0;

    loop {
        for i in 0..pos.len() {
            for j in (i + 1)..pos.len() {
                if pos[i] < pos[j] {
                    vel[i] += 1;
                    vel[j] -= 1;
                } else if pos[i] > pos[j] {
                    vel[i] -= 1;
                    vel[j] += 1;
                }
            }
        }

        for i in 0..pos.len() {
            pos[i] += vel[i];
        }
        steps += 1;

        if pos == initial_pos && vel == initial_vel {
            return steps;
        }
    }
}

fn part2(moons: &Vec<Moon>) -> i64 {
    let xs: Vec<i32> = moons.iter().map(|m| m.pos.x).collect();
    let ys: Vec<i32> = moons.iter().map(|m| m.pos.y).collect();
    let zs: Vec<i32> = moons.iter().map(|m| m.pos.z).collect();

    let cx = find_cycle(&xs);
    let cy = find_cycle(&ys);
    let cz = find_cycle(&zs);

    lcm(lcm(cx, cy), cz)
}

fn gcd(a: i64, b: i64) -> i64 {
    if b == 0 {
        a
    } else {
        gcd(b, a % b)
    }
}
fn lcm(a: i64, b: i64) -> i64 {
    a / gcd(a, b) * b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1_sample1() {
        let input = "\
            <x=-1, y=0, z=2>\n\
            <x=2, y=-10, z=-7>\n\
            <x=4, y=-8, z=8>\n\
            <x=3, y=5, z=-1>\
            ";
        assert_eq!(
            simulate(&parse(input), 10)
                .iter()
                .map(|m| energy(m))
                .sum::<i32>(),
            179
        );
    }

    #[test]
    fn test_part1_sample2() {
        let input = "\
            <x=-8, y=-10, z=0>\n\
            <x=5, y=5, z=10>\n\
            <x=2, y=-7, z=3>\n\
            <x=9, y=-8, z=-3>\
            ";
        assert_eq!(
            simulate(&parse(input), 100)
                .iter()
                .map(|m| energy(m))
                .sum::<i32>(),
            1940
        );
    }

    #[test]
    fn test_part2_sample1() {
        let input = "\
            <x=-1, y=0, z=2>\n\
            <x=2, y=-10, z=-7>\n\
            <x=4, y=-8, z=8>\n\
            <x=3, y=5, z=-1>\
            ";
        assert_eq!(part2(&parse(input)), 2772);
    }

    #[test]
    fn test_part2_sample2() {
        let input = "\
            <x=-8, y=-10, z=0>\n\
            <x=5, y=5, z=10>\n\
            <x=2, y=-7, z=3>\n\
            <x=9, y=-8, z=-3>\
            ";
        assert_eq!(part2(&parse(input)), 4686774924);
    }
}
