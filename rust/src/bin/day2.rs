use std::collections::HashMap;

fn get_maxes(s: &str) -> HashMap<&str, i32> {
    let mut result: HashMap<&str, i32> = HashMap::from([
        ("red", 0),
        ("green", 0),
        ("blue", 0),
    ]);
    for (color, number) in s.split([',', ';'])
        .map(|cube| {
            let cube = cube.trim();
            let (number, color) = cube.split_once(" ").unwrap();
            let number = number.parse::<i32>().unwrap();
            (color, number)
        }) {
        result.insert(color, result[color].max(number));
    }

    return result
}

fn part_1(input: &str) -> i32 {
    input.lines().into_iter().map(|line| {
        let (game, info) = line.split_once(": ").unwrap();
        if info.split([',', ';']).map(|cubes| {
            let cubes = cubes.trim();
            let (number, color) = cubes.split_once(" ").unwrap();
            let number = number.parse::<i32>().unwrap();
            (color, number)
        }).all(|(color, number)| {
            match color {
                "red" => number <= 12,
                "green" => number <= 13,
                "blue" => number <= 14,
                _ => true
            }
        }) {
            let (_, id) = game.split_once(" ").unwrap();
            id.parse::<i32>().unwrap()
        } else {
            0
        }
    }).sum()
}

fn part_2(input: &str) -> i32 {
    input.lines().into_iter().map(|line| {
        let (_, info) = line.split_once(": ").unwrap();
        let map = get_maxes(info);
        map.values().product::<i32>()
    }).sum()
}

#[test]
fn day_2_part_1() {
    let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
";
    assert_eq!(8, part_1(input))
}

#[test]
fn day_2_part_2() {
    let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
";
    assert_eq!(2286, part_2(input))
}

fn main() {
    let input = aoc::include_aoc!("input/day2");
    println!("Part 1: {}", part_1(input));
    println!("Part 2: {}", part_2(input));
}