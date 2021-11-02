package main

/*Đề 2:
Một đơn vị sản xuất gồm có các cán bộ là công nhân, kỹ sư, nhân viên.
+ Mỗi cán bộ cần quản lý lý các thuộc tính: Họ tên, năm sinh, giới tính, địa chỉ
+ Các công nhân cần quản lý: Bậc (công nhân bậc 3/7, bậc 4/7 ...)
+ Các kỹ sư cần quản lý: Ngành đào tạo
+ Các nhân viên phục vụ cần quản lý thông tin: công việc
1. Xây dựng các lớp NhanVien, CongNhan, KySu kế thừa từ lớp CanBo
2. Xây dựng các hàm để truy nhập, hiển thị thông tin và kiểm tra về các thuộc tính của các lớp.
3. Xây dựng lớp QLCB cài đặt các phương thức thực hiện các chức năng sau: - Nhập thông tin mới cho cán bộ
- Tìm kiếm theo họ tên
- Hiển thị thông tin về danh sách các cán bộ
- Thoát khỏi chương trình.*/
import (
	"fmt"
)

//Cán bộ
type Officer struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

func (o Officer) inforOfficer() {
	fmt.Printf("inforOfficer Name: '%s', Age: %d, Gender:%s ,Address:%s\n", o.Name, o.Age, o.Gender, o.Address)
}
func (o Officer) getName() string {
	return fmt.Sprintf(" Name: '%s'\n", o.Name)
}
func (o *Officer) setName(newName string) {
	o.Name = newName
}
func (o Officer) getAge() string {
	return fmt.Sprintf(" Age: '%d'\n", o.Age)
}
func (o *Officer) setAge(newAge int) {
	o.Age = newAge
}
func (o Officer) getGender() string {
	return fmt.Sprintf(" Gender: '%s'\n", o.Gender)
}
func (o *Officer) setGender(newGender string) {
	o.Gender = newGender
}
func (o Officer) getAddress() string {
	return fmt.Sprintf(" Address: '%s'\n", o.Address)
}
func (o *Officer) setAddress(newAddress string) {
	o.Address = newAddress
}

//Nhân Viên văn Phòng
type Staff struct {
	Task string
	Officer
}

func (s Staff) getTask() string {
	return fmt.Sprintf(" Task: '%s'\n", s.Task)
}
func (s *Staff) setTask(newTask string) {
	s.Task = newTask
}
func (s Staff) inforStaff() {
	fmt.Printf("inforStaff Name: '%s', Age: %d, Gender:%s ,Address:%s,Task:%s\n", s.Name, s.Age, s.Gender, s.Address, s.Task)
}

//Người lao động
type Worker struct {
	Level int
	Officer
}

func (w Worker) inforWorker() {
	fmt.Printf("inforWorker Name: '%s', Age: %d, Gender:%s ,Address:%s,Level:%d\n", w.Name, w.Age, w.Gender, w.Address, w.Level)
}

//Kỹ sư
type Engineer struct {
	Branch int
	Officer
}

func (e *Engineer) inforEngineer() {
	fmt.Printf("inforEngineer Name: '%s', Age: %d, Gender:%s ,Address:%s,Branch:%d\n", e.Name, e.Age, e.Gender, e.Address, e.Branch)
}

//Quản lý cán bộ
type ManagerOfficer struct {
	ListOfficer []Officer
}

func main() {
	officer := Officer{"abc", 18, "nam", "TH"}
	officer.inforOfficer()
	fmt.Println("TEN SAU KHI DOI:")
	officer.setName("ABC")
	officer.inforOfficer()
	staff := Staff{
		Task:    "Van Phong",
		Officer: officer,
	}
	staff.inforStaff()
	fmt.Println("Name Staff", staff.Officer.getName())
	fmt.Println("Name Staff", staff.getName())
	worker := Worker{
		Level:   1,
		Officer: Officer{"C", 18, "nam", "TB"},
	}
	worker.inforWorker()
	fmt.Println("Name Worker", worker.getName())
	engineer := Engineer{
		Branch:  1,
		Officer: Officer{"D", 18, "nam", "PT"},
	}
	engineer.inforEngineer()
	fmt.Println("Name Engineer", engineer.getName())
}
