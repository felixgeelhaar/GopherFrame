// Package interfaces provides the public API facade for GopherFrame.
// This layer exposes a simplified, user-friendly interface that hides the complexity
// of the underlying domain model and coordinates with application services.
package interfaces

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/felixgeelhaar/GopherFrame/pkg/application"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/aggregation"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
)

// DataFrame is the public API wrapper around the domain DataFrame.
// It provides a user-friendly interface with error accumulation for method chaining.
type DataFrame struct {
	domainDF *dataframe.DataFrame
	service  *application.DataFrameService
	err      error
}

// NewDataFrame creates a new public DataFrame from an Arrow Record.
func NewDataFrame(record arrow.Record) *DataFrame {
	return &DataFrame{
		domainDF: dataframe.NewDataFrame(record),
		service:  application.NewDataFrameService(),
	}
}

// LoadParquet loads a DataFrame from a Parquet file.
func LoadParquet(filename string) *DataFrame {
	service := application.NewDataFrameService()
	df, err := service.LoadFromParquet(filename)

	return &DataFrame{
		domainDF: df,
		service:  service,
		err:      err,
	}
}

// NumRows returns the number of rows in the DataFrame.
func (df *DataFrame) NumRows() int64 {
	if df.err != nil || df.domainDF == nil {
		return 0
	}
	return df.domainDF.NumRows()
}

// NumCols returns the number of columns in the DataFrame.
func (df *DataFrame) NumCols() int64 {
	if df.err != nil || df.domainDF == nil {
		return 0
	}
	return df.domainDF.NumCols()
}

// ColumnNames returns the names of all columns.
func (df *DataFrame) ColumnNames() []string {
	if df.err != nil || df.domainDF == nil {
		return nil
	}
	return df.domainDF.ColumnNames()
}

// HasColumn checks if a column exists in the DataFrame.
func (df *DataFrame) HasColumn(name string) bool {
	if df.err != nil || df.domainDF == nil {
		return false
	}
	return df.domainDF.HasColumn(name)
}

// SaveParquet saves the DataFrame to a Parquet file.
func (df *DataFrame) SaveParquet(filename string) error {
	if df.err != nil {
		return df.err
	}
	return df.service.SaveToParquet(df.domainDF, filename)
}

// GroupBy creates a grouped DataFrame for aggregation operations.
func (df *DataFrame) GroupBy(columns ...string) *GroupedDataFrame {
	if df.err != nil {
		return &GroupedDataFrame{err: df.err}
	}

	return &GroupedDataFrame{
		df:      df,
		columns: columns,
	}
}

// Err returns any accumulated error from chained operations.
func (df *DataFrame) Err() error {
	return df.err
}

// Release releases the underlying resources.
func (df *DataFrame) Release() {
	if df.domainDF != nil {
		df.domainDF.Release()
	}
}

// GroupedDataFrame represents a grouped DataFrame ready for aggregation.
type GroupedDataFrame struct {
	df      *DataFrame
	columns []string
	err     error
}

// Sum performs a sum aggregation on the specified column.
func (gdf *GroupedDataFrame) Sum(column string) *DataFrame {
	if gdf.err != nil {
		return &DataFrame{err: gdf.err}
	}

	builder := application.NewGroupByBuilder(gdf.columns...)
	request := builder.Sum(column).Build()

	result := gdf.df.service.GroupBy(gdf.df.domainDF, request)

	return &DataFrame{
		domainDF: result.DataFrame,
		service:  gdf.df.service,
		err:      result.Error,
	}
}

// Mean performs a mean aggregation on the specified column.
func (gdf *GroupedDataFrame) Mean(column string) *DataFrame {
	if gdf.err != nil {
		return &DataFrame{err: gdf.err}
	}

	builder := application.NewGroupByBuilder(gdf.columns...)
	request := builder.Mean(column).Build()

	result := gdf.df.service.GroupBy(gdf.df.domainDF, request)

	return &DataFrame{
		domainDF: result.DataFrame,
		service:  gdf.df.service,
		err:      result.Error,
	}
}

// Count performs a count aggregation on the specified column.
func (gdf *GroupedDataFrame) Count(column string) *DataFrame {
	if gdf.err != nil {
		return &DataFrame{err: gdf.err}
	}

	builder := application.NewGroupByBuilder(gdf.columns...)
	request := builder.Count(column).Build()

	result := gdf.df.service.GroupBy(gdf.df.domainDF, request)

	return &DataFrame{
		domainDF: result.DataFrame,
		service:  gdf.df.service,
		err:      result.Error,
	}
}

