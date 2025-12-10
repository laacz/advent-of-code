fn main() {
    let input = std::fs::read_to_string("data/10.txt").expect("Failed to read input file");
    let input = parse(&input);
    println!("part1: {}", part1(&input));
    println!("part2: {}", part2(&input));
}

struct Machine {
    target_bits: u64,
    buttons_bits: Vec<u64>,
    joltage: Vec<i64>,
    buttons_indices: Vec<Vec<usize>>,
}

fn parse(input: &str) -> Vec<Machine> {
    input
        .trim()
        .lines()
        .map(|line| {
            let bracket_start = line.find('[').unwrap();
            let bracket_end = line.find(']').unwrap();
            let target_str = &line[bracket_start + 1..bracket_end];

            let target_bits = target_str.chars().enumerate().fold(0u64, |acc, (i, c)| {
                if c == '#' {
                    acc | (1 << i)
                } else {
                    acc
                }
            });

            let buttons_bits: Vec<u64> = line
                .split_whitespace()
                .filter(|s| s.starts_with('('))
                .map(|part| {
                    part.trim_matches(&['(', ')'][..])
                        .split(',')
                        .filter_map(|s| s.parse::<usize>().ok())
                        .fold(0u64, |acc, idx| acc | (1 << idx))
                })
                .collect();

            let buttons_indices: Vec<Vec<usize>> = line
                .split_whitespace()
                .filter(|s| s.starts_with('('))
                .map(|part| {
                    part.trim_matches(&['(', ')'][..])
                        .split(',')
                        .filter_map(|s| s.parse::<usize>().ok())
                        .collect()
                })
                .collect();

            let joltage: Vec<i64> = line
                .split_whitespace()
                .find(|s| s.starts_with('{'))
                .map(|part| {
                    part.trim_matches(&['{', '}'][..])
                        .split(',')
                        .filter_map(|s| s.parse::<i64>().ok())
                        .collect()
                })
                .unwrap_or_default();

            Machine {
                target_bits,
                buttons_bits,
                joltage,
                buttons_indices,
            }
        })
        .collect()
}

fn part1(input: &[Machine]) -> usize {
    let mut ret = 0;
    for machine in input {
        let n = machine.buttons_bits.len();
        let mut min = usize::MAX;

        for subset in 0..(1u32 << n) {
            let mut state = 0u64;
            for (i, &btn) in machine.buttons_bits.iter().enumerate() {
                if subset & (1 << i) != 0 {
                    state ^= btn;
                }
            }
            if state == machine.target_bits {
                min = min.min(subset.count_ones() as usize);
            }
        }
        ret += min;
    }

    ret
}

fn part2(input: &[Machine]) -> usize {
    input.iter().map(|m| min_joltage_presses(m)).sum()
}

fn min_joltage_presses(machine: &Machine) -> usize {
    let num_buttons = machine.buttons_indices.len();
    let num_counters = machine.joltage.len();

    let mut matrix: Vec<Vec<(i64, i64)>> = vec![vec![(0, 1); num_buttons + 1]; num_counters];

    // button cols
    for (btn, counters_it_affects) in machine.buttons_indices.iter().enumerate() {
        for &counter in counters_it_affects {
            if counter < num_counters {
                matrix[counter][btn] = (1, 1);
            }
        }
    }

    // last col
    for (counter, &target) in machine.joltage.iter().enumerate() {
        matrix[counter][num_buttons] = (target, 1);
    }

    let mut pivot_positions: Vec<(usize, usize)> = vec![];
    let mut current_row = 0;

    // reduce to echelon
    for col in 0..num_buttons {
        // pivot
        let mut pivot_row = None;
        for row in current_row..num_counters {
            if matrix[row][col].0 != 0 {
                pivot_row = Some(row);
                break;
            }
        }

        let Some(pivot_row) = pivot_row else {
            continue; // no pivots here
        };

        matrix.swap(current_row, pivot_row);
        pivot_positions.push((current_row, col));

        let pivot_val = matrix[current_row][col];
        for c in 0..=num_buttons {
            matrix[current_row][c] = div_frac(matrix[current_row][c], pivot_val);
        }

        for row in 0..num_counters {
            if row != current_row && matrix[row][col].0 != 0 {
                let factor = matrix[row][col];
                for c in 0..=num_buttons {
                    let subtrahend = mul_frac(factor, matrix[current_row][c]);
                    matrix[row][c] = sub_frac(matrix[row][c], subtrahend);
                }
            }
        }

        current_row += 1;
    }

    // free vars
    let pivot_cols: std::collections::HashSet<usize> =
        pivot_positions.iter().map(|&(_, col)| col).collect();
    let free_variables: Vec<usize> = (0..num_buttons)
        .filter(|col| !pivot_cols.contains(col))
        .collect();

    if free_variables.is_empty() {
        let mut solution = vec![0i64; num_buttons];
        for &(row, col) in &pivot_positions {
            let (num, denom) = matrix[row][num_buttons];

            if denom == 0 || num % denom != 0 || num / denom < 0 {
                return usize::MAX; 
            }
            solution[col] = num / denom;
        }
        return solution.iter().sum::<i64>() as usize;
    }

    let max_value = *machine.joltage.iter().max().unwrap_or(&0);
    let mut best = usize::MAX;

    search_free_variables(
        &matrix,
        &pivot_positions,
        &free_variables,
        num_buttons,
        0,
        &mut vec![0i64; free_variables.len()],
        max_value,
        &mut best,
    );

    best
}

