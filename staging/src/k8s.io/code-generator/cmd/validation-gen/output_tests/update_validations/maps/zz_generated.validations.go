//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by validation-gen. DO NOT EDIT.

package maps

import (
	context "context"
	fmt "fmt"

	operation "k8s.io/apimachinery/pkg/api/operation"
	safe "k8s.io/apimachinery/pkg/api/safe"
	validate "k8s.io/apimachinery/pkg/api/validate"
	field "k8s.io/apimachinery/pkg/util/validation/field"
	testscheme "k8s.io/code-generator/cmd/validation-gen/testscheme"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *testscheme.Scheme) error {
	scheme.AddValidationFunc((*M1)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_M1(ctx, op, nil /* fldPath */, obj.(*M1), safe.Cast[*M1](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc((*T1)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T1(ctx, op, nil /* fldPath */, obj.(*T1), safe.Cast[*T1](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_M1(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *M1) (errs field.ErrorList) {
	// field M1.S
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *string) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "M1.S")...)
			return
		}(fldPath.Child("s"), &obj.S, safe.Field(oldObj, func(oldObj *M1) *string { return &oldObj.S }))...)

	return errs
}

func Validate_T1(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *T1) (errs field.ErrorList) {
	// field T1.MSM1
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj map[string]M1) (errs field.ErrorList) {
			errs = append(errs, validate.EachMapVal(ctx, op, fldPath, obj, oldObj, Validate_M1)...)
			return
		}(fldPath.Child("msm1"), obj.MSM1, safe.Field(oldObj, func(oldObj *T1) map[string]M1 { return oldObj.MSM1 }))...)

	return errs
}
