package person

import (
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestPb3(t *testing.T) {
	t.Log("====== Test proto3 ======")

	//Create empty object
	personEmpty := &Person{}

	t.Logf("person of empty:%v", personEmpty.String())

	//Create object with fields
	personWithFields := &Person{
		Name:  "techie",
		Id:    200508628,
		Email: "200508628@qq.com",
		Phone: []*Person_PhoneNumber{
			{
				Type:   Person_MOBILE,
				Number: "56774",
			},
			{
				Type:   Person_WORK,
				Number: "456789",
			},
		},
	}

	t.Logf("person with fields:%v\r\n", personWithFields.String())

	//Create object and set fields
	person := &Person{}
	person.Name = "techie"
	person.Id = 200508628
	person.Email = "200508628@qq.com"

	phoneNumbers := []*Person_PhoneNumber{}
	phoneNumbers = append(phoneNumbers, &Person_PhoneNumber{Type: Person_HOME, Number: "1234346"})

	person.Phone = phoneNumbers

	t.Logf("person set fields:%v\r\n", person.String())

	//Get object fields
	t.Log("GetId() return: ", person.GetId())
	t.Log("Id return: ", person.Id)

	//Marshal
	buffer, err := proto.Marshal(person)

	if err != nil {
		t.Log("Marshal failed! error:", err.Error())
		return
	} else {
		t.Log("Marshal successs! buffer:", buffer)
	}

	//Unmarshal
	personUmmarshal := &Person{}

	err = proto.Unmarshal(buffer, personUmmarshal)

	if err != nil {
		t.Log("Unmarshal failed! error:", err.Error())
		return
	} else {
		t.Log("Unmarshal success! person:", personUmmarshal.String())
	}
}
