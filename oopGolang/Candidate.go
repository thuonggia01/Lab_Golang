/*
Đề 3:
Các thí sinh dự thi đại học bao gồm các thí sinh thi khối A, thí sinh thi khối B, thí sinh thi khối C
+ Các thí sinh cần quản lý các thuộc tính: Số báo danh, họ tên, địa chỉ, ưu tiên.
+ Thí sinh thi khối A thi các môn: Toán, lý, hoá
+ Thí sinh thi khối B thi các môn: Toán, Hoá, Sinh
+ Thí sinh thi khối C thi các môn: văn, Sử, Địa
1. Xây dựng các lớp để quản lý các thí sinh sao cho sử dụng lại được nhiều nhất.
2. Xây dựng lớp TuyenSinh cài đặt các phương thức thực hiện các nhiệm vụ sau:
- Nhập thông tin về các thí sinh dự thi
- Hiển thị thông tin về một thí sinh
- Tìm kiếm theo số báo danh
- Kết thúc chương trình.
*/
package main

import "fmt"

type Candidate struct {
	Id        string
	Name      string
	Address   string
	Priotity  int
	Subjects1 string
	Subjects2 string
	Subjects3 string
}

func (c Candidate) getId() string {
	return c.Id
}

func (c Candidate) inForCandidate() {
	fmt.Printf("ID: %v , Name: %s ,Address: %s , priority: %v , Subjects1: %v, Subjects2 :%v, Subjects3 :%v ,\n ", c.Id, c.Name, c.Address, c.Priotity, c.Subjects1, c.Subjects2, c.Subjects3)
}

func newCandidateA(Id string, Name string, Address string, Priotity int) *Candidate {
	return &Candidate{Id, Name, Address, Priotity, "Maths", "physics", "chemistry"}
}
func newCandidateB(Id string, Name string, Address string, Priotity int) *Candidate {
	return &Candidate{Id, Name, Address, Priotity, "Maths", "chemistry", "biology"}
}
func newCandidateC(Id string, Name string, Address string, Priotity int) *Candidate {
	return &Candidate{Id, Name, Address, Priotity, "Literature", " physics", "geography"}
}
func showCandidate(list []*Candidate) {
	for _, sa := range list {
		sa.inForCandidate()
	}
}

func findCandidate(list []*Candidate, name string) {
	for _, sa := range list {
		if sa.Name == name {
			sa.inForCandidate()
		}
	}
}
func main() {
	listCandidate := make([]*Candidate, 0)

	a := newCandidateA("1", "Long", "BG", 1)
	b := newCandidateB("2", "Lan", "HN", 1)
	c := newCandidateC("3", "Phung", "HCM", 1)

	listCandidate = append(listCandidate, a)
	listCandidate = append(listCandidate, b)
	listCandidate = append(listCandidate, c)

	showCandidate(listCandidate)
	fmt.Println("Find Candidate:")
	findCandidate(listCandidate, "Long")
}
