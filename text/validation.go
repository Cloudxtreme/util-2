// Copyright 2013 Felipe Alves Cavani. All rights reserved.
// Start date:		2014-08-08
// Last modification:	2014-x

package text

import (
	"github.com/fcavani/e"
	uni "projects/util/unicode"
	"projects/util/utf8string"
	"regexp"
	"unicode"
)

var MinPassLen = 8
var MaxPassLen = 100
var MinEmailLen = 6
var MaxEmailLen = 60

const ErrInvLengthNumber = "wrong number of digits"
const ErrInvDigit = "caracter isn't a digit"
const ErrInvNumberChars = "invalid number of characters"
const ErrInvCharacter = "invalid character"
const ErrInvalidPassLength = "password length is invalid"
const ErrInvalidPassChar = "invalid password character"
const ErrInvEmailLength = "email length is invalid"
const ErrCantCheckEmail = "can't check the e-mail address"
const ErrInvEmailString = "invalid e-mail address"
const ErrInvalidChar = "character is invalid"

func CheckNumber(number string, min, max int) error {
	if len(number) < min || len(number) > max {
		return e.New(ErrInvLengthNumber)
	}
	for _, v := range number {
		if !unicode.IsDigit(v) {
			return e.New(ErrInvDigit)
		}
	}
	return nil
}

func CheckLettersNumber(text string, min, max int) error {
	if len(text) < min || len(text) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range text {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func CheckText(text string, min, max int) error {
	if len(text) < min || len(text) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range text {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != '\n' && v != ' ' && v != '`' && v != '~' && v != '!' && v != '@' && v != '#' && v != '$' && v != '%' && v != '^' && v != '&' && v != '*' && v != '(' && v != ')' && v != '_' && v != '-' && v != '+' && v != '=' && v != '{' && v != '}' && v != '[' && v != ']' && v != '|' && v != '\\' && v != ':' && v != ';' && v != '"' && v != '\'' && v != '?' && v != '/' && v != ',' && v != '.' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

// Check the user password. Graphics character are allowed. See unicode.IsGraphic.
func CheckPassword(pass string, min, max int) error {
	if len(pass) < min || len(pass) > max {
		return e.New(ErrInvalidPassLength)
	}
	for _, r := range pass {
		if !unicode.IsGraphic(r) {
			return e.New(ErrInvalidPassChar)
		}
	}
	return nil
}

func CheckEmail(email string) error {
	if len(email) < MinEmailLen || len(email) > MaxEmailLen {
		return e.New(ErrInvEmailLength)
	}
	r, err := regexp.Compile(`([a-zA-Z0-9]+)([.-_][a-zA-Z0-9]+)*@([a-zA-Z0-9]+)([.-_][a-zA-Z0-9]+)*`)
	if err != nil {
		return e.Push(e.New(err), ErrCantCheckEmail)
	}
	if email != r.FindString(email) {
		return e.New(ErrInvEmailString)
	}
	return nil
}

func CheckNome(nome string, min, max int) error {
	if len(nome) < min || len(nome) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range nome {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != ' ' && v != '`' && v != '~' && v != '!' && v != '@' && v != '#' && v != '$' && v != '%' && v != '^' && v != '&' && v != '*' && v != '(' && v != ')' && v != '_' && v != '-' && v != '+' && v != '=' && v != '{' && v != '}' && v != '[' && v != ']' && v != '|' && v != '\\' && v != ':' && v != ';' && v != '"' && v != '\'' && v != '?' && v != '/' && v != ',' && v != '.' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func CheckNomeWithoutSpecials(nome string, min, max int) error {
	if len(nome) < min || len(nome) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range nome {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != ' ' && v != '&' && v != '(' && v != ')' && v != '-' && v != ':' && v != '/' && v != ',' && v != '.' && v != '_' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func CheckFileName(nome string, min, max int) error {
	if len(nome) < min || len(nome) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range nome {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != ' ' && v != '-' && v != ':' && v != ',' && v != '.' && v != '_' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' in filename is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func ValidateRedirect(redirect string, min, max int) error {
	utf8name := utf8string.NewString(redirect)
	len := utf8name.RuneCount()
	if len < min || len > max {
		return e.New(e.ErrInvalidLength)
	}
	for _, s := range utf8name.Slice(1, len) {
		if !uni.IsLetter(s) && !unicode.IsDigit(s) && s != '/' && s != '-' && s != '_' && s != '?' && s != '&' && s != '=' && s != '%' && s != '*' && s != '+' && s != ' ' && s != ',' {
			println("redirect:", string([]byte{byte(s)}))
			return e.Push(e.New("the character '%v' in redirect is invalid", string([]byte{byte(s)})), e.New("invalid redirect"))
		}
	}
	return nil
}

func CheckSearch(nome string, min, max int) error {
	if len(nome) < min || len(nome) > max {
		return e.New("invalid length")
	}
	for _, v := range nome {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != '@' && v != '.' && v != '-' && v != '_' && v != ' ' && v != '+' && v != '"' {
			return e.New("invalid character")
		}
	}
	return nil
}