fn search_free_variables(
    matrix: &[Vec<(i64, i64)>],
    pivots: &[(usize, usize)],
    free_vars: &[usize],
    num_buttons: usize,
    idx: usize,
    free_vals: &mut [i64],
    max_val: i64,
    best: &mut usize,
) {
    if idx == free_vars.len() {
        let mut solution = vec![0i64; num_buttons];

        for (i, &var) in free_vars.iter().enumerate() {
            solution[var] = free_vals[i];
        }

        for &(row, col) in pivots.iter().rev() {
            let (mut num, mut denom) = matrix[row][num_buttons];

            for other_col in 0..num_buttons {
                if other_col != col {
                    let coef = matrix[row][other_col];
                    let contrib = mul_frac(coef, (solution[other_col], 1));
                    (num, denom) = sub_frac((num, denom), contrib);
                }
            }

            if denom == 0 || num % denom != 0 {
                return;
            }
            solution[col] = num / denom;

            if solution[col] < 0 {
                return;
            }
        }

        if solution.iter().all(|&x| x >= 0) {
            let total = solution.iter().sum::<i64>() as usize;
            if total < *best {
                *best = total;
            }
        }
        return;
    }

    // prune here
    let current_sum: i64 = free_vals[..idx].iter().sum();
    if current_sum as usize >= *best {
        return;
    }

    for val in 0..=max_val {
        free_vals[idx] = val;
        search_free_variables(matrix, pivots, free_vars, num_buttons, idx + 1, free_vals, max_val, best);
    }
}

// fractions helpers
fn gcd(a: i64, b: i64) -> i64 {
    if b == 0 { a.abs() } else { gcd(b, a % b) }
}

fn simplify_frac((n, d): (i64, i64)) -> (i64, i64) {
    if n == 0 {
        return (0, 1);
    }
    let g = gcd(n, d);
    let (n, d) = (n / g, d / g);

    if d < 0 { (-n, -d) } else { (n, d) }
}

fn mul_frac((n1, d1): (i64, i64), (n2, d2): (i64, i64)) -> (i64, i64) {
    simplify_frac((n1 * n2, d1 * d2))
}

fn div_frac((n1, d1): (i64, i64), (n2, d2): (i64, i64)) -> (i64, i64) {
    simplify_frac((n1 * d2, d1 * n2))
}

fn sub_frac((n1, d1): (i64, i64), (n2, d2): (i64, i64)) -> (i64, i64) {
    simplify_frac((n1 * d2 - n2 * d1, d1 * d2))
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "\
    [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n\
    [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n\
    [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}\
    ";

    #[test]
    fn test_part1() {
        let input = parse(INPUT);
        assert_eq!(part1(&input), 7);
    }

    #[test]
    fn test_part2() {
        let input = parse(INPUT);
        assert_eq!(part2(&input), 33);
    }
}