// Agg performs multiple aggregations using the builder pattern.
func (gdf *GroupedDataFrame) Agg(specs ...AggregationSpec) *DataFrame {
	if gdf.err != nil {
		return &DataFrame{err: gdf.err}
	}

	// Convert public API specs to domain specs
	domainSpecs := make([]aggregation.AggregationSpec, len(specs))
	for i, spec := range specs {
		domainSpecs[i] = aggregation.AggregationSpec{
			Column:       spec.Column,
			Type:         aggregation.AggregationType(spec.Type),
			Alias:        spec.Alias,
			Percentile:   spec.Percentile,
			SecondColumn: spec.SecondColumn,
		}
	}

	request := aggregation.GroupByRequest{
		GroupColumns: gdf.columns,
		Aggregations: domainSpecs,
	}

	result := gdf.df.service.GroupBy(gdf.df.domainDF, request)

	return &DataFrame{
		domainDF: result.DataFrame,
		service:  gdf.df.service,
		err:      result.Error,
	}
}

// AggregationType represents the type of aggregation operation.
type AggregationType int

// Aggregation type constants
const (
	// SumAgg represents sum aggregation
	SumAgg AggregationType = iota
	MeanAgg
	CountAgg
	MinAgg
	MaxAgg
	PercentileAgg
	MedianAgg
	ModeAgg
	CorrelationAgg
)

// AggregationSpec specifies an aggregation operation for the public API.
type AggregationSpec struct {
	Column       string
	Type         AggregationType
	Alias        string
	Percentile   float64 // For Percentile aggregation (0.0-1.0)
	SecondColumn string  // For Correlation aggregation
}

// Sum creates a sum aggregation specification.
func Sum(column string) AggregationSpec {
	return AggregationSpec{
		Column: column,
		Type:   SumAgg,
		Alias:  column + "_sum",
	}
}

// Mean creates a mean aggregation specification.
func Mean(column string) AggregationSpec {
	return AggregationSpec{
		Column: column,
		Type:   MeanAgg,
		Alias:  column + "_mean",
	}
}

// Count creates a count aggregation specification.
func Count(column string) AggregationSpec {
	return AggregationSpec{
		Column: column,
		Type:   CountAgg,
		Alias:  column + "_count",
	}
}

// Min creates a min aggregation specification.
func Min(column string) AggregationSpec {
	return AggregationSpec{
		Column: column,
		Type:   MinAgg,
		Alias:  column + "_min",
	}
}

// Max creates a max aggregation specification.
func Max(column string) AggregationSpec {
	return AggregationSpec{
		Column: column,
		Type:   MaxAgg,
		Alias:  column + "_max",
	}
}

// Percentile creates a percentile aggregation specification.
// The p parameter should be between 0.0 and 1.0 (e.g., 0.95 for 95th percentile).
func Percentile(column string, p float64) AggregationSpec {
	return AggregationSpec{
		Column:     column,
		Type:       PercentileAgg,
		Alias:      fmt.Sprintf("%s_p%.0f", column, p*100),
		Percentile: p,
	}
}

// Median creates a median aggregation specification (equivalent to 50th percentile).
func Median(column string) AggregationSpec {
	return AggregationSpec{
		Column:     column,
		Type:       MedianAgg,
		Alias:      column + "_median",
		Percentile: 0.5, // Median is 50th percentile
	}
}

// Mode creates a mode aggregation specification (most frequent value).
func Mode(column string) AggregationSpec {
	return AggregationSpec{
		Column: column,
		Type:   ModeAgg,
		Alias:  column + "_mode",
	}
}

// Correlation creates a correlation aggregation specification.
// Calculates the Pearson correlation coefficient between two columns.
func Correlation(column1, column2 string) AggregationSpec {
	return AggregationSpec{
		Column:       column1,
		SecondColumn: column2,
		Type:         CorrelationAgg,
		Alias:        fmt.Sprintf("corr_%s_%s", column1, column2),
	}
}

// As sets a custom alias for the aggregation result.
func (spec AggregationSpec) As(alias string) AggregationSpec {
	spec.Alias = alias
	return spec
}
