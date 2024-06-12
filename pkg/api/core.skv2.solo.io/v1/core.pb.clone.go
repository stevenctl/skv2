// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/core/v1/core.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *ObjectRef) Clone() proto.Message {
	var target *ObjectRef
	if m == nil {
		return target
	}
	target = &ObjectRef{}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	return target
}

// Clone function
func (m *ObjectRefList) Clone() proto.Message {
	var target *ObjectRefList
	if m == nil {
		return target
	}
	target = &ObjectRefList{}

	if m.GetRefs() != nil {
		target.Refs = make([]*ObjectRef, len(m.GetRefs()))
		for idx, v := range m.GetRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Refs[idx] = h.Clone().(*ObjectRef)
			} else {
				target.Refs[idx] = proto.Clone(v).(*ObjectRef)
			}

		}
	}

	return target
}

// Clone function
func (m *ClusterObjectRef) Clone() proto.Message {
	var target *ClusterObjectRef
	if m == nil {
		return target
	}
	target = &ClusterObjectRef{}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	target.ClusterName = m.GetClusterName()

	return target
}

// Clone function
func (m *TypedObjectRef) Clone() proto.Message {
	var target *TypedObjectRef
	if m == nil {
		return target
	}
	target = &TypedObjectRef{}

	if h, ok := interface{}(m.GetApiGroup()).(clone.Cloner); ok {
		target.ApiGroup = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.ApiGroup = proto.Clone(m.GetApiGroup()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetKind()).(clone.Cloner); ok {
		target.Kind = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Kind = proto.Clone(m.GetKind()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	return target
}

// Clone function
func (m *TypedClusterObjectRef) Clone() proto.Message {
	var target *TypedClusterObjectRef
	if m == nil {
		return target
	}
	target = &TypedClusterObjectRef{}

	if h, ok := interface{}(m.GetApiGroup()).(clone.Cloner); ok {
		target.ApiGroup = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.ApiGroup = proto.Clone(m.GetApiGroup()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetKind()).(clone.Cloner); ok {
		target.Kind = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Kind = proto.Clone(m.GetKind()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	target.Name = m.GetName()

	target.Namespace = m.GetNamespace()

	target.ClusterName = m.GetClusterName()

	return target
}

// Clone function
func (m *Status) Clone() proto.Message {
	var target *Status
	if m == nil {
		return target
	}
	target = &Status{}

	target.State = m.GetState()

	target.Message = m.GetMessage()

	target.ObservedGeneration = m.GetObservedGeneration()

	if h, ok := interface{}(m.GetProcessingTime()).(clone.Cloner); ok {
		target.ProcessingTime = h.Clone().(*google_golang_org_protobuf_types_known_timestamppb.Timestamp)
	} else {
		target.ProcessingTime = proto.Clone(m.GetProcessingTime()).(*google_golang_org_protobuf_types_known_timestamppb.Timestamp)
	}

	if h, ok := interface{}(m.GetOwner()).(clone.Cloner); ok {
		target.Owner = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Owner = proto.Clone(m.GetOwner()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	return target
}

// Clone function
func (m *ObjectSelector) Clone() proto.Message {
	var target *ObjectSelector
	if m == nil {
		return target
	}
	target = &ObjectSelector{}

	if m.GetNamespaces() != nil {
		target.Namespaces = make([]string, len(m.GetNamespaces()))
		for idx, v := range m.GetNamespaces() {

			target.Namespaces[idx] = v

		}
	}

	if m.GetLabels() != nil {
		target.Labels = make(map[string]string, len(m.GetLabels()))
		for k, v := range m.GetLabels() {

			target.Labels[k] = v

		}
	}

	if m.GetExpressions() != nil {
		target.Expressions = make([]*ObjectSelector_Expression, len(m.GetExpressions()))
		for idx, v := range m.GetExpressions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Expressions[idx] = h.Clone().(*ObjectSelector_Expression)
			} else {
				target.Expressions[idx] = proto.Clone(v).(*ObjectSelector_Expression)
			}

		}
	}

	return target
}

// Clone function
func (m *PolicyTargetReference) Clone() proto.Message {
	var target *PolicyTargetReference
	if m == nil {
		return target
	}
	target = &PolicyTargetReference{}

	target.Group = m.GetGroup()

	target.Kind = m.GetKind()

	target.Name = m.GetName()

	if h, ok := interface{}(m.GetNamespace()).(clone.Cloner); ok {
		target.Namespace = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Namespace = proto.Clone(m.GetNamespace()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	return target
}

// Clone function
func (m *PolicyTargetReferenceWithSectionName) Clone() proto.Message {
	var target *PolicyTargetReferenceWithSectionName
	if m == nil {
		return target
	}
	target = &PolicyTargetReferenceWithSectionName{}

	target.Group = m.GetGroup()

	target.Kind = m.GetKind()

	target.Name = m.GetName()

	if h, ok := interface{}(m.GetNamespace()).(clone.Cloner); ok {
		target.Namespace = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Namespace = proto.Clone(m.GetNamespace()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetSectionName()).(clone.Cloner); ok {
		target.SectionName = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.SectionName = proto.Clone(m.GetSectionName()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	return target
}

// Clone function
func (m *ObjectSelector_Expression) Clone() proto.Message {
	var target *ObjectSelector_Expression
	if m == nil {
		return target
	}
	target = &ObjectSelector_Expression{}

	target.Key = m.GetKey()

	target.Operator = m.GetOperator()

	if m.GetValues() != nil {
		target.Values = make([]string, len(m.GetValues()))
		for idx, v := range m.GetValues() {

			target.Values[idx] = v

		}
	}

	return target
}
