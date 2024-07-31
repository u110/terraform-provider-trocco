package provider

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &datamartDefinitionResource{}
var _ resource.ResourceWithImportState = &datamartDefinitionResource{}

func newDatamartDefinitionResource() resource.Resource {
	return &datamartDefinitionResource{}
}

type datamartDefinitionResource struct {
	client *client.TroccoClient
}

type datamartDefinitionModel struct {
	ID                     types.Int64                  `tfsdk:"id"`
	Name                   types.String                 `tfsdk:"name"`
	Description            types.String                 `tfsdk:"description"`
	DataWarehouseType      types.String                 `tfsdk:"data_warehouse_type"`
	IsRunnableConcurrently types.Bool                   `tfsdk:"is_runnable_concurrently"`
	ResourceGroupID        types.Int64                  `tfsdk:"resource_group_id"`
	CustomVariableSettings []customVariableSettingModel `tfsdk:"custom_variable_settings"`
	DatamartBigqueryOption *datamartBigqueryOptionModel `tfsdk:"datamart_bigquery_option"`
	Notifications          []datamartNotificationModel  `tfsdk:"notifications"`
	Schedules              []scheduleModel              `tfsdk:"schedules"`
	Labels                 []labelModel                 `tfsdk:"labels"`
}

type customVariableSettingModel struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int32  `tfsdk:"quantity"`
	Unit      types.String `tfsdk:"unit"`
	Direction types.String `tfsdk:"direction"`
	Format    types.String `tfsdk:"format"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

type datamartBigqueryOptionModel struct {
	BigqueryConnectionID types.Int64    `tfsdk:"bigquery_connection_id"`
	QueryMode            types.String   `tfsdk:"query_mode"`
	Query                types.String   `tfsdk:"query"`
	DestinationDataset   types.String   `tfsdk:"destination_dataset"`
	DestinationTable     types.String   `tfsdk:"destination_table"`
	WriteDisposition     types.String   `tfsdk:"write_disposition"`
	BeforeLoad           types.String   `tfsdk:"before_load"`
	Partitioning         types.String   `tfsdk:"partitioning"`
	PartitioningTime     types.String   `tfsdk:"partitioning_time"`
	PartitioningField    types.String   `tfsdk:"partitioning_field"`
	ClusteringFields     []types.String `tfsdk:"clustering_fields"`
	Location             types.String   `tfsdk:"location"`
}

type datamartNotificationModel struct {
	DestinationType  types.String `tfsdk:"destination_type"`
	SlackChannelID   types.Int64  `tfsdk:"slack_channel_id"`
	EmailID          types.Int64  `tfsdk:"email_id"`
	NotificationType types.String `tfsdk:"notification_type"`
	NotifyWhen       types.String `tfsdk:"notify_when"`
	RecordCount      types.Int64  `tfsdk:"record_count"`
	RecordOperator   types.String `tfsdk:"record_operator"`
	Message          types.String `tfsdk:"message"`
}

type scheduleModel struct {
	Frequency types.String `tfsdk:"frequency"`
	Minute    types.Int32  `tfsdk:"minute"`
	Hour      types.Int32  `tfsdk:"hour"`
	Day       types.Int32  `tfsdk:"day"`
	DayOfWeek types.Int32  `tfsdk:"day_of_week"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

type labelModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *datamartDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	// resp.TypeName = req.ProviderTypeName + "_datamart_definition"
	resp.TypeName = "datamart_definition"
}

