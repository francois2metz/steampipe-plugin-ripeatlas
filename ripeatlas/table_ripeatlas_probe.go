package ripeatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableRipeatlasProbe() *plugin.Table {
	return &plugin.Table{
		Name:        "ripeatlas_probe",
		Description: "The list of RIPE Atlas probes.",
		List: &plugin.ListConfig{
			Hydrate:    listProbe,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "country_code",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProbe,
		},
		Columns: []*plugin.Column{
			{
				Name: "id",
				Type: proto.ColumnType_INT,
				Description: "Unique identifer of the probe.",
			},
			{
				Name: "address_v4",
				Type: proto.ColumnType_STRING,
				Description: "The last IPv4 address that was known to be held by this probe, or null if there is no known address. Note: a probe that connects over IPv6 may fail to report its IPv4 address, meaning that this field can sometimes be null even though the probe may have working IPv4.",
			},
			{
				Name: "address_v6",
				Type: proto.ColumnType_STRING,
				Description: "The last IPv6 address that was known to be held by this probe, or null if there is no known address.",
			},
			{
				Name: "asn_v4",
				Type: proto.ColumnType_INT,
				Description: "The IPv4 ASN if any.",
			},
			{
				Name: "asn_v6",
				Type: proto.ColumnType_INT,
				Description: "The IPv6 ASN if any.",
			},
			{
				Name: "country_code",
				Type: proto.ColumnType_STRING,
				Description: "An ISO-3166-1 alpha-2 code indicating the country that this probe is located in, as derived from the user supplied longitude and latitude.",
			},
			{
				Name: "description",
				Type: proto.ColumnType_STRING,
				Description: "User defined description of the probe.",
			},
			{
				Name: "first_connected",
				Type: proto.ColumnType_INT,
				Description: "When the probe connected for the first time (UTC Time).",
			},
			{
				Name: "last_connected",
				Type: proto.ColumnType_INT,
				Description: "When the probe connected for the last time (UTC Time).",
			},
			{
				Name: "geometry_longitude",
				Type: proto.ColumnType_DOUBLE,
				Description: "User-supplied longitude of the probe.",
			},
			{
				Name: "geometry_latitude",
				Type: proto.ColumnType_DOUBLE,
				Description: "User-supplied latitude of the probe.",
			},
			{
				Name: "is_anchor",
				Type: proto.ColumnType_BOOL,
				Description: "Whether or not this probe is a RIPE Atlas Anchor.",
			},
			{
				Name: "is_public",
				Type: proto.ColumnType_BOOL,
				Description: "If a probe is not public then certain details, including exact IP addresses, are not returned.",
			},
			{
				Name: "prefix_v4",
				Type: proto.ColumnType_STRING,
				Description: "The IPv4 prefix if any.",
			},
			{
				Name: "prefix_v6",
				Type: proto.ColumnType_STRING,
				Description: "The IPv6 prefix if any.",
			},
			{
				Name: "tags",
				Type: proto.ColumnType_JSON,
				Description: "Tags.",
			},
		},
	}
}

func listProbe(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_probe.listProbe", "connection_error", err)
		return nil, err
	}
	opts := make(map[string]string)
	quals := d.KeyColumnQuals
	if quals["country_code"] != nil {
		opts["country_code"] = quals["country_code"].GetStringValue()
	}
	result, err := client.GetProbes(opts)
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_probe.listProbe", err)
		return nil, err
	}
	for _, i := range result {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getProbe(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_probe.getProbe", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	result, err := client.GetProbe(int(id))
	if err != nil {
		plugin.Logger(ctx).Error("ripeatlas_probe.getProbe", err)
		return nil, err
	}
	return result, nil
}
