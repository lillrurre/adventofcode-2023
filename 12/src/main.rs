use std::fs::read_to_string;

fn main() {
    let content = read_to_string("src/input.txt").unwrap();
    let lines: Vec<&str> = content.lines().collect();

    let records: Vec<&str> = lines.iter().map(|line| line.split_whitespace().next().unwrap()).collect();

    let groups: Vec<Vec<u8>> = lines.iter().map(|line| {
        line.split_whitespace()
            .nth(1).unwrap()
            .split(',')
            .map(|num| num.parse().unwrap())
            .collect()
    }).collect();

    let mut sum: u32 = 0;
    for (record, group) in records.iter().zip(groups.iter()) {
        sum += solve(record, group);
    }

    println!("Total result: {}", sum);
}

fn solve(record: &str, group: &[u8]) -> u32 {
    if record.is_empty() {
        if group.is_empty() {
            return 1;
        }
        return 0;
    }

    if group.is_empty() {
        if record.chars().all(|c| c != '#') {
            return 1;
        }
        return 0;
    }

    let first_spring = record.chars().next().unwrap().to_ascii_lowercase();
    let first_num = group[0] as usize;
    let mut res = 0;

    if first_spring == '.' {
        return solve(&record[1..], group);
    }

    if first_spring == '?' {
        res += solve(&record[1..], group);
    }

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
