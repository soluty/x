package xrand

import (
	cRand "crypto/rand"
	"errors"
	"fmt"
	"math/big"
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
func Shuffle(src interface{}) interface{} {
	t := reflect.TypeOf(src)
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		v := reflect.ValueOf(src)
		l := v.Len()
		randIndexes := rand.Perm(l)
		var ret []interface{}
		for i := 0; i < l; i++ {
			ret = append(ret, v.Index(randIndexes[i]).Interface())
		}
		return ret
	default:
		panic("Shuffle need slice or array params")
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

// 从一个环形中随机取n个点
func RandomPointsInSquareRing(randCount int, r1, r2 struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}) (points []struct {
	X int
	Y int
}, err error) {
	if r1.X1 > r2.X1 || r1.Y1 > r2.Y1 || r1.X2 < r2.X2 || r1.Y2 < r2.Y2 {
		return nil, errors.New("传入坐标参数错误，r1是外圈，r2是内圈")
	}
	totalCount := 0
	candidate := make([]struct {
		X int
		Y int
	}, 0)
	for i := r1.X1; i <= r1.X2; i++ {
		for j := r1.Y1; j <= r1.Y2; j++ {
			if i > r2.X1 && i < r2.X2 && j > r2.Y1 && j < r2.Y2 {
				continue
			} else {
				totalCount++
				candidate = append(candidate, struct {
					X int
					Y int
				}{X: i, Y: j})
			}
		}
	}
	if randCount > totalCount {
		return nil, fmt.Errorf("randCount太大，总共只有%v个可随机的点，传入的randCount=%v", totalCount, randCount)
	}
	randIndexes := GetOrderedRandArray(randCount, totalCount)
	for _, value := range randIndexes {
		points = append(points, candidate[value])
	}
	return
}

func sum(arr []int) int {
	var sum int
	for _, i := range arr {
		sum += i
	}
	return sum
}

var maxList = []*big.Int{
	new(big.Int).SetInt64(10),
	new(big.Int).SetInt64(100),
	new(big.Int).SetInt64(1000),
	new(big.Int).SetInt64(10000),
	new(big.Int).SetInt64(100000),
	new(big.Int).SetInt64(1000000),
	new(big.Int).SetInt64(10000000),
	new(big.Int).SetInt64(100000000),
	new(big.Int).SetInt64(1000000000),
	new(big.Int).SetInt64(10000000000),
}

func RandNumberString(strLen int) (string, error) {
	if strLen <= 0 || strLen > len(maxList) {
		return "", fmt.Errorf("strLen must in [1, %v]", len(maxList))
	}
	max := maxList[strLen-1]
	i, err := cRand.Int(cRand.Reader, max)
	if err != nil {
		return "", fmt.Errorf("RandNumberString: %w", err)
	}
	x := i.String()
	for len(x) < strLen {
		x = "0" + x
	}
	return x, nil
}
