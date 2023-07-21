// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/eesoymilk/health-statistic-api/ent/answer"
	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/price"
	"github.com/eesoymilk/health-statistic-api/ent/question"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaire"
	"github.com/eesoymilk/health-statistic-api/ent/questionnaireresponse"
	"github.com/eesoymilk/health-statistic-api/ent/schema"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescCreatedAt is the schema descriptor for created_at field.
	answerDescCreatedAt := answerFields[1].Descriptor()
	// answer.DefaultCreatedAt holds the default value on creation for the created_at field.
	answer.DefaultCreatedAt = answerDescCreatedAt.Default.(func() time.Time)
	// answerDescID is the schema descriptor for id field.
	answerDescID := answerFields[0].Descriptor()
	// answer.DefaultID holds the default value on creation for the id field.
	answer.DefaultID = answerDescID.Default.(func() uuid.UUID)
	mycardFields := schema.MyCard{}.Fields()
	_ = mycardFields
	// mycardDescCardNumber is the schema descriptor for card_number field.
	mycardDescCardNumber := mycardFields[0].Descriptor()
	// mycard.CardNumberValidator is a validator for the "card_number" field. It is called by the builders before save.
	mycard.CardNumberValidator = func() func(string) error {
		validators := mycardDescCardNumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(card_number string) error {
			for _, fn := range fns {
				if err := fn(card_number); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mycardDescCardPassword is the schema descriptor for card_password field.
	mycardDescCardPassword := mycardFields[1].Descriptor()
	// mycard.CardPasswordValidator is a validator for the "card_password" field. It is called by the builders before save.
	mycard.CardPasswordValidator = func() func(string) error {
		validators := mycardDescCardPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(card_password string) error {
			for _, fn := range fns {
				if err := fn(card_password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mycardDescCreatedAt is the schema descriptor for created_at field.
	mycardDescCreatedAt := mycardFields[2].Descriptor()
	// mycard.DefaultCreatedAt holds the default value on creation for the created_at field.
	mycard.DefaultCreatedAt = mycardDescCreatedAt.Default.(func() time.Time)
	priceFields := schema.Price{}.Fields()
	_ = priceFields
	// priceDescCreatedAt is the schema descriptor for created_at field.
	priceDescCreatedAt := priceFields[2].Descriptor()
	// price.DefaultCreatedAt holds the default value on creation for the created_at field.
	price.DefaultCreatedAt = priceDescCreatedAt.Default.(func() time.Time)
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescID is the schema descriptor for id field.
	questionDescID := questionFields[0].Descriptor()
	// question.DefaultID holds the default value on creation for the id field.
	question.DefaultID = questionDescID.Default.(func() uuid.UUID)
	questionnaireFields := schema.Questionnaire{}.Fields()
	_ = questionnaireFields
	// questionnaireDescCreatedAt is the schema descriptor for created_at field.
	questionnaireDescCreatedAt := questionnaireFields[2].Descriptor()
	// questionnaire.DefaultCreatedAt holds the default value on creation for the created_at field.
	questionnaire.DefaultCreatedAt = questionnaireDescCreatedAt.Default.(func() time.Time)
	// questionnaireDescID is the schema descriptor for id field.
	questionnaireDescID := questionnaireFields[0].Descriptor()
	// questionnaire.DefaultID holds the default value on creation for the id field.
	questionnaire.DefaultID = questionnaireDescID.Default.(func() uuid.UUID)
	questionnaireresponseFields := schema.QuestionnaireResponse{}.Fields()
	_ = questionnaireresponseFields
	// questionnaireresponseDescCreatedAt is the schema descriptor for created_at field.
	questionnaireresponseDescCreatedAt := questionnaireresponseFields[1].Descriptor()
	// questionnaireresponse.DefaultCreatedAt holds the default value on creation for the created_at field.
	questionnaireresponse.DefaultCreatedAt = questionnaireresponseDescCreatedAt.Default.(func() time.Time)
	// questionnaireresponseDescID is the schema descriptor for id field.
	questionnaireresponseDescID := questionnaireresponseFields[0].Descriptor()
	// questionnaireresponse.DefaultID holds the default value on creation for the id field.
	questionnaireresponse.DefaultID = questionnaireresponseDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[3].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[4].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescBirthYear is the schema descriptor for birth_year field.
	userDescBirthYear := userFields[5].Descriptor()
	// user.BirthYearValidator is a validator for the "birth_year" field. It is called by the builders before save.
	user.BirthYearValidator = userDescBirthYear.Validators[0].(func(int) error)
	// userDescHeight is the schema descriptor for height field.
	userDescHeight := userFields[6].Descriptor()
	// user.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	user.HeightValidator = userDescHeight.Validators[0].(func(float64) error)
	// userDescWeight is the schema descriptor for weight field.
	userDescWeight := userFields[7].Descriptor()
	// user.WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	user.WeightValidator = userDescWeight.Validators[0].(func(float64) error)
	// userDescMedicalHistory is the schema descriptor for medical_history field.
	userDescMedicalHistory := userFields[12].Descriptor()
	// user.MedicalHistoryValidator is a validator for the "medical_history" field. It is called by the builders before save.
	user.MedicalHistoryValidator = userDescMedicalHistory.Validators[0].(func(string) error)
	// userDescMedicationStatus is the schema descriptor for medication_status field.
	userDescMedicationStatus := userFields[13].Descriptor()
	// user.MedicationStatusValidator is a validator for the "medication_status" field. It is called by the builders before save.
	user.MedicationStatusValidator = userDescMedicationStatus.Validators[0].(func(string) error)
}
