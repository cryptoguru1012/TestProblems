use time::PreciseTime;
//problem 48
//The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
//Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

//Get n^n
pub fn mod_of_power(value: u64, power: u64, modulus: u64) -> u64 {
    let mut x: u64 = 1;
    for _ in 0..power {
        x = (x * value) % modulus;
    }
    x
}

//Get last 10 digits of 1^1 + 2^2 + 3^3 + ... + n^n.
pub fn prob48(value: u64) -> u64 {
    let mut sum_value = 0;
    for n in 1..value+1 {
        sum_value += mod_of_power(n, n, 10_000_000_000);
    }
    sum_value % 10_000_000_000
}

#[cfg_attr(rustfmt, rustfmt_skip)]
fn main() {
    let input  = 1000;// You can change this value as you want.

    let start_all = PreciseTime::now();

    let start = PreciseTime::now(); 
	let result_sum = prob48(input) as u64; 
	let end = PreciseTime::now(); 
	println!("problem 48 seconds: {} answer: {:?}", start.to(end), result_sum);

    let end_all = PreciseTime::now();
    println!("benchmark time {}", start_all.to(end_all))
}


#[cfg(test)]
mod tests {
    
    use super::*;
    //OK case
    #[test]
    fn test_case() {
        assert_eq!(prob48(1000), 9110846700);
    }

    //Bad case
    #[test]
    fn test_bad_case() {
    
        assert_eq!(prob48(100), 123);
    }
}