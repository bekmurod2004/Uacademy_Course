package main

func LongestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}

	for i:=0;i<len(strs[0]); i++ {
		c := strs[0][i]
		for j:= 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return res
			}
		}
		res += string(c)
	}
	return res
}
