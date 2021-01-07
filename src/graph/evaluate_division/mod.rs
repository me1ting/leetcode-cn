pub struct Solution {}
// https://leetcode-cn.com/problems/evaluate-division/
// 注意该题容易让人误解：a/b,b/c,ab/bc，其中ab是a*b吗？答案是不是，`每个 Ai 或 Bi 是一个表示单个变量的字符串`，
// 否则的话，就涉及到分数的化简...，很是复杂
//
// 解法1：图，本题考察对多点最短路径的考察，即Floyd算法
// 与标准的Floyd算法不同的是，这里将每一个变量作为一个顶点，将a/b作为a->b的值，根据数学我们可以得出，无论从哪一条路径走下去，只要不出现环
// ，其路径总值总是相等的，因此不需要比较大小。
//
// 这里使用Rust的原因是因为个人使用的参考题解是用Go写的，使用Rust可以避免纯粹的复制粘贴
//
// Golang使用47行，这里使用了55行，相比Golang写起来确实麻烦一点，特别是很久没写Rust，语法都忘记了。
impl Solution {
    pub fn calc_equation(
        equations: Vec<Vec<String>>,
        values: Vec<f64>,
        queries: Vec<Vec<String>>,
    ) -> Vec<f64> {
        use std::collections::HashMap;
        //为每个变量分配数字id，供图使用
        let mut id: HashMap<String, usize> = HashMap::new();
        for equ in &equations {
            let (a, b) = (&equ[0], &equ[1]);
            if !id.contains_key(a) {
                id.insert(a.clone(), id.len());
            }
            if !id.contains_key(b) {
                id.insert(b.clone(), id.len());
            }
        }

        //创建图
        let mut graph: Vec<Vec<f64>> = vec![vec![0.0; id.len()]; id.len()];

        for (i, equ) in equations.iter().enumerate() {
            let (a, b) = (&equ[0], &equ[1]);
            let (v, w) = (id[a], id[b]);
            let val = values[i];
            graph[v][w] = val;
            graph[w][v] = 1.0 / val;
        }

        //Floyd算法
        for k in 0..id.len() {
            for i in 0..id.len() {
                for j in 0..id.len() {
                    let from = graph[i][k];
                    let to = graph[k][j];
                    if from > 0.0 && to > 0.0 {
                        graph[i][j] = from * to;
                    }
                }
            }
        }

        let mut ans: Vec<f64> = vec![0.0; queries.len()];
        for (i, q) in queries.iter().enumerate() {
            let (a, b) = (&q[0], &q[1]);
            let (v, m) = (id.get(a), id.get(b));
            if v.is_some() && m.is_some() {
                let (v, m) = (*v.unwrap(), *m.unwrap());
                if graph[v][m] > 0.0 {
                    ans[i] = graph[v][m];
                } else {
                    ans[i] = -1.0;
                }
            } else {
                ans[i] = -1.0;
            }
        }
        return ans;
    }
}

#[cfg(test)]
mod tests {
    use crate::graph::evaluate_division;
    #[test]
    fn it_works() {
        let equations = vec![
            vec!["a".to_string(), "b".to_string()],
            vec!["b".to_string(), "c".to_string()],
        ];
        let values = vec![2.0, 3.0];
        let queries = vec![
            vec!["a".to_string(), "c".to_string()],
            vec!["b".to_string(), "a".to_string()],
        ];
        let expected = vec![6.0, 0.5];
        assert_eq!(
            expected,
            evaluate_division::Solution::calc_equation(equations, values, queries)
        );
    }
}
