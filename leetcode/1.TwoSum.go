package leet
//sum
func twoSum(nums []int, target int) []int {
    an:=make([]int,2)
for k, i:= range nums{
    for j, n:=range nums{
        if i+n ==target && k != j{
an[0],an[1]=k,j
break
        }
    }
}
return an
}