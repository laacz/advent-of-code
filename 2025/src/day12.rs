use std::collections::HashSet;

fn main() {
    let input = std::fs::read_to_string("data/12.txt").expect("Failed to read input file");
    let input = parse(&input);
    println!("part1: {}", part1(&input));
}

#[derive(Debug, Clone)]
struct Region {
    w: usize,
    h: usize,
    req: Vec<usize>,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
struct Shape([[bool; 3]; 3]);

impl Shape {
    fn cell_count(&self) -> usize {
        self.0.iter().flatten().filter(|&&c| c).count()
    }

    fn rotate(&self) -> Shape {
        let mut rotated = [[false; 3]; 3];
        for y in 0..3 {
            for x in 0..3 {
                rotated[x][2 - y] = self.0[y][x];
            }
        }
        Shape(rotated)
    }

    fn flip(&self) -> Shape {
        let mut flipped = [[false; 3]; 3];
        for y in 0..3 {
            for x in 0..3 {
                flipped[y][2 - x] = self.0[y][x];
            }
        }
        Shape(flipped)
    }

    fn all_orientations(&self) -> Vec<Shape> {
        let mut orientations = Vec::new();
        let mut seen: HashSet<Shape> = HashSet::new();

        let mut current = *self;
        for _ in 0..4 {
            if seen.insert(current) {
                orientations.push(current);
            }
            let flipped = current.flip();
            if seen.insert(flipped) {
                orientations.push(flipped);
            }
            current = current.rotate();
        }
        orientations
    }

    #[allow(dead_code)]
    fn to_str(&self) -> String {
        let mut s = String::new();
        for row in &self.0 {
            for &cell in row {
                s.push(if cell { '#' } else { '.' });
            }
            s.push('\n');
        }
        s
    }
}

#[derive(Debug)]
struct Input {
    shapes: Vec<Shape>,
    regions: Vec<Region>,
}

fn parse(input: &str) -> Input {
    let parts = input.trim().split("\n\n");

    let mut ret = Input {
        shapes: Vec::new(),
        regions: Vec::new(),
    };

    for p in parts.map(|p| p.trim()) {
        if !p.contains('x') {
            let mut cells = [[false; 3]; 3];
            for (y, line) in p.lines().skip(1).enumerate() {
                for (x, ch) in line.trim().chars().enumerate() {
                    if x < 3 && y < 3 {
                        cells[y][x] = ch == '#';
                    }
                }
            }
            ret.shapes.push(Shape(cells));
        } else {
            for line in p.trim().lines() {
                let mut parts = line.split(':');
                let size_part = parts.next().unwrap().trim();
                let required_part = parts.next().unwrap().trim();

                let mut size_parts = size_part.split('x');
                let width: usize = size_parts.next().unwrap().trim().parse().unwrap();
                let height: usize = size_parts.next().unwrap().trim().parse().unwrap();

                let required: Vec<usize> = required_part
                    .split_whitespace()
                    .map(|n| n.parse().unwrap())
                    .collect();

                ret.regions.push(Region {
                    w: width,
                    h: height,
                    req: required,
                });
            }
        }
    }

    ret
}

#[derive(Debug, Clone)]

struct Grid {
    cells: Vec<Vec<Option<usize>>>, // Option<shape_idx>
    w: usize,
    h: usize,
}

impl Grid {

    fn new(w: usize, h: usize) -> Self {
        Grid {
            cells: vec![vec![None; w]; h],
            w,
            h,
        }
    }


    fn can_place(&self, shape: &Shape, sx: usize, sy: usize) -> bool {
        if sy + 3 > self.h || sx + 3 > self.w {
            return false;
        }
        for y in 0..3 {
            for x in 0..3 {
                if shape.0[y][x] && self.cells[sy + y][sx + x].is_some() {
                    return false;
                }
            }
        }
        true
    }


    fn place(&mut self, shape: &Shape, sx: usize, sy: usize, shape_idx: usize) {
        for y in 0..3 {
            for x in 0..3 {
                if shape.0[y][x] {
                    self.cells[sy + y][sx + x] = Some(shape_idx);
                }
            }
        }
    }


    fn remove(&mut self, shape: &Shape, sx: usize, sy: usize) {
        for y in 0..3 {
            for x in 0..3 {
                if shape.0[y][x] {
                    self.cells[sy + y][sx + x] = None;
                }
            }
        }
    }

