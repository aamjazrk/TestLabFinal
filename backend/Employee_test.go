package backend

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName  string    `valid:"required~FirstName cannot be null,alpha~FirstName must have only character,minstringlength(10)~FirstName need to have min length 10"`
	MiddleName string    `valid:"character_only~MiddleName must be only character,required~MiddleName cannot be null,maxstringlength(20)~FirstName need to have max length 20"`
	LastName   string    `valid:"matches([a-zA-Z]$)~LastName must be only character,required~LastName cannot be null,stringlength(10|20)~LastName must be in length 10-20"`
	Email      string    `valid:"email~Emial is not rigth format,required~Email cannot be null"`
	Age        uint      `valid:"numeric~Age must be number,range(18|120)~Age must be in range 1-120"`
	Phone      string    `valid:"matches(^(0)([0-9]{9})$)~Phone must have only character and length is 10"`
	Low string `valid:"lowercase~Lower Case Only"`
	Upper string `valid:"uppercase~Upper Case Only"` 
	BirthDate  time.Time `valid:"datenotfuture~BirthDate cannot be future"`
}
func TestUpAndDown(t *testing.T){
	g := gomega.NewGomegaWithT(t)
	date1 := time.Date(1999,2,5,0,0,0,0,time.Local)
	t.Run("Lowercase",func(t *testing.T){
	emp := Employee{
		FirstName:  "sdfghjasdfgh",
			MiddleName: "Aj",
			LastName:   "kotpanyakkk",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
			Low: "HKg",
			Upper: "PPPPP",
	}
	ok, err := govalidator.ValidateStruct(emp)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Lower Case Only"))
	})

	t.Run("Uppercase",func(t *testing.T){
		emp := Employee{
			FirstName:  "sdfghjasdfgh",
			MiddleName: "Aj",
			LastName:   "kotpanyakkk",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
			Low: "kkkg",
			Upper: "lllllllluuuuuo",
		}
		ok,err := govalidator.ValidateStruct(emp)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Upper Case Only"))
	})
	
}
func Test_EmployeePass_Test(t *testing.T) {
	date1 := time.Date(2002, 4, 27, 14, 50, 45, 659000, time.Local)
	g := gomega.NewGomegaWithT(t)
	emp := Employee{
		FirstName:  "sirinyakkkkkk",
		MiddleName: "Aj",
		LastName:   "kotpanyakkkkkkkk",
		Email:      "aam@mail.com",
		Age:        20,
		Phone:      "0624563333",
		BirthDate:  date1,
	}
	ok, err := govalidator.ValidateStruct(emp)
	//g.Expect(err.Error()).To(gomega.Equal("FirstName cannot be blank"))
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}
func TestEmpNotNull(t *testing.T) {
	date1 := time.Date(2002, 4, 27, 14, 50, 45, 659000, time.Local)
	g := gomega.NewGomegaWithT(t)
	t.Run("FirstName cannot be null", func(t *testing.T) {
		emp := Employee{
			FirstName:  "",
			MiddleName: "Aj",
			LastName:   "kotpanyakkk",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("FirstName cannot be null"))
	})

	t.Run("MiddleName cannot be null", func(t *testing.T) {
		emp := Employee{
			FirstName:  "Sirinyakkkk",
			MiddleName: "",
			LastName:   "Kotpanyakkkkk",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("MiddleName cannot be null"))
	})

	t.Run("LastName cannot be null", func(t *testing.T) {
		emp := Employee{
			FirstName:  "Sirinyajjjjj",
			MiddleName: "AJ",
			LastName:   "",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("LastName cannot be null"))
	})

	t.Run("Email cannot be null", func(t *testing.T) {
		emp := Employee{
			FirstName:  "Sirinyllllllla",
			MiddleName: "AJ",
			LastName:   "KKkkkkkkkkkkkkk",
			Email:      "",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Email cannot be null"))
	})
	// t.Run("BirthDate cannot be null", func(t *testing.T) {
	// 	emp := Employee{
	// 		FirstName:  "Sirinya",
	// 		MiddleName: "AJ",
	// 		LastName:   "KK",
	// 		Email:      "amm@mail.com",
	// 		Age:        20,
	// 	}
	// 	ok, err := govalidator.ValidateStruct(emp)
	// 	g.Expect(ok).ToNot(gomega.BeTrue())
	// 	g.Expect(err).ToNot(gomega.BeNil())
	// 	g.Expect(err.Error).To(gomega.Equal("BirthDate cannot be null"))
	// })

}
func TestEmployeeNotvalit(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	date1 := time.Date(2007, 3, 27, 18, 40, 20, 345000, time.Local)
	t.Run("FirstName must be only character", func(t *testing.T) {
		emp := Employee{
			FirstName:  "122kkkdkkkk",
			MiddleName: "Aj",
			LastName:   "kotpanyakkkk",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
		}
		ok, err := govalidator.ValidateStruct(emp)

		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("FirstName must have only character"))
	})
	t.Run("MiddleName must be only character", func(t *testing.T) {
		emp := Employee{
			FirstName:  "GGdlllllllll",
			MiddleName: "Aj11112445",
			LastName:   "kotpanyalllllllll",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("MiddleName must be only character"))
	})

	t.Run("LastName must be only character", func(t *testing.T){
		emp := Employee{
			FirstName:  "GGdlllllllll",
			MiddleName: "Ajppppjjjjjj",
			LastName:   "kotpanyalllllllllสสสส",
			Email:      "aam@mail.com",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("LastName must be only character"))
	})
}
func TestEmail(t *testing.T){
	g := gomega.NewGomegaWithT(t)
	date1 := time.Date(1999,12,05,00,00,00,00000000,time.Local)
	emp := Employee{
		FirstName:  "GGdlllllllll",
			MiddleName: "Ajppppppkkk",
			LastName:   "kotpanyalllllllll",
			Email:      "aam@",
			Age:        20,
			Phone:      "0624563333",
			BirthDate:  date1,
	}
	ok, err := govalidator.ValidateStruct(emp)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Emial is not rigth format"))
}
func TestPhone(t *testing.T){
	g := gomega.NewGomegaWithT(t)
	date1 := time.Date(1999,10,5,0,0,0,0,time.Local)
	emp := Employee{
		FirstName:  "GGdlllllllll",
		MiddleName: "Ajpppppp",
		LastName:   "kotpanyalllllllll",
		Email:      "aam@mail.com",
		Age:        20,
		Phone:      "06245633",
		BirthDate:  date1,
	}
	ok, err := govalidator.ValidateStruct(emp)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Phone must have only character and length is 10"))
}
func init() {
	govalidator.CustomTypeTagMap.Set("character_only", func(i interface{}, context interface{}) bool {
		s, ok := i.(string)
		if !ok {
			return false
		}
		for _, c := range s {
			if !(('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('ก' <= c && 'ฮ' <= c) || ('ะ' <= c && c <= 'ู') || ('เ' <= c && c <= '์')) {
				return false
			}
		}
		return true
	})
	govalidator.CustomTypeTagMap.Set("datenotfuture", func(i interface{}, context interface{}) bool {
		date := i.(time.Time)
		datenow := time.Now()
		if date.After(datenow) {
			return false
		} else {
			return true
		}
	})

}
