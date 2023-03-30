package leet

//	func numIdenticalPairs(nums []int) int {
//	    var answer int
//	    for i:=0; i < len(nums); i++{
//	        for j:= i + 1; j < len(nums); j++{
//	            if nums[i]== nums[j]{
//	                answer++
//	            }
//	        }
//	    }
//	    return answer
//	}
func numIdenticalPairs(nums []int) int {
	key := map[int]int{}
	var res int
	for _, v := range nums {
		res += key[v]
		key[v]++
	}
	return res
}
