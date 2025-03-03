package setup

import (
	"log"
	"reflect"
	"strings"

	"github.com/ahpehgit/golang-boilerplate/config"
	validators "github.com/ahpehgit/golang-boilerplate/validators"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/spf13/viper"
)

func SetupConfig() {
	//* Set the file name of the configurations file
	viper.SetConfigName("config")

	//* Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config.Configuration)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}
}

func RegisterValidators() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("email-format", validators.ValidateEmailAddressFormat)
	validate.RegisterValidation("script-injection", validators.ValidateStringForScriptInjection)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}

func RegisterTranslations(validate *validator.Validate) *ut.Translator {
	translator := en.New()
	uni := ut.New(translator, translator)

	transTemp, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("Translator not found")
	}

	translation := &transTemp
	if err := en_translations.RegisterDefaultTranslations(validate, *translation); err != nil {
		log.Fatal(err)
	}

	validators.AddTranslations(validate, translation)

	return translation
}
