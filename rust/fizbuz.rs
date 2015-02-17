fn main() {
  for i in range(1, 101) {
    match (i % 3, i % 5) {
      (0, 0) => println!("FizBuz"),
      (0, _) => println!("Fiz"),
      (_, 0) => println!("Buz"),
      _ => println!("{}", i)
    }
  }
}
