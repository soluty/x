package xrand

import (
	"math/rand"
	"reflect"
)

// input->[]type   output->type
func RandomByWeight(src interface{}, weight []int) interface{} {
	sum := sum(weight)
	if sum <= 0 {
		return nil
	}
	t := reflect.TypeOf(src)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		panic("RandomByWeight need a slice or array")
	}
	v := reflect.ValueOf(src)
	length := v.Len()
	if len(weight) != length {
		panic("weight length must equal to src length")
	}
	r := rand.Intn(sum)
	for i := 0; i < length; i++ {
		if r < weight[i] {
			return v.Index(i).Interface()
		}
		r -= weight[i]
	}
	return nil
}

// input->[]type   output->[]type
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

// 0 < randLen <= maxSize
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
		sum += i
	}
	return sum
}
