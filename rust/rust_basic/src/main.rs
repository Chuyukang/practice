// practice demo
// link: https://doc.rust-lang.org/book/

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
    // let mut spaces_mut = "   ";
    // spaces_mut = spaces_mut.len(); type is not same
    // let mut spaces_mut = spaces_mut.len();
    // println!("The length of spaces_mut is: {}", spaces_mut);


    // Data Types
    // u32
    let a: u32 = 2;
    let b: u32 = 4;
    let c = a/b;
    println!("{}",c);
    // char
    let ch = "😻";
    println!("{}",ch);
    // tuple
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("({}, {}, {})", x, y, z);
    // array
    let arr: [i32; 5] = [1, 2, 3, 4, 5];
    println!("{}", arr[0]);

    // function
    another_function();
}


fn another_function() {
    println!("Another function.");
}