    fn print_colored(&self, shape_indices: &[usize], orientation_indices: &[usize], shape_colors: &[&str]) {
        let w = self.w;
        // Top border
        print!("\x1b[90m┌{}┐\x1b[0m\n", "─".repeat(w));
        for row in &self.cells {
            print!("\x1b[90m│\x1b[0m");
            for &cell in row {
                if let Some(idx) = cell {
                    let orientation = orientation_indices[idx];
                    let color = shape_colors[orientation % shape_colors.len()];
                    let number = shape_indices[idx];
                    print!("\x1b[{}m{}\x1b[0m", color, number);
                } else {
                    // Grey middot (·), ANSI 90 is bright black/grey
                    print!("\x1b[90m·\x1b[0m");
                }
            }
            print!("\x1b[90m│\x1b[0m\n");
        }
        // Bottom border
        print!("\x1b[90m└{}┘\x1b[0m\n", "─".repeat(w));
    }
}

fn solve(
    grid: &mut Grid,
    shapes_to_place: &[Vec<Shape>],
    idx: usize,
    min_pos: usize,
    shape_indices: &[usize],
    orientation_indices: &mut Vec<usize>,
    shape_colors: &[&str],
) -> bool {
    if idx == shapes_to_place.len() {
        grid.print_colored(shape_indices, orientation_indices, shape_colors);
        return true;
    }

    let orientations = &shapes_to_place[idx];

    for (o_idx, orientation) in orientations.iter().enumerate() {
        for pos in min_pos..grid.w * grid.h {
            let x = pos % grid.w;
            let y = pos / grid.w;
            if x + 3 > grid.w || y + 3 > grid.h {
                continue;
            }
            if grid.can_place(orientation, x, y) {
                grid.place(orientation, x, y, idx);
                orientation_indices[idx] = o_idx;
                if solve(grid, shapes_to_place, idx + 1, pos, shape_indices, orientation_indices, shape_colors) {
                    return true;
                }
                grid.remove(orientation, x, y);
            }
        }
    }
    false
}

fn can_fit(shapes: &[Shape], region: &Region) -> bool {
    let mut shapes_to_place: Vec<Vec<Shape>> = Vec::new();
    let mut shape_indices: Vec<usize> = Vec::new();

    for (shape_idx, &count) in region.req.iter().enumerate() {
        if shape_idx < shapes.len() {
            let orientations = shapes[shape_idx].all_orientations();
            for _ in 0..count {
                shapes_to_place.push(orientations.clone());
                shape_indices.push(shape_idx);
            }
        }
    }

    if shapes_to_place.is_empty() {
        return true;
    }

    let total_cells: usize = shapes_to_place.iter().map(|s| s[0].cell_count()).sum();
    if total_cells > region.w * region.h {
        return false;
    }

    // Sort by cell count descending, but keep shape_indices in sync
    let mut zipped: Vec<_> = shapes_to_place.into_iter().zip(shape_indices.into_iter()).collect();
    zipped.sort_by(|a, b| b.0[0].cell_count().cmp(&a.0[0].cell_count()));
    let (shapes_to_place, shape_indices): (Vec<_>, Vec<_>) = zipped.into_iter().unzip();

    let mut grid = Grid::new(region.w, region.h);
    // Static mapping for up to 8 orientations
    let shape_colors = ["31", "32", "33", "34", "35", "36", "91", "92"];
    let mut orientation_indices = vec![0; shape_indices.len()];
    solve(&mut grid, &shapes_to_place, 0, 0, &shape_indices, &mut orientation_indices, &shape_colors)
}

fn part1(input: &Input) -> usize {
    input
        .regions
        .iter()
        .filter(|region| can_fit(&input.shapes, region))
        .count()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "0:\n\
    ###\n\
    ##.\n\
    ##.\n\
    \n\
    1:\n\
    ###\n\
    ##.\n\
    .##\n\
    \n\
    2:\n\
    .##\n\
    ###\n\
    ##.\n\
    \n\
    3:\n\
    ##.\n\
    ###\n\
    ##.\n\
    \n\
    4:\n\
    ###\n\
    #..\n\
    ###\n\
    \n\
    5:\n\
    ###\n\
    .#.\n\
    ###\n\
    \n\
    4x4: 0 0 0 0 2 0\n\
    12x5: 1 0 1 0 2 2\n\
    12x5: 1 0 1 0 3 2\
    ";

    #[test]
    fn test_part1() {
        assert_eq!(part1(&parse(INPUT)), 2);
    }
}
