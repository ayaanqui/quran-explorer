pub fn parse_ayah(ayah: &str) -> (i32, i32, &str) {
    let parsed_ayah: Vec<&str> = ayah.split('|').collect();

    let surah_num: i32 = parsed_ayah[0].parse::<i32>().unwrap();
    let ayah_num: i32 = parsed_ayah[1].parse::<i32>().unwrap();

    return (surah_num, ayah_num, parsed_ayah[2]);
}