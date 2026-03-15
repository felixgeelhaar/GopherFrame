# GopherCon Talk Proposal

## Title

**GopherFrame: Building Production Data Pipelines Without Leaving Go**

## Abstract (300 words)

Go backend services increasingly need non-trivial data processing — ETL pipelines, ML feature engineering, real-time analytics. Teams typically bridge to Python for these tasks, creating operational complexity with dual-language deployments and serialization boundaries.

GopherFrame is an Apache Arrow-backed DataFrame library that enables production-grade data processing entirely in Go. By leveraging Arrow's columnar memory format, it achieves 2-428x better performance than existing Go alternatives while maintaining type safety and idiomatic API design.

This talk covers:

1. **Why Arrow matters for Go**: How columnar memory layout enables O(1) column selection, cache-friendly iteration, and zero-copy interoperability with the modern data ecosystem (Parquet, Flight, Spark).

2. **Production-first design decisions**: Why we chose immutability over mutability, hash joins over sort-merge as the default, and expression trees over callback-based filtering. Each decision optimized for production workloads, not notebook exploration.

3. **Real-world patterns**: Live demos of ETL pipelines, window functions for time-series analysis, and ML inference preprocessing — showing how GopherFrame eliminates the Go-Python boundary.

4. **Performance deep dive**: How we achieve 67.8x faster column selection and 428x faster iteration than Gota, with memory profiling showing 2-200x less memory usage.

Attendees will leave understanding how to build data-intensive Go services without Python, when GopherFrame is the right choice versus Pandas/Polars, and how Arrow's columnar format unlocks performance patterns impossible with row-oriented designs.

## Outline

1. The Problem: Go + Python impedance mismatch (5 min)
2. Apache Arrow fundamentals for Go developers (10 min)
3. GopherFrame architecture and design decisions (10 min)
4. Live demo: ETL pipeline in Go (10 min)
5. Performance analysis and benchmarks (5 min)
6. Q&A (5 min)

## Speaker Bio

Felix Geelhaar is a software engineer focused on data infrastructure and Go systems programming. He created GopherFrame to enable Go teams to build complete data pipelines without leaving the Go ecosystem.

## Target Audience

Go developers building backend services that process structured data. No prior data science or DataFrame experience required.
