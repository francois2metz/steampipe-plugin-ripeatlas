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
				Name:        "af",
				Type:        proto.ColumnType_INT,
				Description: "IPv4 of IPv6 Address family of the measurement.",
			},
			{
				Name:        "creation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation date and time of the measurement (Defaults to unix timestamp format).",
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
				Description: "Actual end time of measurement (as a unix timestamp).",
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
