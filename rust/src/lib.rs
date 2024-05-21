#[macro_export]
macro_rules! include_aoc {
    ($file:expr $(,)?) => {
        include_str!(concat!("../../../", $file))
    };
}