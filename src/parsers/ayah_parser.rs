pub fn parse_ayah(ayah: &str) -> (i32, i32, &str) {
    let parsed_ayah: Vec<&str> = ayah.split('|').collect();

    let surah_num: i32 = parsed_ayah[0].parse::<i32>().unwrap();
    let ayah_num: i32 = parsed_ayah[1].parse::<i32>().unwrap();
    let ayah_text: &str = parsed_ayah[2];

    return (surah_num, ayah_num, ayah_text);
}