package davudpasha

import "github.com/mtnmunuklu/davudpasha-api-client-go/api/common"

type ReportsChartVisualization struct {
	ChartType      *string               `json:"ChartType,omitempty"`
	SerieColors    []string              `json:"SerieColors,omitempty"`
	LineWidth      *int64                `json:"LineWidth,omitempty"`
	InnerRadius    *float64              `json:"InnerRadius,omitempty"`
	LineType       *string               `json:"LineType,omitempty"`
	ShowValues     *bool                 `json:"ShowValues,omitempty"`
	ShowNullValues *bool                 `json:"ShowNullValues,omitempty"`
	XAxis          *ReportsChartXAxis    `json:"XAxis,omitempty"`
	YAxis          *ReportsChartYAxis    `json:"YAxis,omitempty"`
	Legend         *ReportsChartLegend   `json:"Legend,omitempty"`
	UiDPModuleId   *string               `json:"UiDPModuleId,omitempty"`
	ColorScheme    common.NullableString `json:"ColorScheme,omitempty"`
	Settings       *string               `json:"Settings,omitempty"`
	MaxRowCount    *int64                `json:"MaxRowCount,omitempty"`
	SlicesNumber   *int64                `json:"SlicesNumber,omitempty"`
	ShowLabel      *bool                 `json:"ShowLabel,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct.
	UnparsedObject map[string]interface{}
	// AdditionalProperties stores any additional properties not explicitly defined in the struct
	AdditionalProperties map[string]interface{}
}
