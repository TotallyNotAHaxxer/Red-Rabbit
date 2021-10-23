extern crate colored; // not needed in Rust 2018

    use colored::*;

    // test the example with `cargo run --example most_simple`
    fn main() {
        // TADAA!
        println!("{} {} !", "it".green(), "works".blue().bold());
    }