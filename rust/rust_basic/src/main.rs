fn main() {
    // imutable and muttable
    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);

    // variables and constants
    const MAX_POINTS: u32 = 100_000; // naming hardcoded value
    println!("The max points you can get is: {}", MAX_POINTS);

    // shadowing
    let x = x+1;
    let x = x+2;
    let x = x*2;
    println!("The value of x is: {}", x);
    let spaces = "   ";
    let spaces = spaces.len();
    println!("The length of spaces is: {}", spaces);
    let mut spaces_mut = "   ";
    // spaces_mut = spaces_mut.len(); type is not same
    let mut spaces_mut = spaces_mut.len();
    println!("The length of spaces_mut is: {}", spaces_mut);


    // Data Types
    // 

}
