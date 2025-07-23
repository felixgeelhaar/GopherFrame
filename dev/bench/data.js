window.BENCHMARK_DATA = {
  "lastUpdate": 1753303143067,
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
      }
    ]
  }
}