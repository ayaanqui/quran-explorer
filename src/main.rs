use std::fs;
#[path = "parsers/ayah_parser.rs"] mod parsers;

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
        let (surah_num, ayah_num, ayah_arabic) = parsers::parse_ayah(ayah);
        println!("{}:{} - {}", surah_num, ayah_num, ayah_arabic);
    }
}
