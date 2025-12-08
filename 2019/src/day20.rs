use std::collections::{HashMap, HashSet, VecDeque};

fn main() {
    let content = std::fs::read_to_string("data/20.txt").expect("Failed to read file");
    let maze = Maze::parse(&content);
    println!("part1: {}", part1(&maze));
    println!("part2: {}", part2(&maze));
}

struct Maze {
    grid: Vec<Vec<char>>,
    start: (usize, usize),
    end: (usize, usize),
    portals: HashMap<(usize, usize), (String, bool)>,
    links: HashMap<(usize, usize), (usize, usize)>,
}

impl Maze {
    fn parse(content: &str) -> Self {
        let grid: Vec<Vec<char>> = content.lines().map(|l| l.chars().collect()).collect();
        let height = grid.len();
        let width = grid.iter().map(|r| r.len()).max().unwrap_or(0);

        let get = |x: usize, y: usize| -> char {
            if y < grid.len() && x < grid[y].len() {
                grid[y][x]
            } else {
                ' '
            }
        };

        let is_outer =
            |x: usize, y: usize| -> bool { x <= 2 || y <= 2 || x >= width - 3 || y >= height - 3 };

        let mut portals: HashMap<(usize, usize), (String, bool)> = HashMap::new();
        let mut start = (0, 0);
        let mut end = (0, 0);

        for y in 0..height {
            for x in 0..width {
                let c = get(x, y);
                if !c.is_ascii_uppercase() {
                    continue;
                }

                if get(x + 1, y).is_ascii_uppercase() {
                    let label = format!("{}{}", c, get(x + 1, y));
                    let pos = if get(x + 2, y) == '.' {
                        Some((x + 2, y))
                    } else if x > 0 && get(x - 1, y) == '.' {
                        Some((x - 1, y))
                    } else {
                        None
                    };
                    if let Some(pos) = pos {
                        portals.insert(pos, (label.clone(), is_outer(pos.0, pos.1)));
                        if label == "AA" {
                            start = pos;
                        }
                        if label == "ZZ" {
                            end = pos;
                        }
                    }
                }

                if get(x, y + 1).is_ascii_uppercase() {
                    let label = format!("{}{}", c, get(x, y + 1));
                    let pos = if get(x, y + 2) == '.' {
                        Some((x, y + 2))
                    } else if y > 0 && get(x, y - 1) == '.' {
                        Some((x, y - 1))
                    } else {
                        None
                    };
                    if let Some(pos) = pos {
                        portals.insert(pos, (label.clone(), is_outer(pos.0, pos.1)));
                        if label == "AA" {
                            start = pos;
                        }
                        if label == "ZZ" {
                            end = pos;
                        }
                    }
                }
            }
        }

        let mut by_label: HashMap<String, Vec<(usize, usize)>> = HashMap::new();
        for (pos, (label, _)) in &portals {
            if label != "AA" && label != "ZZ" {
                by_label.entry(label.clone()).or_default().push(*pos);
            }
        }

        let mut links: HashMap<(usize, usize), (usize, usize)> = HashMap::new();
        for positions in by_label.values() {
            if positions.len() == 2 {
                links.insert(positions[0], positions[1]);
                links.insert(positions[1], positions[0]);
            }
        }

        Maze {
            grid,
            start,
            end,
            portals,
            links,
        }
    }

    fn get(&self, x: usize, y: usize) -> char {
        if y < self.grid.len() && x < self.grid[y].len() {
            self.grid[y][x]
        } else {
            ' '
        }
    }
}

fn part1(maze: &Maze) -> i32 {
    let dirs: [(i32, i32); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];
    let mut visited: HashSet<(usize, usize)> = HashSet::new();
    let mut q: VecDeque<((usize, usize), i32)> = VecDeque::new();

    q.push_back((maze.start, 0));
    visited.insert(maze.start);

    while let Some((pos, steps)) = q.pop_front() {
        if pos == maze.end {
            return steps;
        }

        let (x, y) = pos;

        for (dx, dy) in &dirs {
            let nx = x as i32 + dx;
            let ny = y as i32 + dy;
            if nx < 0 || ny < 0 {
                continue;
            }
            let (nx, ny) = (nx as usize, ny as usize);
            if maze.get(nx, ny) == '.' && !visited.contains(&(nx, ny)) {
                visited.insert((nx, ny));
                q.push_back(((nx, ny), steps + 1));
            }
        }

        if let Some(&dest) = maze.links.get(&pos) {
            if !visited.contains(&dest) {
                visited.insert(dest);
                q.push_back((dest, steps + 1));
            }
        }
    }

    -1
}

