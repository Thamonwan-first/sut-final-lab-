package entity

import (
	"testing"
	."github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func TestNegative1(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Negative1",func (t *testing.T){
		employees := Employees{
			Name : "John",
			Salary: 10000,
			EmployeeCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}