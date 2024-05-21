enum Direction {
    Forwards,
    Backwards
}

fn parse_digit(s: &str) -> Option<char> {
    match s {
        "one" => Some('1'),
        "two" => Some('2'),
        "three" => Some('3'),
        "four" => Some('4'),
        "five" => Some('5'),
        "six" => Some('6'),
        "seven" => Some('7'),
        "eight" => Some('8'),
        "nine" => Some('9'),
        _ => None
    }
}

fn find_digit(s: &str, direction: Direction) -> Option<char> {
    let s = match direction {
        Direction::Forwards => s.to_string(),
        Direction::Backwards => s.to_string()
            .chars()
            .rev()
            .collect()
    };

    for (i, c) in s.chars().enumerate() {
        if c.is_numeric() {
            return Some(c)
        }
        
        for k in 0..i {
            let sub = match direction {
                Direction::Forwards => s[k..=i].to_string(),
                Direction::Backwards => s[k..=i].chars().rev().collect()
            };
            if let res @ Some(_digit) = parse_digit(&sub) {
                return res
            }
        }
    }

    return None
}

fn part_1(s: &str) -> i32 {
    s.lines().into_iter()
        .map(|line| {
            line.replace(char::is_alphabetic, "")
        })
        .map(|filtered| {
            (filtered[0..1].to_string() + &filtered[filtered.len()-1..filtered.len()])
                .parse::<i32>().expect("should only contain numbers")
        })
        .sum()
}

fn part_2(s: &str) -> i32 {
    s.lines().into_iter()
        .map(|line| {
            let left = find_digit(line, Direction::Forwards).unwrap();
            let right = find_digit(line, Direction::Backwards).unwrap();
            [left, right].into_iter().collect::<String>()
                .parse::<i32>().expect("should be valid number")
        })
        .sum()
}

#[test]
fn day_1_part_1() {
    let input = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
";
    assert_eq!(142, part_1(input))
}

#[test]
fn day_1_part_2() {
    let input = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
";
    assert_eq!(281, part_2(input))
}

fn main() {
    let input = aoc::include_aoc!("input/day1");
    println!("Part 1: {}", part_1(input));
    println!("Part 2: {}", part_2(input));
}