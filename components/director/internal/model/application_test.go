package model_test

import (
	"fmt"
	"testing"

	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplication_AddLabel(t *testing.T) {
	// given
	testCases := []struct {
		Name               string
		InitialApplication model.Application
		InputKey           string
		InputValues        []string
		ExpectedLabels     map[string][]string
	}{
		{
			Name: "New Label",
			InitialApplication: model.Application{
				Labels: map[string][]string{
					"test": {"testVal"},
				},
			},
			InputKey:    "foo",
			InputValues: []string{"bar", "baz", "bar"},
			ExpectedLabels: map[string][]string{
				"test": {"testVal"},
				"foo":  {"bar", "baz"},
			},
		},
		{
			Name: "Nil map",
			InitialApplication: model.Application{
				Labels: nil,
			},
			InputKey:    "foo",
			InputValues: []string{"bar", "baz"},
			ExpectedLabels: map[string][]string{
				"foo": {"bar", "baz"},
			},
		},
		{
			Name: "Append Values",
			InitialApplication: model.Application{
				Labels: map[string][]string{
					"foo": {"bar", "baz"},
				},
			},
			InputKey:    "foo",
			InputValues: []string{"zzz", "bar"},
			ExpectedLabels: map[string][]string{
				"foo": {"bar", "baz", "zzz"},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, testCase.Name), func(t *testing.T) {
			app := testCase.InitialApplication

			// when

			app.AddLabel(testCase.InputKey, testCase.InputValues)

			// then

			for key, val := range testCase.ExpectedLabels {
				assert.ElementsMatch(t, val, app.Labels[key])
			}
		})
	}

}

func TestApplication_DeleteLabel(t *testing.T) {
	// given
	testCases := []struct {
		Name                string
		InputApplication    model.Application
		InputKey            string
		InputValuesToDelete []string
		ExpectedLabels      map[string][]string
		ExpectedErr         error
	}{
		{
			Name:     "Whole Label",
			InputKey: "foo",
			InputApplication: model.Application{
				Labels: map[string][]string{
					"no":  {"delete"},
					"foo": {"bar", "baz"},
				},
			},
			InputValuesToDelete: []string{},
			ExpectedErr:         nil,
			ExpectedLabels: map[string][]string{
				"no": {"delete"},
			},
		},
		{
			Name:     "Label Values",
			InputKey: "foo",
			InputApplication: model.Application{
				Labels: map[string][]string{
					"no":  {"delete"},
					"foo": {"foo", "bar", "baz"},
				},
			},
			InputValuesToDelete: []string{"bar", "baz"},
			ExpectedErr:         nil,
			ExpectedLabels: map[string][]string{
				"no":  {"delete"},
				"foo": {"foo"},
			},
		},
		{
			Name:     "Error",
			InputKey: "foobar",
			InputApplication: model.Application{
				Labels: map[string][]string{
					"no": {"delete"},
				},
			},
			InputValuesToDelete: []string{"bar", "baz"},
			ExpectedErr:         fmt.Errorf("label %s doesn't exist", "foobar"),
			ExpectedLabels: map[string][]string{
				"no": {"delete"},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, testCase.Name), func(t *testing.T) {
			app := testCase.InputApplication

			// when

			err := app.DeleteLabel(testCase.InputKey, testCase.InputValuesToDelete)

			// then

			require.Equal(t, testCase.ExpectedErr, err)

			for key, val := range testCase.ExpectedLabels {
				assert.ElementsMatch(t, val, app.Labels[key])
			}
		})
	}
}

func TestApplication_AddAnnotation(t *testing.T) {
	// given
	testCases := []struct {
		Name                string
		InputApplication    model.Application
		InputKey            string
		InputValue          string
		ExpectedAnnotations map[string]interface{}
		ExpectedErr         error
	}{
		{
			Name:       "Success",
			InputKey:   "foo",
			InputValue: "bar",
			InputApplication: model.Application{
				Annotations: map[string]interface{}{
					"test": "val",
				},
			},
			ExpectedErr: nil,
			ExpectedAnnotations: map[string]interface{}{
				"test": "val",
				"foo":  "bar",
			},
		},
		{
			Name:       "Nil map",
			InputKey:   "foo",
			InputValue: "bar",
			InputApplication: model.Application{
				Annotations: nil,
			},
			ExpectedErr: nil,
			ExpectedAnnotations: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			Name:       "Error",
			InputKey:   "foo",
			InputValue: "bar",
			InputApplication: model.Application{
				Annotations: map[string]interface{}{
					"foo": "val",
				},
			},
			ExpectedErr: fmt.Errorf("annotation %s does already exist", "foo"),
			ExpectedAnnotations: map[string]interface{}{
				"foo": "val",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, testCase.Name), func(t *testing.T) {
			app := testCase.InputApplication

			// when

			err := app.AddAnnotation(testCase.InputKey, testCase.InputValue)

			// then

			require.Equal(t, testCase.ExpectedErr, err)
			assert.Equal(t, testCase.ExpectedAnnotations, app.Annotations)
		})
	}
}

func TestApplication_DeleteAnnotation(t *testing.T) {
	// given
	testCases := []struct {
		Name                string
		InputApplication    model.Application
		InputKey            string
		ExpectedAnnotations map[string]interface{}
		ExpectedErr         error
	}{
		{
			Name:     "Success",
			InputKey: "foo",
			InputApplication: model.Application{
				Annotations: map[string]interface{}{
					"no":  "delete",
					"foo": "bar",
				},
			},
			ExpectedErr: nil,
			ExpectedAnnotations: map[string]interface{}{
				"no": "delete",
			},
		},
		{
			Name:     "Error",
			InputKey: "foobar",
			InputApplication: model.Application{
				Annotations: map[string]interface{}{
					"no": "delete",
				},
			},
			ExpectedErr: fmt.Errorf("annotation %s doesn't exist", "foobar"),
			ExpectedAnnotations: map[string]interface{}{
				"no": "delete",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, testCase.Name), func(t *testing.T) {
			app := testCase.InputApplication

			// when

			err := app.DeleteAnnotation(testCase.InputKey)

			// then

			require.Equal(t, testCase.ExpectedErr, err)
			assert.Equal(t, testCase.ExpectedAnnotations, app.Annotations)
		})
	}
}

func TestApplicationInput_ToApplication(t *testing.T) {
	// given
	url := "https://foo.bar"
	desc := "Sample"
	id := "foo"
	tenant := "sample"
	testCases := []struct {
		Name     string
		Input    *model.ApplicationInput
		Expected *model.Application
	}{
		{
			Name: "All properties given",
			Input: &model.ApplicationInput{
				Name:        "Foo",
				Description: &desc,
				Annotations: map[string]interface{}{
					"key": "value",
				},
				Labels: map[string][]string{
					"test": {"val", "val2"},
				},
				HealthCheckURL: &url,
			},
			Expected: &model.Application{
				Name:        "Foo",
				ID:          id,
				Tenant:      tenant,
				Description: &desc,
				Annotations: map[string]interface{}{
					"key": "value",
				},
				Labels: map[string][]string{
					"test": {"val", "val2"},
				},
				HealthCheckURL: &url,
			},
		},
		{
			Name:     "Nil",
			Input:    nil,
			Expected: nil,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, testCase.Name), func(t *testing.T) {

			// when
			result := testCase.Input.ToApplication(id, tenant)

			// then
			assert.Equal(t, testCase.Expected, result)
		})
	}
}
