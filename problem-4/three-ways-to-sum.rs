// First solution: adds in a for loop 
// Time complexity: O(n)
fn sum_to_n_a(n: i32) -> i32 {
    let mut sum = 0;
    for i in 1..n + 1 {
        sum += i;
    }
        
    sum
}    
    
// Second solution: adds in a recursion
// Time complexity: O(n)
fn sum_to_n_b(n: i32) -> i32 {
    if n == 0 {
        0
    } else {
        n + sum_to_n_b(n - 1)
    }
}
    
// Third solution: adds using gauss summation formula n(n+1)/2
// Time complexity: O(1)
fn sum_to_n_c(n: i32) -> i32 {
    n * (n + 1) / 2
}
