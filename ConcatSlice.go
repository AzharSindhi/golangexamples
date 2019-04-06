package golangexamples

func ConcatSlice(sliceToConcat[] byte) string {
	concat := ""
	for i:=0; i<len(sliceToConcat); i++ {
		concat = concat + string(sliceToConcat[i])+ "-"
	}
	
	return concat	
}

