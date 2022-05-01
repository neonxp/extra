package levenshtein

import (
	"go.neonxp.dev/extra/math"
)

// Levenshtein возвращает последовательность правок для превращения последовательности элементов s1 в s2
// с учетом стоимостей операций возвращаемых функцией cost
// TODO в алгоритме не предусмотрена экономия памяти ¯\_(ツ)_/¯
func Levenshtein[T comparable](s1, s2 []T, cost EditCost[T]) []Edit {
	len1 := len(s1)
	len2 := len(s2)
	if cost == nil {
		cost = EditCost[T](func(t EditType, s1 *T, s2 *T) int {
			return 1
		})
	}

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	dp[0][0] = 0
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			switch {
			case i == 0:
				dp[0][j] = dp[0][j-1] + cost(Insert, nil, &s2[j])
			case j == 0:
				dp[i][0] = dp[i-1][0] + cost(Delete, &s1[i], nil)
			default:
				m1 := dp[i-1][j] + cost(Delete, &s1[i], nil)
				m2 := dp[i][j-1] + cost(Insert, nil, &s2[j])
				m3 := dp[i-1][j-1] + cost(Replace, &s1[i], &s2[j])
				dp[i][j] = math.MinScalar(m1, m2, m3)
			}
		}
	}
	i, j := len1-1, len2-1
	edits := []Edit{}
	for i > 0 && j > 0 {
		m1 := dp[i-1][j] + cost(Delete, &s1[i], nil)
		m2 := dp[i][j-1] + cost(Insert, nil, &s2[j])
		m3 := dp[i-1][j-1] + cost(Replace, &s1[i], &s2[j])
		min := math.MinScalar(m1, m2, m3)
		if min == m1 {
			// добавляем удаление символа S1[i]
			edits = append(edits, Edit{
				Type: Delete,
				Idx1: i,
			})
			i--
		}
		if min == m2 {
			// добавляем вставку символа S2[j] после S1[i]
			edits = append(edits, Edit{
				Type: Insert,
				Idx1: i,
				Idx2: j,
			})
			j--
		}
		if min == m3 {
			// добавляем замену S1[i] на S2[j]
			i--
			j--
			if s1[i] != s2[j] {
				edits = append(edits, Edit{
					Type: Replace,
					Idx1: i,
					Idx2: j,
				})
			}
		}
	}
	return edits
}

// EditCost функция возвращающая стоимость действия t над элементами from, to
type EditCost[T comparable] func(t EditType, from *T, to *T) int

// Edit редакторская правка описывающее одно действие над исходной последовательностью
type Edit struct {
	Type EditType // Тип правки: Вставка/Замена/Удаление
	Idx1 int      // Индекс элемента из первой последовательности
	Idx2 int      // Индекс элемента из второй последовательности
}

// EditType тип правки
type EditType int

const (
	Insert  EditType = iota // Вставка
	Replace                 // Замена
	Delete                  // Удаление
)
