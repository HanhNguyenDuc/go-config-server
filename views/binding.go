package views

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
)

func ToString(obj interface{}) string {
	dispArr := []string{}
	v := reflect.ValueOf(obj)
	fields := make([]interface{}, v.NumField())
	for i := range fields {
		if v.Field(i).Interface() != nil {
			fieldName := v.Type().Field(i).Name
			switch v.Field(i).Interface().(type) {
			case *int:
				ptrToObj := v.Field(i).Interface().(*int)
				if ptrToObj != nil {
					dispArr = append(dispArr, fmt.Sprintf("%s: %v", fieldName, *v.Field(i).Interface().(*int)))
				}
			case *string:
				ptrToObj := v.Field(i).Interface().(*string)
				if ptrToObj != nil {
					dispArr = append(dispArr, fmt.Sprintf("%s: %v", fieldName, *v.Field(i).Interface().(*string)))
				}
			case *multipart.FileHeader:
				ptrToObj := v.Field(i).Interface().(*multipart.FileHeader)

				if ptrToObj != nil {
					fileName := ptrToObj.Filename
					if ptrToObj != nil {
						dispArr = append(dispArr, fmt.Sprintf("%s: %v", fieldName, fileName))
					}
				}

			}
		}
	}
	return "[" + strings.Join(dispArr, " | ") + "]"
}

type GetConfigRequest struct {
	ID          *string `form:"id" binding:"required"`
	CurrentTime *int64  `form:"current_time" binding:"required"`
	Service     *string `form:"service" binding:"required"`
}
