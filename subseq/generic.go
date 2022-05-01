package subseq

func MaxSubset[T comparable](s1, s2 []T) []T {
	len1 := len(s1)
	len2 := len(s2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	max := 0
	u := 0
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			switch {
			case i == 0:
				dp[0][j] = 0
			case j == 0:
				dp[i][0] = 0
			case s1[i] != s2[j]:
				dp[i][j] = 0
			default:
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
					u = i
				}
			}
		}
	}
	ret := []T{}
	for i := max; i > 0; i-- {
		ret = append(ret, s1[u-i+1])
	}
	return ret
}

func MaxSubsetSequence[T comparable](s1, s2 []T) []T {
	len1 := len(s1)
	len2 := len(s2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := len1; i >= 0; i-- {
		for j := len2; j >= 0; j-- {
			switch {
			case i == len1, j == len2:
				dp[i][j] = 0
			case s1[i] == s2[j]:
				dp[i][j] = 1 + dp[i+1][j+1]
			default:
				max := dp[i+1][j]
				if dp[i][j+1] > max {
					max = dp[i][j+1]
				}
				dp[i][j] = max
			}
		}
	}

	ret := []T{}
	i, j := 0, 0
	for i < len1 && j < len2 {
		switch {
		case s1[i] == s2[j]:
			ret = append(ret, s1[i])
			i++
			j++
		case dp[i+1][j] >= dp[i][j+1]:
			i++
		default:
			j++
		}
	}

	return ret
}
