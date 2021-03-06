// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/customer.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A customer.
type Customer struct {
	// The resource name of the customer.
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the customer.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional, non-unique descriptive name of the customer.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,4,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// The currency in which the account operates.
	// A subset of the currency codes from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The local timezone ID of the customer.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// The URL template for constructing a tracking URL out of parameters.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,7,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The URL template for appending params to the final URL
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,11,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Whether auto-tagging is enabled for the customer.
	AutoTaggingEnabled *wrappers.BoolValue `protobuf:"bytes,8,opt,name=auto_tagging_enabled,json=autoTaggingEnabled,proto3" json:"auto_tagging_enabled,omitempty"`
	// Whether the Customer has a Partners program badge. If the Customer is not
	// associated with the Partners program, this will be false. For more
	// information, see https://support.google.com/partners/answer/3125774.
	HasPartnersBadge *wrappers.BoolValue `protobuf:"bytes,9,opt,name=has_partners_badge,json=hasPartnersBadge,proto3" json:"has_partners_badge,omitempty"`
	// Call reporting setting for a customer.
	CallReportingSetting *CallReportingSetting `protobuf:"bytes,10,opt,name=call_reporting_setting,json=callReportingSetting,proto3" json:"call_reporting_setting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_1a67549631f2cc56, []int{0}
}
func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (dst *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(dst, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Customer) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Customer) GetDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.DescriptiveName
	}
	return nil
}

func (m *Customer) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *Customer) GetTimeZone() *wrappers.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func (m *Customer) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Customer) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *Customer) GetAutoTaggingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.AutoTaggingEnabled
	}
	return nil
}

func (m *Customer) GetHasPartnersBadge() *wrappers.BoolValue {
	if m != nil {
		return m.HasPartnersBadge
	}
	return nil
}

func (m *Customer) GetCallReportingSetting() *CallReportingSetting {
	if m != nil {
		return m.CallReportingSetting
	}
	return nil
}

// Call reporting setting for a customer.
type CallReportingSetting struct {
	// Enable reporting of phone call events by redirecting them via Google
	// System.
	CallReportingEnabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=call_reporting_enabled,json=callReportingEnabled,proto3" json:"call_reporting_enabled,omitempty"`
	// Whether to enable call conversion reporting.
	CallConversionReportingEnabled *wrappers.BoolValue `protobuf:"bytes,2,opt,name=call_conversion_reporting_enabled,json=callConversionReportingEnabled,proto3" json:"call_conversion_reporting_enabled,omitempty"`
	// Customer-level call conversion action to attribute a call conversion to.
	// If not set a default conversion action is used. Only in effect when
	// call_conversion_reporting_enabled is set to true.
	CallConversionAction *wrappers.StringValue `protobuf:"bytes,9,opt,name=call_conversion_action,json=callConversionAction,proto3" json:"call_conversion_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallReportingSetting) Reset()         { *m = CallReportingSetting{} }
func (m *CallReportingSetting) String() string { return proto.CompactTextString(m) }
func (*CallReportingSetting) ProtoMessage()    {}
func (*CallReportingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_1a67549631f2cc56, []int{1}
}
func (m *CallReportingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallReportingSetting.Unmarshal(m, b)
}
func (m *CallReportingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallReportingSetting.Marshal(b, m, deterministic)
}
func (dst *CallReportingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallReportingSetting.Merge(dst, src)
}
func (m *CallReportingSetting) XXX_Size() int {
	return xxx_messageInfo_CallReportingSetting.Size(m)
}
func (m *CallReportingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CallReportingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CallReportingSetting proto.InternalMessageInfo

func (m *CallReportingSetting) GetCallReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallConversionReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.CallConversionAction
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "google.ads.googleads.v0.resources.Customer")
	proto.RegisterType((*CallReportingSetting)(nil), "google.ads.googleads.v0.resources.CallReportingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/customer.proto", fileDescriptor_customer_1a67549631f2cc56)
}

var fileDescriptor_customer_1a67549631f2cc56 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0x49, 0x6a, 0x6b, 0x3b, 0x6d, 0xb5, 0x8c, 0x51, 0x96, 0x2a, 0xa5, 0xad, 0x08, 0x05,
	0x61, 0x12, 0x54, 0x14, 0xf1, 0x6a, 0x13, 0xb4, 0x2a, 0x22, 0x61, 0xfb, 0x71, 0x51, 0x02, 0xcb,
	0x64, 0xf6, 0x64, 0x3b, 0x38, 0x3b, 0xb3, 0xcc, 0xcc, 0xc6, 0x8f, 0xc7, 0xd1, 0x3b, 0x9f, 0xc3,
	0x2b, 0x2f, 0x7d, 0x22, 0xd9, 0xd9, 0x9d, 0xa5, 0x25, 0xa1, 0xc9, 0x55, 0x0e, 0xc3, 0xff, 0xf7,
	0x3b, 0x27, 0xb3, 0xc3, 0x41, 0xbd, 0x54, 0xa9, 0x54, 0x40, 0x97, 0x26, 0xa6, 0x5b, 0x95, 0x65,
	0x35, 0xed, 0x75, 0x35, 0x18, 0x55, 0x68, 0x06, 0xa6, 0xcb, 0x0a, 0x63, 0x55, 0x06, 0x9a, 0xe4,
	0x5a, 0x59, 0x85, 0x0f, 0xaa, 0x18, 0xa1, 0x89, 0x21, 0x0d, 0x41, 0xa6, 0x3d, 0xd2, 0x10, 0xbb,
	0x7b, 0xb5, 0xd4, 0x01, 0xe3, 0x62, 0xd2, 0xfd, 0xaa, 0x69, 0x9e, 0x83, 0x36, 0x95, 0xe2, 0xf0,
	0xcf, 0x2a, 0x5a, 0x1f, 0xd4, 0x56, 0xfc, 0x18, 0x6d, 0x7b, 0x32, 0x96, 0x34, 0x83, 0xa0, 0xb5,
	0xdf, 0x3a, 0xda, 0x88, 0xb6, 0xfc, 0xe1, 0x67, 0x9a, 0x01, 0x7e, 0x8a, 0xda, 0x3c, 0x09, 0x56,
	0xf6, 0x5b, 0x47, 0x9b, 0xcf, 0x1e, 0xd6, 0x6d, 0x89, 0xd7, 0x93, 0x0f, 0xd2, 0xbe, 0x7c, 0x71,
	0x4e, 0x45, 0x01, 0x51, 0x9b, 0x27, 0xf8, 0x18, 0xed, 0x24, 0x60, 0x98, 0xe6, 0xb9, 0xe5, 0xd3,
	0x5a, 0x7a, 0xcb, 0xa1, 0x8f, 0x66, 0xd0, 0x13, 0xab, 0xb9, 0x4c, 0x2b, 0xf6, 0xee, 0x15, 0xca,
	0x75, 0x0d, 0xd1, 0x36, 0x2b, 0xb4, 0x06, 0xc9, 0xbe, 0xc7, 0x4c, 0x25, 0x10, 0xac, 0x2e, 0x61,
	0xd9, 0xf2, 0xc8, 0x40, 0x25, 0x80, 0x5f, 0xa3, 0x0d, 0xcb, 0x33, 0x88, 0x7f, 0x28, 0x09, 0xc1,
	0xda, 0x12, 0xf8, 0x7a, 0x19, 0xbf, 0x50, 0x12, 0xf0, 0x10, 0xdd, 0xb7, 0x9a, 0xb2, 0x2f, 0x5c,
	0xa6, 0x71, 0xa1, 0x45, 0x6c, 0x21, 0xcb, 0x05, 0xb5, 0x10, 0xdc, 0x5e, 0x42, 0x73, 0xcf, 0xa3,
	0x67, 0x5a, 0x9c, 0xd6, 0x20, 0x7e, 0x87, 0x76, 0x26, 0x5c, 0x52, 0xe1, 0x74, 0xa6, 0x98, 0x4c,
	0xf8, 0xb7, 0x60, 0x73, 0x09, 0xd9, 0x1d, 0x47, 0x9d, 0x69, 0x71, 0xe2, 0x18, 0xfc, 0x09, 0x75,
	0x68, 0x61, 0x55, 0x6c, 0x69, 0x9a, 0x96, 0xd3, 0x81, 0xa4, 0x63, 0x01, 0x49, 0xb0, 0xee, 0x5c,
	0xbb, 0x33, 0xae, 0xbe, 0x52, 0xa2, 0x32, 0xe1, 0x92, 0x3b, 0xad, 0xb0, 0xb7, 0x15, 0x85, 0xdf,
	0x23, 0x7c, 0x49, 0x4d, 0x9c, 0x53, 0x6d, 0x25, 0x68, 0x13, 0x8f, 0x69, 0x92, 0x42, 0xb0, 0xb1,
	0xd0, 0xb5, 0x73, 0x49, 0xcd, 0xb0, 0x86, 0xfa, 0x25, 0x83, 0x33, 0xf4, 0x80, 0x51, 0x21, 0x62,
	0x0d, 0xb9, 0xd2, 0xb6, 0x9c, 0xcc, 0x80, 0x2d, 0x7f, 0x03, 0xe4, 0x6c, 0xaf, 0xc8, 0xc2, 0xb7,
	0x4b, 0x06, 0x54, 0x88, 0xc8, 0xf3, 0x27, 0x15, 0x1e, 0x75, 0xd8, 0x9c, 0xd3, 0xc3, 0x5f, 0x6d,
	0xd4, 0x99, 0x17, 0xc7, 0xc3, 0x99, 0x39, 0xfc, 0x0d, 0xb5, 0x16, 0xfe, 0xab, 0xeb, 0xad, 0xfc,
	0x1d, 0x01, 0x3a, 0x70, 0x46, 0xa6, 0xe4, 0x14, 0xb4, 0xe1, 0x4a, 0xce, 0x91, 0xb7, 0x17, 0xca,
	0xf7, 0x4a, 0xc9, 0xa0, 0x71, 0xcc, 0xb4, 0x89, 0xea, 0xc1, 0xaf, 0xb4, 0xa1, 0xcc, 0x72, 0x25,
	0xeb, 0xcf, 0x71, 0xf3, 0x33, 0xe9, 0x5c, 0xb7, 0x87, 0x8e, 0xec, 0xff, 0x6b, 0xa1, 0x27, 0x4c,
	0x65, 0x8b, 0xaf, 0xbe, 0xbf, 0xed, 0x77, 0xc2, 0xb0, 0xb4, 0x0f, 0x5b, 0x17, 0x1f, 0x6b, 0x26,
	0x55, 0x82, 0xca, 0x94, 0x28, 0x9d, 0x76, 0x53, 0x90, 0xae, 0xb7, 0x5f, 0x56, 0x39, 0x37, 0x37,
	0xec, 0xae, 0x37, 0x4d, 0xf5, 0xb3, 0xbd, 0x72, 0x1c, 0x86, 0xbf, 0xdb, 0x07, 0xc7, 0x95, 0x32,
	0x4c, 0x0c, 0xa9, 0xca, 0xb2, 0x3a, 0xef, 0x91, 0xc8, 0x27, 0xff, 0xfa, 0xcc, 0x28, 0x4c, 0xcc,
	0xa8, 0xc9, 0x8c, 0xce, 0x7b, 0xa3, 0x26, 0x33, 0x5e, 0x73, 0x43, 0x3c, 0xff, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0x9c, 0xcb, 0x74, 0x3f, 0x05, 0x00, 0x00,
}
