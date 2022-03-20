package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])                 // 4
	t.Logf("len m1=%d", len(m1)) // len m1=3
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2)) // len m2 1
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3)) // len m3=0
}

// 访问的 Key 不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 0
	m1[2] = 0
	t.Log(m1[2])   // 0
	t.Log(len(m1)) // 1
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.") // Key 3 is not existing.
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for i, v := range m1 {
		t.Log(i, v) // 2 4 // 3 9 //1 1
	}
}
