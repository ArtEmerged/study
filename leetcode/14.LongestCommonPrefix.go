package leet

func longestCommonPrefix(strs []string) string {
    result :=strs[0]
    for _,n:=range strs[1:]{
        if len(n)<len(result){
            result = result[:len(n)]
        }
        for i,v:=range n{
            if  i >len(result)-1 ||result[i]!= byte(v){
                result= result[:i]
                break
            }
        }
    }
return result
}