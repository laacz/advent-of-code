use std::usize;

fn main() {
    let input = std::fs::read_to_string("data/09.txt").expect("Failed to read input file");
    let input = parse(&input);
    println!("part1: {}", part1(&input));
    println!("part2: {}", part2(&input));
}

fn parse(input: &str) -> Vec<(usize, usize)> {
    input
        .trim()
        .lines()
        .map(|line| {
            let parts: Vec<usize> = line
                .split(',')
                .map(|s| s.parse().expect("Failed to parse number"))
                .collect();
            (parts[0], parts[1])
        })
        .collect()
}

fn part1(input: &[(usize, usize)]) -> usize {
    let mut ret: usize = 0;
    let mut cnt = 0;

    for &(x1, y1) in input {
        for &(x2, y2) in input {
            ret = ret.max((x1.max(x2) - x1.min(x2) + 1) * (y1.max(y2) - y1.min(y2) + 1));
            cnt+=1;
        }
    }

    println!("{}", cnt);

    ret
}


fn print_map(grid: &Vec<Vec<usize>>) {
    for y in 0..grid[0].len() {
        for x in 0..grid.len() {
            match grid[x][y] {
                0 => print!("·"),
                1 => print!("O"),
                2 => print!("X"),
                _ => print!(" "),
            }
        }
        println!();
    }
}

fn flood_fill(grid: &Vec<Vec<usize>>) -> Vec<Vec<usize>> {
    let mut new_grid = grid.clone();
    let maxx = grid.len();
    let maxy = grid[0].len();

    let mut stack = Vec::new();
    
    for x in 0..maxx {
        if new_grid[x][0] == 0 {
            stack.push((x, 0));
            new_grid[x][0] = 3;
        }
        if new_grid[x][maxy - 1] == 0 {
            stack.push((x, maxy - 1));
            new_grid[x][maxy - 1] = 3;
        }
    }
    for y in 0..maxy {
        if new_grid[0][y] == 0 {
            stack.push((0, y));
            new_grid[0][y] = 3;
        }
        if new_grid[maxx - 1][y] == 0 {
            stack.push((maxx - 1, y));
            new_grid[maxx - 1][y] = 3;
        }
    }

    while let Some((x, y)) = stack.pop() {
        if x > 0 && new_grid[x - 1][y] == 0 {
            new_grid[x - 1][y] = 3;
            stack.push((x - 1, y));
        }
        if x + 1 < maxx && new_grid[x + 1][y] == 0 {
            new_grid[x + 1][y] = 3;
            stack.push((x + 1, y));
        }
        if y > 0 && new_grid[x][y - 1] == 0 {
            new_grid[x][y - 1] = 3;
            stack.push((x, y - 1));
        }
        if y + 1 < maxy && new_grid[x][y + 1] == 0 {
            new_grid[x][y + 1] = 3;
            stack.push((x, y + 1));
        }
    }

    new_grid
}

fn part2(input: &[(usize, usize)]) -> usize {
    let mut xs = input.iter().map(|&(x, _)| x).collect::<Vec<_>>();
    xs.sort();
    xs.dedup();
    let mut ys = input.iter().map(|&(_, y)| y).collect::<Vec<_>>();
    ys.sort();
    ys.dedup();

    let mut grid = vec![vec![0usize; ys.len()]; xs.len()];

    for i in 0..input.len() {
        let (x1, y1) = input[i];
        let (x2, y2) = input[(i + 1) % input.len()];
        
        let grid_x1 = xs.binary_search(&x1).unwrap();
        let grid_y1 = ys.binary_search(&y1).unwrap();
        let grid_x2 = xs.binary_search(&x2).unwrap();
        let grid_y2 = ys.binary_search(&y2).unwrap();

        grid[grid_x1][grid_y1] = 1;
        if grid_x1 == grid_x2 {
            let (start, end) = if grid_y1 < grid_y2 { (grid_y1, grid_y2) } else { (grid_y2, grid_y1) };
            for y in start+1..end {
                grid[grid_x1][y] = 2;
            }
            grid[grid_x1][grid_y2] = 1;
        } else if grid_y1 == grid_y2 {
            let (start, end) = if grid_x1 < grid_x2 { (grid_x1, grid_x2) } else { (grid_x2, grid_x1) };
            for x in start+1..end {
                grid[x][grid_y1] = 2;
            }
            grid[grid_x2][grid_y1] = 1;
        }
    }

    grid = flood_fill(&grid);
    print_map(&grid);

    let mut ret = 0;

    for &(x1, y1) in input {
        for &(x2, y2) in input {
            let grid_x1 = xs.binary_search(&x1).unwrap();
            let grid_y1 = ys.binary_search(&y1).unwrap();
            let grid_x2 = xs.binary_search(&x2).unwrap();
            let grid_y2 = ys.binary_search(&y2).unwrap();

            let mut ok = true;
            for x in grid_x1.min(grid_x2)..=grid_x1.max(grid_x2) {
                for y in grid_y1.min(grid_y2)..=grid_y1.max(grid_y2) {
                    if grid[x][y] == 3 {
                        ok = false;
                        break;
                    }
                }
                
                if !ok {
                    break;
                }
            }
            if ok {
                ret = ret.max((x1.max(x2) - x1.min(x2) + 1) * (y1.max(y2) - y1.min(y2) + 1));
            }
        }
    }

    ret
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "\
    7,1\n\
    11,1\n\
    11,7\n\
    9,7\n\
    9,5\n\
    2,5\n\
    2,3\n\
    7,3\
    ";

    #[test]
    fn test_part1() {
        let input = parse(INPUT);
        assert_eq!(part1(&input), 50);
    }

    #[test]
    fn test_part2() {
        let input = parse(INPUT);
        assert_eq!(part2(&input), 24);
    }
}
