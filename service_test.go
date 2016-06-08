package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestWhenIInsertValidAlert_ShouldReturnOK(t *testing.T) {
	save = MockSave

	a, err := CreateAlert(SampleAlert)
	assert.Equal(t, SampleAlert.Name, a.Name, "Create should return OK for valid return")
	assert.Equal(t, nil, err, "Create should not return OK for valid return")
}

func TestWhenIInsertAlertWithoutName_ShouldReturnError(t *testing.T) {
	save = MockSave
	alertWithoutName := SampleAlert
	alertWithoutName.Name = ""

	a, err := CreateAlert(alertWithoutName)
	assert.Equal(t, VALIDATION_MESSAGE_NAME_EMPTY, err.Error(), "Create alert without name should return error")
	assert.Equal(t, alertWithoutName, a, "Create alert without name should return error")
}

func TestWhenIInsertAlertWithNameLengthIsMoreThanMax_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.Name = "ayse jkajshdkjashdjsahd kashdjashdkjahsdkjhaskjdhaksjhdkjashd jashdkjahsdkjahsdjhsd"

	a, err := CreateAlert(alert)
	assert.Equal(t, VALIDATION_MESSAGE_NAME_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_NAME), err.Error(), "Create alert with name length is more than max should return error")
	assert.Equal(t, alert, a, "Create alert with name length is more than max should return error")
}

func TestWhenIInsertAlertWithoutOwnerID_ShouldReturnError(t *testing.T) {
	save = MockSave
	alertWithoutOwnerID := SampleAlert
	alertWithoutOwnerID.OwnerID = -2

	a, err := CreateAlert(alertWithoutOwnerID)
	assert.Equal(t, VALIDATION_MESSAGE_OWNER_ID, err.Error(), "Create alert without owner id should return error")
	assert.Equal(t, alertWithoutOwnerID, a, "Create alert without owner id should return error")
}

func TestWhenIInsertTurkishCharacterForCriteria_ShouldReturnOk(t *testing.T) {
	save = MockSave

	alert := SampleAlert
	alert.RequiredCriteria = "ayçe çç öö ğ ü ı şşşşşşş"

	a, err := CreateAlert(SampleAlert)
	assert.Equal(t, SampleAlert.Name, a.Name, "Create alert with turkish character for criteria should return OK")
	assert.Nil(t, err, nil, "Create alert with turkish character for criteria should not return OK")
}

func TestWhenIInsertNonAlphanumericCharacterForCriteria_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.RequiredCriteria = ">>> < | ~~~ ]"

	a, err := CreateAlert(alert)
	assert.Equal(t, VALIDATION_MESSAGE_CRITERIA_PHRASES_ALPHANUMERIC, err.Error(), "Create alert with non alphanumeric creiteria should return error")
	assert.Equal(t, alert, a, "Create alert with  non alphanumeric creiteria should return error")
}

func TestWhenIInsertRequiredCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.RequiredCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.Equal(t, VALIDATION_MESSAGE_REQUIRED_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long required criteria should return error")
	assert.Equal(t, alert, a, "Create alert with long required criteria should return error")
}

func TestWhenIInsertNiceToHaveCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.NiceToHaveCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.Equal(t, VALIDATION_MESSAGE_NICE_TO_HAVE_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long nice-to-have criteria should return error")
	assert.Equal(t, alert, a, "Create alert with long nice-to-have criteria should return error")
}

func TestWhenIInsertExcludedCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.ExcludedCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.Equal(t, VALIDATION_MESSAGE_EXCLUDED_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long excluded criteria should return error")
	assert.Equal(t, alert, a, "Create alert with long excluded criteria should return error")
}

func TestWhenIInsertAlertWithInvalidThreshold_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.Threshold = 2000000

	a, err := CreateAlert(alert)
	assert.Equal(t, VALIDATION_MESSAGE_THRESHOLD, err.Error(), "Create alert with greater than max valued threshold should return error")
	assert.Equal(t, alert, a, "Create alert with greater than max valued threshold should return error")
}

func TestWhenIProvideAlertId_ShouldReturnSampleAlert(t *testing.T) {
	find = MockFind
	id := SampleAlert.ID
	a, err := FindAlert(id.String())

	assert.Equal(t, nil, err, "Find alert returns nil error when id provided")
	assert.Equal(t, SampleAlert.Name, a.Name, "Find alert by id should return alert when id provided")
}

func TestWhenIProvideEmptyAlertId_ShouldReturnError(t *testing.T) {
	find = MockFind
	id := ""
	_, err := FindAlert(id)

	assert.Equal(t, "id should be provided", err.Error(), "Find alert by id should return error when id not provided")
}
