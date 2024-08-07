use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut tmp_map = HashMap::new();
        let mut i = 0;
        for num in nums.iter() {
            match tmp_map.get(&(target - num)) {
                Some(&chord) => return vec![chord, i],
                None => {tmp_map.insert(num, i);}
            };
            i += 1;
        }
        vec![]
    }
}
