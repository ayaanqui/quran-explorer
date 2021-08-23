use std::fs;
use std::collections::HashMap;
#[path = "parsers/ayah_parser.rs"] mod parsers;
#[path = "types/ayah.rs"] mod ayah;

fn main() {
    let arabic_file_contents = fs::read_to_string("res/quran-simple.txt")
        .expect("Could not open quran-simple.txt");

    let mut arabic: Vec<&str> = arabic_file_contents.split('\n').collect();
    let mut verses: Vec<Box<ayah::Ayah>> = Vec::with_capacity(arabic.len());
    for line in &arabic {
        if line.is_empty() || line.chars().nth(0).unwrap() == '#' {
            continue;
        }

        let (surah_num, ayah_num, ayah_arabic) = parsers::parse_ayah(line);
        verses.push(Box::new(ayah::Ayah{
            chapter: surah_num,
            verse: ayah_num,
            text: String::from(ayah_arabic)
        }));
    }
    // Clear string verses to save memory...
    arabic.clear();
}
