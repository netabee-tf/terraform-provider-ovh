// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudProjectLoadbalancerDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"created_at": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "The UTC date and timestamp when the loadbalancer was created",
			MarkdownDescription: "The UTC date and timestamp when the loadbalancer was created",
		},
		"flavor_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "ID of the flavor",
			MarkdownDescription: "ID of the flavor",
		},
		"floating_ip": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "ID of the floating IP",
					MarkdownDescription: "ID of the floating IP",
				},
				"ip": schema.StringAttribute{
					CustomType:          ovhtypes.TfStringType{},
					Computed:            true,
					Description:         "IP Address of the floating IP",
					MarkdownDescription: "IP Address of the floating IP",
				},
			},
			CustomType: FloatingIpType{
				ObjectType: types.ObjectType{
					AttrTypes: FloatingIpValue{}.AttributeTypes(ctx),
				},
			},
			Computed:            true,
			Description:         "Information about floating IP",
			MarkdownDescription: "Information about floating IP",
		},
		"id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Load balancer ID",
			MarkdownDescription: "Load balancer ID",
		},
		"name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Name of the loadbalancer",
			MarkdownDescription: "Name of the loadbalancer",
		},
		"operating_status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Operating status of the loadbalancer",
			MarkdownDescription: "Operating status of the loadbalancer",
		},
		"provisioning_status": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Provisioning status of the loadbalancer",
			MarkdownDescription: "Provisioning status of the loadbalancer",
		},
		"region_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Region of the loadbalancer",
			MarkdownDescription: "Region of the loadbalancer",
		},
		"service_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Service name",
			MarkdownDescription: "Service name",
		},
		"updated_at": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "UTC date and timestamp when the loadbalancer was updated",
			MarkdownDescription: "UTC date and timestamp when the loadbalancer was updated",
		},
		"vip_address": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "IP address of the Virtual IP",
			MarkdownDescription: "IP address of the Virtual IP",
		},
		"vip_network_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "Openstack ID of the network for the Virtual IP",
			MarkdownDescription: "Openstack ID of the network for the Virtual IP",
		},
		"vip_subnet_id": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Computed:            true,
			Description:         "ID of the subnet for the Virtual IP",
			MarkdownDescription: "ID of the subnet for the Virtual IP",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type CloudProjectLoadbalancerModel struct {
	CreatedAt          ovhtypes.TfStringValue `tfsdk:"created_at" json:"createdAt"`
	FlavorId           ovhtypes.TfStringValue `tfsdk:"flavor_id" json:"flavorId"`
	FloatingIp         FloatingIpValue        `tfsdk:"floating_ip" json:"floatingIp"`
	Id                 ovhtypes.TfStringValue `tfsdk:"id" json:"id"`
	Name               ovhtypes.TfStringValue `tfsdk:"name" json:"name"`
	OperatingStatus    ovhtypes.TfStringValue `tfsdk:"operating_status" json:"operatingStatus"`
	ProvisioningStatus ovhtypes.TfStringValue `tfsdk:"provisioning_status" json:"provisioningStatus"`
	RegionName         ovhtypes.TfStringValue `tfsdk:"region_name" json:"regionName"`
	ServiceName        ovhtypes.TfStringValue `tfsdk:"service_name" json:"serviceName"`
	UpdatedAt          ovhtypes.TfStringValue `tfsdk:"updated_at" json:"updatedAt"`
	VipAddress         ovhtypes.TfStringValue `tfsdk:"vip_address" json:"vipAddress"`
	VipNetworkId       ovhtypes.TfStringValue `tfsdk:"vip_network_id" json:"vipNetworkId"`
	VipSubnetId        ovhtypes.TfStringValue `tfsdk:"vip_subnet_id" json:"vipSubnetId"`
}

func (v *CloudProjectLoadbalancerModel) MergeWith(other *CloudProjectLoadbalancerModel) {

	if (v.CreatedAt.IsUnknown() || v.CreatedAt.IsNull()) && !other.CreatedAt.IsUnknown() {
		v.CreatedAt = other.CreatedAt
	}

	if (v.FlavorId.IsUnknown() || v.FlavorId.IsNull()) && !other.FlavorId.IsUnknown() {
		v.FlavorId = other.FlavorId
	}

	if (v.FloatingIp.IsUnknown() || v.FloatingIp.IsNull()) && !other.FloatingIp.IsUnknown() {
		v.FloatingIp = other.FloatingIp
	}

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.OperatingStatus.IsUnknown() || v.OperatingStatus.IsNull()) && !other.OperatingStatus.IsUnknown() {
		v.OperatingStatus = other.OperatingStatus
	}

	if (v.ProvisioningStatus.IsUnknown() || v.ProvisioningStatus.IsNull()) && !other.ProvisioningStatus.IsUnknown() {
		v.ProvisioningStatus = other.ProvisioningStatus
	}

	if (v.RegionName.IsUnknown() || v.RegionName.IsNull()) && !other.RegionName.IsUnknown() {
		v.RegionName = other.RegionName
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}

	if (v.UpdatedAt.IsUnknown() || v.UpdatedAt.IsNull()) && !other.UpdatedAt.IsUnknown() {
		v.UpdatedAt = other.UpdatedAt
	}

	if (v.VipAddress.IsUnknown() || v.VipAddress.IsNull()) && !other.VipAddress.IsUnknown() {
		v.VipAddress = other.VipAddress
	}

	if (v.VipNetworkId.IsUnknown() || v.VipNetworkId.IsNull()) && !other.VipNetworkId.IsUnknown() {
		v.VipNetworkId = other.VipNetworkId
	}

	if (v.VipSubnetId.IsUnknown() || v.VipSubnetId.IsNull()) && !other.VipSubnetId.IsUnknown() {
		v.VipSubnetId = other.VipSubnetId
	}

}

