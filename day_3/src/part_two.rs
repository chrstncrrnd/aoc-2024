use regex::Regex;
use std::{env, fs, str::FromStr};

#[derive(Debug)]
enum Token{
    Do,
    DoNot,
    Mul(i32, i32)
}

impl FromStr for Token{
    type Err = env::VarError; 
    
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s{
            "do()" => Ok(Token::Do),
            "don't()" => Ok(Token::DoNot), 
            _ => {
                let end = s.len() - 1;
                let params = &s[4..end];
                let (left, right) = params.split_once(",").unwrap();
                let a: i32 = left.parse().unwrap();
                let b: i32 = right.parse().unwrap();
                Ok(Token::Mul(a,b)) 
            }
        }
    } 
}

pub fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();
    let re = Regex::new(r"(?:mul\([0-9]{1,3},[0-9]{1,3}\))|(?:do(n't){0,1}\(\))").unwrap();
    let token_stream = re.find_iter(input.as_str()).map(|m| Token::from_str(m.as_str()).unwrap()); 
    
    let mut total = 0;
    let mut enabled = true; 
    for token in token_stream{ 
        match token{
            Token::Do => {
                enabled = true
            },
            Token::DoNot => {
                enabled = false
            }
            Token::Mul(a, b) => {
                if enabled{
                    total += a * b
                }
            }
        } 
    }
    println!("The total is: {}", total)
}
