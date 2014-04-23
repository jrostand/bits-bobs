// Iterative fibonacci up to 64-bit unsigned int
fn fib(n : u64) -> u64 {
  if n <= 1 { return n }
  else {
    // Assign these separately because of dead assignment warnings
    let mut tmp:u64;
    let (mut cur, mut next) = (0u64, 1u64);

    for _ in range(0, n) {
      tmp = cur + next;

      // Detect int overflow and halt
      if next < cur {
        println!("u64 overflow at n {:u}", n);
        return 0;
      }

      cur = next;
      next = tmp;
    }

    return cur;
  }
}

fn main() {
  // Just output fibonacci numbers 1-100
  // Spoiler alert: overflows at i=94
  for i in range(1u64,100u64) {
    println!("{} | {}", i, fib(i));
  }
}
