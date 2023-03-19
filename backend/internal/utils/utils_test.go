package utils

import "testing"

type randomTestType struct {
	N int    `json:"n"`
	S string `json:"s"`
}

func TestUnmarshalJsonResponse(t *testing.T) {
	jsonString := `{"data":"test","errorText":""}`
	stringResponse, err := UnmarshalJsonResponse[string]([]byte(jsonString))
	CheckTestError(t, err, "Error while unmarshaling string json")
	AssertEqual(t, "test", stringResponse.Data, "UnmarshalJsonResponse did return the correct value")
	AssertEqual(t, "", stringResponse.ErrorText, "UnmarshalJsonResponse did return the correct value")

	jsonString = `{"data":[5, 9, 6],"errorText":"a"}`
	intResponse, err := UnmarshalJsonResponse[[]int]([]byte(jsonString))
	CheckTestError(t, err, "Error while unmarshaling array int json")
	AssertEqual(t, 5, intResponse.Data[0], "UnmarshalJsonResponse did return the correct value")
	AssertEqual(t, 9, intResponse.Data[1], "UnmarshalJsonResponse did return the correct value")
	AssertEqual(t, 6, intResponse.Data[2], "UnmarshalJsonResponse did return the correct value")
	AssertEqual(t, "a", intResponse.ErrorText, "UnmarshalJsonResponse did return the correct value")

	jsonString = `{"data":{"n":0,"s":"test"},"errorText":""}`
	randomTestTypeResponse, err := UnmarshalJsonResponse[randomTestType]([]byte(jsonString))
	CheckTestError(t, err, "Error while unmarshaling randomTestType json")
	AssertEqual(t, 0, randomTestTypeResponse.Data.N, "UnmarshalJsonResponse did return the correct value")
	AssertEqual(t, "test", randomTestTypeResponse.Data.S, "UnmarshalJsonResponse did return the correct value")
	AssertEqual(t, "", randomTestTypeResponse.ErrorText, "UnmarshalJsonResponse did return the correct value")
}

func TestMarshalJsonResponse(t *testing.T) {
	expected := `{"data":"test","errorText":""}`
	jsonString, err := MarshalJsonResponse("test")
	CheckTestError(t, err, "Error while marshaling string json")
	AssertEqual(t, expected, string(jsonString), "MarshalJsonResponse did return the correct value")

	expected = `{"data":5,"errorText":""}`
	jsonString, err = MarshalJsonResponse(5)
	CheckTestError(t, err, "Error while marshaling string json")
	AssertEqual(t, expected, string(jsonString), "MarshalJsonResponse did return the correct value")

	expected = `{"data":{"n":17,"s":"test"},"errorText":""}`
	jsonString, err = MarshalJsonResponse(randomTestType{
		N: 17,
		S: "test",
	})
	CheckTestError(t, err, "Error while marshaling string json")
	AssertEqual(t, expected, string(jsonString), "MarshalJsonResponse did return the correct value")
}

func TestMarshalJsonErrorResponse(t *testing.T) {
	expected := `{"data":"","errorText":"Error text"}`
	jsonString, err := MarshalJsonErrorResponse("Error text")
	CheckTestError(t, err, "Error while marshaling error json")
	AssertEqual(t, expected, string(jsonString), "MarshalJsonErrorResponse did return the correct value")

	expected = `{"data":"","errorText":"Error text 2"}`
	jsonString, err = MarshalJsonErrorResponse("Error text 2")
	CheckTestError(t, err, "Error while marshaling error json")
	AssertEqual(t, expected, string(jsonString), "MarshalJsonErrorResponse did return the correct value")
}