func (r *datamartDefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *datamartDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"data_warehouse_type": schema.StringAttribute{
				Required: true,
			},
			"is_runnable_concurrently": schema.BoolAttribute{
				Required: true,
			},
			"resource_group_id": schema.Int64Attribute{
				Optional: true,
			},
			"custom_variable_settings": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Optional: true,
						},
						"quantity": schema.Int32Attribute{
							Optional: true,
						},
						"unit": schema.StringAttribute{
							Optional: true,
						},
						"direction": schema.StringAttribute{
							Optional: true,
						},
						"format": schema.StringAttribute{
							Optional: true,
						},
						"time_zone": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"datamart_bigquery_option": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"bigquery_connection_id": schema.Int64Attribute{
						Required: true,
					},
					"query_mode": schema.StringAttribute{
						Required: true,
					},
					"query": schema.StringAttribute{
						Required: true,
					},
					"destination_dataset": schema.StringAttribute{
						Optional: true,
					},
					"destination_table": schema.StringAttribute{
						Optional: true,
					},
					"write_disposition": schema.StringAttribute{
						Optional: true,
					},
					"before_load": schema.StringAttribute{
						Optional: true,
					},
					"partitioning": schema.StringAttribute{
						Optional: true,
					},
					"partitioning_time": schema.StringAttribute{
						Optional: true,
					},
					"partitioning_field": schema.StringAttribute{
						Optional: true,
					},
					"clustering_fields": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"location": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"notifications": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination_type": schema.StringAttribute{
							Required: true,
						},
						"slack_channel_id": schema.Int64Attribute{
							Optional: true,
						},
						"email_id": schema.Int64Attribute{
							Optional: true,
						},
						"notification_type": schema.StringAttribute{
							Required: true,
						},
						"notify_when": schema.StringAttribute{
							Optional: true,
						},
						"record_count": schema.Int64Attribute{
							Optional: true,
						},
						"record_operator": schema.StringAttribute{
							Optional: true,
						},
						"message": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"schedules": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"frequency": schema.StringAttribute{
							Required: true,
						},
						"minute": schema.Int32Attribute{
							Optional: true,
						},
						"hour": schema.Int32Attribute{
							Optional: true,
						},
						"day": schema.Int32Attribute{
							Optional: true,
						},
						"day_of_week": schema.Int32Attribute{
							Optional: true,
						},
						"time_zone": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"labels": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}

