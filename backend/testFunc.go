package backend

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.CustomTypeTagMap.Set("charactorandnum", func(i interface{}, complex interface{}) bool {
		s, ok := i.(string)
		if !ok {
			return false
		}
		for _, c := range s {
			if !(('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('ก' <= c && c <= 'ฮ') || ('ะ' <= c && c <= 'ู') || ('เ' <= c && c <= '์') ||
				('0' <= c && c <= '9') || ('๐' <= c && c <= '๙') || (c == ' ')) {
				return false
			}

		}
		return true
	})

	govalidator.CustomTypeTagMap.Set("notFuter", func(i interface{}, complex interface{} )bool{
		timenew,ok := i.(time.Time)
		if !ok {
			return false
		}
		return timenew.After(time.Now())
	} )

	govalidator.CustomTypeTagMap.Set("TimeNotPast", func(i interface{}, complex interface{}) bool{
		timenew,ok := i.(time.Time)

		if !ok {
			return false
		}
		return timenew.Before(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("CurrentTime" , func(i interface{}, complex interface{}) bool {
		timenew, ok := i.(time.Time)
		if !ok{
			return false
		}
		return timenew.Equal(time.Now())
	})
}
