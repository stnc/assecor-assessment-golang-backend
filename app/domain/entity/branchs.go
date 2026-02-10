package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

// Branches strcut
type Branches struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null;" json:"name" validate:"required"`
	LastName  string     `gorm:"type:varchar(255);not null;" json:"last_name" validate:"required"`
	ZipCode   string     `gorm:"type:text;" json:"zip_code"`
	City      string     `gorm:"type:varchar(255);not null;" json:"city" validate:"required"`
	Color     string     `gorm:"type:varchar(255);not null;" json:"color" validate:"required"`
	CreatedAt time.Time  ` json:"created_at"`
	UpdatedAt time.Time  ` json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	ZipCode  string `json:"zipcode"`
	City     string `json:"city"`
	Color    string `json:"color"`
}

/*
type BranchSaveDTO struct {
	ID               uint64 `gorm:"primary_key;auto_increment" json:"id"`
	ParentCategoryID uint64 `gorm:"not null;DEFAULT:'0'" json:"parent_category_Id"`
	Title            string `gorm:"size:100 ;not null;" json:"Title" validate:"required"`
	Description      string `gorm:"type:text ;" json:"description"`
	Slug             string `gorm:"size:255 ;null;" json:"type"`
	SelectedID       uint64
	PostType         int `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required"`
}
*/
// BeforeSave init
func (f *Branches) BeforeSave() {
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
}

// Prepare init
func (f *Branches) Prepare() {
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

// Validate fluent validation
func (f *Branches) Validate() map[string]string {
	var (
		validate *validator.Validate
		uni      *ut.UniversalTranslator
	)
	tr := en.New()
	uni = ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")
	validate = validator.New()
	tr_translations.RegisterDefaultTranslations(validate, trans)

	errorLog := make(map[string]string)

	err := validate.Struct(f)
	fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			// can translate each error one at a time.
			lng := strings.Replace(e.Translate(trans), e.Field(), "Burası", 1)
			errorLog[e.Field()+"_error"] = e.Translate(trans)
			// errorLog[e.Field()] = e.Translate(trans)
			errorLog[e.Field()] = lng
			errorLog[e.Field()+"_valid"] = "is-invalid"
		}
	}
	return errorLog
}
