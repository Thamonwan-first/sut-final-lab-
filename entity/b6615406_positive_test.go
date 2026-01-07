package entity

import (
	"testing"
	."github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func TestPositive(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("positive",func (t *testing.T){
		employees := Employees{
			Name : "John",
			Salary: 20000,
			EmployeeCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}