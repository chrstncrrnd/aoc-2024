use std::{collections::HashMap, fs};

fn load_file(filename: &str) -> Vec<Vec<i32>> {
    let contents = fs::read_to_string(filename).unwrap();
    contents
        .split_terminator("\n")
        .map(|s| {
            s.chars()
                .map(|ch: char| ch.to_digit(10).unwrap() as i32)
                .collect::<Vec<i32>>()
        })
        .collect()
}

fn find_trail_heads(input: &[Vec<i32>]) -> Vec<(i32, i32)> {
    let height = input.len();
    let width = input.first().unwrap().len();
    let mut out = Vec::<(i32, i32)>::new();
    for x in 0..width {
        for y in 0..height {
            if *input.get(y).unwrap().get(x).unwrap() == 0 {
                out.push((x as i32, y as i32));
            }
        }
    }
    out
}

fn is_in_area(position: (i32, i32), size: (i32, i32)) -> bool {
    position.0 >= 0 && position.1 >= 0 && position.0 < size.0 && position.1 < size.1
}

fn get_at_position(input: &[Vec<i32>], position: (i32, i32)) -> i32 {
    *input
        .get(position.1 as usize)
        .unwrap()
        .get(position.0 as usize)
        .unwrap()
}

fn explore_trail(
    prev_level: i32,
    input: &[Vec<i32>],
    position: (i32, i32),
    size: (i32, i32),
    unique: &mut HashMap<(i32, i32), bool>,
) {
    if !is_in_area(position, size) {
        return;
    }
    let level = get_at_position(input, position);
    if level != prev_level + 1 {
        return;
    }
    if level == 9 {
        (*unique).entry(position).or_insert(true);
        return;
    }
    explore_trail(level, input, (position.0, position.1 + 1), size, unique);
    explore_trail(level, input, (position.0 + 1, position.1), size, unique);
    explore_trail(level, input, (position.0, position.1 - 1), size, unique);
    explore_trail(level, input, (position.0 - 1, position.1), size, unique);
}

pub fn main() {
    let input = load_file("input.txt");
    let size = (input.len() as i32, input.first().unwrap().len() as i32);
    let mut unique = HashMap::new();
    let total: i32 = find_trail_heads(&input)
        .iter()
        .map(|value| {
            explore_trail(-1, &input, *value, size, &mut unique);
            let score = unique.len() as i32;
            unique.clear();
            score
        })
        .sum();
    println!("The total is: {}", total);
}
