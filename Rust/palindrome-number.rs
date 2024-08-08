impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        let mut x_copy = x;
        let mut x_reversed = 0;

        while x_copy > 0 {
            x_reversed *= 10;
            x_reversed += x_copy % 10;
            x_copy /= 10;
        }
        x_reversed == x
    }
}
