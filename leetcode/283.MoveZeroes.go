package leet

func moveZeroes(nums []int)  {
	f := 0
		for n := 0; n < len(nums); n++ {
			if nums[n] != 0 {
				nums[n-f] = nums[n]
				continue
			}
			f++
		}
		for g := len(nums) - f; g < len(nums); g++ {
			nums[g] = 0
		}
	}
