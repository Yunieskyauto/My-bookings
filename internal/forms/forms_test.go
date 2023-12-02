package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form show valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "b")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form shows does not has field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("X", 3)
	if form.Valid() {
		t.Error("form shows minlength for non-existent field ")
	}

	isError := form.Errors.Get("X")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength  of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should have an error but did not get one")
	}

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("X")
	if form.Valid() {
		t.Error("form shows valid email for non existing field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}

}
