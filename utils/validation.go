//=============================================================================
// developer: boxlesslabsng@gmail.com
// Pagination library
//=============================================================================
 
/**
 **
 * @struct ValidateUtil
 **
 * @init() intiliaze translators for gp-validator
 * @Validate() Returns an error object
**/

package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")

	validate = validator.New()
	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Println(err)
	}
}

type ValidateUtil struct {}

func (util ValidateUtil) Validate(data interface{}) map[string]string {
	err := validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	err = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} is not a valid email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})

	err = validate.Struct(data)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		var result = map[string]string{}

		for _, e := range errs {
			result[e.Field()] = e.Translate(trans)

		}

		return result
	}
	return nil
}

func (util ValidateUtil) ValidateParam(ctx echo.Context, param string, result Result) (primitive.ObjectID, error) {
	if (ctx.Param(param) == "") {
		return primitive.ObjectID{}, errors.New(fmt.Sprintf("%s is required", param))
	}

	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id, nil
}

func (util ValidateUtil) ValidateQueryParam(ctx echo.Context, param string, result Result) (primitive.ObjectID, error) {
	if (ctx.QueryParam(param) == "") {
		return primitive.ObjectID{}, errors.New(fmt.Sprintf("%s is required", param))
	}

	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id, nil
}