var _ basetypes.ObjectTypable = FloatingIpType{}

type FloatingIpType struct {
	basetypes.ObjectType
}

func (t FloatingIpType) Equal(o attr.Type) bool {
	other, ok := o.(FloatingIpType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t FloatingIpType) String() string {
	return "FloatingIpType"
}

func (t FloatingIpType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	ipAttribute, ok := attributes["ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip is missing from object`)

		return nil, diags
	}

	ipVal, ok := ipAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip expected to be ovhtypes.TfStringValue, was: %T`, ipAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return FloatingIpValue{
		Id:    idVal,
		Ip:    ipVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewFloatingIpValueNull() FloatingIpValue {
	return FloatingIpValue{
		state: attr.ValueStateNull,
	}
}

func NewFloatingIpValueUnknown() FloatingIpValue {
	return FloatingIpValue{
		state: attr.ValueStateUnknown,
	}
}

func NewFloatingIpValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (FloatingIpValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing FloatingIpValue Attribute Value",
				"While creating a FloatingIpValue value, a missing attribute value was detected. "+
					"A FloatingIpValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("FloatingIpValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid FloatingIpValue Attribute Type",
				"While creating a FloatingIpValue value, an invalid attribute value was detected. "+
					"A FloatingIpValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("FloatingIpValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("FloatingIpValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra FloatingIpValue Attribute Value",
				"While creating a FloatingIpValue value, an extra attribute value was detected. "+
					"A FloatingIpValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra FloatingIpValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewFloatingIpValueUnknown(), diags
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewFloatingIpValueUnknown(), diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	ipAttribute, ok := attributes["ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip is missing from object`)

		return NewFloatingIpValueUnknown(), diags
	}

	ipVal, ok := ipAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip expected to be ovhtypes.TfStringValue, was: %T`, ipAttribute))
	}

	if diags.HasError() {
		return NewFloatingIpValueUnknown(), diags
	}

	return FloatingIpValue{
		Id:    idVal,
		Ip:    ipVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewFloatingIpValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) FloatingIpValue {
	object, diags := NewFloatingIpValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewFloatingIpValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t FloatingIpType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewFloatingIpValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewFloatingIpValueUnknown(), nil
	}

	if in.IsNull() {
		return NewFloatingIpValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewFloatingIpValueMust(FloatingIpValue{}.AttributeTypes(ctx), attributes), nil
}

func (t FloatingIpType) ValueType(ctx context.Context) attr.Value {
	return FloatingIpValue{}
}

var _ basetypes.ObjectValuable = FloatingIpValue{}

type FloatingIpValue struct {
	Id    ovhtypes.TfStringValue `tfsdk:"id" json:"id"`
	Ip    ovhtypes.TfStringValue `tfsdk:"ip" json:"ip"`
	state attr.ValueState
}

func (v *FloatingIpValue) UnmarshalJSON(data []byte) error {
	type JsonFloatingIpValue FloatingIpValue

	var tmp JsonFloatingIpValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.Id = tmp.Id
	v.Ip = tmp.Ip

	v.state = attr.ValueStateKnown

	return nil
}

func (v *FloatingIpValue) MergeWith(other *FloatingIpValue) {

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.Ip.IsUnknown() || v.Ip.IsNull()) && !other.Ip.IsUnknown() {
		v.Ip = other.Ip
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v FloatingIpValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"id": v.Id,
		"ip": v.Ip,
	}
}
func (v FloatingIpValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["ip"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.Ip.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ip"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v FloatingIpValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v FloatingIpValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v FloatingIpValue) String() string {
	return "FloatingIpValue"
}

func (v FloatingIpValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"id": ovhtypes.TfStringType{},
			"ip": ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"id": v.Id,
			"ip": v.Ip,
		})

	return objVal, diags
}

func (v FloatingIpValue) Equal(o attr.Value) bool {
	other, ok := o.(FloatingIpValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.Ip.Equal(other.Ip) {
		return false
	}

	return true
}

func (v FloatingIpValue) Type(ctx context.Context) attr.Type {
	return FloatingIpType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v FloatingIpValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"id": ovhtypes.TfStringType{},
		"ip": ovhtypes.TfStringType{},
	}
}
