use std::fs::read_to_string;

fn main() {
    let sum = part1("src/input.txt");
    println!("[1] Result: {sum}");
}

fn part1(input_file: &str) -> u32 {
    let content = read_to_string(input_file).unwrap();
    let lines: Vec<&str> = content.lines().collect();

    let records: Vec<&str> = lines.iter().map(|line| line.split_whitespace().next().unwrap()).collect();

    let groups: Vec<Vec<u8>> = lines.iter()
        .map(|line| line.split_whitespace()
            .nth(1).unwrap()
            .split(',')
            .map(|num| num.parse().unwrap())
            .collect()
        ).collect();

    let sum: u32 = records.iter()
        .zip(groups)
        .map(|(record, group)| solve(record, group.as_slice()))
        .sum();

    return sum
}

fn solve(record: &str, group: &[u8]) -> u32 {
    return match (record.is_empty(), group.is_empty()) {
        (true, true) => 1,
        (false, true) => if record.contains('#') { 0 } else { 1 },
        (true, false) => 0,
        _ => {
            let first_spring = record.chars().next().unwrap().to_ascii_lowercase();
            let first_num = group[0] as usize;

            let res = match first_spring {
                '?' => solve(&record[1..], group),
                _ => 0,
            };
            return match first_spring {
                '.' => solve(&record[1..], group),
                _ => {
                    if record.len() < first_num
                        || record[..first_num].contains('.')
                        || first_num != record.len() && record.chars().nth(first_num).unwrap() == '#' {
                        return res;
                    }

                    if first_num == record.len() {
                        return res + solve("", &group[1..]);
                    }

                    return res + solve(&record[first_num + 1..], &group[1..]);
                }
            };
        }
    };
}

#[cfg(test)]
mod tests {
    use crate::part1;

    #[test]
    fn test() {
        let sum = part1("src/input.txt");
        assert_eq!(sum, 7163);
    }

    #[test]
    fn test_sample() {
        let sum = part1("src/input_sample.txt");
        assert_eq!(sum, 21);
    }

}