fn part2(maze: &Maze) -> i32 {
    let dirs: [(i32, i32); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];
    let mut visited: HashSet<(usize, usize, i32)> = HashSet::new();
    let mut q: VecDeque<((usize, usize, i32), i32)> = VecDeque::new();

    q.push_back(((maze.start.0, maze.start.1, 0), 0));
    visited.insert((maze.start.0, maze.start.1, 0));

    while let Some(((x, y, level), steps)) = q.pop_front() {
        if (x, y) == maze.end && level == 0 {
            return steps;
        }

        if level > 50 {
            continue;
        }

        for (dx, dy) in &dirs {
            let nx = x as i32 + dx;
            let ny = y as i32 + dy;
            if nx < 0 || ny < 0 {
                continue;
            }
            let (nx, ny) = (nx as usize, ny as usize);
            if maze.get(nx, ny) == '.' && !visited.contains(&(nx, ny, level)) {
                visited.insert((nx, ny, level));
                q.push_back(((nx, ny, level), steps + 1));
            }
        }

        if let Some(&dest) = maze.links.get(&(x, y)) {
            if let Some((_, is_outer)) = maze.portals.get(&(x, y)) {
                let new_level = if *is_outer { level - 1 } else { level + 1 };
                if new_level >= 0 && !visited.contains(&(dest.0, dest.1, new_level)) {
                    visited.insert((dest.0, dest.1, new_level));
                    q.push_back(((dest.0, dest.1, new_level), steps + 1));
                }
            }
        }
    }

    -1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example1() {
        let input = "         A
         A
  #######.#########
  #######.........#
  #######.#######.#
  #######.#######.#
  #######.#######.#
  #####  B    ###.#
BC...##  C    ###.#
  ##.##       ###.#
  ##...DE  F  ###.#
  #####    G  ###.#
  #########.#####.#
DE..#######...###.#
  #.#########.###.#
FG..#########.....#
  ###########.#####
             Z
             Z       ";
        assert_eq!(part1(&Maze::parse(input)), 23);
    }

    #[test]
    fn test_example2() {
        let input = "                   A
                   A
  #################.#############
  #.#...#...................#.#.#
  #.#.#.###.###.###.#########.#.#
  #.#.#.......#...#.....#.#.#...#
  #.#########.###.#####.#.#.###.#
  #.............#.#.....#.......#
  ###.###########.###.#####.#.#.#
  #.....#        A   C    #.#.#.#
  #######        S   P    #####.#
  #.#...#                 #......VT
  #.#.#.#                 #.#####
  #...#.#               YN....#.#
  #.###.#                 #####.#
DI....#.#                 #.....#
  #####.#                 #.###.#
ZZ......#               QG....#..AS
  ###.###                 #######
JO..#.#.#                 #.....#
  #.#.#.#                 ###.#.#
  #...#..DI             BU....#..LF
  #####.#                 #.#####
YN......#               VT..#....QG
  #.###.#                 #.###.#
  #.#...#                 #.....#
  ###.###    J L     J    #.#.###
  #.....#    O F     P    #.#...#
  #.###.#####.#.#####.#####.###.#
  #...#.#.#...#.....#.....#.#...#
  #.#####.###.###.#.#.#########.#
  #...#.#.....#...#.#.#.#.....#.#
  #.###.#####.###.###.#.#.#######
  #.#.........#...#.............#
  #########.###.###.#############
           B   J   C
           U   P   P               ";
        assert_eq!(part1(&Maze::parse(input)), 58);
    }

    #[test]
    fn test_part2_example() {
        let input = "             Z L X W       C
             Z P Q B       K
  ###########.#.#.#.#######.###############
  #...#.......#.#.......#.#.......#.#.#...#
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###
  #.#...#.#.#...#.#.#...#...#...#.#.......#
  #.###.#######.###.###.#.###.###.#.#######
  #...#.......#.#...#...#.............#...#
  #.#########.#######.#.#######.#######.###
  #...#.#    F       R I       Z    #.#.#.#
  #.###.#    D       E C       H    #.#.#.#
  #.#...#                           #...#.#
  #.###.#                           #.###.#
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#
CJ......#                           #.....#
  #######                           #######
  #.#....CK                         #......IC
  #.###.#                           #.###.#
  #.....#                           #...#.#
  ###.###                           #.#.#.#
XF....#.#                         RF..#.#.#
  #####.#                           #######
  #......CJ                       NM..#...#
  ###.#.#                           #.###.#
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#
  #.....#        F   Q       P      #.#.#.#
  ###.###########.###.#######.#########.###
  #.....#...#.....#.......#...#.....#.#...#
  #####.#.###.#######.#######.###.###.#.#.#
  #.......#.......#.#.#.#.#...#...#...#.#.#
  #####.###.#####.#.#.#.#.###.###.#.###.###
  #.......#.....#.#...#...............#...#
  #############.#.#.###.###################
               A O F   N
               A A D   M                     ";
        assert_eq!(part2(&Maze::parse(input)), 396);
    }
}
