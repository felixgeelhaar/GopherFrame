window.BENCHMARK_DATA = {
  "lastUpdate": 1753303320565,
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
      }
    ]
  }
}