func (r *datamartDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan datamartDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.CreateDatamartDefinitionInput{
		Name:                   plan.Name.ValueString(),
		DatawarehouseType:      plan.DataWarehouseType.ValueString(),
		IsRunnableConcurrently: plan.IsRunnableConcurrently.ValueBool(),
	}
	if !plan.Description.IsNull() {
		input.SetDescription(plan.Description.ValueString())
	}
	if !plan.ResourceGroupID.IsNull() {
		input.SetResourceGroupID(plan.ResourceGroupID.ValueInt64())
	}
	if plan.CustomVariableSettings != nil {
		customVariableSettingInputs := make([]client.CustomVariableSettingInput, len(plan.CustomVariableSettings))
		for i, v := range plan.CustomVariableSettings {
			if v.Type.ValueString() == "string" {
				customVariableSettingInputs[i] = client.NewStringTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Value.ValueString(),
				)
			} else {
				customVariableSettingInputs[i] = client.NewTimestampTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Type.ValueString(),
					int(v.Quantity.ValueInt32()),
					v.Unit.ValueString(),
					v.Direction.ValueString(),
					v.Format.ValueString(),
					v.TimeZone.ValueString(),
				)
			}
		}
		input.SetCustomVariableSettings(customVariableSettingInputs)
	}
	if plan.DatamartBigqueryOption != nil {
		if plan.DatamartBigqueryOption.QueryMode.ValueString() == "insert" {
			optionInput := client.NewInsertModeCreateDatamartBigqueryOptionInput(
				plan.DatamartBigqueryOption.BigqueryConnectionID.ValueInt64(),
				plan.DatamartBigqueryOption.Query.ValueString(),
				plan.DatamartBigqueryOption.DestinationDataset.ValueString(),
				plan.DatamartBigqueryOption.DestinationTable.ValueString(),
				plan.DatamartBigqueryOption.WriteDisposition.ValueString(),
			)
			if !plan.DatamartBigqueryOption.BeforeLoad.IsNull() {
				optionInput.SetBeforeLoad(plan.DatamartBigqueryOption.BeforeLoad.ValueString())
			}
			if !plan.DatamartBigqueryOption.Partitioning.IsNull() {
				optionInput.SetPartitioning(plan.DatamartBigqueryOption.Partitioning.ValueString())
			}
			if !plan.DatamartBigqueryOption.PartitioningTime.IsNull() {
				optionInput.SetPartitioningTime(plan.DatamartBigqueryOption.PartitioningTime.ValueString())
			}
			if !plan.DatamartBigqueryOption.PartitioningField.IsNull() {
				optionInput.SetPartitioningField(plan.DatamartBigqueryOption.PartitioningField.ValueString())
			}
			if plan.DatamartBigqueryOption.ClusteringFields != nil {
				clusteringFields := make([]string, len(plan.DatamartBigqueryOption.ClusteringFields))
				for i, v := range plan.DatamartBigqueryOption.ClusteringFields {
					clusteringFields[i] = v.ValueString()
				}
				optionInput.SetClusteringFields(clusteringFields)
			}
			input.SetDatamartBigqueryOption(optionInput)
		} else {
			input.SetDatamartBigqueryOption(
				client.NewQueryModeCreateDatamartBigqueryOptionInput(
					plan.DatamartBigqueryOption.BigqueryConnectionID.ValueInt64(),
					plan.DatamartBigqueryOption.Query.ValueString(),
					plan.DatamartBigqueryOption.Location.ValueString(),
				),
			)
		}
	}
	res, err := r.client.CreateDatamartDefinition(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating datamart_definition",
			fmt.Sprintf("Unable to create datamart_definition, got error: %s", err),
		)
		return
	}

	data, err := r.fetchModel(res.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("Unable to read datamart_definition (id: %d), got error: %s", res.ID, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *datamartDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state datamartDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueInt64()
	data, err := r.fetchModel(id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("Unable to read datamart_definition (id: %d), got error: %s", id, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *datamartDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state datamartDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := client.UpdateDatamartDefinitionInput{}
	input.SetName(plan.Name.ValueString())
	if !plan.Description.IsNull() {
		input.SetDescription(plan.Description.ValueString())
	}
	input.SetIsRunnableConcurrently(plan.IsRunnableConcurrently.ValueBool())
	if !plan.ResourceGroupID.IsNull() {
		input.SetResourceGroupID(plan.ResourceGroupID.ValueInt64())
	}
	if plan.CustomVariableSettings != nil {
		customVariableSettingInputs := make([]client.CustomVariableSettingInput, len(plan.CustomVariableSettings))
		for i, v := range plan.CustomVariableSettings {
			if v.Type.ValueString() == "string" {
				customVariableSettingInputs[i] = client.NewStringTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Value.ValueString(),
				)
			} else {
				customVariableSettingInputs[i] = client.NewTimestampTypeCustomVariableSettingInput(
					v.Name.ValueString(),
					v.Type.ValueString(),
					int(v.Quantity.ValueInt32()),
					v.Unit.ValueString(),
					v.Direction.ValueString(),
					v.Format.ValueString(),
					v.TimeZone.ValueString(),
				)
			}
		}
		input.SetCustomVariableSettings(customVariableSettingInputs)
	} else {
		input.SetCustomVariableSettings([]client.CustomVariableSettingInput{})
	}
	if plan.DatamartBigqueryOption != nil {
		optionInput := client.UpdateDatamartBigqueryOptionInput{}
		optionInput.SetBigqueryConnectionID(plan.DatamartBigqueryOption.BigqueryConnectionID.ValueInt64())
		optionInput.SetQueryMode(plan.DatamartBigqueryOption.QueryMode.ValueString())
		optionInput.SetQuery(plan.DatamartBigqueryOption.Query.ValueString())
		if !plan.DatamartBigqueryOption.DestinationDataset.IsNull() {
			optionInput.SetDestinationDataset(plan.DatamartBigqueryOption.DestinationDataset.ValueString())
		}
		if !plan.DatamartBigqueryOption.DestinationTable.IsNull() {
			optionInput.SetDestinationTable(plan.DatamartBigqueryOption.DestinationTable.ValueString())
		}
		if !plan.DatamartBigqueryOption.WriteDisposition.IsNull() {
			optionInput.SetWriteDisposition(plan.DatamartBigqueryOption.WriteDisposition.ValueString())
		}
		if !plan.DatamartBigqueryOption.BeforeLoad.IsNull() {
			optionInput.SetBeforeLoad(plan.DatamartBigqueryOption.BeforeLoad.ValueString())
		}
		if !plan.DatamartBigqueryOption.Partitioning.IsNull() {
			optionInput.SetPartitioning(plan.DatamartBigqueryOption.Partitioning.ValueString())
		}
		if !plan.DatamartBigqueryOption.PartitioningTime.IsNull() {
			optionInput.SetPartitioningTime(plan.DatamartBigqueryOption.PartitioningTime.ValueString())
		}
		if !plan.DatamartBigqueryOption.PartitioningField.IsNull() {
			optionInput.SetPartitioningField(plan.DatamartBigqueryOption.PartitioningField.ValueString())
		}
		if plan.DatamartBigqueryOption.ClusteringFields != nil {
			clusteringFields := make([]string, len(plan.DatamartBigqueryOption.ClusteringFields))
			for i, v := range plan.DatamartBigqueryOption.ClusteringFields {
				clusteringFields[i] = v.ValueString()
			}
			optionInput.SetClusteringFields(clusteringFields)
		}
		if !plan.DatamartBigqueryOption.Location.IsNull() {
			optionInput.SetLocation(plan.DatamartBigqueryOption.Location.ValueString())
		}
		input.SetDatamartBigqueryOption(optionInput)
	}
	if plan.Schedules != nil {
		scheduleInputs := make([]client.ScheduleInput, len(plan.Schedules))
		for i, v := range plan.Schedules {
			switch v.Frequency.ValueString() {
			case "hourly":
				{
					scheduleInputs[i] = client.NewHourlyScheduleInput(
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "daily":
				{
					scheduleInputs[i] = client.NewDailyScheduleInput(
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "weekly":
				{
					scheduleInputs[i] = client.NewWeeklyScheduleInput(
						int(v.DayOfWeek.ValueInt32()),
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			case "monthly":
				{
					scheduleInputs[i] = client.NewMonthlyScheduleInput(
						int(v.Day.ValueInt32()),
						int(v.Hour.ValueInt32()),
						int(v.Minute.ValueInt32()),
						v.TimeZone.ValueString(),
					)
				}
			}
		}
		input.SetSchedules(scheduleInputs)
	} else {
		input.SetSchedules([]client.ScheduleInput{})
	}
	if plan.Notifications != nil {
		notificationInputs := make([]client.DatamartNotificationInput, len(plan.Notifications))
		for i, v := range plan.Notifications {
			if v.DestinationType.ValueString() == "slack" {
				if v.NotificationType.ValueString() == "job" {
					notificationInputs[i] = client.NewSlackJobDatamartNotificationInput(
						v.SlackChannelID.ValueInt64(),
						v.NotifyWhen.ValueString(),
						v.Message.ValueString(),
					)
				} else {
					notificationInputs[i] = client.NewSlackRecordDatamartNotificationInput(
						v.SlackChannelID.ValueInt64(),
						v.RecordCount.ValueInt64(),
						v.RecordOperator.ValueString(),
						v.Message.ValueString(),
					)
				}
			} else {
				if v.NotificationType.ValueString() == "job" {
					notificationInputs[i] = client.NewEmailJobDatamartNotificationInput(
						v.EmailID.ValueInt64(),
						v.NotifyWhen.ValueString(),
						v.Message.ValueString(),
					)
				} else {
					notificationInputs[i] = client.NewEmailRecordDatamartNotificationInput(
						v.EmailID.ValueInt64(),
						v.RecordCount.ValueInt64(),
						v.RecordOperator.ValueString(),
						v.Message.ValueString(),
					)
				}
			}
		}
		input.SetNotifications(notificationInputs)
	} else {
		input.SetNotifications([]client.DatamartNotificationInput{})
	}
	if plan.Labels != nil {
		labelInputs := make([]string, len(plan.Labels))
		for i, v := range plan.Labels {
			labelInputs[i] = v.Name.ValueString()
		}
		input.SetLabels(labelInputs)
	} else {
		input.SetLabels([]string{})
	}

	err := r.client.UpdateDatamartDefinition(state.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating datamart definition",
			fmt.Sprintf("Unable to update datamart definition (id: %d), got error: %s", state.ID.ValueInt64(), err),
		)
		return
	}

	data, err := r.fetchModel(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading datamart_definition",
			fmt.Sprintf("failed to get datamart definition: %v", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *datamartDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state datamartDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueInt64()
	err := r.client.DeleteDatamartDefinition(id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Deleting datamart_definition",
			fmt.Sprintf("Unable to delete datamart_definition (id: %d), got error: %s", id, err),
		)
		return
	}
}

func (r *datamartDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing datamart_definition",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *datamartDefinitionResource) fetchModel(id int64) (*datamartDefinitionModel, error) {
	datamartDefinition, err := r.client.GetDatamartDefinition(id)
	if err != nil {
		return nil, err
	}
	model := datamartDefinitionModel{
		ID:                     types.Int64Value(datamartDefinition.ID),
		Name:                   types.StringValue(datamartDefinition.Name),
		Description:            types.StringValue(datamartDefinition.Description),
		DataWarehouseType:      types.StringValue(datamartDefinition.DataWarehouseType),
		IsRunnableConcurrently: types.BoolValue(datamartDefinition.IsRunnableConcurrently),
	}
	if datamartDefinition.ResourceGroup != nil {
		model.ResourceGroupID = types.Int64Value(datamartDefinition.ResourceGroup.ID)
	}
	if datamartDefinition.CustomVariableSettings != nil {
		customVariableSettings := make([]customVariableSettingModel, len(datamartDefinition.CustomVariableSettings))
		for i, v := range datamartDefinition.CustomVariableSettings {
			customVariableSettings[i] = customVariableSettingModel{
				Name: types.StringValue(v.Name),
				Type: types.StringValue(v.Type),
			}
			if v.Value != nil {
				customVariableSettings[i].Value = types.StringValue(*v.Value)
			}
			if v.Quantity != nil {
				customVariableSettings[i].Quantity = types.Int32Value(int32(*v.Quantity))
			}
			if v.Unit != nil {
				customVariableSettings[i].Unit = types.StringValue(*v.Unit)
			}
			if v.Direction != nil {
				customVariableSettings[i].Direction = types.StringValue(*v.Direction)
			}
			if v.Format != nil {
				customVariableSettings[i].Format = types.StringValue(*v.Format)
			}
			if v.TimeZone != nil {
				customVariableSettings[i].TimeZone = types.StringValue(*v.TimeZone)
			}
		}
		model.CustomVariableSettings = customVariableSettings
	}
	if datamartDefinition.DatamartBigqueryOption != nil {
		datamartBigqueryOption := &datamartBigqueryOptionModel{
			BigqueryConnectionID: types.Int64Value(datamartDefinition.DatamartBigqueryOption.BigqueryConnectionID),
			QueryMode:            types.StringValue(datamartDefinition.DatamartBigqueryOption.QueryMode),
			Query:                types.StringValue(datamartDefinition.DatamartBigqueryOption.Query),
		}
		if datamartDefinition.DatamartBigqueryOption.DestinationDataset != nil {
			datamartBigqueryOption.DestinationDataset = types.StringValue(*datamartDefinition.DatamartBigqueryOption.DestinationDataset)
		}
		if datamartDefinition.DatamartBigqueryOption.DestinationTable != nil {
			datamartBigqueryOption.DestinationTable = types.StringValue(*datamartDefinition.DatamartBigqueryOption.DestinationTable)
		}
		if datamartDefinition.DatamartBigqueryOption.WriteDisposition != nil {
			datamartBigqueryOption.WriteDisposition = types.StringValue(*datamartDefinition.DatamartBigqueryOption.WriteDisposition)
		}
		if datamartDefinition.DatamartBigqueryOption.BeforeLoad != nil {
			datamartBigqueryOption.BeforeLoad = types.StringValue(*datamartDefinition.DatamartBigqueryOption.BeforeLoad)
		}
		if datamartDefinition.DatamartBigqueryOption.Partitioning != nil {
			datamartBigqueryOption.Partitioning = types.StringValue(*datamartDefinition.DatamartBigqueryOption.Partitioning)
		}
		if datamartDefinition.DatamartBigqueryOption.PartitioningTime != nil {
			datamartBigqueryOption.PartitioningTime = types.StringValue(*datamartDefinition.DatamartBigqueryOption.PartitioningTime)
		}
		if datamartDefinition.DatamartBigqueryOption.PartitioningField != nil {
			datamartBigqueryOption.PartitioningField = types.StringValue(*datamartDefinition.DatamartBigqueryOption.PartitioningField)
		}
		if datamartDefinition.DatamartBigqueryOption.ClusteringFields != nil {
			clusteringFields := make([]types.String, len(datamartDefinition.DatamartBigqueryOption.ClusteringFields))
			for i, v := range datamartDefinition.DatamartBigqueryOption.ClusteringFields {
				clusteringFields[i] = types.StringValue(v)
			}
			datamartBigqueryOption.ClusteringFields = clusteringFields
		}
		if datamartDefinition.DatamartBigqueryOption.Location != nil {
			datamartBigqueryOption.Location = types.StringValue(*datamartDefinition.DatamartBigqueryOption.Location)
		}
		model.DatamartBigqueryOption = datamartBigqueryOption
	}
	if datamartDefinition.Notifications != nil {
		notifications := make([]datamartNotificationModel, len(datamartDefinition.Notifications))
		for i, v := range datamartDefinition.Notifications {
			notifications[i] = datamartNotificationModel{
				DestinationType:  types.StringValue(v.DestinationType),
				NotificationType: types.StringValue(v.NotificationType),
				Message:          types.StringValue(v.Message),
			}
			if v.SlackChannelID != nil {
				notifications[i].SlackChannelID = types.Int64Value(*v.SlackChannelID)
			}
			if v.EmailID != nil {
				notifications[i].EmailID = types.Int64Value(*v.EmailID)
			}
			if v.NotifyWhen != nil {
				notifications[i].NotifyWhen = types.StringValue(*v.NotifyWhen)
			}
			if v.RecordCount != nil {
				notifications[i].RecordCount = types.Int64Value(*v.RecordCount)
			}
			if v.RecordOperator != nil {
				notifications[i].RecordOperator = types.StringValue(*v.RecordOperator)
			}
		}
		model.Notifications = notifications
	}
	if datamartDefinition.Schedules != nil {
		schedules := make([]scheduleModel, len(datamartDefinition.Schedules))
		for i, v := range datamartDefinition.Schedules {
			schedules[i] = scheduleModel{
				Frequency: types.StringValue(v.Frequency),
				TimeZone:  types.StringValue(v.TimeZone),
			}
			if v.Minute != nil {
				schedules[i].Minute = types.Int32Value(int32(*v.Minute))
			}
			if v.Hour != nil {
				schedules[i].Hour = types.Int32Value(int32(*v.Hour))
			}
			if v.DayOfWeek != nil {
				schedules[i].DayOfWeek = types.Int32Value(int32(*v.DayOfWeek))
			}
			if v.Day != nil {
				schedules[i].Day = types.Int32Value(int32(*v.Day))
			}
		}
		model.Schedules = schedules
	}
	if datamartDefinition.Labels != nil {
		labels := make([]labelModel, len(datamartDefinition.Labels))
		for i, v := range datamartDefinition.Labels {
			labels[i] = labelModel{
				ID:   types.Int64Value(v.ID),
				Name: types.StringValue(v.Name),
			}
		}
		model.Labels = labels
	}

	return &model, nil
}
