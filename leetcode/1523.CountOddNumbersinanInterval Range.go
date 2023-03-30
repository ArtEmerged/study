package leet

func countOdds(low int, high int) int {
    answ:= 0
   for i:=low;i <= high; i++{
        if i%2 != 0{
            answ++
        }
   }
   return answ 
}