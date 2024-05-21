use std::collections::{HashSet, HashMap};

fn count_matches(line: &str) -> usize {
    let (_left, right) = line.split_once(": ").unwrap();
    let (winning, numbers) = right.split_once(" | ").unwrap();
    let winning = winning.split_whitespace().collect::<HashSet<&str>>();
    numbers.split_whitespace().into_iter().filter(|s| winning.contains(s)).count()
}

fn calc_points(input: &str) -> usize {
    input.lines().map(|line| {
        let count = count_matches(line);
        if count > 0 {
            return 2_usize.pow(count as u32 - 1)
        }
        return count
    }).sum()
}

fn calc_cards(input: &str) -> usize {
    let mut map: HashMap<usize, usize> = HashMap::new();
    for (i, line) in input.lines().enumerate() {
        let count = count_matches(line);
        let copies = map.get(&i).unwrap_or(&0) + 1;
        map.insert(i, copies);
        if count > 0 {
            for k in 1..=count {
                let j = i + k;
                map.insert(j, map.get(&j).unwrap_or(&0) + copies);
            }
        }
    }

    return map.values().sum()
}

#[test]
fn day_4_part_1() {
    let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
";
    assert_eq!(13, calc_points(input))
}

#[test]
fn day_4_part_2() {
    let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
";
    assert_eq!(30, calc_cards(input))
}

fn main() {
    let input = aoc::include_aoc!("input/day4");
    println!("Part 1: {}", calc_points(input));
    println!("Part 2: {}", calc_cards(input));
}