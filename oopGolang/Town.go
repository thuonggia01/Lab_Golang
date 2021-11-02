/*
Đề 4:
Để quản lý các hộ dân trong một khu phố, người ta quản lý các thông tin
như sau:
- Với mỗi hộ dân, có các thuộc tính:
+ Số thành viên trong hộ ( số người)
+ Số nhà của hộ dân đó. ( Số nhà được gắn cho mỗi hộ dân)
+ Thông tin về mỗi cá nhân trong hộ gia đình.
- Với mỗi cá nhân, người ta quản lý các thông tin như: họ và tên, tuổi, năm sinh, nghề nghiệp.
1. Hãy xây dựng lớp Nguoi để quản lý thông tin về mỗi cá nhân.
2. Xây dựng lớp KhuPho để quản lý thông tin về các hộ gia đình.
3. Viết các phương thức nhập, hiển thị thông tin cho mỗi cá nhân.
4. Cài đặt chương trình thực hiện các công việc sau:
- Nhập vào một dãy gồm n hộ dân (n - nhập từ bàn phím). - Hiển thị ra màn hình thông tin về các hộ trong khu phố.
*/
package main

import "fmt"

//person
type Person struct {
	Name     string
	Age      int
	Job      string
	Passport string
}

func (p Person) getName() string {
	return fmt.Sprintf("Name:%s", p.Name)
}
func (p *Person) setName(newName string) {
	p.Name = newName
}
func (p Person) getAge() string {
	return fmt.Sprintf("Age:%d", p.Age)
}
func (p *Person) setAge(newAge int) {
	p.Age = newAge
}
func (p Person) getJob() string {
	return fmt.Sprintf("Job:%s", p.Job)
}
func (p *Person) setJob(newJob string) {
	p.Job = newJob
}
func (p Person) getPassport() string {
	return fmt.Sprintf("Passport:%s", p.Passport)
}
func (p *Person) setPassport(newPassport string) {
	p.Passport = newPassport
}

func (p *Person) inForPerson() string {
	return fmt.Sprintf("inForPerson Name: %s, Age: %d, Job:%s ,Passport:%s\n", p.Name, p.Age, p.Job, p.Passport)
}
func newPerson(name string, age int, job string, passport string) *Person {
	return &Person{Name: name, Age: age, Job: job, Passport: passport}
}

// family
type Family struct {
	ListPerson []Person
	Address    string
}

func (f Family) getAddress() string {
	return fmt.Sprintf("Passport:%s", f.Address)
}
func (f *Family) setAddress(newAddress string) {
	f.Address = newAddress
}

func (f *Family) inForFamily() string {
	return fmt.Sprintf(" Address %v ,ListPerson %v:", f.Address, f.ListPerson)
}

func (f *Family) getListPerson() []Person {
	list := make([]Person, 0, 0)
	for _, person := range f.ListPerson {
		list = append(list, person)
	}
	return list
}
func newFamily(Address string) *Family {
	list := make([]Person, 0)
	return &Family{list, Address}
}

type Town struct {
	ListFamily []Family
}

func newTown(listFamily []Family) *Town {
	return &Town{listFamily}
}

func (t *Town) inForTown() {
	for _, family := range t.ListFamily {
		inForFamily := family.inForFamily()
		fmt.Println("====================================")
		fmt.Println(inForFamily)
		for _, person := range family.getListPerson() {
			fmt.Printf("\t %v \n", person.inForPerson())
		}
	}
}
func main() {
	p1 := &Person{
		Name:     "A",
		Age:      18,
		Job:      "abc",
		Passport: "HN",
	}
	p2 := &Person{
		Name:     "B",
		Age:      18,
		Job:      "abc",
		Passport: "HN",
	}

	p1.inForPerson()
	fmt.Println("Passport Person:", p1.getPassport())
	p2.setPassport("TH")
	fmt.Println("Passport Person:", p2.getPassport())
	Family1 := newFamily("cau giay")
	Family1.ListPerson = append(Family1.ListPerson, *p1)
	Family1.ListPerson = append(Family1.ListPerson, *p2)

	Family2 := newFamily("Xuan Phuong")
	Family2.ListPerson = append(Family2.ListPerson, *p1)
	Family2.ListPerson = append(Family2.ListPerson, *p2)

	listFamily := make([]Family, 0)
	listFamily = append(listFamily, *Family1)
	listFamily = append(listFamily, *Family2)

	Town := newTown(listFamily)
	Town.inForTown()
}
