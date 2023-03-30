package leet

func numberOfSteps(num int) int {
    answer:=0
  for num > 0 {
      if num%2 == 0{
          num/=2
          answer++
      }else{
          num-=1
          answer++
      }
  }  
  return answer
}