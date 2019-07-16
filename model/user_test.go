package model

import (
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"gopkg.in/mgo.v2/bson"
)

var testModel = []struct {
	in            User
	out           bool
	outName       string
	outEmail      string
	outImgProfile string
}{
	{
		in: User{
			bson.NewObjectId(),
			"",
			"Testname1 Testlastname1",
			"test1@testmail.test",
			"testImageProfile1.jpg",
			"slack",
			"0812345678",
			time.Now(),
		},
		out:           false,
		outName:       "Testname1 Testlastname1",
		outEmail:      "test1@testmail.test",
		outImgProfile: "testImageProfile1.jpg",
	},
	{
		in: User{
			bson.NewObjectId(),
			"individual",
			"Testname2 Testlastname2",
			"test2@testmail.test",
			"testImageProfile2.jpg",
			"slack",
			"0812345678",
			time.Now(),
		},
		out:           false,
		outName:       "Testname2 Testlastname2",
		outEmail:      "test2@testmail.test",
		outImgProfile: "testImageProfile2.jpg",
	},
	{
		in: User{
			bson.NewObjectId(),
			"admin",
			"Testname3 Testlastname3",
			"test3@testmail.test",
			"testImageProfile3.jpg",
			"slack",
			"0812345678",
			time.Now(),
		},
		out:           true,
		outName:       "Testname3 Testlastname3",
		outEmail:      "test3@testmail.test",
		outImgProfile: "testImageProfile3.jpg",
	},
}

func TestIsAdmin(t *testing.T) {
	for _, _test := range testModel {
		assert.Equal(t, _test.out, _test.in.IsAdmin())
	}
}

func TestGetName(t *testing.T) {
	for _, _test := range testModel {
		assert.Equal(t, _test.outName, _test.in.GetName())
	}
}

func TestGetEmail(t *testing.T) {
	for _, _test := range testModel {
		assert.Equal(t, _test.outEmail, _test.in.GetEmail())
	}
}

func TestGetImgProfile(t *testing.T) {
	for _, _test := range testModel {
		assert.Equal(t, _test.outImgProfile, _test.in.GetImgProfile())
	}
}
