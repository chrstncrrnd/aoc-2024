use regex::Regex;
use std::fs;

fn eval_mul(input: String) -> i32{
    let end = input.len() - 1;
    let params = &input[4..end];
    let (left, right) = params.split_once(",").unwrap();
    let a: i32 = left.parse().unwrap();
    let b: i32 = right.parse().unwrap();
    a * b
}

pub fn main() {
    let mut total = 0;
    let input = fs::read_to_string("./input.txt").unwrap();
    let re = Regex::new(r"mul\([0-9]{1,3},[0-9]{1,3}\)").unwrap();
    for mult in re.find_iter(input.as_str()){
        total += eval_mul(mult.as_str().to_string());
    }
    println!("The total is: {}", total)
}
