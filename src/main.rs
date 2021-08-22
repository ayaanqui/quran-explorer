use std::fs;

fn parse_ayah(ayah: &str) -> (i32, i32, &str) {
    let parsed_ayah: Vec<&str> = ayah.split('|').collect();

    let surah_num: i32 = parsed_ayah[0].parse::<i32>().unwrap();
    let ayah_num: i32 = parsed_ayah[1].parse::<i32>().unwrap();
    let ayah_text: &str = parsed_ayah[2];

    return (surah_num, ayah_num, ayah_text);
}

fn main() {
    let arabic_file_contents = fs::read_to_string("res/quran-simple.txt")
        .expect("Could not open quran-simple.txt");

    let arabic: Vec<&str> = arabic_file_contents.split('\n').collect();
    for ayah in arabic {
        let ch = ayah.chars().next();
        if ch == None {
            break;
        }

        if ch.unwrap() == '#' || ch.unwrap().is_whitespace() {
            continue;
        }
        let (surah_num, ayah_num, ayah_arabic) = parse_ayah(ayah);
        println!("{}:{} - {}", surah_num, ayah_num, ayah_arabic);
    }
}
