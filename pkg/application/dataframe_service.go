// Package application contains application services that orchestrate domain operations.
// These services coordinate between domain objects and infrastructure concerns.
package application

import (
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/aggregation"
	"github.com/felixgeelhaar/GopherFrame/pkg/domain/dataframe"
	"github.com/felixgeelhaar/GopherFrame/pkg/infrastructure/io"
)

// DataFrameService coordinates DataFrame operations across domain and infrastructure layers.
type DataFrameService struct {
	groupByService *aggregation.GroupByService
	parquetReader  *io.ParquetReader
	parquetWriter  *io.ParquetWriter
}

// NewDataFrameService creates a new DataFrameService with its dependencies.
func NewDataFrameService() *DataFrameService {
	return &DataFrameService{
		groupByService: aggregation.NewGroupByService(),
		parquetReader:  io.NewParquetReader(),
		parquetWriter:  io.NewParquetWriter(),
	}
}

// LoadFromParquet loads a DataFrame from a Parquet file.
func (s *DataFrameService) LoadFromParquet(filename string) (*dataframe.DataFrame, error) {
	return s.parquetReader.ReadFile(filename)
}

// SaveToParquet saves a DataFrame to a Parquet file.
func (s *DataFrameService) SaveToParquet(df *dataframe.DataFrame, filename string) error {
	return s.parquetWriter.WriteFile(df, filename)
}

// GroupBy performs a group-by operation using the domain service.
func (s *DataFrameService) GroupBy(df *dataframe.DataFrame, request aggregation.GroupByRequest) aggregation.GroupByResult {
	return s.groupByService.Execute(df, request)
}

// GroupByBuilder provides a fluent interface for building group-by requests.
type GroupByBuilder struct {
	groupColumns []string
	aggregations []aggregation.AggregationSpec
}

// NewGroupByBuilder creates a new GroupByBuilder.
func NewGroupByBuilder(columns ...string) *GroupByBuilder {
	return &GroupByBuilder{
		groupColumns: columns,
		aggregations: make([]aggregation.AggregationSpec, 0),
	}
}

// Sum adds a sum aggregation to the group-by operation.
func (b *GroupByBuilder) Sum(column string) *GroupByBuilder {
	b.aggregations = append(b.aggregations, aggregation.AggregationSpec{
		Column: column,
		Type:   aggregation.Sum,
		Alias:  column + "_sum",
	})
	return b
}

// Mean adds a mean aggregation to the group-by operation.
func (b *GroupByBuilder) Mean(column string) *GroupByBuilder {
	b.aggregations = append(b.aggregations, aggregation.AggregationSpec{
		Column: column,
		Type:   aggregation.Mean,
		Alias:  column + "_mean",
	})
	return b
}

// Count adds a count aggregation to the group-by operation.
func (b *GroupByBuilder) Count(column string) *GroupByBuilder {
	b.aggregations = append(b.aggregations, aggregation.AggregationSpec{
		Column: column,
		Type:   aggregation.Count,
		Alias:  column + "_count",
	})
	return b
}

// Min adds a min aggregation to the group-by operation.
func (b *GroupByBuilder) Min(column string) *GroupByBuilder {
	b.aggregations = append(b.aggregations, aggregation.AggregationSpec{
		Column: column,
		Type:   aggregation.Min,
		Alias:  column + "_min",
	})
	return b
}

// Max adds a max aggregation to the group-by operation.
func (b *GroupByBuilder) Max(column string) *GroupByBuilder {
	b.aggregations = append(b.aggregations, aggregation.AggregationSpec{
		Column: column,
		Type:   aggregation.Max,
		Alias:  column + "_max",
	})
	return b
}

// As sets a custom alias for the last added aggregation.
func (b *GroupByBuilder) As(alias string) *GroupByBuilder {
	if len(b.aggregations) > 0 {
		b.aggregations[len(b.aggregations)-1].Alias = alias
	}
	return b
}

// Build creates the final GroupByRequest.
func (b *GroupByBuilder) Build() aggregation.GroupByRequest {
	return aggregation.GroupByRequest{
		GroupColumns: b.groupColumns,
		Aggregations: b.aggregations,
	}
}
