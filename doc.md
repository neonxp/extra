<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# channels

```go
import "go.neonxp.dev/extra/channels"
```

Функции работы над каналами

## Index

- [func FanIn[T any](chans ...chan T) chan T](<#func-fanin>)
- [func FanOut[T any](in chan T, workers int) []chan T](<#func-fanout>)


## func FanIn

```go
func FanIn[T any](chans ...chan T) chan T
```

FanIn сливает несколько каналов в один

## func FanOut

```go
func FanOut[T any](in chan T, workers int) []chan T
```

FanOut раскидывает очередь канала на несколько каналов равномерно

# collections

```go
import "go.neonxp.dev/extra/collections"
```

Функции работы над обобщенными коллекциями

## Index

- [func Merge[T utils.Sortable[T]](s1, s2 []T) []T](<#func-merge>)
- [func MergeScalar[T utils.Scalar](s1, s2 []T) []T](<#func-mergescalar>)


## func Merge

```go
func Merge[T utils.Sortable[T]](s1, s2 []T) []T
```

Merge стабильное слияние двух слайсов

## func MergeScalar

```go
func MergeScalar[T utils.Scalar](s1, s2 []T) []T
```

MergeScalar стабильное слияние двух слайсов скаляров

# levenshtein

```go
import "go.neonxp.dev/extra/levenshtein"
```

Пакет с функциями получения редакторских правок по Левенштейну

## Index

- [type Edit](<#type-edit>)
  - [func Levenshtein[T comparable](s1, s2 []T, cost EditCost[T]) []Edit](<#func-levenshtein>)
  - [func String(s1, s2 string) []Edit](<#func-string>)
  - [func Strings(s1, s2 []string) []Edit](<#func-strings>)
- [type EditCost](<#type-editcost>)
- [type EditType](<#type-edittype>)


## type Edit

Edit редакторская правка описывающее одно действие над исходной последовательностью

```go
type Edit struct {
    Type EditType // Тип правки: Вставка/Замена/Удаление
    Idx1 int      // Индекс элемента из первой последовательности
    Idx2 int      // Индекс элемента из второй последовательности
}
```

### func Levenshtein

```go
func Levenshtein[T comparable](s1, s2 []T, cost EditCost[T]) []Edit
```

Levenshtein возвращает последовательность правок для превращения последовательности элементов s1 в s2 с учетом стоимостей операций возвращаемых функцией cost TODO в алгоритме не предусмотрена экономия памяти ¯\\\_\(ツ\)\_/¯

### func String

```go
func String(s1, s2 string) []Edit
```

String возвращает последовательность правок для превращения строки s1 в строку s2

### func Strings

```go
func Strings(s1, s2 []string) []Edit
```

Strings возвращает последовательность правок для превращения слайса строк s1 в слайс строк s2

## type EditCost

EditCost функция возвращающая стоимость действия t над элементами from\, to

```go
type EditCost[T comparable] func(t EditType, from *T, to *T) int
```

## type EditType

EditType тип правки

```go
type EditType int
```

```go
const (
    Insert  EditType = iota // Вставка
    Replace                 // Замена
    Delete                  // Удаление
)
```

# math

```go
import "go.neonxp.dev/extra/math"
```

Пакет с математическими функциями над обобщенными типами

## Index

- [func Max[T utils.Sortable[T]](a ...T) T](<#func-max>)
- [func MaxScalar[T utils.Scalar](a ...T) T](<#func-maxscalar>)
- [func Min[T utils.Sortable[T]](a ...T) T](<#func-min>)
- [func MinScalar[T utils.Scalar](a ...T) T](<#func-minscalar>)


## func Max

```go
func Max[T utils.Sortable[T]](a ...T) T
```

Max возвращает максимальное из переданных Sortable значений

## func MaxScalar

```go
func MaxScalar[T utils.Scalar](a ...T) T
```

MaxScalar возвращает максимальное из переданных скалярных значений

## func Min

```go
func Min[T utils.Sortable[T]](a ...T) T
```

Min возвращает минимальное из переданных Sortable значений

## func MinScalar

```go
func MinScalar[T utils.Scalar](a ...T) T
```

MinScalar возвращает минимальное из переданных скалярных значений

# subseq

```go
import "go.neonxp.dev/extra/subseq"
```

Пакет с функциями получения подпоследовательностей \(например\, подстроки\)

## Index

- [func MaxSubsequence(s1, s2 string) string](<#func-maxsubsequence>)
- [func MaxSubset[T comparable](s1, s2 []T) []T](<#func-maxsubset>)
- [func MaxSubsetSequence[T comparable](s1, s2 []T) []T](<#func-maxsubsetsequence>)
- [func MaxSubstring(s1, s2 string) string](<#func-maxsubstring>)


## func MaxSubsequence

```go
func MaxSubsequence(s1, s2 string) string
```

MaxSubsequence возвращает максимальную общую подпоследовательность символов

## func MaxSubset

```go
func MaxSubset[T comparable](s1, s2 []T) []T
```

## func MaxSubsetSequence

```go
func MaxSubsetSequence[T comparable](s1, s2 []T) []T
```

## func MaxSubstring

```go
func MaxSubstring(s1, s2 string) string
```

MaxSubstring возвращаает максимальную общую подстроку

# utils

```go
import "go.neonxp.dev/extra/utils"
```

Всякое разное

## Index

- [func Ptr[T any](t T) *T](<#func-ptr>)
- [type Scalar](<#type-scalar>)
- [type Sortable](<#type-sortable>)


## func Ptr

```go
func Ptr[T any](t T) *T
```

Ptr получение указателя от произвольного значения\, например\, ToPtr\(true\) \-\> \*bool\(true\)

## type Scalar

Scalar соответствует скалярам над которыми определены операции сравнения

```go
type Scalar interface {
    // contains filtered or unexported methods
}
```

## type Sortable

Sortable соответствует типам над которыми определена операция Less

```go
type Sortable[T any] interface {
    Less(T) bool
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
