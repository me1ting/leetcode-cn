package sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	N := len(nums)

	ans := [][]int{}

	for i := 0; i+3 < N; {
		iVal := nums[i]
		for j := i + 1; j+2 < N; {
			jVal := nums[j]
			need := target - iVal - jVal
		inner:
			for m, n := j+1, N-1; m < n; {
				mVal, nVal := nums[m], nums[n]
				if need < 2*mVal || need > 2*nVal {
					break inner
				}

				sum := mVal + nVal
				if sum == need {
					ans = append(ans, []int{iVal, mVal, nVal, jVal})
					for nums[m] == mVal && m < n {
						m++
					}
					for nums[n] == nVal && m < n {
						n--
					}
				} else if sum < need {
					m++
				} else {
					n--
				}
			}
			j++
			for j < N && nums[j] == jVal {
				j++
			}
		}
		i++
		for i < N && nums[i] == iVal {
			i++
		}
	}
	return ans
}
