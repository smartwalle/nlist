package combination

func Combination(params [][]interface{}) (results [][]interface{}) {
	var pCount = 1
	for _, p := range params {
		pCount = pCount * len(p)
	}
	results = make([][]interface{}, 0, 0)

	var psLen = len(params)
	for i:=0; i<pCount; i++ {
		var s = make([]interface{}, 0, 0)
		var temp = i
		for m:=0; m<psLen; m++ {
			var pLen = len(params[m])
			if (temp / pLen) >= 0 {
				s = append(s, params[m][temp % pLen])
				temp /= pLen
			}
		}
		results = append(results, s)
	}
	return results
}
