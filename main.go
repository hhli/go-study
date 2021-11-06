package main

import "fmt"

// main
func main() {
	//compile.Walk()

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < 10; i++ {
	//	fmt.Println(r.Intn(10))
	//}

	fmt.Println(len(PersonUnique{}.FocusOnList))
}

type PersonUnique struct {
	FocusOnList   []string `json:"focus_on_list,omitempty"`  // 特别关注列表，只保留id
	PrivilegeType int32    `json:"privilege_type,omitempty"` // 100 role_id  200 role_id、monitor_type  300 role_id、monitor_type、source 400 role_id、monitor_type、source、operator
}
