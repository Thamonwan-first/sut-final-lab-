package entity

import (
	"testing"
	."github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func TestNegative2(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Negative2",func (t *testing.T){
		employees := Employees{
			Name : "John",
			Salary: 20000,
			EmployeeCode: "HR1024",
		}
		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
	})
}