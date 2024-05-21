fn next_number(s: &str, start: usize, stop: usize) -> Option<(usize, usize)> {
    for (i, c) in s[start..stop].chars().enumerate() {
        let i = i + start;
        if c.is_numeric() {
            for (k, other) in s[i..s.len()].chars().enumerate() {
                if !other.is_numeric() {
                    return Some((i, k))
                }
            }
        }
    }

    return None
}

struct NumberIterator<'a> {
    s: &'a str,
    start: usize,
    stop: usize
}

impl <'a>Iterator for NumberIterator<'a> {
    type Item = (usize, usize);
    fn next(&mut self) -> Option<Self::Item> {
        if self.start > self.stop {
            return None
        }
        if let Some((idx, size)) = next_number(self.s, self.start, self.stop) {
            self.start = idx + size;
            return Some((idx, size))
        }
        return None
    }
} 

fn part_1(input: &str) -> i32 {
    let line_length = input.chars().take_while(|&c| c != '\n').count() + 1;
    input.chars().enumerate()
        .filter(|&(_i, c)| !c.is_numeric() && c != '.' && c != '\n')
        .map(|(i, _char)| {
            let line_pos = i % line_length;
            let start = i - line_length - 3;
            let stop = i + line_length + 3;
            (NumberIterator{s: input, start, stop}).filter_map(|(start, size)|{
                let start_pos = start % line_length;
                let end_pos = start_pos + size;
                if (start_pos > line_pos - 2 && start_pos < line_pos + 2)
                || (end_pos > line_pos - 1 && end_pos < line_pos + 2) {
                    return Some(input[start..start+size].parse::<i32>().unwrap())
                }
                return None
            }).sum::<i32>()
        }).sum()
}

fn part_2(input: &str) -> i32 {
    let line_length = input.chars().take_while(|&c| c != '\n').count() + 1;
    input.chars().enumerate()
        .filter(|&(_i, c)| !c.is_numeric() && c != '.' && c != '\n')
        .map(|(i, _char)| {
            let line_pos = i % line_length;
            let start = i - line_length - 3;
            let stop = i + line_length + 3;
            let numbers = (NumberIterator{s: input, start, stop}).filter_map(|(start, size)|{
                let start_pos = start % line_length;
                let end_pos = start_pos + size;
                if (start_pos > line_pos - 2 && start_pos < line_pos + 2)
                || (end_pos > line_pos - 1 && end_pos < line_pos + 2) {
                    return Some(input[start..start+size].parse::<i32>().unwrap())
                }
                return None
            }).collect::<Vec<_>>();
            if numbers.len() == 2 {
                return numbers.iter().product()
            }
            return 0
        }).sum()
}

#[allow(dead_code)]
const TEST_SCHEMATIC: &'static str = "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
";
#[test]
fn day_3_part_1() {
    let input = TEST_SCHEMATIC;
    assert_eq!(4361, part_1(input));
}

#[test]
fn day_3_part_2() {
    let input = TEST_SCHEMATIC;
    assert_eq!(467835, part_2(input));
}

fn main() {
    let input = aoc::include_aoc!("input/day3");
    
    println!("Part 1: {}", part_1(input));
    println!("Part 2: {}", part_2(input));
}