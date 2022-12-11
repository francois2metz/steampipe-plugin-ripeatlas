package ripeatlas

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func convertTimestamp(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	return time.Unix(int64(input.Value.(int)), 0), nil
}

func tableRipeatlasMeasurement() *plugin.Table {
	return &plugin.Table{
		Name:        "ripeatlas_measurement",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listMeasurement,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "mine",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getMeasurement,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique identifer of the measurement.",
			},
			{
				Name:        "mine",
				Type:        proto.ColumnType_BOOL,
				Description: "",
				Transform:   transform.FromQual("mine"),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the measurement.",
			},
			{
				Name:        "status_id",
				Type:        proto.ColumnType_INT,
				Description: "Status ID.",
				Transform:   transform.FromField("Status.ID"),
			},
			{
				Name:        "status_name",
				Type:        proto.ColumnType_STRING,
				Description: "Status name.",
				Transform:   transform.FromField("Status.Name"),
			},
			{
				Name:        "af",
				Type:        proto.ColumnType_INT,
				Description: "IPv4 of IPv6 Address family of the measurement.",
			},
			{
				Name:        "creation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation date and time of the measurement.",
				Transform:   transform.FromField("CreationTime").NullIfZero().Transform(convertTimestamp),
			},
			{
				Name:        "start_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Configured start time.",
				Transform:   transform.FromField("StartTime").NullIfZero().Transform(convertTimestamp),
			},
			{
				Name:        "stop_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Actual end time of measurement.",
				Transform:   transform.FromField("StopTime").NullIfZero().Transform(convertTimestamp),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "User-defined description of the measurement.",
			},
			{
				Name:        "is_all_scheduled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates if all probe requests have made it through the scheduling process.",
			},
			{
				Name:        "is_oneoff",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates this is a one-off or a recurring measurement.",
			},
			{
				Name:        "is_public",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates this measurement is publicly available.",
			},
			{
				Name:        "is_reachability_test",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates this measurement is a reachability test.",
			},
			{
				Name:        "spread",
				Type:        proto.ColumnType_INT,
				Description: "Distribution of probes' measurements throughout the interval (default is half the interval, maximum 400 seconds).",
			},
			{
				Name:        "resolve_on_probe",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates that a name should be resolved (using DNS) on the probe. Otherwise it will be resolved on the RIPE Atlas servers.",
			},
			{
				Name:        "participant_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of participating probes.",
			},
			{
				Name:        "target_asn",
				Type:        proto.ColumnType_INT,
				Description: "The number of the Autonomous System the IP address of the target belongs to.",
			},
			{
				Name:        "target",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "target_ip",
				Type:        proto.ColumnType_STRING,
				Description: "The IP Address of the target of the measurement.",
			},
			{
				Name:        "target_prefix",
				Type:        proto.ColumnType_STRING,
				Description: "Enclosing prefix of the IP address of the target.",
			},
			{
				Name:        "in_wifi_group",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates this measurement belongs to a wifi measurement group.",
			},
			{
				Name:        "probes_requested",
				Type:        proto.ColumnType_INT,
				Description: "Number of probes requested, but not necessarily granted to this measurement.",
			},
			{
				Name:        "probes_scheduled",
				Type:        proto.ColumnType_INT,
				Description: "Number of probes actually scheduled for this measurement.",
			},
		},
	}
}

func listMeasurement(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_measurement.listMeasurement", "connection_error", err)
		return nil, err
	}
	opts := make(map[string]string)
	quals := d.KeyColumnQuals

	if quals["mine"] != nil {
		opts["mine"] = "true"
	}
	result, err := client.GetMeasurements(opts)
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_measurement.listMeasurement", err)
		return nil, err
	}
	for _, i := range result {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getMeasurement(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_measurement.getMeasurement", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	result, err := client.GetMeasurement(int(id))
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_measurement.getMeasurement", err)
		return nil, err
	}
	return result, nil
}
