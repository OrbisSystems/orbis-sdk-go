package validator

import (
	"github.com/OrbisSystems/orbis-sdk-go/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Validator is a helpful service for data validation.
// It uses ozzo-validation library.
type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateLoginByEmailRequest(input model.LoginByEmailRequest) error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Email, validation.Required, is.Email),
		validation.Field(&input.Password, validation.Required),
	)
}

func (v *Validator) ValidateLoginByAPIKeyRequest(input model.LoginByAPIKeyRequest) error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.APIKey, validation.Required),
		validation.Field(&input.APISecret, validation.Required),
	)
}

func (v *Validator) ValidateNewsFilterRequest(filter model.NewsFilterRequest) error { // nolint:gocognit
	return validation.ValidateStruct(&filter,
		validation.Field(&filter.Symbol, validation.When(filter.Symbol != nil, validation.Required)),
		validation.Field(&filter.NewsSource, validation.Each(validation.Required)),
		validation.Field(&filter.StartDate, validation.When(filter.StartDate != nil, validation.Required)),
		validation.Field(&filter.EndDate, validation.When(filter.EndDate != nil, validation.Required)),
		validation.Field(&filter.Language, validation.When(filter.Language != nil, validation.Required)),
		validation.Field(&filter.RelevanceLevel, validation.When(filter.RelevanceLevel != nil, validation.By(
			func(_ interface{}) error {
				return validation.ValidateStruct(&filter,
					validation.Field(&filter.RelevanceLevel, validation.Required),
					validation.Field(&filter.Subject, validation.Required, validation.Each(validation.Required)))
			}))),
		validation.Field(&filter.Subject, validation.When(filter.Subject != nil, validation.By(
			func(_ interface{}) error {
				return validation.ValidateStruct(&filter,
					validation.Field(&filter.RelevanceLevel, validation.Required),
					validation.Field(&filter.Subject, validation.Required))
			}))),
		validation.Field(&filter.Paging, validation.When(filter.Paging != nil, validation.By(
			func(_ interface{}) error {
				p := *filter.Paging
				return validation.ValidateStruct(&p,
					validation.Field(&p.Limit, validation.When(p.Limit != nil, validation.Required, validation.Min(1))),
					validation.Field(&p.Offset, validation.When(p.Offset != nil, validation.Min(0))))
			}))),
	)
}

func (v *Validator) IsUUID(id string) error {
	return validation.Validate(&id, is.UUIDv4)
}

func (v *Validator) NotEmptyString(value string) error {
	return validation.Validate(&value, validation.Required)
}
