window.BENCHMARK_DATA = {
  "lastUpdate": 1753525503263,
  "repoUrl": "https://github.com/felixgeelhaar/GopherFrame",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "b065a205fd667d2954deaa455ce0d551bb9f237d",
          "message": "Fix CI permissions for GitHub Pro features\n\n- Add security-events write permission for SARIF upload\n- Add contents write and pages write for benchmark storage\n- Set fail-on-alert to false to prevent build failures on performance regression\n- Add if: always() to ensure SARIF upload runs even on gosec failures",
          "timestamp": "2025-07-23T22:27:20+02:00",
          "tree_id": "6c5a5572cf27823255cfe0ad316d8daf4c243ea2",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/b065a205fd667d2954deaa455ce0d551bb9f237d"
        },
        "date": 1753302752269,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 44836,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "27895 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 44836,
            "unit": "ns/op",
            "extra": "27895 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "27895 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "27895 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 489283,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2426 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 489283,
            "unit": "ns/op",
            "extra": "2426 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2426 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2426 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4053552,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4053552,
            "unit": "ns/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54663,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "21280 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54663,
            "unit": "ns/op",
            "extra": "21280 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "21280 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21280 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 534987,
            "unit": "ns/op\t  717615 B/op\t     133 allocs/op",
            "extra": "2203 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 534987,
            "unit": "ns/op",
            "extra": "2203 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717615,
            "unit": "B/op",
            "extra": "2203 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2203 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4523981,
            "unit": "ns/op\t 5226408 B/op\t     168 allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4523981,
            "unit": "ns/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226408,
            "unit": "B/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1402,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "822876 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1402,
            "unit": "ns/op",
            "extra": "822876 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "822876 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "822876 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1452,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "806346 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1452,
            "unit": "ns/op",
            "extra": "806346 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "806346 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "806346 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1251,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "833708 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1251,
            "unit": "ns/op",
            "extra": "833708 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "833708 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "833708 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 30085,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "39777 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 30085,
            "unit": "ns/op",
            "extra": "39777 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "39777 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "39777 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 314475,
            "unit": "ns/op\t  596201 B/op\t      74 allocs/op",
            "extra": "3517 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 314475,
            "unit": "ns/op",
            "extra": "3517 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596201,
            "unit": "B/op",
            "extra": "3517 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3517 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2136389,
            "unit": "ns/op\t 4381153 B/op\t      88 allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2136389,
            "unit": "ns/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381153,
            "unit": "B/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 66959,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "17988 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 66959,
            "unit": "ns/op",
            "extra": "17988 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "17988 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "17988 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 552335,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 552335,
            "unit": "ns/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5584904,
            "unit": "ns/op\t 3583646 B/op\t     255 allocs/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5584904,
            "unit": "ns/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583646,
            "unit": "B/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 151704,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8338 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 151704,
            "unit": "ns/op",
            "extra": "8338 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8338 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8338 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1206749,
            "unit": "ns/op\t  429817 B/op\t   20273 allocs/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1206749,
            "unit": "ns/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429817,
            "unit": "B/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 105460,
            "unit": "ns/op\t  119184 B/op\t     173 allocs/op",
            "extra": "9640 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 105460,
            "unit": "ns/op",
            "extra": "9640 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119184,
            "unit": "B/op",
            "extra": "9640 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "9640 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 870625,
            "unit": "ns/op\t 1033043 B/op\t     223 allocs/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 870625,
            "unit": "ns/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033043,
            "unit": "B/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 673390,
            "unit": "ns/op\t  318775 B/op\t    3277 allocs/op",
            "extra": "1747 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 673390,
            "unit": "ns/op",
            "extra": "1747 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318775,
            "unit": "B/op",
            "extra": "1747 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1747 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4043564,
            "unit": "ns/op\t 3481377 B/op\t   30228 allocs/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4043564,
            "unit": "ns/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3481377,
            "unit": "B/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "290 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 285262,
            "unit": "ns/op\t  286080 B/op\t     576 allocs/op",
            "extra": "3992 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 285262,
            "unit": "ns/op",
            "extra": "3992 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 286080,
            "unit": "B/op",
            "extra": "3992 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "3992 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 927514,
            "unit": "ns/op\t 1331220 B/op\t     693 allocs/op",
            "extra": "1279 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 927514,
            "unit": "ns/op",
            "extra": "1279 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1331220,
            "unit": "B/op",
            "extra": "1279 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1279 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1684464,
            "unit": "ns/op\t 1366027 B/op\t   13605 allocs/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1684464,
            "unit": "ns/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366027,
            "unit": "B/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13605,
            "unit": "allocs/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17972691,
            "unit": "ns/op\t14395447 B/op\t  152970 allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17972691,
            "unit": "ns/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14395447,
            "unit": "B/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152970,
            "unit": "allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 596370,
            "unit": "ns/op\t  717631 B/op\t     133 allocs/op",
            "extra": "1827 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 596370,
            "unit": "ns/op",
            "extra": "1827 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717631,
            "unit": "B/op",
            "extra": "1827 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "1827 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 554627,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2102 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 554627,
            "unit": "ns/op",
            "extra": "2102 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2102 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2102 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "108024999ca7119325aaf73bf5be31d31ad5723b",
          "message": "Test workflow permissions update\n\nTrigger CI to verify gh-pages push permissions are working",
          "timestamp": "2025-07-23T22:37:55+02:00",
          "tree_id": "6c5a5572cf27823255cfe0ad316d8daf4c243ea2",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/108024999ca7119325aaf73bf5be31d31ad5723b"
        },
        "date": 1753303142695,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 42551,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "30220 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 42551,
            "unit": "ns/op",
            "extra": "30220 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "30220 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30220 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 523141,
            "unit": "ns/op\t  787913 B/op\t      88 allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 523141,
            "unit": "ns/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787913,
            "unit": "B/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2168 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3883882,
            "unit": "ns/op\t 5801469 B/op\t     111 allocs/op",
            "extra": "318 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3883882,
            "unit": "ns/op",
            "extra": "318 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801469,
            "unit": "B/op",
            "extra": "318 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "318 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54336,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54336,
            "unit": "ns/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 573953,
            "unit": "ns/op\t  717621 B/op\t     133 allocs/op",
            "extra": "2023 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 573953,
            "unit": "ns/op",
            "extra": "2023 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717621,
            "unit": "B/op",
            "extra": "2023 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2023 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4388405,
            "unit": "ns/op\t 5226412 B/op\t     168 allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4388405,
            "unit": "ns/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226412,
            "unit": "B/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1352,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "773161 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1352,
            "unit": "ns/op",
            "extra": "773161 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "773161 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "773161 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1356,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "779776 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1356,
            "unit": "ns/op",
            "extra": "779776 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "779776 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "779776 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1193,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "993544 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1193,
            "unit": "ns/op",
            "extra": "993544 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "993544 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "993544 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28207,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "43436 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28207,
            "unit": "ns/op",
            "extra": "43436 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "43436 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "43436 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 302549,
            "unit": "ns/op\t  596202 B/op\t      74 allocs/op",
            "extra": "3980 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 302549,
            "unit": "ns/op",
            "extra": "3980 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596202,
            "unit": "B/op",
            "extra": "3980 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3980 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1985119,
            "unit": "ns/op\t 4381135 B/op\t      88 allocs/op",
            "extra": "576 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1985119,
            "unit": "ns/op",
            "extra": "576 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381135,
            "unit": "B/op",
            "extra": "576 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "576 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 64958,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18680 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 64958,
            "unit": "ns/op",
            "extra": "18680 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18680 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18680 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 534201,
            "unit": "ns/op\t  259463 B/op\t     184 allocs/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 534201,
            "unit": "ns/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259463,
            "unit": "B/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5213185,
            "unit": "ns/op\t 3583646 B/op\t     255 allocs/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5213185,
            "unit": "ns/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583646,
            "unit": "B/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "228 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 148173,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8271 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 148173,
            "unit": "ns/op",
            "extra": "8271 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8271 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8271 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1223468,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "945 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1223468,
            "unit": "ns/op",
            "extra": "945 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "945 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "945 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 101480,
            "unit": "ns/op\t  119184 B/op\t     173 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 101480,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119184,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 834109,
            "unit": "ns/op\t 1033037 B/op\t     223 allocs/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 834109,
            "unit": "ns/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033037,
            "unit": "B/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 728911,
            "unit": "ns/op\t  319161 B/op\t    3277 allocs/op",
            "extra": "1717 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 728911,
            "unit": "ns/op",
            "extra": "1717 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 319161,
            "unit": "B/op",
            "extra": "1717 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1717 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4670883,
            "unit": "ns/op\t 3477798 B/op\t   30226 allocs/op",
            "extra": "292 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4670883,
            "unit": "ns/op",
            "extra": "292 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3477798,
            "unit": "B/op",
            "extra": "292 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30226,
            "unit": "allocs/op",
            "extra": "292 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 275438,
            "unit": "ns/op\t  285499 B/op\t     576 allocs/op",
            "extra": "4184 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 275438,
            "unit": "ns/op",
            "extra": "4184 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285499,
            "unit": "B/op",
            "extra": "4184 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4184 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 860685,
            "unit": "ns/op\t 1330565 B/op\t     693 allocs/op",
            "extra": "1366 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 860685,
            "unit": "ns/op",
            "extra": "1366 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330565,
            "unit": "B/op",
            "extra": "1366 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1366 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1738012,
            "unit": "ns/op\t 1365428 B/op\t   13592 allocs/op",
            "extra": "710 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1738012,
            "unit": "ns/op",
            "extra": "710 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1365428,
            "unit": "B/op",
            "extra": "710 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13592,
            "unit": "allocs/op",
            "extra": "710 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17257607,
            "unit": "ns/op\t14375164 B/op\t  152547 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17257607,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14375164,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152547,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 567940,
            "unit": "ns/op\t  717626 B/op\t     133 allocs/op",
            "extra": "2046 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 567940,
            "unit": "ns/op",
            "extra": "2046 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717626,
            "unit": "B/op",
            "extra": "2046 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2046 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 547865,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 547865,
            "unit": "ns/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2143 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "108024999ca7119325aaf73bf5be31d31ad5723b",
          "message": "Test workflow permissions update\n\nTrigger CI to verify gh-pages push permissions are working",
          "timestamp": "2025-07-23T22:37:55+02:00",
          "tree_id": "6c5a5572cf27823255cfe0ad316d8daf4c243ea2",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/108024999ca7119325aaf73bf5be31d31ad5723b"
        },
        "date": 1753303320205,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 41016,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "28030 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 41016,
            "unit": "ns/op",
            "extra": "28030 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "28030 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "28030 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 40590,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "29758 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 40590,
            "unit": "ns/op",
            "extra": "29758 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "29758 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "29758 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 40350,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "28737 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 40350,
            "unit": "ns/op",
            "extra": "28737 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "28737 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "28737 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 470717,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 470717,
            "unit": "ns/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2575 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 470468,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 470468,
            "unit": "ns/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2608 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 470689,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2498 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 470689,
            "unit": "ns/op",
            "extra": "2498 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2498 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2498 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3907847,
            "unit": "ns/op\t 5801464 B/op\t     111 allocs/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3907847,
            "unit": "ns/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801464,
            "unit": "B/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3799494,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3799494,
            "unit": "ns/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3851136,
            "unit": "ns/op\t 5801466 B/op\t     111 allocs/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3851136,
            "unit": "ns/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801466,
            "unit": "B/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55539,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21496 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55539,
            "unit": "ns/op",
            "extra": "21496 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21496 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21496 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55265,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21618 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55265,
            "unit": "ns/op",
            "extra": "21618 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21618 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21618 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54505,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54505,
            "unit": "ns/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 552726,
            "unit": "ns/op\t  717617 B/op\t     133 allocs/op",
            "extra": "2154 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 552726,
            "unit": "ns/op",
            "extra": "2154 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717617,
            "unit": "B/op",
            "extra": "2154 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2154 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 556349,
            "unit": "ns/op\t  717618 B/op\t     133 allocs/op",
            "extra": "2257 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 556349,
            "unit": "ns/op",
            "extra": "2257 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717618,
            "unit": "B/op",
            "extra": "2257 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2257 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 556236,
            "unit": "ns/op\t  717619 B/op\t     133 allocs/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 556236,
            "unit": "ns/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717619,
            "unit": "B/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4547321,
            "unit": "ns/op\t 5226436 B/op\t     169 allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4547321,
            "unit": "ns/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226436,
            "unit": "B/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "258 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4533832,
            "unit": "ns/op\t 5226433 B/op\t     169 allocs/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4533832,
            "unit": "ns/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226433,
            "unit": "B/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4487255,
            "unit": "ns/op\t 5226413 B/op\t     168 allocs/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4487255,
            "unit": "ns/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226413,
            "unit": "B/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1361,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "885718 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1361,
            "unit": "ns/op",
            "extra": "885718 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "885718 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "885718 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1436,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "751934 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1436,
            "unit": "ns/op",
            "extra": "751934 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "751934 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "751934 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1388,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "731433 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1388,
            "unit": "ns/op",
            "extra": "731433 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "731433 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "731433 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1404,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "786830 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1404,
            "unit": "ns/op",
            "extra": "786830 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "786830 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "786830 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1405,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "733641 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1405,
            "unit": "ns/op",
            "extra": "733641 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "733641 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "733641 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1394,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "844924 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1394,
            "unit": "ns/op",
            "extra": "844924 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "844924 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "844924 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1214,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "825160 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1214,
            "unit": "ns/op",
            "extra": "825160 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "825160 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "825160 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1204,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "850461 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1204,
            "unit": "ns/op",
            "extra": "850461 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "850461 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "850461 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1215,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "854072 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1215,
            "unit": "ns/op",
            "extra": "854072 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "854072 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "854072 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28926,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "39960 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28926,
            "unit": "ns/op",
            "extra": "39960 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "39960 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "39960 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28434,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40351 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28434,
            "unit": "ns/op",
            "extra": "40351 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40351 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40351 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28710,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40977 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28710,
            "unit": "ns/op",
            "extra": "40977 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40977 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40977 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 298797,
            "unit": "ns/op\t  596199 B/op\t      74 allocs/op",
            "extra": "3898 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 298797,
            "unit": "ns/op",
            "extra": "3898 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596199,
            "unit": "B/op",
            "extra": "3898 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3898 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 320886,
            "unit": "ns/op\t  596196 B/op\t      74 allocs/op",
            "extra": "3724 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 320886,
            "unit": "ns/op",
            "extra": "3724 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596196,
            "unit": "B/op",
            "extra": "3724 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3724 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 320321,
            "unit": "ns/op\t  596200 B/op\t      74 allocs/op",
            "extra": "3698 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 320321,
            "unit": "ns/op",
            "extra": "3698 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596200,
            "unit": "B/op",
            "extra": "3698 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3698 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2094809,
            "unit": "ns/op\t 4381147 B/op\t      88 allocs/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2094809,
            "unit": "ns/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381147,
            "unit": "B/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "598 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2067218,
            "unit": "ns/op\t 4381148 B/op\t      88 allocs/op",
            "extra": "586 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2067218,
            "unit": "ns/op",
            "extra": "586 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381148,
            "unit": "B/op",
            "extra": "586 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "586 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2048540,
            "unit": "ns/op\t 4381144 B/op\t      88 allocs/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2048540,
            "unit": "ns/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381144,
            "unit": "B/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 66355,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18288 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 66355,
            "unit": "ns/op",
            "extra": "18288 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18288 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18288 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 66184,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18096 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 66184,
            "unit": "ns/op",
            "extra": "18096 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18096 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18096 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 66084,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "17991 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 66084,
            "unit": "ns/op",
            "extra": "17991 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "17991 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "17991 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 552105,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 552105,
            "unit": "ns/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 554009,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 554009,
            "unit": "ns/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2143 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 559542,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2215 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 559542,
            "unit": "ns/op",
            "extra": "2215 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2215 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2215 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5419231,
            "unit": "ns/op\t 3583648 B/op\t     255 allocs/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5419231,
            "unit": "ns/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583648,
            "unit": "B/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5386353,
            "unit": "ns/op\t 3583646 B/op\t     255 allocs/op",
            "extra": "220 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5386353,
            "unit": "ns/op",
            "extra": "220 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583646,
            "unit": "B/op",
            "extra": "220 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "220 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5428961,
            "unit": "ns/op\t 3583645 B/op\t     255 allocs/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5428961,
            "unit": "ns/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583645,
            "unit": "B/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 146181,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8242 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 146181,
            "unit": "ns/op",
            "extra": "8242 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8242 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8242 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 148810,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7778 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 148810,
            "unit": "ns/op",
            "extra": "7778 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7778 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7778 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 147652,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7887 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 147652,
            "unit": "ns/op",
            "extra": "7887 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7887 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7887 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1213253,
            "unit": "ns/op\t  429820 B/op\t   20273 allocs/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1213253,
            "unit": "ns/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429820,
            "unit": "B/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1207960,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1207960,
            "unit": "ns/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1186948,
            "unit": "ns/op\t  429820 B/op\t   20273 allocs/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1186948,
            "unit": "ns/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429820,
            "unit": "B/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 102172,
            "unit": "ns/op\t  119184 B/op\t     173 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 102172,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119184,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 102256,
            "unit": "ns/op\t  119183 B/op\t     173 allocs/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 102256,
            "unit": "ns/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119183,
            "unit": "B/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "9919 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 101372,
            "unit": "ns/op\t  119183 B/op\t     173 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 101372,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119183,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 868745,
            "unit": "ns/op\t 1033043 B/op\t     223 allocs/op",
            "extra": "1383 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 868745,
            "unit": "ns/op",
            "extra": "1383 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033043,
            "unit": "B/op",
            "extra": "1383 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1383 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 868335,
            "unit": "ns/op\t 1033043 B/op\t     223 allocs/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 868335,
            "unit": "ns/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033043,
            "unit": "B/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1400 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 883156,
            "unit": "ns/op\t 1033045 B/op\t     223 allocs/op",
            "extra": "1378 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 883156,
            "unit": "ns/op",
            "extra": "1378 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033045,
            "unit": "B/op",
            "extra": "1378 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1378 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 671835,
            "unit": "ns/op\t  318975 B/op\t    3277 allocs/op",
            "extra": "1491 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 671835,
            "unit": "ns/op",
            "extra": "1491 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318975,
            "unit": "B/op",
            "extra": "1491 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1491 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 662519,
            "unit": "ns/op\t  318623 B/op\t    3277 allocs/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 662519,
            "unit": "ns/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318623,
            "unit": "B/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 659395,
            "unit": "ns/op\t  318903 B/op\t    3277 allocs/op",
            "extra": "1765 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 659395,
            "unit": "ns/op",
            "extra": "1765 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318903,
            "unit": "B/op",
            "extra": "1765 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1765 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4033496,
            "unit": "ns/op\t 3483433 B/op\t   30228 allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4033496,
            "unit": "ns/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3483433,
            "unit": "B/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4206136,
            "unit": "ns/op\t 3481120 B/op\t   30227 allocs/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4206136,
            "unit": "ns/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3481120,
            "unit": "B/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "295 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4186900,
            "unit": "ns/op\t 3481286 B/op\t   30228 allocs/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4186900,
            "unit": "ns/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3481286,
            "unit": "B/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 279228,
            "unit": "ns/op\t  285833 B/op\t     576 allocs/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 279228,
            "unit": "ns/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285833,
            "unit": "B/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 278259,
            "unit": "ns/op\t  285499 B/op\t     576 allocs/op",
            "extra": "4023 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 278259,
            "unit": "ns/op",
            "extra": "4023 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285499,
            "unit": "B/op",
            "extra": "4023 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4023 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 283093,
            "unit": "ns/op\t  285709 B/op\t     576 allocs/op",
            "extra": "4218 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 283093,
            "unit": "ns/op",
            "extra": "4218 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285709,
            "unit": "B/op",
            "extra": "4218 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4218 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 932143,
            "unit": "ns/op\t 1330676 B/op\t     693 allocs/op",
            "extra": "1312 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 932143,
            "unit": "ns/op",
            "extra": "1312 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330676,
            "unit": "B/op",
            "extra": "1312 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1312 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 945552,
            "unit": "ns/op\t 1329812 B/op\t     693 allocs/op",
            "extra": "1243 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 945552,
            "unit": "ns/op",
            "extra": "1243 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1329812,
            "unit": "B/op",
            "extra": "1243 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1243 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 938774,
            "unit": "ns/op\t 1330149 B/op\t     693 allocs/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 938774,
            "unit": "ns/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330149,
            "unit": "B/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1639954,
            "unit": "ns/op\t 1366313 B/op\t   13610 allocs/op",
            "extra": "724 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1639954,
            "unit": "ns/op",
            "extra": "724 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366313,
            "unit": "B/op",
            "extra": "724 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13610,
            "unit": "allocs/op",
            "extra": "724 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1680891,
            "unit": "ns/op\t 1366184 B/op\t   13608 allocs/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1680891,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366184,
            "unit": "B/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13608,
            "unit": "allocs/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1693629,
            "unit": "ns/op\t 1367469 B/op\t   13635 allocs/op",
            "extra": "709 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1693629,
            "unit": "ns/op",
            "extra": "709 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1367469,
            "unit": "B/op",
            "extra": "709 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13635,
            "unit": "allocs/op",
            "extra": "709 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17484511,
            "unit": "ns/op\t14349307 B/op\t  152009 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17484511,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14349307,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152009,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17139759,
            "unit": "ns/op\t14374696 B/op\t  152538 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17139759,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14374696,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152538,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17252394,
            "unit": "ns/op\t14388018 B/op\t  152815 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17252394,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14388018,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152815,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 594221,
            "unit": "ns/op\t  717624 B/op\t     133 allocs/op",
            "extra": "1996 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 594221,
            "unit": "ns/op",
            "extra": "1996 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717624,
            "unit": "B/op",
            "extra": "1996 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "1996 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 594408,
            "unit": "ns/op\t  717625 B/op\t     133 allocs/op",
            "extra": "2065 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 594408,
            "unit": "ns/op",
            "extra": "2065 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717625,
            "unit": "B/op",
            "extra": "2065 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2065 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 597051,
            "unit": "ns/op\t  717622 B/op\t     133 allocs/op",
            "extra": "2054 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 597051,
            "unit": "ns/op",
            "extra": "2054 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717622,
            "unit": "B/op",
            "extra": "2054 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2054 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 553859,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2132 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 553859,
            "unit": "ns/op",
            "extra": "2132 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2132 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2132 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 548566,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2145 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 548566,
            "unit": "ns/op",
            "extra": "2145 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2145 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2145 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 556462,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2097 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 556462,
            "unit": "ns/op",
            "extra": "2097 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2097 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2097 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "committer": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "id": "108024999ca7119325aaf73bf5be31d31ad5723b",
          "message": "Test workflow permissions update\n\nTrigger CI to verify gh-pages push permissions are working",
          "timestamp": "2025-07-23T20:37:55Z",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/108024999ca7119325aaf73bf5be31d31ad5723b"
        },
        "date": 1753325206826,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39501,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "30308 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39501,
            "unit": "ns/op",
            "extra": "30308 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "30308 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30308 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39750,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "30301 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39750,
            "unit": "ns/op",
            "extra": "30301 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "30301 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30301 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39157,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "30146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39157,
            "unit": "ns/op",
            "extra": "30146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "30146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 485358,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2443 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 485358,
            "unit": "ns/op",
            "extra": "2443 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2443 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2443 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 567575,
            "unit": "ns/op\t  787910 B/op\t      88 allocs/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 567575,
            "unit": "ns/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787910,
            "unit": "B/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 464274,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 464274,
            "unit": "ns/op",
            "extra": "2354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3607584,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "330 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3607584,
            "unit": "ns/op",
            "extra": "330 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "330 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "330 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3619613,
            "unit": "ns/op\t 5801476 B/op\t     111 allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3619613,
            "unit": "ns/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801476,
            "unit": "B/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4039954,
            "unit": "ns/op\t 5801472 B/op\t     111 allocs/op",
            "extra": "315 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4039954,
            "unit": "ns/op",
            "extra": "315 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801472,
            "unit": "B/op",
            "extra": "315 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "315 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55115,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55115,
            "unit": "ns/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21940 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54729,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21692 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54729,
            "unit": "ns/op",
            "extra": "21692 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21692 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21692 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55043,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21708 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55043,
            "unit": "ns/op",
            "extra": "21708 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21708 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21708 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 577064,
            "unit": "ns/op\t  717618 B/op\t     133 allocs/op",
            "extra": "1975 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 577064,
            "unit": "ns/op",
            "extra": "1975 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717618,
            "unit": "B/op",
            "extra": "1975 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "1975 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 580890,
            "unit": "ns/op\t  717619 B/op\t     133 allocs/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 580890,
            "unit": "ns/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717619,
            "unit": "B/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 576035,
            "unit": "ns/op\t  717622 B/op\t     133 allocs/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 576035,
            "unit": "ns/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717622,
            "unit": "B/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4478115,
            "unit": "ns/op\t 5226409 B/op\t     168 allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4478115,
            "unit": "ns/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226409,
            "unit": "B/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4455956,
            "unit": "ns/op\t 5226417 B/op\t     168 allocs/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4455956,
            "unit": "ns/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226417,
            "unit": "B/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4460449,
            "unit": "ns/op\t 5226410 B/op\t     168 allocs/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4460449,
            "unit": "ns/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226410,
            "unit": "B/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1404,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "775180 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1404,
            "unit": "ns/op",
            "extra": "775180 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "775180 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "775180 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1401,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "791214 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1401,
            "unit": "ns/op",
            "extra": "791214 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "791214 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "791214 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1400,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "757120 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1400,
            "unit": "ns/op",
            "extra": "757120 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "757120 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "757120 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1467,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "822429 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1467,
            "unit": "ns/op",
            "extra": "822429 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "822429 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "822429 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1462,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "754629 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1462,
            "unit": "ns/op",
            "extra": "754629 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "754629 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "754629 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1464,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "763351 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1464,
            "unit": "ns/op",
            "extra": "763351 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "763351 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "763351 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1238,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "859509 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1238,
            "unit": "ns/op",
            "extra": "859509 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "859509 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "859509 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1240,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "843037 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1240,
            "unit": "ns/op",
            "extra": "843037 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "843037 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "843037 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1242,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "862972 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1242,
            "unit": "ns/op",
            "extra": "862972 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "862972 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "862972 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 30035,
            "unit": "ns/op\t   42624 B/op\t      58 allocs/op",
            "extra": "38936 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 30035,
            "unit": "ns/op",
            "extra": "38936 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42624,
            "unit": "B/op",
            "extra": "38936 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "38936 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 29954,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "39764 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 29954,
            "unit": "ns/op",
            "extra": "39764 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "39764 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "39764 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 30088,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40110 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 30088,
            "unit": "ns/op",
            "extra": "40110 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40110 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40110 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 308192,
            "unit": "ns/op\t  596202 B/op\t      74 allocs/op",
            "extra": "3508 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 308192,
            "unit": "ns/op",
            "extra": "3508 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596202,
            "unit": "B/op",
            "extra": "3508 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3508 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 319547,
            "unit": "ns/op\t  596199 B/op\t      74 allocs/op",
            "extra": "3885 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 319547,
            "unit": "ns/op",
            "extra": "3885 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596199,
            "unit": "B/op",
            "extra": "3885 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3885 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 328749,
            "unit": "ns/op\t  596200 B/op\t      74 allocs/op",
            "extra": "3602 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 328749,
            "unit": "ns/op",
            "extra": "3602 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596200,
            "unit": "B/op",
            "extra": "3602 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3602 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2042045,
            "unit": "ns/op\t 4381141 B/op\t      88 allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2042045,
            "unit": "ns/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381141,
            "unit": "B/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1997242,
            "unit": "ns/op\t 4381141 B/op\t      88 allocs/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1997242,
            "unit": "ns/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381141,
            "unit": "B/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2044043,
            "unit": "ns/op\t 4381154 B/op\t      88 allocs/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2044043,
            "unit": "ns/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381154,
            "unit": "B/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 60133,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "19851 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 60133,
            "unit": "ns/op",
            "extra": "19851 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "19851 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "19851 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 60810,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "19488 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 60810,
            "unit": "ns/op",
            "extra": "19488 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "19488 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "19488 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 60461,
            "unit": "ns/op\t   27777 B/op\t     144 allocs/op",
            "extra": "19797 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 60461,
            "unit": "ns/op",
            "extra": "19797 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27777,
            "unit": "B/op",
            "extra": "19797 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "19797 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 525479,
            "unit": "ns/op\t  259463 B/op\t     184 allocs/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 525479,
            "unit": "ns/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259463,
            "unit": "B/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 521138,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 521138,
            "unit": "ns/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 523160,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2310 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 523160,
            "unit": "ns/op",
            "extra": "2310 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2310 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2310 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5102368,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5102368,
            "unit": "ns/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "232 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5140493,
            "unit": "ns/op\t 3583646 B/op\t     255 allocs/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5140493,
            "unit": "ns/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583646,
            "unit": "B/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5243506,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5243506,
            "unit": "ns/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "231 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 133659,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 133659,
            "unit": "ns/op",
            "extra": "8223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8223 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 134021,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8114 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 134021,
            "unit": "ns/op",
            "extra": "8114 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8114 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8114 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 135758,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8380 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 135758,
            "unit": "ns/op",
            "extra": "8380 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8380 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8380 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1153177,
            "unit": "ns/op\t  429818 B/op\t   20273 allocs/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1153177,
            "unit": "ns/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429818,
            "unit": "B/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1149372,
            "unit": "ns/op\t  429820 B/op\t   20273 allocs/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1149372,
            "unit": "ns/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429820,
            "unit": "B/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1041 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1145681,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "1036 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1145681,
            "unit": "ns/op",
            "extra": "1036 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "1036 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1036 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 97551,
            "unit": "ns/op\t  119182 B/op\t     173 allocs/op",
            "extra": "12481 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 97551,
            "unit": "ns/op",
            "extra": "12481 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119182,
            "unit": "B/op",
            "extra": "12481 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12481 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 95550,
            "unit": "ns/op\t  119182 B/op\t     173 allocs/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 95550,
            "unit": "ns/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119182,
            "unit": "B/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 96351,
            "unit": "ns/op\t  119182 B/op\t     173 allocs/op",
            "extra": "12547 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 96351,
            "unit": "ns/op",
            "extra": "12547 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119182,
            "unit": "B/op",
            "extra": "12547 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12547 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 778868,
            "unit": "ns/op\t 1033043 B/op\t     223 allocs/op",
            "extra": "1507 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 778868,
            "unit": "ns/op",
            "extra": "1507 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033043,
            "unit": "B/op",
            "extra": "1507 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1507 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 776863,
            "unit": "ns/op\t 1033036 B/op\t     223 allocs/op",
            "extra": "1509 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 776863,
            "unit": "ns/op",
            "extra": "1509 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033036,
            "unit": "B/op",
            "extra": "1509 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1509 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 777541,
            "unit": "ns/op\t 1033036 B/op\t     223 allocs/op",
            "extra": "1466 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 777541,
            "unit": "ns/op",
            "extra": "1466 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033036,
            "unit": "B/op",
            "extra": "1466 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1466 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 600754,
            "unit": "ns/op\t  318251 B/op\t    3277 allocs/op",
            "extra": "1897 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 600754,
            "unit": "ns/op",
            "extra": "1897 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318251,
            "unit": "B/op",
            "extra": "1897 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1897 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 611929,
            "unit": "ns/op\t  318494 B/op\t    3277 allocs/op",
            "extra": "1935 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 611929,
            "unit": "ns/op",
            "extra": "1935 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318494,
            "unit": "B/op",
            "extra": "1935 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1935 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 607294,
            "unit": "ns/op\t  318048 B/op\t    3277 allocs/op",
            "extra": "1947 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 607294,
            "unit": "ns/op",
            "extra": "1947 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318048,
            "unit": "B/op",
            "extra": "1947 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1947 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3879128,
            "unit": "ns/op\t 3482865 B/op\t   30228 allocs/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3879128,
            "unit": "ns/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3482865,
            "unit": "B/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3827373,
            "unit": "ns/op\t 3483003 B/op\t   30227 allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3827373,
            "unit": "ns/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3483003,
            "unit": "B/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "304 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3912161,
            "unit": "ns/op\t 3482019 B/op\t   30228 allocs/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3912161,
            "unit": "ns/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3482019,
            "unit": "B/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 275643,
            "unit": "ns/op\t  285632 B/op\t     576 allocs/op",
            "extra": "4189 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 275643,
            "unit": "ns/op",
            "extra": "4189 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285632,
            "unit": "B/op",
            "extra": "4189 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4189 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 278011,
            "unit": "ns/op\t  285470 B/op\t     576 allocs/op",
            "extra": "4066 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 278011,
            "unit": "ns/op",
            "extra": "4066 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285470,
            "unit": "B/op",
            "extra": "4066 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4066 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 278081,
            "unit": "ns/op\t  285526 B/op\t     576 allocs/op",
            "extra": "4206 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 278081,
            "unit": "ns/op",
            "extra": "4206 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285526,
            "unit": "B/op",
            "extra": "4206 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4206 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 823555,
            "unit": "ns/op\t 1329925 B/op\t     693 allocs/op",
            "extra": "1454 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 823555,
            "unit": "ns/op",
            "extra": "1454 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1329925,
            "unit": "B/op",
            "extra": "1454 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1454 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 835599,
            "unit": "ns/op\t 1330326 B/op\t     693 allocs/op",
            "extra": "1443 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 835599,
            "unit": "ns/op",
            "extra": "1443 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330326,
            "unit": "B/op",
            "extra": "1443 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1443 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 825373,
            "unit": "ns/op\t 1331197 B/op\t     693 allocs/op",
            "extra": "1474 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 825373,
            "unit": "ns/op",
            "extra": "1474 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1331197,
            "unit": "B/op",
            "extra": "1474 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1474 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1625784,
            "unit": "ns/op\t 1366466 B/op\t   13614 allocs/op",
            "extra": "723 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1625784,
            "unit": "ns/op",
            "extra": "723 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366466,
            "unit": "B/op",
            "extra": "723 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13614,
            "unit": "allocs/op",
            "extra": "723 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1602447,
            "unit": "ns/op\t 1366663 B/op\t   13618 allocs/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1602447,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366663,
            "unit": "B/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13618,
            "unit": "allocs/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1565939,
            "unit": "ns/op\t 1363893 B/op\t   13560 allocs/op",
            "extra": "744 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1565939,
            "unit": "ns/op",
            "extra": "744 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1363893,
            "unit": "B/op",
            "extra": "744 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13560,
            "unit": "allocs/op",
            "extra": "744 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 16814156,
            "unit": "ns/op\t14353998 B/op\t  152106 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 16814156,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14353998,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152106,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 16744545,
            "unit": "ns/op\t14403763 B/op\t  153143 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 16744545,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14403763,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 153143,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 16560947,
            "unit": "ns/op\t14353819 B/op\t  152103 allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 16560947,
            "unit": "ns/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14353819,
            "unit": "B/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152103,
            "unit": "allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 538557,
            "unit": "ns/op\t  717617 B/op\t     133 allocs/op",
            "extra": "2250 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 538557,
            "unit": "ns/op",
            "extra": "2250 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717617,
            "unit": "B/op",
            "extra": "2250 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2250 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 532988,
            "unit": "ns/op\t  717611 B/op\t     133 allocs/op",
            "extra": "2235 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 532988,
            "unit": "ns/op",
            "extra": "2235 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717611,
            "unit": "B/op",
            "extra": "2235 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2235 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 532187,
            "unit": "ns/op\t  717615 B/op\t     133 allocs/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 532187,
            "unit": "ns/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717615,
            "unit": "B/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 530305,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 530305,
            "unit": "ns/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2218 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 532331,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 532331,
            "unit": "ns/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2205 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 534574,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2140 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 534574,
            "unit": "ns/op",
            "extra": "2140 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2140 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2140 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "ac6e1ecdda96eb48516395d277d19964648e11b6",
          "message": "Implement comprehensive join functionality for DataFrames\n\n- Add InnerJoin and LeftJoin methods to public DataFrame API\n- Implement hash-based join algorithm in core DataFrame\n- Support for all Arrow data types (int64, float64, string, boolean)\n- Handle duplicate column names with automatic prefixing\n- Comprehensive null value handling for join keys\n- Memory-efficient with proper resource management\n- Add extensive test coverage with 8 comprehensive test cases\n- Performance tested with 10K+ row datasets\n- Exceeds v0.1 MVP scope - joins were planned for v0.2\n\nFeatures implemented:\n- df.InnerJoin(other, leftKey, rightKey) - rows with matches in both\n- df.LeftJoin(other, leftKey, rightKey) - all left rows + matching right\n- Column conflict resolution (right_columnname for duplicates)\n- Type-safe value extraction and comparison\n- Proper bounds checking and error handling\n\nAll tests passing with comprehensive edge case coverage.",
          "timestamp": "2025-07-24T10:51:59+02:00",
          "tree_id": "a8fa5d2c431efabb93a07afc2ed71cb1d34a6b9f",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/ac6e1ecdda96eb48516395d277d19964648e11b6"
        },
        "date": 1753347737194,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39417,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "30685 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39417,
            "unit": "ns/op",
            "extra": "30685 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "30685 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30685 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 463749,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2444 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 463749,
            "unit": "ns/op",
            "extra": "2444 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2444 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2444 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4424428,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4424428,
            "unit": "ns/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "298 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 53719,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "21523 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 53719,
            "unit": "ns/op",
            "extra": "21523 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "21523 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21523 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 541960,
            "unit": "ns/op\t  717612 B/op\t     133 allocs/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 541960,
            "unit": "ns/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717612,
            "unit": "B/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4443907,
            "unit": "ns/op\t 5226401 B/op\t     168 allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4443907,
            "unit": "ns/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226401,
            "unit": "B/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1301,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "787513 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1301,
            "unit": "ns/op",
            "extra": "787513 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "787513 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "787513 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1322,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "865016 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1322,
            "unit": "ns/op",
            "extra": "865016 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "865016 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "865016 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1151,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "908563 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1151,
            "unit": "ns/op",
            "extra": "908563 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "908563 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "908563 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 27706,
            "unit": "ns/op\t   42622 B/op\t      58 allocs/op",
            "extra": "42667 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 27706,
            "unit": "ns/op",
            "extra": "42667 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42622,
            "unit": "B/op",
            "extra": "42667 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "42667 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 305502,
            "unit": "ns/op\t  596199 B/op\t      74 allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 305502,
            "unit": "ns/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596199,
            "unit": "B/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2081478,
            "unit": "ns/op\t 4381140 B/op\t      88 allocs/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2081478,
            "unit": "ns/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381140,
            "unit": "B/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 64997,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18625 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 64997,
            "unit": "ns/op",
            "extra": "18625 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18625 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18625 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 532931,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 532931,
            "unit": "ns/op",
            "extra": "2224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5255294,
            "unit": "ns/op\t 3583645 B/op\t     255 allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5255294,
            "unit": "ns/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583645,
            "unit": "B/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 143341,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8996 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 143341,
            "unit": "ns/op",
            "extra": "8996 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8996 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8996 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1160250,
            "unit": "ns/op\t  429818 B/op\t   20273 allocs/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1160250,
            "unit": "ns/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429818,
            "unit": "B/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1017 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 99751,
            "unit": "ns/op\t  119182 B/op\t     173 allocs/op",
            "extra": "12177 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 99751,
            "unit": "ns/op",
            "extra": "12177 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119182,
            "unit": "B/op",
            "extra": "12177 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12177 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 825794,
            "unit": "ns/op\t 1033035 B/op\t     223 allocs/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 825794,
            "unit": "ns/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033035,
            "unit": "B/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 717123,
            "unit": "ns/op\t  318370 B/op\t    3277 allocs/op",
            "extra": "1665 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 717123,
            "unit": "ns/op",
            "extra": "1665 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318370,
            "unit": "B/op",
            "extra": "1665 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1665 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3879780,
            "unit": "ns/op\t 3479497 B/op\t   30226 allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3879780,
            "unit": "ns/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3479497,
            "unit": "B/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30226,
            "unit": "allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 276431,
            "unit": "ns/op\t  285403 B/op\t     576 allocs/op",
            "extra": "4239 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 276431,
            "unit": "ns/op",
            "extra": "4239 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285403,
            "unit": "B/op",
            "extra": "4239 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4239 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 871984,
            "unit": "ns/op\t 1330202 B/op\t     693 allocs/op",
            "extra": "1350 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 871984,
            "unit": "ns/op",
            "extra": "1350 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330202,
            "unit": "B/op",
            "extra": "1350 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1350 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1707487,
            "unit": "ns/op\t 1368109 B/op\t   13648 allocs/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1707487,
            "unit": "ns/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1368109,
            "unit": "B/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13648,
            "unit": "allocs/op",
            "extra": "674 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17023577,
            "unit": "ns/op\t14378718 B/op\t  152621 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17023577,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14378718,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152621,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 569598,
            "unit": "ns/op\t  717621 B/op\t     133 allocs/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 569598,
            "unit": "ns/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717621,
            "unit": "B/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 544296,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2186 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 544296,
            "unit": "ns/op",
            "extra": "2186 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2186 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2186 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "ac6e1ecdda96eb48516395d277d19964648e11b6",
          "message": "Implement comprehensive join functionality for DataFrames\n\n- Add InnerJoin and LeftJoin methods to public DataFrame API\n- Implement hash-based join algorithm in core DataFrame\n- Support for all Arrow data types (int64, float64, string, boolean)\n- Handle duplicate column names with automatic prefixing\n- Comprehensive null value handling for join keys\n- Memory-efficient with proper resource management\n- Add extensive test coverage with 8 comprehensive test cases\n- Performance tested with 10K+ row datasets\n- Exceeds v0.1 MVP scope - joins were planned for v0.2\n\nFeatures implemented:\n- df.InnerJoin(other, leftKey, rightKey) - rows with matches in both\n- df.LeftJoin(other, leftKey, rightKey) - all left rows + matching right\n- Column conflict resolution (right_columnname for duplicates)\n- Type-safe value extraction and comparison\n- Proper bounds checking and error handling\n\nAll tests passing with comprehensive edge case coverage.",
          "timestamp": "2025-07-24T10:51:59+02:00",
          "tree_id": "a8fa5d2c431efabb93a07afc2ed71cb1d34a6b9f",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/ac6e1ecdda96eb48516395d277d19964648e11b6"
        },
        "date": 1753347912012,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39204,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "30486 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39204,
            "unit": "ns/op",
            "extra": "30486 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "30486 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30486 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39591,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "29740 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39591,
            "unit": "ns/op",
            "extra": "29740 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "29740 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "29740 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39826,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "29727 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39826,
            "unit": "ns/op",
            "extra": "29727 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "29727 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "29727 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 463373,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2499 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 463373,
            "unit": "ns/op",
            "extra": "2499 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2499 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2499 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 471636,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 471636,
            "unit": "ns/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2430 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 473845,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2502 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 473845,
            "unit": "ns/op",
            "extra": "2502 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2502 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2502 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4146356,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4146356,
            "unit": "ns/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3685338,
            "unit": "ns/op\t 5801466 B/op\t     111 allocs/op",
            "extra": "319 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3685338,
            "unit": "ns/op",
            "extra": "319 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801466,
            "unit": "B/op",
            "extra": "319 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "319 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3938926,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3938926,
            "unit": "ns/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 53769,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "21742 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 53769,
            "unit": "ns/op",
            "extra": "21742 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "21742 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21742 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54556,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21814 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54556,
            "unit": "ns/op",
            "extra": "21814 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21814 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21814 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54195,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54195,
            "unit": "ns/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21903 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 578940,
            "unit": "ns/op\t  717619 B/op\t     133 allocs/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 578940,
            "unit": "ns/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717619,
            "unit": "B/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2074 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 578803,
            "unit": "ns/op\t  717616 B/op\t     133 allocs/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 578803,
            "unit": "ns/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717616,
            "unit": "B/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2138 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 572285,
            "unit": "ns/op\t  717621 B/op\t     133 allocs/op",
            "extra": "2169 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 572285,
            "unit": "ns/op",
            "extra": "2169 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717621,
            "unit": "B/op",
            "extra": "2169 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2169 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4453459,
            "unit": "ns/op\t 5226404 B/op\t     168 allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4453459,
            "unit": "ns/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226404,
            "unit": "B/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "271 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4451520,
            "unit": "ns/op\t 5226415 B/op\t     168 allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4451520,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226415,
            "unit": "B/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4457319,
            "unit": "ns/op\t 5226379 B/op\t     168 allocs/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4457319,
            "unit": "ns/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226379,
            "unit": "B/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "268 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1351,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "784652 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1351,
            "unit": "ns/op",
            "extra": "784652 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "784652 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "784652 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1360,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "816162 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1360,
            "unit": "ns/op",
            "extra": "816162 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "816162 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "816162 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1372,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "825349 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1372,
            "unit": "ns/op",
            "extra": "825349 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "825349 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "825349 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1398,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "841040 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1398,
            "unit": "ns/op",
            "extra": "841040 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "841040 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "841040 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1392,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "816488 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1392,
            "unit": "ns/op",
            "extra": "816488 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "816488 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "816488 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1389,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "795252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1389,
            "unit": "ns/op",
            "extra": "795252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "795252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "795252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1205,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "898081 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1205,
            "unit": "ns/op",
            "extra": "898081 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "898081 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "898081 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1199,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "908472 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1199,
            "unit": "ns/op",
            "extra": "908472 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "908472 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "908472 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1197,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "863486 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1197,
            "unit": "ns/op",
            "extra": "863486 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "863486 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "863486 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 29778,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40628 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 29778,
            "unit": "ns/op",
            "extra": "40628 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40628 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40628 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 29696,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "41144 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 29696,
            "unit": "ns/op",
            "extra": "41144 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "41144 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "41144 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 29955,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40989 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 29955,
            "unit": "ns/op",
            "extra": "40989 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40989 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40989 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 318038,
            "unit": "ns/op\t  596194 B/op\t      74 allocs/op",
            "extra": "3621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 318038,
            "unit": "ns/op",
            "extra": "3621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596194,
            "unit": "B/op",
            "extra": "3621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 312569,
            "unit": "ns/op\t  596197 B/op\t      74 allocs/op",
            "extra": "3876 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 312569,
            "unit": "ns/op",
            "extra": "3876 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596197,
            "unit": "B/op",
            "extra": "3876 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3876 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 311131,
            "unit": "ns/op\t  596198 B/op\t      74 allocs/op",
            "extra": "3718 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 311131,
            "unit": "ns/op",
            "extra": "3718 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596198,
            "unit": "B/op",
            "extra": "3718 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3718 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1964098,
            "unit": "ns/op\t 4381144 B/op\t      88 allocs/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1964098,
            "unit": "ns/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381144,
            "unit": "B/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "615 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1932306,
            "unit": "ns/op\t 4381134 B/op\t      88 allocs/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1932306,
            "unit": "ns/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381134,
            "unit": "B/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1988744,
            "unit": "ns/op\t 4381143 B/op\t      88 allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1988744,
            "unit": "ns/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381143,
            "unit": "B/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 60366,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "19681 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 60366,
            "unit": "ns/op",
            "extra": "19681 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "19681 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "19681 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 60518,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "19981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 60518,
            "unit": "ns/op",
            "extra": "19981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "19981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "19981 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 60615,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "19760 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 60615,
            "unit": "ns/op",
            "extra": "19760 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "19760 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "19760 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 518375,
            "unit": "ns/op\t  259463 B/op\t     184 allocs/op",
            "extra": "2263 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 518375,
            "unit": "ns/op",
            "extra": "2263 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259463,
            "unit": "B/op",
            "extra": "2263 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2263 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 513756,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2300 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 513756,
            "unit": "ns/op",
            "extra": "2300 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2300 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2300 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 517055,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2326 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 517055,
            "unit": "ns/op",
            "extra": "2326 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2326 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2326 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5105568,
            "unit": "ns/op\t 3583650 B/op\t     255 allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5105568,
            "unit": "ns/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583650,
            "unit": "B/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5095522,
            "unit": "ns/op\t 3583648 B/op\t     255 allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5095522,
            "unit": "ns/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583648,
            "unit": "B/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5100712,
            "unit": "ns/op\t 3583653 B/op\t     255 allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5100712,
            "unit": "ns/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583653,
            "unit": "B/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 133818,
            "unit": "ns/op\t   54130 B/op\t    2233 allocs/op",
            "extra": "8667 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 133818,
            "unit": "ns/op",
            "extra": "8667 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54130,
            "unit": "B/op",
            "extra": "8667 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8667 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 132473,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 132473,
            "unit": "ns/op",
            "extra": "8182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 135640,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8750 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 135640,
            "unit": "ns/op",
            "extra": "8750 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8750 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8750 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1135110,
            "unit": "ns/op\t  429818 B/op\t   20273 allocs/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1135110,
            "unit": "ns/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429818,
            "unit": "B/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1068 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1124762,
            "unit": "ns/op\t  429818 B/op\t   20273 allocs/op",
            "extra": "1038 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1124762,
            "unit": "ns/op",
            "extra": "1038 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429818,
            "unit": "B/op",
            "extra": "1038 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1038 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1132712,
            "unit": "ns/op\t  429821 B/op\t   20273 allocs/op",
            "extra": "1040 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1132712,
            "unit": "ns/op",
            "extra": "1040 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429821,
            "unit": "B/op",
            "extra": "1040 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1040 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 94791,
            "unit": "ns/op\t  119181 B/op\t     173 allocs/op",
            "extra": "12535 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 94791,
            "unit": "ns/op",
            "extra": "12535 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119181,
            "unit": "B/op",
            "extra": "12535 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12535 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 94829,
            "unit": "ns/op\t  119181 B/op\t     173 allocs/op",
            "extra": "12603 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 94829,
            "unit": "ns/op",
            "extra": "12603 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119181,
            "unit": "B/op",
            "extra": "12603 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12603 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 96324,
            "unit": "ns/op\t  119181 B/op\t     173 allocs/op",
            "extra": "12627 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 96324,
            "unit": "ns/op",
            "extra": "12627 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119181,
            "unit": "B/op",
            "extra": "12627 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12627 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 797515,
            "unit": "ns/op\t 1033033 B/op\t     223 allocs/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 797515,
            "unit": "ns/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033033,
            "unit": "B/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 785410,
            "unit": "ns/op\t 1033032 B/op\t     223 allocs/op",
            "extra": "1532 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 785410,
            "unit": "ns/op",
            "extra": "1532 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033032,
            "unit": "B/op",
            "extra": "1532 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1532 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 787334,
            "unit": "ns/op\t 1033031 B/op\t     223 allocs/op",
            "extra": "1497 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 787334,
            "unit": "ns/op",
            "extra": "1497 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033031,
            "unit": "B/op",
            "extra": "1497 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1497 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 721122,
            "unit": "ns/op\t  318379 B/op\t    3277 allocs/op",
            "extra": "1656 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 721122,
            "unit": "ns/op",
            "extra": "1656 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318379,
            "unit": "B/op",
            "extra": "1656 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1656 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 677900,
            "unit": "ns/op\t  318342 B/op\t    3277 allocs/op",
            "extra": "1635 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 677900,
            "unit": "ns/op",
            "extra": "1635 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318342,
            "unit": "B/op",
            "extra": "1635 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1635 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 680049,
            "unit": "ns/op\t  318233 B/op\t    3277 allocs/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 680049,
            "unit": "ns/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318233,
            "unit": "B/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3831771,
            "unit": "ns/op\t 3480991 B/op\t   30227 allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3831771,
            "unit": "ns/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3480991,
            "unit": "B/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3904138,
            "unit": "ns/op\t 3481202 B/op\t   30227 allocs/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3904138,
            "unit": "ns/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3481202,
            "unit": "B/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "294 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3734257,
            "unit": "ns/op\t 3478192 B/op\t   30225 allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3734257,
            "unit": "ns/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3478192,
            "unit": "B/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30225,
            "unit": "allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 272168,
            "unit": "ns/op\t  285264 B/op\t     576 allocs/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 272168,
            "unit": "ns/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285264,
            "unit": "B/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4254 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 274077,
            "unit": "ns/op\t  285661 B/op\t     576 allocs/op",
            "extra": "4214 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 274077,
            "unit": "ns/op",
            "extra": "4214 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285661,
            "unit": "B/op",
            "extra": "4214 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4214 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 278237,
            "unit": "ns/op\t  285606 B/op\t     576 allocs/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 278237,
            "unit": "ns/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285606,
            "unit": "B/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 821421,
            "unit": "ns/op\t 1330857 B/op\t     693 allocs/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 821421,
            "unit": "ns/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330857,
            "unit": "B/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 818066,
            "unit": "ns/op\t 1330006 B/op\t     693 allocs/op",
            "extra": "1392 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 818066,
            "unit": "ns/op",
            "extra": "1392 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330006,
            "unit": "B/op",
            "extra": "1392 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1392 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 813900,
            "unit": "ns/op\t 1330100 B/op\t     693 allocs/op",
            "extra": "1498 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 813900,
            "unit": "ns/op",
            "extra": "1498 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330100,
            "unit": "B/op",
            "extra": "1498 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1498 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1710957,
            "unit": "ns/op\t 1366106 B/op\t   13606 allocs/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1710957,
            "unit": "ns/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366106,
            "unit": "B/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13606,
            "unit": "allocs/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1621700,
            "unit": "ns/op\t 1364802 B/op\t   13579 allocs/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1621700,
            "unit": "ns/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1364802,
            "unit": "B/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13579,
            "unit": "allocs/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1608663,
            "unit": "ns/op\t 1364610 B/op\t   13575 allocs/op",
            "extra": "745 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1608663,
            "unit": "ns/op",
            "extra": "745 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1364610,
            "unit": "B/op",
            "extra": "745 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13575,
            "unit": "allocs/op",
            "extra": "745 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 16888991,
            "unit": "ns/op\t14360065 B/op\t  152233 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 16888991,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14360065,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152233,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 16578627,
            "unit": "ns/op\t14330873 B/op\t  151625 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 16578627,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14330873,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 151625,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 16654790,
            "unit": "ns/op\t14366889 B/op\t  152375 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 16654790,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14366889,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152375,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 534498,
            "unit": "ns/op\t  717617 B/op\t     133 allocs/op",
            "extra": "2180 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 534498,
            "unit": "ns/op",
            "extra": "2180 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717617,
            "unit": "B/op",
            "extra": "2180 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2180 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 539005,
            "unit": "ns/op\t  717614 B/op\t     133 allocs/op",
            "extra": "2113 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 539005,
            "unit": "ns/op",
            "extra": "2113 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717614,
            "unit": "B/op",
            "extra": "2113 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2113 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 540087,
            "unit": "ns/op\t  717611 B/op\t     133 allocs/op",
            "extra": "2137 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 540087,
            "unit": "ns/op",
            "extra": "2137 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717611,
            "unit": "B/op",
            "extra": "2137 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2137 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 536694,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2107 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 536694,
            "unit": "ns/op",
            "extra": "2107 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2107 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2107 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 539640,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 539640,
            "unit": "ns/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2210 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 542258,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2181 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 542258,
            "unit": "ns/op",
            "extra": "2181 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2181 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2181 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "627ff698d0f09b1b3e4c018f3564577a8fb4976b",
          "message": "Add comprehensive string manipulation functions to expression system\n\n- Extend Expr interface with Contains, StartsWith, EndsWith methods\n- Implement string operations for all expression types (Column, Literal, Binary)\n- Add evaluation logic for string comparisons using Go strings package\n- Support for proper null value handling in string operations\n- Type safety with clear error messages for non-string operands\n\nFeatures added:\n- df.Col('name').Contains(Lit('substring')) - check if string contains substring\n- df.Col('email').StartsWith(Lit('prefix')) - check if string starts with prefix\n- df.Col('file').EndsWith(Lit('suffix')) - check if string ends with suffix\n- Chainable with other operations for complex filtering\n- Case-sensitive string matching (as expected in production systems)\n\nComprehensive test coverage with 10 test scenarios:\n- Basic functionality for all three operations\n- WithColumn integration for boolean result columns\n- Case sensitivity validation\n- Empty string handling\n- Null value propagation\n- Type error validation\n- Chained expression filtering\n- Performance testing with 10K records\n\nAll tests passing with proper Arrow boolean array results.\nFix join test column indexing for reliable test execution.",
          "timestamp": "2025-07-24T12:50:59+02:00",
          "tree_id": "60c5736a3a01c6fac696b74e3a878aefcca943cd",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/627ff698d0f09b1b3e4c018f3564577a8fb4976b"
        },
        "date": 1753354360869,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 40201,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "29300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 40201,
            "unit": "ns/op",
            "extra": "29300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "29300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "29300 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 464775,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2361 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 464775,
            "unit": "ns/op",
            "extra": "2361 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2361 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2361 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3820124,
            "unit": "ns/op\t 5801467 B/op\t     111 allocs/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3820124,
            "unit": "ns/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801467,
            "unit": "B/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 53413,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "22503 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 53413,
            "unit": "ns/op",
            "extra": "22503 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "22503 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "22503 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 535374,
            "unit": "ns/op\t  717612 B/op\t     133 allocs/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 535374,
            "unit": "ns/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717612,
            "unit": "B/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2160 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4601528,
            "unit": "ns/op\t 5226387 B/op\t     168 allocs/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4601528,
            "unit": "ns/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226387,
            "unit": "B/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "274 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1496,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "674055 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1496,
            "unit": "ns/op",
            "extra": "674055 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "674055 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "674055 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1363,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "832911 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1363,
            "unit": "ns/op",
            "extra": "832911 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "832911 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "832911 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1183,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "897002 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1183,
            "unit": "ns/op",
            "extra": "897002 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "897002 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "897002 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28594,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "42112 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28594,
            "unit": "ns/op",
            "extra": "42112 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "42112 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "42112 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 311524,
            "unit": "ns/op\t  596197 B/op\t      74 allocs/op",
            "extra": "3762 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 311524,
            "unit": "ns/op",
            "extra": "3762 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596197,
            "unit": "B/op",
            "extra": "3762 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3762 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1937714,
            "unit": "ns/op\t 4381138 B/op\t      88 allocs/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1937714,
            "unit": "ns/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381138,
            "unit": "B/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 61057,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "19525 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 61057,
            "unit": "ns/op",
            "extra": "19525 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "19525 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "19525 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 523851,
            "unit": "ns/op\t  259463 B/op\t     184 allocs/op",
            "extra": "2290 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 523851,
            "unit": "ns/op",
            "extra": "2290 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259463,
            "unit": "B/op",
            "extra": "2290 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2290 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5249775,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5249775,
            "unit": "ns/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 142668,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8052 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 142668,
            "unit": "ns/op",
            "extra": "8052 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8052 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8052 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1151715,
            "unit": "ns/op\t  429818 B/op\t   20273 allocs/op",
            "extra": "1053 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1151715,
            "unit": "ns/op",
            "extra": "1053 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429818,
            "unit": "B/op",
            "extra": "1053 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1053 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 99819,
            "unit": "ns/op\t  119182 B/op\t     173 allocs/op",
            "extra": "12111 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 99819,
            "unit": "ns/op",
            "extra": "12111 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119182,
            "unit": "B/op",
            "extra": "12111 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12111 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 808608,
            "unit": "ns/op\t 1033034 B/op\t     223 allocs/op",
            "extra": "1485 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 808608,
            "unit": "ns/op",
            "extra": "1485 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033034,
            "unit": "B/op",
            "extra": "1485 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1485 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 713419,
            "unit": "ns/op\t  318374 B/op\t    3277 allocs/op",
            "extra": "1704 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 713419,
            "unit": "ns/op",
            "extra": "1704 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318374,
            "unit": "B/op",
            "extra": "1704 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1704 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4020138,
            "unit": "ns/op\t 3479392 B/op\t   30227 allocs/op",
            "extra": "303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4020138,
            "unit": "ns/op",
            "extra": "303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3479392,
            "unit": "B/op",
            "extra": "303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 275275,
            "unit": "ns/op\t  285445 B/op\t     576 allocs/op",
            "extra": "4147 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 275275,
            "unit": "ns/op",
            "extra": "4147 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285445,
            "unit": "B/op",
            "extra": "4147 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4147 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 869182,
            "unit": "ns/op\t 1330302 B/op\t     693 allocs/op",
            "extra": "1352 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 869182,
            "unit": "ns/op",
            "extra": "1352 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330302,
            "unit": "B/op",
            "extra": "1352 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1352 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1731603,
            "unit": "ns/op\t 1367593 B/op\t   13637 allocs/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1731603,
            "unit": "ns/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1367593,
            "unit": "B/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13637,
            "unit": "allocs/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 16635128,
            "unit": "ns/op\t14393108 B/op\t  152921 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 16635128,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14393108,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152921,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 532648,
            "unit": "ns/op\t  717623 B/op\t     133 allocs/op",
            "extra": "2181 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 532648,
            "unit": "ns/op",
            "extra": "2181 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717623,
            "unit": "B/op",
            "extra": "2181 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2181 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 533885,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2182 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 533885,
            "unit": "ns/op",
            "extra": "2182 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2182 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2182 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "627ff698d0f09b1b3e4c018f3564577a8fb4976b",
          "message": "Add comprehensive string manipulation functions to expression system\n\n- Extend Expr interface with Contains, StartsWith, EndsWith methods\n- Implement string operations for all expression types (Column, Literal, Binary)\n- Add evaluation logic for string comparisons using Go strings package\n- Support for proper null value handling in string operations\n- Type safety with clear error messages for non-string operands\n\nFeatures added:\n- df.Col('name').Contains(Lit('substring')) - check if string contains substring\n- df.Col('email').StartsWith(Lit('prefix')) - check if string starts with prefix\n- df.Col('file').EndsWith(Lit('suffix')) - check if string ends with suffix\n- Chainable with other operations for complex filtering\n- Case-sensitive string matching (as expected in production systems)\n\nComprehensive test coverage with 10 test scenarios:\n- Basic functionality for all three operations\n- WithColumn integration for boolean result columns\n- Case sensitivity validation\n- Empty string handling\n- Null value propagation\n- Type error validation\n- Chained expression filtering\n- Performance testing with 10K records\n\nAll tests passing with proper Arrow boolean array results.\nFix join test column indexing for reliable test execution.",
          "timestamp": "2025-07-24T12:50:59+02:00",
          "tree_id": "60c5736a3a01c6fac696b74e3a878aefcca943cd",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/627ff698d0f09b1b3e4c018f3564577a8fb4976b"
        },
        "date": 1753354537125,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 42159,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "28275 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 42159,
            "unit": "ns/op",
            "extra": "28275 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "28275 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "28275 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 42448,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "28378 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 42448,
            "unit": "ns/op",
            "extra": "28378 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "28378 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "28378 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 42721,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "28146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 42721,
            "unit": "ns/op",
            "extra": "28146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "28146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "28146 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 495029,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 495029,
            "unit": "ns/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 496017,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 496017,
            "unit": "ns/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2427 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 513302,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2390 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 513302,
            "unit": "ns/op",
            "extra": "2390 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2390 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2390 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4292960,
            "unit": "ns/op\t 5801464 B/op\t     111 allocs/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4292960,
            "unit": "ns/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801464,
            "unit": "B/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "272 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4159053,
            "unit": "ns/op\t 5801464 B/op\t     111 allocs/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4159053,
            "unit": "ns/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801464,
            "unit": "B/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "278 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4014117,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4014117,
            "unit": "ns/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55138,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "22188 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55138,
            "unit": "ns/op",
            "extra": "22188 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "22188 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "22188 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54923,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "21601 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54923,
            "unit": "ns/op",
            "extra": "21601 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "21601 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21601 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55467,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "21360 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55467,
            "unit": "ns/op",
            "extra": "21360 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "21360 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21360 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 559877,
            "unit": "ns/op\t  717610 B/op\t     133 allocs/op",
            "extra": "2001 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 559877,
            "unit": "ns/op",
            "extra": "2001 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717610,
            "unit": "B/op",
            "extra": "2001 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2001 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 561813,
            "unit": "ns/op\t  717616 B/op\t     133 allocs/op",
            "extra": "2161 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 561813,
            "unit": "ns/op",
            "extra": "2161 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717616,
            "unit": "B/op",
            "extra": "2161 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2161 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 558720,
            "unit": "ns/op\t  717613 B/op\t     133 allocs/op",
            "extra": "2072 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 558720,
            "unit": "ns/op",
            "extra": "2072 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717613,
            "unit": "B/op",
            "extra": "2072 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2072 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4625120,
            "unit": "ns/op\t 5226421 B/op\t     168 allocs/op",
            "extra": "259 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4625120,
            "unit": "ns/op",
            "extra": "259 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226421,
            "unit": "B/op",
            "extra": "259 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "259 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4641550,
            "unit": "ns/op\t 5226423 B/op\t     169 allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4641550,
            "unit": "ns/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226423,
            "unit": "B/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4680351,
            "unit": "ns/op\t 5226415 B/op\t     168 allocs/op",
            "extra": "255 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4680351,
            "unit": "ns/op",
            "extra": "255 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226415,
            "unit": "B/op",
            "extra": "255 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "255 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1366,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "897034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1366,
            "unit": "ns/op",
            "extra": "897034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "897034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "897034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1348,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "884804 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1348,
            "unit": "ns/op",
            "extra": "884804 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "884804 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "884804 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1346,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "768598 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1346,
            "unit": "ns/op",
            "extra": "768598 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "768598 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "768598 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1360,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1360,
            "unit": "ns/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "783528 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1372,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "801703 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1372,
            "unit": "ns/op",
            "extra": "801703 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "801703 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "801703 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1370,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "812580 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1370,
            "unit": "ns/op",
            "extra": "812580 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "812580 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "812580 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1210,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "991425 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1210,
            "unit": "ns/op",
            "extra": "991425 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "991425 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "991425 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1209,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "848878 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1209,
            "unit": "ns/op",
            "extra": "848878 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "848878 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "848878 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1203,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "1001871 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1203,
            "unit": "ns/op",
            "extra": "1001871 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "1001871 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "1001871 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 30268,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "39696 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 30268,
            "unit": "ns/op",
            "extra": "39696 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "39696 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "39696 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 30204,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "39651 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 30204,
            "unit": "ns/op",
            "extra": "39651 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "39651 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "39651 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 30404,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "39970 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 30404,
            "unit": "ns/op",
            "extra": "39970 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "39970 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "39970 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 332109,
            "unit": "ns/op\t  596198 B/op\t      74 allocs/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 332109,
            "unit": "ns/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596198,
            "unit": "B/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3012 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 339441,
            "unit": "ns/op\t  596198 B/op\t      74 allocs/op",
            "extra": "3405 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 339441,
            "unit": "ns/op",
            "extra": "3405 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596198,
            "unit": "B/op",
            "extra": "3405 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3405 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 335748,
            "unit": "ns/op\t  596199 B/op\t      74 allocs/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 335748,
            "unit": "ns/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596199,
            "unit": "B/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3499 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2118163,
            "unit": "ns/op\t 4381152 B/op\t      88 allocs/op",
            "extra": "552 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2118163,
            "unit": "ns/op",
            "extra": "552 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381152,
            "unit": "B/op",
            "extra": "552 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "552 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2106006,
            "unit": "ns/op\t 4381147 B/op\t      88 allocs/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2106006,
            "unit": "ns/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381147,
            "unit": "B/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2050829,
            "unit": "ns/op\t 4381135 B/op\t      88 allocs/op",
            "extra": "547 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2050829,
            "unit": "ns/op",
            "extra": "547 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381135,
            "unit": "B/op",
            "extra": "547 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "547 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 67002,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "17930 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 67002,
            "unit": "ns/op",
            "extra": "17930 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "17930 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "17930 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 66559,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "17954 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 66559,
            "unit": "ns/op",
            "extra": "17954 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "17954 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "17954 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 65781,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "17906 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 65781,
            "unit": "ns/op",
            "extra": "17906 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "17906 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "17906 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 552461,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 552461,
            "unit": "ns/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2172 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 548770,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 548770,
            "unit": "ns/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 548807,
            "unit": "ns/op\t  259463 B/op\t     184 allocs/op",
            "extra": "2103 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 548807,
            "unit": "ns/op",
            "extra": "2103 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259463,
            "unit": "B/op",
            "extra": "2103 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2103 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5432036,
            "unit": "ns/op\t 3583648 B/op\t     255 allocs/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5432036,
            "unit": "ns/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583648,
            "unit": "B/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5459959,
            "unit": "ns/op\t 3583648 B/op\t     255 allocs/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5459959,
            "unit": "ns/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583648,
            "unit": "B/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "214 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5558291,
            "unit": "ns/op\t 3583648 B/op\t     255 allocs/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5558291,
            "unit": "ns/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583648,
            "unit": "B/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "219 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 155992,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8049 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 155992,
            "unit": "ns/op",
            "extra": "8049 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8049 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8049 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 155442,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7438 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 155442,
            "unit": "ns/op",
            "extra": "7438 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7438 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7438 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 156658,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7645 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 156658,
            "unit": "ns/op",
            "extra": "7645 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7645 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7645 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1249219,
            "unit": "ns/op\t  429820 B/op\t   20273 allocs/op",
            "extra": "907 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1249219,
            "unit": "ns/op",
            "extra": "907 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429820,
            "unit": "B/op",
            "extra": "907 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "907 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1245666,
            "unit": "ns/op\t  429821 B/op\t   20273 allocs/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1245666,
            "unit": "ns/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429821,
            "unit": "B/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1264548,
            "unit": "ns/op\t  429820 B/op\t   20273 allocs/op",
            "extra": "1000 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1264548,
            "unit": "ns/op",
            "extra": "1000 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429820,
            "unit": "B/op",
            "extra": "1000 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 106435,
            "unit": "ns/op\t  119184 B/op\t     173 allocs/op",
            "extra": "11077 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 106435,
            "unit": "ns/op",
            "extra": "11077 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119184,
            "unit": "B/op",
            "extra": "11077 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "11077 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 106961,
            "unit": "ns/op\t  119183 B/op\t     173 allocs/op",
            "extra": "9909 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 106961,
            "unit": "ns/op",
            "extra": "9909 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119183,
            "unit": "B/op",
            "extra": "9909 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "9909 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 104790,
            "unit": "ns/op\t  119182 B/op\t     173 allocs/op",
            "extra": "9590 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 104790,
            "unit": "ns/op",
            "extra": "9590 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119182,
            "unit": "B/op",
            "extra": "9590 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "9590 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 898517,
            "unit": "ns/op\t 1033039 B/op\t     223 allocs/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 898517,
            "unit": "ns/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033039,
            "unit": "B/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 888742,
            "unit": "ns/op\t 1033035 B/op\t     223 allocs/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 888742,
            "unit": "ns/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033035,
            "unit": "B/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 896414,
            "unit": "ns/op\t 1033043 B/op\t     223 allocs/op",
            "extra": "1291 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 896414,
            "unit": "ns/op",
            "extra": "1291 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033043,
            "unit": "B/op",
            "extra": "1291 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1291 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 660408,
            "unit": "ns/op\t  317849 B/op\t    3277 allocs/op",
            "extra": "1675 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 660408,
            "unit": "ns/op",
            "extra": "1675 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 317849,
            "unit": "B/op",
            "extra": "1675 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1675 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 654661,
            "unit": "ns/op\t  318306 B/op\t    3277 allocs/op",
            "extra": "1735 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 654661,
            "unit": "ns/op",
            "extra": "1735 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318306,
            "unit": "B/op",
            "extra": "1735 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1735 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 656186,
            "unit": "ns/op\t  318701 B/op\t    3277 allocs/op",
            "extra": "1851 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 656186,
            "unit": "ns/op",
            "extra": "1851 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318701,
            "unit": "B/op",
            "extra": "1851 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1851 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4033409,
            "unit": "ns/op\t 3480687 B/op\t   30226 allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4033409,
            "unit": "ns/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3480687,
            "unit": "B/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30226,
            "unit": "allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4124874,
            "unit": "ns/op\t 3479793 B/op\t   30226 allocs/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4124874,
            "unit": "ns/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3479793,
            "unit": "B/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30226,
            "unit": "allocs/op",
            "extra": "291 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4197109,
            "unit": "ns/op\t 3477579 B/op\t   30226 allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4197109,
            "unit": "ns/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3477579,
            "unit": "B/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30226,
            "unit": "allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 280059,
            "unit": "ns/op\t  285342 B/op\t     576 allocs/op",
            "extra": "4178 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 280059,
            "unit": "ns/op",
            "extra": "4178 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285342,
            "unit": "B/op",
            "extra": "4178 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4178 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 287096,
            "unit": "ns/op\t  285662 B/op\t     576 allocs/op",
            "extra": "3843 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 287096,
            "unit": "ns/op",
            "extra": "3843 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285662,
            "unit": "B/op",
            "extra": "3843 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "3843 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 286399,
            "unit": "ns/op\t  285780 B/op\t     576 allocs/op",
            "extra": "4256 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 286399,
            "unit": "ns/op",
            "extra": "4256 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285780,
            "unit": "B/op",
            "extra": "4256 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4256 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 912226,
            "unit": "ns/op\t 1330134 B/op\t     693 allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 912226,
            "unit": "ns/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330134,
            "unit": "B/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 911407,
            "unit": "ns/op\t 1330169 B/op\t     693 allocs/op",
            "extra": "1264 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 911407,
            "unit": "ns/op",
            "extra": "1264 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330169,
            "unit": "B/op",
            "extra": "1264 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1264 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 911329,
            "unit": "ns/op\t 1329717 B/op\t     693 allocs/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 911329,
            "unit": "ns/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1329717,
            "unit": "B/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1334 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1691232,
            "unit": "ns/op\t 1364908 B/op\t   13581 allocs/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1691232,
            "unit": "ns/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1364908,
            "unit": "B/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13581,
            "unit": "allocs/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1662465,
            "unit": "ns/op\t 1366007 B/op\t   13604 allocs/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1662465,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366007,
            "unit": "B/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13604,
            "unit": "allocs/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1782410,
            "unit": "ns/op\t 1366897 B/op\t   13623 allocs/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1782410,
            "unit": "ns/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366897,
            "unit": "B/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13623,
            "unit": "allocs/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 18037244,
            "unit": "ns/op\t14359986 B/op\t  152231 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 18037244,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14359986,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152231,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17495788,
            "unit": "ns/op\t14335455 B/op\t  151720 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17495788,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14335455,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 151720,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 18094167,
            "unit": "ns/op\t14369765 B/op\t  152435 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 18094167,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14369765,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152435,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 594788,
            "unit": "ns/op\t  717615 B/op\t     133 allocs/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 594788,
            "unit": "ns/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717615,
            "unit": "B/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "1812 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 595044,
            "unit": "ns/op\t  717617 B/op\t     133 allocs/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 595044,
            "unit": "ns/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717617,
            "unit": "B/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "1825 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 600650,
            "unit": "ns/op\t  717617 B/op\t     133 allocs/op",
            "extra": "1934 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 600650,
            "unit": "ns/op",
            "extra": "1934 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717617,
            "unit": "B/op",
            "extra": "1934 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "1934 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 570621,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2090 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 570621,
            "unit": "ns/op",
            "extra": "2090 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2090 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2090 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 575353,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2073 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 575353,
            "unit": "ns/op",
            "extra": "2073 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2073 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2073 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 568535,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2094 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 568535,
            "unit": "ns/op",
            "extra": "2094 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2094 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2094 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "committer": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "id": "627ff698d0f09b1b3e4c018f3564577a8fb4976b",
          "message": "Add comprehensive string manipulation functions to expression system\n\n- Extend Expr interface with Contains, StartsWith, EndsWith methods\n- Implement string operations for all expression types (Column, Literal, Binary)\n- Add evaluation logic for string comparisons using Go strings package\n- Support for proper null value handling in string operations\n- Type safety with clear error messages for non-string operands\n\nFeatures added:\n- df.Col('name').Contains(Lit('substring')) - check if string contains substring\n- df.Col('email').StartsWith(Lit('prefix')) - check if string starts with prefix\n- df.Col('file').EndsWith(Lit('suffix')) - check if string ends with suffix\n- Chainable with other operations for complex filtering\n- Case-sensitive string matching (as expected in production systems)\n\nComprehensive test coverage with 10 test scenarios:\n- Basic functionality for all three operations\n- WithColumn integration for boolean result columns\n- Case sensitivity validation\n- Empty string handling\n- Null value propagation\n- Type error validation\n- Chained expression filtering\n- Performance testing with 10K records\n\nAll tests passing with proper Arrow boolean array results.\nFix join test column indexing for reliable test execution.",
          "timestamp": "2025-07-24T10:50:59Z",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/627ff698d0f09b1b3e4c018f3564577a8fb4976b"
        },
        "date": 1753411592458,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 40167,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "29822 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 40167,
            "unit": "ns/op",
            "extra": "29822 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "29822 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "29822 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39779,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "29354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39779,
            "unit": "ns/op",
            "extra": "29354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "29354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "29354 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 40892,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "31074 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 40892,
            "unit": "ns/op",
            "extra": "31074 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "31074 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "31074 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 552914,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2056 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 552914,
            "unit": "ns/op",
            "extra": "2056 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2056 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2056 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 476272,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2541 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 476272,
            "unit": "ns/op",
            "extra": "2541 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2541 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2541 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 469136,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2397 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 469136,
            "unit": "ns/op",
            "extra": "2397 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2397 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2397 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3755444,
            "unit": "ns/op\t 5801466 B/op\t     111 allocs/op",
            "extra": "325 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3755444,
            "unit": "ns/op",
            "extra": "325 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801466,
            "unit": "B/op",
            "extra": "325 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "325 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 4028083,
            "unit": "ns/op\t 5801464 B/op\t     111 allocs/op",
            "extra": "314 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 4028083,
            "unit": "ns/op",
            "extra": "314 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801464,
            "unit": "B/op",
            "extra": "314 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "314 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3584667,
            "unit": "ns/op\t 5801465 B/op\t     111 allocs/op",
            "extra": "321 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3584667,
            "unit": "ns/op",
            "extra": "321 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801465,
            "unit": "B/op",
            "extra": "321 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "321 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 53578,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "22568 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 53578,
            "unit": "ns/op",
            "extra": "22568 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "22568 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "22568 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 52920,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "22225 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 52920,
            "unit": "ns/op",
            "extra": "22225 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "22225 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "22225 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 54864,
            "unit": "ns/op\t   48617 B/op\t      88 allocs/op",
            "extra": "22701 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 54864,
            "unit": "ns/op",
            "extra": "22701 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48617,
            "unit": "B/op",
            "extra": "22701 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "22701 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 531602,
            "unit": "ns/op\t  717614 B/op\t     133 allocs/op",
            "extra": "2124 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 531602,
            "unit": "ns/op",
            "extra": "2124 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717614,
            "unit": "B/op",
            "extra": "2124 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2124 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 549368,
            "unit": "ns/op\t  717616 B/op\t     133 allocs/op",
            "extra": "2146 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 549368,
            "unit": "ns/op",
            "extra": "2146 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717616,
            "unit": "B/op",
            "extra": "2146 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2146 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 522755,
            "unit": "ns/op\t  717611 B/op\t     133 allocs/op",
            "extra": "2200 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 522755,
            "unit": "ns/op",
            "extra": "2200 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717611,
            "unit": "B/op",
            "extra": "2200 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2200 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4548808,
            "unit": "ns/op\t 5226410 B/op\t     168 allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4548808,
            "unit": "ns/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226410,
            "unit": "B/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 168,
            "unit": "allocs/op",
            "extra": "267 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4523895,
            "unit": "ns/op\t 5226425 B/op\t     169 allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4523895,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226425,
            "unit": "B/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4520861,
            "unit": "ns/op\t 5226436 B/op\t     169 allocs/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4520861,
            "unit": "ns/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226436,
            "unit": "B/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1329,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "885366 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1329,
            "unit": "ns/op",
            "extra": "885366 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "885366 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "885366 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1336,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "786496 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1336,
            "unit": "ns/op",
            "extra": "786496 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "786496 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "786496 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1344,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "751952 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1344,
            "unit": "ns/op",
            "extra": "751952 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "751952 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "751952 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1363,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "811192 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1363,
            "unit": "ns/op",
            "extra": "811192 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "811192 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "811192 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1346,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "850807 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1346,
            "unit": "ns/op",
            "extra": "850807 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "850807 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "850807 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1394,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "835494 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1394,
            "unit": "ns/op",
            "extra": "835494 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "835494 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "835494 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1193,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "890252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1193,
            "unit": "ns/op",
            "extra": "890252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "890252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "890252 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1188,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "844960 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1188,
            "unit": "ns/op",
            "extra": "844960 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "844960 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "844960 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1201,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "843954 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1201,
            "unit": "ns/op",
            "extra": "843954 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "843954 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "843954 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28896,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28896,
            "unit": "ns/op",
            "extra": "40621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40621 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28441,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "41268 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28441,
            "unit": "ns/op",
            "extra": "41268 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "41268 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "41268 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28847,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "42147 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28847,
            "unit": "ns/op",
            "extra": "42147 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "42147 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "42147 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 296249,
            "unit": "ns/op\t  596201 B/op\t      74 allocs/op",
            "extra": "3938 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 296249,
            "unit": "ns/op",
            "extra": "3938 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596201,
            "unit": "B/op",
            "extra": "3938 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3938 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 299215,
            "unit": "ns/op\t  596199 B/op\t      74 allocs/op",
            "extra": "4028 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 299215,
            "unit": "ns/op",
            "extra": "4028 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596199,
            "unit": "B/op",
            "extra": "4028 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "4028 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 294980,
            "unit": "ns/op\t  596201 B/op\t      74 allocs/op",
            "extra": "3908 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 294980,
            "unit": "ns/op",
            "extra": "3908 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596201,
            "unit": "B/op",
            "extra": "3908 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3908 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2090549,
            "unit": "ns/op\t 4381160 B/op\t      88 allocs/op",
            "extra": "588 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2090549,
            "unit": "ns/op",
            "extra": "588 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381160,
            "unit": "B/op",
            "extra": "588 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "588 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2122590,
            "unit": "ns/op\t 4381169 B/op\t      88 allocs/op",
            "extra": "559 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2122590,
            "unit": "ns/op",
            "extra": "559 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381169,
            "unit": "B/op",
            "extra": "559 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "559 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2230073,
            "unit": "ns/op\t 4381180 B/op\t      88 allocs/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2230073,
            "unit": "ns/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381180,
            "unit": "B/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 66522,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18046 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 66522,
            "unit": "ns/op",
            "extra": "18046 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18046 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18046 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 65930,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 65930,
            "unit": "ns/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 65782,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 65782,
            "unit": "ns/op",
            "extra": "18224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 548702,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2125 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 548702,
            "unit": "ns/op",
            "extra": "2125 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2125 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2125 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 546892,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2142 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 546892,
            "unit": "ns/op",
            "extra": "2142 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2142 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2142 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 557606,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2131 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 557606,
            "unit": "ns/op",
            "extra": "2131 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2131 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2131 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5315298,
            "unit": "ns/op\t 3583645 B/op\t     255 allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5315298,
            "unit": "ns/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583645,
            "unit": "B/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5402104,
            "unit": "ns/op\t 3583647 B/op\t     255 allocs/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5402104,
            "unit": "ns/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583647,
            "unit": "B/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5301045,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5301045,
            "unit": "ns/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 148301,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8331 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 148301,
            "unit": "ns/op",
            "extra": "8331 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8331 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8331 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 148117,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8196 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 148117,
            "unit": "ns/op",
            "extra": "8196 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8196 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8196 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 145297,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7622 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 145297,
            "unit": "ns/op",
            "extra": "7622 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7622 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7622 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1196717,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "992 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1196717,
            "unit": "ns/op",
            "extra": "992 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "992 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "992 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1226185,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "987 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1226185,
            "unit": "ns/op",
            "extra": "987 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "987 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "987 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1199768,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "969 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1199768,
            "unit": "ns/op",
            "extra": "969 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "969 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "969 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 100141,
            "unit": "ns/op\t  119182 B/op\t     173 allocs/op",
            "extra": "12088 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 100141,
            "unit": "ns/op",
            "extra": "12088 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119182,
            "unit": "B/op",
            "extra": "12088 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12088 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 100729,
            "unit": "ns/op\t  119183 B/op\t     173 allocs/op",
            "extra": "12037 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 100729,
            "unit": "ns/op",
            "extra": "12037 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119183,
            "unit": "B/op",
            "extra": "12037 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "12037 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 100422,
            "unit": "ns/op\t  119183 B/op\t     173 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 100422,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119183,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 825364,
            "unit": "ns/op\t 1033034 B/op\t     223 allocs/op",
            "extra": "1432 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 825364,
            "unit": "ns/op",
            "extra": "1432 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033034,
            "unit": "B/op",
            "extra": "1432 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1432 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 829488,
            "unit": "ns/op\t 1033031 B/op\t     223 allocs/op",
            "extra": "1435 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 829488,
            "unit": "ns/op",
            "extra": "1435 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033031,
            "unit": "B/op",
            "extra": "1435 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1435 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 831273,
            "unit": "ns/op\t 1033033 B/op\t     223 allocs/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 831273,
            "unit": "ns/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033033,
            "unit": "B/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1467 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 682157,
            "unit": "ns/op\t  318241 B/op\t    3277 allocs/op",
            "extra": "1713 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 682157,
            "unit": "ns/op",
            "extra": "1713 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318241,
            "unit": "B/op",
            "extra": "1713 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1713 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 681036,
            "unit": "ns/op\t  318932 B/op\t    3277 allocs/op",
            "extra": "1856 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 681036,
            "unit": "ns/op",
            "extra": "1856 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318932,
            "unit": "B/op",
            "extra": "1856 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1856 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 676102,
            "unit": "ns/op\t  318927 B/op\t    3277 allocs/op",
            "extra": "1777 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 676102,
            "unit": "ns/op",
            "extra": "1777 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318927,
            "unit": "B/op",
            "extra": "1777 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1777 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3985525,
            "unit": "ns/op\t 3481852 B/op\t   30228 allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3985525,
            "unit": "ns/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3481852,
            "unit": "B/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3991073,
            "unit": "ns/op\t 3482847 B/op\t   30227 allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3991073,
            "unit": "ns/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3482847,
            "unit": "B/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "288 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3980696,
            "unit": "ns/op\t 3485086 B/op\t   30229 allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3980696,
            "unit": "ns/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3485086,
            "unit": "B/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30229,
            "unit": "allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 283477,
            "unit": "ns/op\t  285690 B/op\t     576 allocs/op",
            "extra": "3871 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 283477,
            "unit": "ns/op",
            "extra": "3871 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285690,
            "unit": "B/op",
            "extra": "3871 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "3871 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 285538,
            "unit": "ns/op\t  285691 B/op\t     576 allocs/op",
            "extra": "3939 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 285538,
            "unit": "ns/op",
            "extra": "3939 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285691,
            "unit": "B/op",
            "extra": "3939 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "3939 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 284258,
            "unit": "ns/op\t  285769 B/op\t     576 allocs/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 284258,
            "unit": "ns/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285769,
            "unit": "B/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "3976 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 888156,
            "unit": "ns/op\t 1330802 B/op\t     693 allocs/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 888156,
            "unit": "ns/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330802,
            "unit": "B/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 889313,
            "unit": "ns/op\t 1330797 B/op\t     693 allocs/op",
            "extra": "1316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 889313,
            "unit": "ns/op",
            "extra": "1316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330797,
            "unit": "B/op",
            "extra": "1316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1316 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 886227,
            "unit": "ns/op\t 1330749 B/op\t     693 allocs/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 886227,
            "unit": "ns/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330749,
            "unit": "B/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1330 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1650188,
            "unit": "ns/op\t 1366499 B/op\t   13614 allocs/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1650188,
            "unit": "ns/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366499,
            "unit": "B/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13614,
            "unit": "allocs/op",
            "extra": "664 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1666137,
            "unit": "ns/op\t 1367637 B/op\t   13638 allocs/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1666137,
            "unit": "ns/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1367637,
            "unit": "B/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13638,
            "unit": "allocs/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1660504,
            "unit": "ns/op\t 1366891 B/op\t   13623 allocs/op",
            "extra": "642 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1660504,
            "unit": "ns/op",
            "extra": "642 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1366891,
            "unit": "B/op",
            "extra": "642 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13623,
            "unit": "allocs/op",
            "extra": "642 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17240105,
            "unit": "ns/op\t14395856 B/op\t  152978 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17240105,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14395856,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152978,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17725114,
            "unit": "ns/op\t14396243 B/op\t  152986 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17725114,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14396243,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152986,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17768757,
            "unit": "ns/op\t14367583 B/op\t  152389 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17768757,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14367583,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152389,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 571027,
            "unit": "ns/op\t  717615 B/op\t     133 allocs/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 571027,
            "unit": "ns/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717615,
            "unit": "B/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 573208,
            "unit": "ns/op\t  717613 B/op\t     133 allocs/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 573208,
            "unit": "ns/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717613,
            "unit": "B/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2077 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 571567,
            "unit": "ns/op\t  717614 B/op\t     133 allocs/op",
            "extra": "2092 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 571567,
            "unit": "ns/op",
            "extra": "2092 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717614,
            "unit": "B/op",
            "extra": "2092 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2092 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 557539,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2120 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 557539,
            "unit": "ns/op",
            "extra": "2120 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2120 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2120 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 556008,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2130 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 556008,
            "unit": "ns/op",
            "extra": "2130 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2130 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2130 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 554259,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2119 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 554259,
            "unit": "ns/op",
            "extra": "2119 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2119 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2119 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "committer": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "id": "627ff698d0f09b1b3e4c018f3564577a8fb4976b",
          "message": "Add comprehensive string manipulation functions to expression system\n\n- Extend Expr interface with Contains, StartsWith, EndsWith methods\n- Implement string operations for all expression types (Column, Literal, Binary)\n- Add evaluation logic for string comparisons using Go strings package\n- Support for proper null value handling in string operations\n- Type safety with clear error messages for non-string operands\n\nFeatures added:\n- df.Col('name').Contains(Lit('substring')) - check if string contains substring\n- df.Col('email').StartsWith(Lit('prefix')) - check if string starts with prefix\n- df.Col('file').EndsWith(Lit('suffix')) - check if string ends with suffix\n- Chainable with other operations for complex filtering\n- Case-sensitive string matching (as expected in production systems)\n\nComprehensive test coverage with 10 test scenarios:\n- Basic functionality for all three operations\n- WithColumn integration for boolean result columns\n- Case sensitivity validation\n- Empty string handling\n- Null value propagation\n- Type error validation\n- Chained expression filtering\n- Performance testing with 10K records\n\nAll tests passing with proper Arrow boolean array results.\nFix join test column indexing for reliable test execution.",
          "timestamp": "2025-07-24T10:50:59Z",
          "url": "https://github.com/felixgeelhaar/GopherFrame/commit/627ff698d0f09b1b3e4c018f3564577a8fb4976b"
        },
        "date": 1753497784738,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 38645,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "30733 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 38645,
            "unit": "ns/op",
            "extra": "30733 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "30733 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30733 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 42423,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "29988 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 42423,
            "unit": "ns/op",
            "extra": "29988 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "29988 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "29988 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39370,
            "unit": "ns/op\t   52411 B/op\t      60 allocs/op",
            "extra": "26133 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39370,
            "unit": "ns/op",
            "extra": "26133 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52411,
            "unit": "B/op",
            "extra": "26133 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "26133 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 463223,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2422 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 463223,
            "unit": "ns/op",
            "extra": "2422 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2422 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2422 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 477656,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2520 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 477656,
            "unit": "ns/op",
            "extra": "2520 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2520 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2520 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 469563,
            "unit": "ns/op\t  787911 B/op\t      88 allocs/op",
            "extra": "2400 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 469563,
            "unit": "ns/op",
            "extra": "2400 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787911,
            "unit": "B/op",
            "extra": "2400 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2400 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3799872,
            "unit": "ns/op\t 5801467 B/op\t     111 allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3799872,
            "unit": "ns/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801467,
            "unit": "B/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3771996,
            "unit": "ns/op\t 5801470 B/op\t     111 allocs/op",
            "extra": "328 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3771996,
            "unit": "ns/op",
            "extra": "328 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801470,
            "unit": "B/op",
            "extra": "328 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "328 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3779353,
            "unit": "ns/op\t 5801468 B/op\t     111 allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3779353,
            "unit": "ns/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801468,
            "unit": "B/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 56166,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21274 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 56166,
            "unit": "ns/op",
            "extra": "21274 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21274 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21274 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55707,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21512 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55707,
            "unit": "ns/op",
            "extra": "21512 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21512 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21512 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 55157,
            "unit": "ns/op\t   48618 B/op\t      88 allocs/op",
            "extra": "21604 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 55157,
            "unit": "ns/op",
            "extra": "21604 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 48618,
            "unit": "B/op",
            "extra": "21604 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "21604 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 549576,
            "unit": "ns/op\t  717619 B/op\t     133 allocs/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 549576,
            "unit": "ns/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717619,
            "unit": "B/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 544631,
            "unit": "ns/op\t  717616 B/op\t     133 allocs/op",
            "extra": "1977 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 544631,
            "unit": "ns/op",
            "extra": "1977 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717616,
            "unit": "B/op",
            "extra": "1977 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "1977 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 548070,
            "unit": "ns/op\t  717618 B/op\t     133 allocs/op",
            "extra": "2170 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 548070,
            "unit": "ns/op",
            "extra": "2170 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 717618,
            "unit": "B/op",
            "extra": "2170 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2170 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4565989,
            "unit": "ns/op\t 5226428 B/op\t     169 allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4565989,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226428,
            "unit": "B/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4523201,
            "unit": "ns/op\t 5226446 B/op\t     169 allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4523201,
            "unit": "ns/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226446,
            "unit": "B/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "264 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4516862,
            "unit": "ns/op\t 5226428 B/op\t     169 allocs/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4516862,
            "unit": "ns/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 5226428,
            "unit": "B/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "262 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1350,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "799743 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1350,
            "unit": "ns/op",
            "extra": "799743 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "799743 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "799743 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1344,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "802540 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1344,
            "unit": "ns/op",
            "extra": "802540 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "802540 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "802540 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1337,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "900343 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1337,
            "unit": "ns/op",
            "extra": "900343 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "900343 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "900343 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1374,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "804963 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1374,
            "unit": "ns/op",
            "extra": "804963 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "804963 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "804963 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1379,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "780034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1379,
            "unit": "ns/op",
            "extra": "780034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "780034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "780034 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1373,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "825954 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1373,
            "unit": "ns/op",
            "extra": "825954 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "825954 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "825954 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1209,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "996388 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1209,
            "unit": "ns/op",
            "extra": "996388 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "996388 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "996388 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1195,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "864628 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1195,
            "unit": "ns/op",
            "extra": "864628 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "864628 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "864628 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1205,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "834241 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1205,
            "unit": "ns/op",
            "extra": "834241 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "834241 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "834241 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 29384,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40424 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 29384,
            "unit": "ns/op",
            "extra": "40424 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40424 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40424 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28905,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40975 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28905,
            "unit": "ns/op",
            "extra": "40975 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40975 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40975 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 28885,
            "unit": "ns/op\t   42623 B/op\t      58 allocs/op",
            "extra": "40900 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 28885,
            "unit": "ns/op",
            "extra": "40900 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 42623,
            "unit": "B/op",
            "extra": "40900 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "40900 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 311101,
            "unit": "ns/op\t  596198 B/op\t      74 allocs/op",
            "extra": "3877 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 311101,
            "unit": "ns/op",
            "extra": "3877 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596198,
            "unit": "B/op",
            "extra": "3877 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3877 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 326091,
            "unit": "ns/op\t  596197 B/op\t      74 allocs/op",
            "extra": "3664 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 326091,
            "unit": "ns/op",
            "extra": "3664 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596197,
            "unit": "B/op",
            "extra": "3664 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3664 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 325458,
            "unit": "ns/op\t  596204 B/op\t      74 allocs/op",
            "extra": "3654 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 325458,
            "unit": "ns/op",
            "extra": "3654 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 596204,
            "unit": "B/op",
            "extra": "3654 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 74,
            "unit": "allocs/op",
            "extra": "3654 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2063635,
            "unit": "ns/op\t 4381149 B/op\t      88 allocs/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2063635,
            "unit": "ns/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381149,
            "unit": "B/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "558 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2101797,
            "unit": "ns/op\t 4381153 B/op\t      88 allocs/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2101797,
            "unit": "ns/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381153,
            "unit": "B/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 2108891,
            "unit": "ns/op\t 4381161 B/op\t      88 allocs/op",
            "extra": "566 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 2108891,
            "unit": "ns/op",
            "extra": "566 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 4381161,
            "unit": "B/op",
            "extra": "566 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "566 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 66193,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18286 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 66193,
            "unit": "ns/op",
            "extra": "18286 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18286 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18286 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 65754,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18168 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 65754,
            "unit": "ns/op",
            "extra": "18168 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18168 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18168 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 64769,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "18321 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 64769,
            "unit": "ns/op",
            "extra": "18321 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "18321 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "18321 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 553544,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2151 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 553544,
            "unit": "ns/op",
            "extra": "2151 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2151 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2151 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 548947,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2116 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 548947,
            "unit": "ns/op",
            "extra": "2116 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2116 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2116 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 556249,
            "unit": "ns/op\t  259461 B/op\t     184 allocs/op",
            "extra": "2182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 556249,
            "unit": "ns/op",
            "extra": "2182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259461,
            "unit": "B/op",
            "extra": "2182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2182 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5304124,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5304124,
            "unit": "ns/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5352250,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5352250,
            "unit": "ns/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5284560,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5284560,
            "unit": "ns/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "222 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 143008,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7815 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 143008,
            "unit": "ns/op",
            "extra": "7815 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7815 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7815 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 142575,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "8395 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 142575,
            "unit": "ns/op",
            "extra": "8395 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "8395 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "8395 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 142991,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7768 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 142991,
            "unit": "ns/op",
            "extra": "7768 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7768 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7768 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1207577,
            "unit": "ns/op\t  429818 B/op\t   20273 allocs/op",
            "extra": "1010 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1207577,
            "unit": "ns/op",
            "extra": "1010 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429818,
            "unit": "B/op",
            "extra": "1010 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1010 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1183852,
            "unit": "ns/op\t  429818 B/op\t   20273 allocs/op",
            "extra": "999 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1183852,
            "unit": "ns/op",
            "extra": "999 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429818,
            "unit": "B/op",
            "extra": "999 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "999 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1195377,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1195377,
            "unit": "ns/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1012 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 103143,
            "unit": "ns/op\t  119185 B/op\t     173 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 103143,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119185,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 101489,
            "unit": "ns/op\t  119184 B/op\t     173 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 101489,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119184,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 100426,
            "unit": "ns/op\t  119183 B/op\t     173 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 100426,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 119183,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 173,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 794131,
            "unit": "ns/op\t 1033033 B/op\t     223 allocs/op",
            "extra": "1513 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 794131,
            "unit": "ns/op",
            "extra": "1513 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033033,
            "unit": "B/op",
            "extra": "1513 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1513 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 801179,
            "unit": "ns/op\t 1033037 B/op\t     223 allocs/op",
            "extra": "1495 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 801179,
            "unit": "ns/op",
            "extra": "1495 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033037,
            "unit": "B/op",
            "extra": "1495 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1495 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 818271,
            "unit": "ns/op\t 1033037 B/op\t     223 allocs/op",
            "extra": "1471 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 818271,
            "unit": "ns/op",
            "extra": "1471 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 1033037,
            "unit": "B/op",
            "extra": "1471 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "1471 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 620243,
            "unit": "ns/op\t  318779 B/op\t    3277 allocs/op",
            "extra": "1810 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 620243,
            "unit": "ns/op",
            "extra": "1810 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318779,
            "unit": "B/op",
            "extra": "1810 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1810 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 633140,
            "unit": "ns/op\t  318836 B/op\t    3277 allocs/op",
            "extra": "1923 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 633140,
            "unit": "ns/op",
            "extra": "1923 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318836,
            "unit": "B/op",
            "extra": "1923 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1923 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 632147,
            "unit": "ns/op\t  319132 B/op\t    3277 allocs/op",
            "extra": "1927 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 632147,
            "unit": "ns/op",
            "extra": "1927 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 319132,
            "unit": "B/op",
            "extra": "1927 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1927 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3914424,
            "unit": "ns/op\t 3482572 B/op\t   30228 allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3914424,
            "unit": "ns/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3482572,
            "unit": "B/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "300 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3965318,
            "unit": "ns/op\t 3483896 B/op\t   30228 allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3965318,
            "unit": "ns/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3483896,
            "unit": "B/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30228,
            "unit": "allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 3959687,
            "unit": "ns/op\t 3481282 B/op\t   30227 allocs/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 3959687,
            "unit": "ns/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3481282,
            "unit": "B/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "306 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 275580,
            "unit": "ns/op\t  285521 B/op\t     576 allocs/op",
            "extra": "3948 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 275580,
            "unit": "ns/op",
            "extra": "3948 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285521,
            "unit": "B/op",
            "extra": "3948 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "3948 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 281934,
            "unit": "ns/op\t  285802 B/op\t     576 allocs/op",
            "extra": "3916 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 281934,
            "unit": "ns/op",
            "extra": "3916 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285802,
            "unit": "B/op",
            "extra": "3916 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "3916 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 277369,
            "unit": "ns/op\t  285396 B/op\t     576 allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 277369,
            "unit": "ns/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285396,
            "unit": "B/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4035 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 861077,
            "unit": "ns/op\t 1329936 B/op\t     693 allocs/op",
            "extra": "1314 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 861077,
            "unit": "ns/op",
            "extra": "1314 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1329936,
            "unit": "B/op",
            "extra": "1314 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1314 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 878225,
            "unit": "ns/op\t 1330419 B/op\t     693 allocs/op",
            "extra": "1346 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 878225,
            "unit": "ns/op",
            "extra": "1346 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330419,
            "unit": "B/op",
            "extra": "1346 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1346 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 862642,
            "unit": "ns/op\t 1330140 B/op\t     693 allocs/op",
            "extra": "1335 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 862642,
            "unit": "ns/op",
            "extra": "1335 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1330140,
            "unit": "B/op",
            "extra": "1335 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1335 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1656505,
            "unit": "ns/op\t 1364499 B/op\t   13573 allocs/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1656505,
            "unit": "ns/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1364499,
            "unit": "B/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13573,
            "unit": "allocs/op",
            "extra": "746 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1658173,
            "unit": "ns/op\t 1365871 B/op\t   13601 allocs/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1658173,
            "unit": "ns/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1365871,
            "unit": "B/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13601,
            "unit": "allocs/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1672925,
            "unit": "ns/op\t 1365096 B/op\t   13585 allocs/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1672925,
            "unit": "ns/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1365096,
            "unit": "B/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13585,
            "unit": "allocs/op",
            "extra": "720 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17416725,
            "unit": "ns/op\t14344900 B/op\t  151917 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17416725,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14344900,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 151917,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17561866,
            "unit": "ns/op\t14348998 B/op\t  152002 allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17561866,
            "unit": "ns/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14348998,
            "unit": "B/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152002,
            "unit": "allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17541425,
            "unit": "ns/op\t14368310 B/op\t  152405 allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17541425,
            "unit": "ns/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14368310,
            "unit": "B/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152405,
            "unit": "allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 567202,
            "unit": "ns/op\t  717612 B/op\t     133 allocs/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 567202,
            "unit": "ns/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717612,
            "unit": "B/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2202 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 577394,
            "unit": "ns/op\t  717613 B/op\t     133 allocs/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 577394,
            "unit": "ns/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717613,
            "unit": "B/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2062 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 580856,
            "unit": "ns/op\t  717613 B/op\t     133 allocs/op",
            "extra": "2047 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 580856,
            "unit": "ns/op",
            "extra": "2047 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 717613,
            "unit": "B/op",
            "extra": "2047 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 133,
            "unit": "allocs/op",
            "extra": "2047 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 559669,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 559669,
            "unit": "ns/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2100 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 564677,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2085 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 564677,
            "unit": "ns/op",
            "extra": "2085 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2085 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2085 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 562805,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2179 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 562805,
            "unit": "ns/op",
            "extra": "2179 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2179 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2179 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "felixgeelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "name": "felixgeelhaar",
            "username": "felixgeelhaar"
          },
          "id": "6adf6db38c2858d2fd9489e24508bfb9d97bdbed",
          "message": "feat: Complete v0.1 MVP with production-ready features",
          "timestamp": "2025-07-24T10:51:35Z",
          "url": "https://github.com/felixgeelhaar/GopherFrame/pull/2/commits/6adf6db38c2858d2fd9489e24508bfb9d97bdbed"
        },
        "date": 1753521252200,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 41323,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "28501 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 41323,
            "unit": "ns/op",
            "extra": "28501 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "28501 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "28501 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 474842,
            "unit": "ns/op\t  787910 B/op\t      88 allocs/op",
            "extra": "2347 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 474842,
            "unit": "ns/op",
            "extra": "2347 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787910,
            "unit": "B/op",
            "extra": "2347 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2347 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3824864,
            "unit": "ns/op\t 5801469 B/op\t     111 allocs/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3824864,
            "unit": "ns/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801469,
            "unit": "B/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 53816,
            "unit": "ns/op\t   39272 B/op\t      82 allocs/op",
            "extra": "22255 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 53816,
            "unit": "ns/op",
            "extra": "22255 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 39272,
            "unit": "B/op",
            "extra": "22255 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 82,
            "unit": "allocs/op",
            "extra": "22255 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 517402,
            "unit": "ns/op\t  563411 B/op\t     118 allocs/op",
            "extra": "2216 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 517402,
            "unit": "ns/op",
            "extra": "2216 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 563411,
            "unit": "B/op",
            "extra": "2216 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "2216 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 5031950,
            "unit": "ns/op\t 4113388 B/op\t     147 allocs/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 5031950,
            "unit": "ns/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 4113388,
            "unit": "B/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 147,
            "unit": "allocs/op",
            "extra": "265 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1361,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "857406 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1361,
            "unit": "ns/op",
            "extra": "857406 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "857406 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "857406 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1369,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "740977 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1369,
            "unit": "ns/op",
            "extra": "740977 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "740977 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "740977 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1196,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "999664 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1196,
            "unit": "ns/op",
            "extra": "999664 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "999664 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "999664 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 25681,
            "unit": "ns/op\t   33278 B/op\t      52 allocs/op",
            "extra": "46372 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 25681,
            "unit": "ns/op",
            "extra": "46372 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 33278,
            "unit": "B/op",
            "extra": "46372 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 52,
            "unit": "allocs/op",
            "extra": "46372 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 247152,
            "unit": "ns/op\t  441989 B/op\t      60 allocs/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 247152,
            "unit": "ns/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 441989,
            "unit": "B/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "4141 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1739663,
            "unit": "ns/op\t 3268179 B/op\t      68 allocs/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1739663,
            "unit": "ns/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 3268179,
            "unit": "B/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 68,
            "unit": "allocs/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 67210,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "17811 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 67210,
            "unit": "ns/op",
            "extra": "17811 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "17811 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "17811 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 559810,
            "unit": "ns/op\t  259463 B/op\t     184 allocs/op",
            "extra": "2102 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 559810,
            "unit": "ns/op",
            "extra": "2102 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259463,
            "unit": "B/op",
            "extra": "2102 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2102 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5543790,
            "unit": "ns/op\t 3583671 B/op\t     255 allocs/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5543790,
            "unit": "ns/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583671,
            "unit": "B/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 155873,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7485 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 155873,
            "unit": "ns/op",
            "extra": "7485 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7485 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7485 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1228354,
            "unit": "ns/op\t  429820 B/op\t   20273 allocs/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1228354,
            "unit": "ns/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429820,
            "unit": "B/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "966 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 98600,
            "unit": "ns/op\t  100493 B/op\t     161 allocs/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 98600,
            "unit": "ns/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 100493,
            "unit": "B/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 161,
            "unit": "allocs/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 825813,
            "unit": "ns/op\t  799528 B/op\t     197 allocs/op",
            "extra": "1399 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 825813,
            "unit": "ns/op",
            "extra": "1399 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 799528,
            "unit": "B/op",
            "extra": "1399 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 197,
            "unit": "allocs/op",
            "extra": "1399 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 763488,
            "unit": "ns/op\t  318818 B/op\t    3277 allocs/op",
            "extra": "1503 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 763488,
            "unit": "ns/op",
            "extra": "1503 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318818,
            "unit": "B/op",
            "extra": "1503 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1503 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4204451,
            "unit": "ns/op\t 3476890 B/op\t   30226 allocs/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4204451,
            "unit": "ns/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3476890,
            "unit": "B/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30226,
            "unit": "allocs/op",
            "extra": "283 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 282854,
            "unit": "ns/op\t  285343 B/op\t     576 allocs/op",
            "extra": "4303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 282854,
            "unit": "ns/op",
            "extra": "4303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285343,
            "unit": "B/op",
            "extra": "4303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4303 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 935637,
            "unit": "ns/op\t 1327943 B/op\t     693 allocs/op",
            "extra": "1281 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 935637,
            "unit": "ns/op",
            "extra": "1281 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1327943,
            "unit": "B/op",
            "extra": "1281 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1281 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1714261,
            "unit": "ns/op\t 1365985 B/op\t   13604 allocs/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1714261,
            "unit": "ns/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1365985,
            "unit": "B/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13604,
            "unit": "allocs/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17811552,
            "unit": "ns/op\t14365238 B/op\t  152340 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17811552,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14365238,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152340,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 575264,
            "unit": "ns/op\t  563421 B/op\t     118 allocs/op",
            "extra": "2108 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 575264,
            "unit": "ns/op",
            "extra": "2108 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 563421,
            "unit": "B/op",
            "extra": "2108 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "2108 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 564273,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2041 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 564273,
            "unit": "ns/op",
            "extra": "2041 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2041 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2041 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000",
            "value": 567586,
            "unit": "ns/op\t  372175 B/op\t    2500 allocs/op",
            "extra": "1915 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000 - ns/op",
            "value": 567586,
            "unit": "ns/op",
            "extra": "1915 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000 - B/op",
            "value": 372175,
            "unit": "B/op",
            "extra": "1915 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000 - allocs/op",
            "value": 2500,
            "unit": "allocs/op",
            "extra": "1915 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000",
            "value": 5156306,
            "unit": "ns/op\t 4525702 B/op\t   20568 allocs/op",
            "extra": "235 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000 - ns/op",
            "value": 5156306,
            "unit": "ns/op",
            "extra": "235 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000 - B/op",
            "value": 4525702,
            "unit": "B/op",
            "extra": "235 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000 - allocs/op",
            "value": 20568,
            "unit": "allocs/op",
            "extra": "235 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000",
            "value": 54272270,
            "unit": "ns/op\t44847092 B/op\t  200627 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000 - ns/op",
            "value": 54272270,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000 - B/op",
            "value": 44847092,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000 - allocs/op",
            "value": 200627,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000",
            "value": 162208,
            "unit": "ns/op\t  170987 B/op\t     169 allocs/op",
            "extra": "7327 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000 - ns/op",
            "value": 162208,
            "unit": "ns/op",
            "extra": "7327 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000 - B/op",
            "value": 170987,
            "unit": "B/op",
            "extra": "7327 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "7327 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000",
            "value": 1542768,
            "unit": "ns/op\t 1385977 B/op\t     223 allocs/op",
            "extra": "769 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000 - ns/op",
            "value": 1542768,
            "unit": "ns/op",
            "extra": "769 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000 - B/op",
            "value": 1385977,
            "unit": "B/op",
            "extra": "769 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "769 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000",
            "value": 14029705,
            "unit": "ns/op\t18619537 B/op\t     293 allocs/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000 - ns/op",
            "value": 14029705,
            "unit": "ns/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000 - B/op",
            "value": 18619537,
            "unit": "B/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000 - allocs/op",
            "value": 293,
            "unit": "allocs/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000",
            "value": 66120,
            "unit": "ns/op\t   27616 B/op\t     112 allocs/op",
            "extra": "17949 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000 - ns/op",
            "value": 66120,
            "unit": "ns/op",
            "extra": "17949 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000 - B/op",
            "value": 27616,
            "unit": "B/op",
            "extra": "17949 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000 - allocs/op",
            "value": 112,
            "unit": "allocs/op",
            "extra": "17949 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000",
            "value": 582079,
            "unit": "ns/op\t  307300 B/op\t     137 allocs/op",
            "extra": "2089 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000 - ns/op",
            "value": 582079,
            "unit": "ns/op",
            "extra": "2089 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000 - B/op",
            "value": 307300,
            "unit": "B/op",
            "extra": "2089 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000 - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "2089 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000",
            "value": 5172385,
            "unit": "ns/op\t 3310829 B/op\t     172 allocs/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000 - ns/op",
            "value": 5172385,
            "unit": "ns/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000 - B/op",
            "value": 3310829,
            "unit": "B/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000 - allocs/op",
            "value": 172,
            "unit": "allocs/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000",
            "value": 4982699,
            "unit": "ns/op\t  168135 B/op\t     157 allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000 - ns/op",
            "value": 4982699,
            "unit": "ns/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000 - B/op",
            "value": 168135,
            "unit": "B/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000 - allocs/op",
            "value": 157,
            "unit": "allocs/op",
            "extra": "236 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000",
            "value": 466243386,
            "unit": "ns/op\t 2181845 B/op\t     216 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000 - ns/op",
            "value": 466243386,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000 - B/op",
            "value": 2181845,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000 - allocs/op",
            "value": 216,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000",
            "value": 51920512539,
            "unit": "ns/op\t18270976 B/op\t     264 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000 - ns/op",
            "value": 51920512539,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000 - B/op",
            "value": 18270976,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000",
            "value": 1359052,
            "unit": "ns/op\t  909351 B/op\t   22068 allocs/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000 - ns/op",
            "value": 1359052,
            "unit": "ns/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000 - B/op",
            "value": 909351,
            "unit": "B/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000 - allocs/op",
            "value": 22068,
            "unit": "allocs/op",
            "extra": "910 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000",
            "value": 13201941,
            "unit": "ns/op\t 9454965 B/op\t  220079 allocs/op",
            "extra": "91 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000 - ns/op",
            "value": 13201941,
            "unit": "ns/op",
            "extra": "91 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000 - B/op",
            "value": 9454965,
            "unit": "B/op",
            "extra": "91 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000 - allocs/op",
            "value": 220079,
            "unit": "allocs/op",
            "extra": "91 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000",
            "value": 135249898,
            "unit": "ns/op\t97308056 B/op\t 2200095 allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000 - ns/op",
            "value": 135249898,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000 - B/op",
            "value": 97308056,
            "unit": "B/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000 - allocs/op",
            "value": 2200095,
            "unit": "allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000",
            "value": 106769,
            "unit": "ns/op\t  169946 B/op\t      88 allocs/op",
            "extra": "9748 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000 - ns/op",
            "value": 106769,
            "unit": "ns/op",
            "extra": "9748 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000 - B/op",
            "value": 169946,
            "unit": "B/op",
            "extra": "9748 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "9748 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000",
            "value": 1155308,
            "unit": "ns/op\t 2359532 B/op\t     130 allocs/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000 - ns/op",
            "value": 1155308,
            "unit": "ns/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000 - B/op",
            "value": 2359532,
            "unit": "B/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000 - allocs/op",
            "value": 130,
            "unit": "allocs/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000",
            "value": 11066352,
            "unit": "ns/op\t27148559 B/op\t     186 allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000 - ns/op",
            "value": 11066352,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000 - B/op",
            "value": 27148559,
            "unit": "B/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000 - allocs/op",
            "value": 186,
            "unit": "allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000",
            "value": 318512,
            "unit": "ns/op\t  268108 B/op\t     400 allocs/op",
            "extra": "3435 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000 - ns/op",
            "value": 318512,
            "unit": "ns/op",
            "extra": "3435 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000 - B/op",
            "value": 268108,
            "unit": "B/op",
            "extra": "3435 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000 - allocs/op",
            "value": 400,
            "unit": "allocs/op",
            "extra": "3435 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000",
            "value": 3348163,
            "unit": "ns/op\t 3177517 B/op\t     551 allocs/op",
            "extra": "363 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000 - ns/op",
            "value": 3348163,
            "unit": "ns/op",
            "extra": "363 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000 - B/op",
            "value": 3177517,
            "unit": "B/op",
            "extra": "363 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000 - allocs/op",
            "value": 551,
            "unit": "allocs/op",
            "extra": "363 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000",
            "value": 34564589,
            "unit": "ns/op\t33425253 B/op\t     763 allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000 - ns/op",
            "value": 34564589,
            "unit": "ns/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000 - B/op",
            "value": 33425253,
            "unit": "B/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000 - allocs/op",
            "value": 763,
            "unit": "allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000",
            "value": 459825,
            "unit": "ns/op\t  218484 B/op\t      44 allocs/op",
            "extra": "2606 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000 - ns/op",
            "value": 459825,
            "unit": "ns/op",
            "extra": "2606 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000 - B/op",
            "value": 218484,
            "unit": "B/op",
            "extra": "2606 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000 - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "2606 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000",
            "value": 6506279,
            "unit": "ns/op\t 2666504 B/op\t      59 allocs/op",
            "extra": "183 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000 - ns/op",
            "value": 6506279,
            "unit": "ns/op",
            "extra": "183 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000 - B/op",
            "value": 2666504,
            "unit": "B/op",
            "extra": "183 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000 - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "183 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000",
            "value": 88298818,
            "unit": "ns/op\t30105658 B/op\t      79 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000 - ns/op",
            "value": 88298818,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000 - B/op",
            "value": 30105658,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000 - allocs/op",
            "value": 79,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage",
            "value": 7500093,
            "unit": "ns/op\t 6117450 B/op\t   20932 allocs/op",
            "extra": "158 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage - ns/op",
            "value": 7500093,
            "unit": "ns/op",
            "extra": "158 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage - B/op",
            "value": 6117450,
            "unit": "B/op",
            "extra": "158 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage - allocs/op",
            "value": 20932,
            "unit": "allocs/op",
            "extra": "158 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage",
            "value": 20170272,
            "unit": "ns/op\t13805171 B/op\t  220266 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage - ns/op",
            "value": 20170272,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage - B/op",
            "value": 13805171,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage - allocs/op",
            "value": 220266,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "felixgeelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "name": "felixgeelhaar",
            "username": "felixgeelhaar"
          },
          "id": "d16d235c46037d7352d64aeb63a8ad6974d4dccd",
          "message": "feat: Complete v0.1 MVP with production-ready features",
          "timestamp": "2025-07-24T10:51:35Z",
          "url": "https://github.com/felixgeelhaar/GopherFrame/pull/2/commits/d16d235c46037d7352d64aeb63a8ad6974d4dccd"
        },
        "date": 1753525502885,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDataFrameCreation/Size_1000",
            "value": 39629,
            "unit": "ns/op\t   52412 B/op\t      60 allocs/op",
            "extra": "30544 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - ns/op",
            "value": 39629,
            "unit": "ns/op",
            "extra": "30544 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - B/op",
            "value": 52412,
            "unit": "B/op",
            "extra": "30544 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_1000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "30544 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000",
            "value": 434033,
            "unit": "ns/op\t  787910 B/op\t      88 allocs/op",
            "extra": "2628 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - ns/op",
            "value": 434033,
            "unit": "ns/op",
            "extra": "2628 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - B/op",
            "value": 787910,
            "unit": "B/op",
            "extra": "2628 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_10000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "2628 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000",
            "value": 3788060,
            "unit": "ns/op\t 5801477 B/op\t     111 allocs/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - ns/op",
            "value": 3788060,
            "unit": "ns/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - B/op",
            "value": 5801477,
            "unit": "B/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkDataFrameCreation/Size_100000 - allocs/op",
            "value": 111,
            "unit": "allocs/op",
            "extra": "313 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000",
            "value": 50736,
            "unit": "ns/op\t   39273 B/op\t      82 allocs/op",
            "extra": "23713 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - ns/op",
            "value": 50736,
            "unit": "ns/op",
            "extra": "23713 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - B/op",
            "value": 39273,
            "unit": "B/op",
            "extra": "23713 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_1000 - allocs/op",
            "value": 82,
            "unit": "allocs/op",
            "extra": "23713 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000",
            "value": 489983,
            "unit": "ns/op\t  563412 B/op\t     118 allocs/op",
            "extra": "2346 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - ns/op",
            "value": 489983,
            "unit": "ns/op",
            "extra": "2346 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - B/op",
            "value": 563412,
            "unit": "B/op",
            "extra": "2346 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_10000 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "2346 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000",
            "value": 4421449,
            "unit": "ns/op\t 4113412 B/op\t     148 allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - ns/op",
            "value": 4421449,
            "unit": "ns/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - B/op",
            "value": 4113412,
            "unit": "B/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkFilter/Size_100000 - allocs/op",
            "value": 148,
            "unit": "allocs/op",
            "extra": "273 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000",
            "value": 1592,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "768993 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - ns/op",
            "value": 1592,
            "unit": "ns/op",
            "extra": "768993 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "768993 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_1000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "768993 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000",
            "value": 1324,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "887374 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - ns/op",
            "value": 1324,
            "unit": "ns/op",
            "extra": "887374 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "887374 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_10000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "887374 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000",
            "value": 1145,
            "unit": "ns/op\t    1648 B/op\t      17 allocs/op",
            "extra": "893070 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - ns/op",
            "value": 1145,
            "unit": "ns/op",
            "extra": "893070 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - B/op",
            "value": 1648,
            "unit": "B/op",
            "extra": "893070 times\n4 procs"
          },
          {
            "name": "BenchmarkSelect/Size_100000 - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "893070 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000",
            "value": 24146,
            "unit": "ns/op\t   33278 B/op\t      52 allocs/op",
            "extra": "49362 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - ns/op",
            "value": 24146,
            "unit": "ns/op",
            "extra": "49362 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - B/op",
            "value": 33278,
            "unit": "B/op",
            "extra": "49362 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_1000 - allocs/op",
            "value": 52,
            "unit": "allocs/op",
            "extra": "49362 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000",
            "value": 241950,
            "unit": "ns/op\t  441991 B/op\t      60 allocs/op",
            "extra": "5106 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - ns/op",
            "value": 241950,
            "unit": "ns/op",
            "extra": "5106 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - B/op",
            "value": 441991,
            "unit": "B/op",
            "extra": "5106 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_10000 - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "5106 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000",
            "value": 1701950,
            "unit": "ns/op\t 3268175 B/op\t      68 allocs/op",
            "extra": "716 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - ns/op",
            "value": 1701950,
            "unit": "ns/op",
            "extra": "716 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - B/op",
            "value": 3268175,
            "unit": "B/op",
            "extra": "716 times\n4 procs"
          },
          {
            "name": "BenchmarkWithColumn/Size_100000 - allocs/op",
            "value": 68,
            "unit": "allocs/op",
            "extra": "716 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000",
            "value": 60263,
            "unit": "ns/op\t   27776 B/op\t     144 allocs/op",
            "extra": "20014 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - ns/op",
            "value": 60263,
            "unit": "ns/op",
            "extra": "20014 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - B/op",
            "value": 27776,
            "unit": "B/op",
            "extra": "20014 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_1000 - allocs/op",
            "value": 144,
            "unit": "allocs/op",
            "extra": "20014 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000",
            "value": 520748,
            "unit": "ns/op\t  259462 B/op\t     184 allocs/op",
            "extra": "2294 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - ns/op",
            "value": 520748,
            "unit": "ns/op",
            "extra": "2294 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - B/op",
            "value": 259462,
            "unit": "B/op",
            "extra": "2294 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_10000 - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2294 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000",
            "value": 5198970,
            "unit": "ns/op\t 3583644 B/op\t     255 allocs/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - ns/op",
            "value": 5198970,
            "unit": "ns/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - B/op",
            "value": 3583644,
            "unit": "B/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupBySum/Size_100000 - allocs/op",
            "value": 255,
            "unit": "allocs/op",
            "extra": "230 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000",
            "value": 136668,
            "unit": "ns/op\t   54129 B/op\t    2233 allocs/op",
            "extra": "7558 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - ns/op",
            "value": 136668,
            "unit": "ns/op",
            "extra": "7558 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - B/op",
            "value": 54129,
            "unit": "B/op",
            "extra": "7558 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_1000 - allocs/op",
            "value": 2233,
            "unit": "allocs/op",
            "extra": "7558 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000",
            "value": 1149366,
            "unit": "ns/op\t  429819 B/op\t   20273 allocs/op",
            "extra": "1039 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - ns/op",
            "value": 1149366,
            "unit": "ns/op",
            "extra": "1039 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - B/op",
            "value": 429819,
            "unit": "B/op",
            "extra": "1039 times\n4 procs"
          },
          {
            "name": "BenchmarkGroupByMultipleAgg/Size_10000 - allocs/op",
            "value": 20273,
            "unit": "allocs/op",
            "extra": "1039 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000",
            "value": 88611,
            "unit": "ns/op\t  100493 B/op\t     161 allocs/op",
            "extra": "13324 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - ns/op",
            "value": 88611,
            "unit": "ns/op",
            "extra": "13324 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - B/op",
            "value": 100493,
            "unit": "B/op",
            "extra": "13324 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_1000 - allocs/op",
            "value": 161,
            "unit": "allocs/op",
            "extra": "13324 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000",
            "value": 736464,
            "unit": "ns/op\t  799527 B/op\t     197 allocs/op",
            "extra": "1638 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - ns/op",
            "value": 736464,
            "unit": "ns/op",
            "extra": "1638 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - B/op",
            "value": 799527,
            "unit": "B/op",
            "extra": "1638 times\n4 procs"
          },
          {
            "name": "BenchmarkChainedOperations/Size_10000 - allocs/op",
            "value": 197,
            "unit": "allocs/op",
            "extra": "1638 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000",
            "value": 638103,
            "unit": "ns/op\t  318446 B/op\t    3277 allocs/op",
            "extra": "1741 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - ns/op",
            "value": 638103,
            "unit": "ns/op",
            "extra": "1741 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - B/op",
            "value": 318446,
            "unit": "B/op",
            "extra": "1741 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_1000 - allocs/op",
            "value": 3277,
            "unit": "allocs/op",
            "extra": "1741 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000",
            "value": 4875780,
            "unit": "ns/op\t 3480209 B/op\t   30227 allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - ns/op",
            "value": 4875780,
            "unit": "ns/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - B/op",
            "value": 3480209,
            "unit": "B/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetWrite/Size_10000 - allocs/op",
            "value": 30227,
            "unit": "allocs/op",
            "extra": "301 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000",
            "value": 276573,
            "unit": "ns/op\t  285655 B/op\t     576 allocs/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - ns/op",
            "value": 276573,
            "unit": "ns/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - B/op",
            "value": 285655,
            "unit": "B/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_1000 - allocs/op",
            "value": 576,
            "unit": "allocs/op",
            "extra": "4185 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000",
            "value": 872293,
            "unit": "ns/op\t 1328852 B/op\t     693 allocs/op",
            "extra": "1351 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - ns/op",
            "value": 872293,
            "unit": "ns/op",
            "extra": "1351 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - B/op",
            "value": 1328852,
            "unit": "B/op",
            "extra": "1351 times\n4 procs"
          },
          {
            "name": "BenchmarkParquetRead/Size_10000 - allocs/op",
            "value": 693,
            "unit": "allocs/op",
            "extra": "1351 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000",
            "value": 1658271,
            "unit": "ns/op\t 1368351 B/op\t   13653 allocs/op",
            "extra": "729 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - ns/op",
            "value": 1658271,
            "unit": "ns/op",
            "extra": "729 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - B/op",
            "value": 1368351,
            "unit": "B/op",
            "extra": "729 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_1000 - allocs/op",
            "value": 13653,
            "unit": "allocs/op",
            "extra": "729 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000",
            "value": 17145816,
            "unit": "ns/op\t14361288 B/op\t  152258 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - ns/op",
            "value": 17145816,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - B/op",
            "value": 14361288,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkCSVWrite/Size_10000 - allocs/op",
            "value": 152258,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory",
            "value": 496419,
            "unit": "ns/op\t  563417 B/op\t     118 allocs/op",
            "extra": "2438 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - ns/op",
            "value": 496419,
            "unit": "ns/op",
            "extra": "2438 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - B/op",
            "value": 563417,
            "unit": "B/op",
            "extra": "2438 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/Filter_Memory - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "2438 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory",
            "value": 535597,
            "unit": "ns/op\t  259460 B/op\t     184 allocs/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - ns/op",
            "value": 535597,
            "unit": "ns/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - B/op",
            "value": 259460,
            "unit": "B/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryUsage/GroupBy_Memory - allocs/op",
            "value": 184,
            "unit": "allocs/op",
            "extra": "2198 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000",
            "value": 520545,
            "unit": "ns/op\t  372175 B/op\t    2500 allocs/op",
            "extra": "2259 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000 - ns/op",
            "value": 520545,
            "unit": "ns/op",
            "extra": "2259 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000 - B/op",
            "value": 372175,
            "unit": "B/op",
            "extra": "2259 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_1000 - allocs/op",
            "value": 2500,
            "unit": "allocs/op",
            "extra": "2259 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000",
            "value": 4987662,
            "unit": "ns/op\t 4525702 B/op\t   20568 allocs/op",
            "extra": "240 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000 - ns/op",
            "value": 4987662,
            "unit": "ns/op",
            "extra": "240 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000 - B/op",
            "value": 4525702,
            "unit": "B/op",
            "extra": "240 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_10000 - allocs/op",
            "value": 20568,
            "unit": "allocs/op",
            "extra": "240 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000",
            "value": 51442032,
            "unit": "ns/op\t44847088 B/op\t  200627 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000 - ns/op",
            "value": 51442032,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000 - B/op",
            "value": 44847088,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameRead/GopherFrame_Read_100000 - allocs/op",
            "value": 200627,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000",
            "value": 158912,
            "unit": "ns/op\t  170986 B/op\t     169 allocs/op",
            "extra": "7138 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000 - ns/op",
            "value": 158912,
            "unit": "ns/op",
            "extra": "7138 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000 - B/op",
            "value": 170986,
            "unit": "B/op",
            "extra": "7138 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_1000 - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "7138 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000",
            "value": 1502074,
            "unit": "ns/op\t 1385965 B/op\t     223 allocs/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000 - ns/op",
            "value": 1502074,
            "unit": "ns/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000 - B/op",
            "value": 1385965,
            "unit": "B/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_10000 - allocs/op",
            "value": 223,
            "unit": "allocs/op",
            "extra": "799 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000",
            "value": 13841288,
            "unit": "ns/op\t18619523 B/op\t     292 allocs/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000 - ns/op",
            "value": 13841288,
            "unit": "ns/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000 - B/op",
            "value": 18619523,
            "unit": "B/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameFilter/GopherFrame_Filter_100000 - allocs/op",
            "value": 292,
            "unit": "allocs/op",
            "extra": "92 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000",
            "value": 64201,
            "unit": "ns/op\t   27616 B/op\t     112 allocs/op",
            "extra": "18727 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000 - ns/op",
            "value": 64201,
            "unit": "ns/op",
            "extra": "18727 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000 - B/op",
            "value": 27616,
            "unit": "B/op",
            "extra": "18727 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_1000 - allocs/op",
            "value": 112,
            "unit": "allocs/op",
            "extra": "18727 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000",
            "value": 570903,
            "unit": "ns/op\t  307300 B/op\t     137 allocs/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000 - ns/op",
            "value": 570903,
            "unit": "ns/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000 - B/op",
            "value": 307300,
            "unit": "B/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_10000 - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "2086 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000",
            "value": 5137130,
            "unit": "ns/op\t 3310829 B/op\t     172 allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000 - ns/op",
            "value": 5137130,
            "unit": "ns/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000 - B/op",
            "value": 3310829,
            "unit": "B/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameGroupBy/GopherFrame_GroupBy_100000 - allocs/op",
            "value": 172,
            "unit": "allocs/op",
            "extra": "234 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000",
            "value": 4738443,
            "unit": "ns/op\t  168139 B/op\t     157 allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000 - ns/op",
            "value": 4738443,
            "unit": "ns/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000 - B/op",
            "value": 168139,
            "unit": "B/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_1000 - allocs/op",
            "value": 157,
            "unit": "allocs/op",
            "extra": "253 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000",
            "value": 461591428,
            "unit": "ns/op\t 2181861 B/op\t     217 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000 - ns/op",
            "value": 461591428,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000 - B/op",
            "value": 2181861,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_10000 - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000",
            "value": 51330183273,
            "unit": "ns/op\t18270976 B/op\t     264 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000 - ns/op",
            "value": 51330183273,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000 - B/op",
            "value": 18270976,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameSort/GopherFrame_Sort_100000 - allocs/op",
            "value": 264,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000",
            "value": 1284968,
            "unit": "ns/op\t  909352 B/op\t   22068 allocs/op",
            "extra": "904 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000 - ns/op",
            "value": 1284968,
            "unit": "ns/op",
            "extra": "904 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000 - B/op",
            "value": 909352,
            "unit": "B/op",
            "extra": "904 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_1000 - allocs/op",
            "value": 22068,
            "unit": "allocs/op",
            "extra": "904 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000",
            "value": 12593943,
            "unit": "ns/op\t 9454966 B/op\t  220079 allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000 - ns/op",
            "value": 12593943,
            "unit": "ns/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000 - B/op",
            "value": 9454966,
            "unit": "B/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_10000 - allocs/op",
            "value": 220079,
            "unit": "allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000",
            "value": 128967126,
            "unit": "ns/op\t97308034 B/op\t 2200094 allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000 - ns/op",
            "value": 128967126,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000 - B/op",
            "value": 97308034,
            "unit": "B/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaRead/Gota_Read_100000 - allocs/op",
            "value": 2200094,
            "unit": "allocs/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000",
            "value": 100207,
            "unit": "ns/op\t  169946 B/op\t      88 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000 - ns/op",
            "value": 100207,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000 - B/op",
            "value": 169946,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_1000 - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000",
            "value": 1080943,
            "unit": "ns/op\t 2359538 B/op\t     130 allocs/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000 - ns/op",
            "value": 1080943,
            "unit": "ns/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000 - B/op",
            "value": 2359538,
            "unit": "B/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_10000 - allocs/op",
            "value": 130,
            "unit": "allocs/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000",
            "value": 10395825,
            "unit": "ns/op\t27148558 B/op\t     186 allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000 - ns/op",
            "value": 10395825,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000 - B/op",
            "value": 27148558,
            "unit": "B/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaFilter/Gota_Filter_100000 - allocs/op",
            "value": 186,
            "unit": "allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000",
            "value": 305729,
            "unit": "ns/op\t  268108 B/op\t     400 allocs/op",
            "extra": "3672 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000 - ns/op",
            "value": 305729,
            "unit": "ns/op",
            "extra": "3672 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000 - B/op",
            "value": 268108,
            "unit": "B/op",
            "extra": "3672 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_1000 - allocs/op",
            "value": 400,
            "unit": "allocs/op",
            "extra": "3672 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000",
            "value": 3277101,
            "unit": "ns/op\t 3177516 B/op\t     551 allocs/op",
            "extra": "361 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000 - ns/op",
            "value": 3277101,
            "unit": "ns/op",
            "extra": "361 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000 - B/op",
            "value": 3177516,
            "unit": "B/op",
            "extra": "361 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_10000 - allocs/op",
            "value": 551,
            "unit": "allocs/op",
            "extra": "361 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000",
            "value": 33374731,
            "unit": "ns/op\t33425253 B/op\t     763 allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000 - ns/op",
            "value": 33374731,
            "unit": "ns/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000 - B/op",
            "value": 33425253,
            "unit": "B/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaGroupBy/Gota_GroupBy_Simulation_100000 - allocs/op",
            "value": 763,
            "unit": "allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000",
            "value": 433565,
            "unit": "ns/op\t  218484 B/op\t      44 allocs/op",
            "extra": "2773 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000 - ns/op",
            "value": 433565,
            "unit": "ns/op",
            "extra": "2773 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000 - B/op",
            "value": 218484,
            "unit": "B/op",
            "extra": "2773 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_1000 - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "2773 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000",
            "value": 6206752,
            "unit": "ns/op\t 2666505 B/op\t      59 allocs/op",
            "extra": "192 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000 - ns/op",
            "value": 6206752,
            "unit": "ns/op",
            "extra": "192 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000 - B/op",
            "value": 2666505,
            "unit": "B/op",
            "extra": "192 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_10000 - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "192 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000",
            "value": 84820901,
            "unit": "ns/op\t30105658 B/op\t      79 allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000 - ns/op",
            "value": 84820901,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000 - B/op",
            "value": 30105658,
            "unit": "B/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaSort/Gota_Sort_100000 - allocs/op",
            "value": 79,
            "unit": "allocs/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage",
            "value": 7256201,
            "unit": "ns/op\t 6117476 B/op\t   20932 allocs/op",
            "extra": "165 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage - ns/op",
            "value": 7256201,
            "unit": "ns/op",
            "extra": "165 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage - B/op",
            "value": 6117476,
            "unit": "B/op",
            "extra": "165 times\n4 procs"
          },
          {
            "name": "BenchmarkGopherFrameMemory/GopherFrame_Memory_Usage - allocs/op",
            "value": 20932,
            "unit": "allocs/op",
            "extra": "165 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage",
            "value": 19090074,
            "unit": "ns/op\t13805172 B/op\t  220266 allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage - ns/op",
            "value": 19090074,
            "unit": "ns/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage - B/op",
            "value": 13805172,
            "unit": "B/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkGotaMemory/Gota_Memory_Usage - allocs/op",
            "value": 220266,
            "unit": "allocs/op",
            "extra": "57 times\n4 procs"
          }
        ]
      }
    ]
  }
}