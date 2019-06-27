/**
* @File: person.go
* @Author: wongxinjie
* @Date: 2019/6/25
*/
package problem_10_2

type Person struct {
	friendIds []int64
	personId int64
}

func CreatePerson(personId int64) *Person {
	return &Person{
		friendIds: make([]int64, 0),
		personId: personId,
	}
}


func (p *Person) GetId() int64 {
	return p.personId
}

func (p *Person) GetFriends() []int64{
	return p.friendIds
}

func (p *Person) AddFriend(id int64) {
	p.friendIds = append(p.friendIds, id)
}
