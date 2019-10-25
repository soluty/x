package xrand

import (
	"math/rand"
	"reflect"
)


type Indexer interface {
	Index() int
}

func RandomByWeight(ids []Indexer, weight []int) Indexer {
	sum := sum(weight)
	if sum <= 0 {
		return nil
	}
	length := len(ids)
	r := rand.Intn(sum)
	for i := 0; i < length; i++ {
		if r < weight[i] {
			return ids[i]
		}
		r -= weight[i]
	}
	return nil
}

func Shuffle(src interface{}, srcLen ...int) interface{} {
	t := reflect.TypeOf(src)
	switch t.Kind() {
	case reflect.Slice:
		l := reflect.ValueOf(src).Len()
		if l <= 1 {
			return src
		}
		var ret []interface{}
		for i := l - 1; i >= 0; i-- {
			v := reflect.ValueOf(src)
			r := rand.Intn(l - len(ret))
			ret = append(ret, v.Index(r).Interface())
			src = reflect.AppendSlice(v.Slice(0, r), v.Slice(r+1, v.Len())).Interface()
		}
		return ret
	default:
		panic("Shuffle need a slice params")
	}
}

func GetOrderedRandArray(randLen int, maxSize int) []int {
	if randLen > maxSize {
		randLen = maxSize
	}
	if randLen <= 0 {
		return nil
	}
	var ret []int
	// rand.Seed(time.Now().Unix())
	for i := randLen - 1; i >= 0; i-- {
		r := rand.Intn(maxSize - len(ret))
		if len(ret) == 0 {
			ret = append(ret, r)
			continue
		}
		for idx, value := range ret {
			if r >= value {
				r++
				if idx != len(ret)-1 {
					continue
				} else {
					idx++
				}
			}
			ret2 := append([]int{r}, ret[idx:]...)
			ret = append(ret[:idx], ret2...)
			break
		}
	}
	return ret
}



func sum(arr []int) int {
	var sum int
	for _, i := range arr {
		sum = sum + (i)
	}
	return sum
}