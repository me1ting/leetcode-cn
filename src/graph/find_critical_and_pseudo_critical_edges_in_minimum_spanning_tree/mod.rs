struct UnionFind {
    parent: Vec<usize>,
    size: Vec<usize>,
    set_count: usize,
}

impl UnionFind {
    pub fn new(n: usize)->UnionFind{
        let mut parent: Vec<usize> = vec![0;n];
        let mut size: Vec<usize> = vec![0;n];
        for i in 0..n {
            parent[i] = i;
            size[i] = i;
        }
        
        UnionFind{
            parent: parent,
            size: size,
            set_count: n,
        }
    }

    pub fn find(&mut self,x: usize)->usize {
        let fx = self.parent[x];
        if fx!=x {
            self.parent[x] = self.find(fx);
        }
        
        self.parent[x]
    }

    pub fn union(&mut self,x:usize,y:usize)->bool {
        let (mut fx,mut fy) = (self.find(x),self.find(y));
        if fx == fy {
            false
        }else{
            if self.size[fx] <self.size[fy] {
                self.size[fx] += self.size[fy];
                self.parent[fy] = fx;
            }else{
                self.size[fy] += self.size[fx];
                self.parent[fx] = fy;
            }
            self.set_count-=1;
            true
        }
    }
}

pub struct Solution {}

impl Solution {
    pub fn find_critical_and_pseudo_critical_edges(n: i32, edges: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut copy: Vec<Vec<i32>> = vec![vec![];edges.len()];
        for (i,e) in edges.iter().enumerate(){
            copy[i] = edges[i].clone();
            copy[i].push(i as i32);
        }

        copy.sort_by_key(|k| k[2]);

        let cal_mst = |mut uf: UnionFind,ignore_id: Option<usize>|->i32 {
            let mut mst_value = 0;
            match ignore_id {
                None => {
                    for e in &copy{
                        if uf.union(e[0] as usize,e[1] as usize) {
                            mst_value+=e[2];
                        }
                    }
                }
                Some(id)=>{
                    for (i,e) in copy.iter().enumerate(){
                        if i!=id&&uf.union(e[0] as usize,e[1] as usize) {
                            mst_value+=e[2];
                        }
                    }
                }
            }

            if uf.set_count >1 {
                return std::i32::MAX;
            }
            mst_value
        };

        let mst_value = cal_mst(UnionFind::new(n as usize),Option::None);
        let mut key_edges:Vec<i32> = vec![];
        let mut pseudokey_edges:Vec<i32> = vec![];

        for (i,e) in copy.iter().enumerate(){
            if cal_mst(UnionFind::new(n as usize),Some(i))>mst_value {
                key_edges.push(e[3]);
                continue
            }

            let mut uf = UnionFind::new(n as usize);
            uf.union(e[0] as usize, e[1] as usize);
            if e[2] + cal_mst(uf,Some(i)) == mst_value {
                pseudokey_edges.push(e[3]);
            }
        }

        vec![key_edges,pseudokey_edges]
    }
}

#[cfg(test)]
mod tests {
    use crate::graph::find_critical_and_pseudo_critical_edges_in_minimum_spanning_tree;
    #[test]
    fn it_works() {
        let n = 5;
        let edges: Vec<Vec<i32>> = vec![vec![0,1,1],vec![1,2,1],vec![2,3,2],vec![0,3,2],vec![0,4,3],vec![3,4,3],vec![1,4,6]];
        let expected = vec![vec![0,1],vec![2,3,4,5]];
        let val = find_critical_and_pseudo_critical_edges_in_minimum_spanning_tree
            ::Solution::find_critical_and_pseudo_critical_edges(n, edges);

        assert_eq!(
            expected.len(),
            val.len()
        );

        for i in 0..expected.len() {
            assert_eq!(
                expected[i].len(),
                val[i].len()
            );
            for j in 0..expected[i].len() {
                assert_eq!(expected[i][j],val[i][j])
            }
        }
    }
}