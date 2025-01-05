package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkGetDomainStat(b *testing.B) {
	zipReader, err := zip.OpenReader("testdata/users.dat.zip")
	if err != nil {
		b.Errorf("data source opening error: %v", err)
	}

	defer func() {
		errC := zipReader.Close()
		if errC != nil {
			b.Errorf("data source closing error: %v", errC)
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		row, err := zipReader.File[0].Open()
		if err != nil {
			b.Errorf("reading content error: %v", err)
		}

		b.StartTimer()
		_, err = GetDomainStat(row, "biz")
		b.StopTimer()
		if err != nil {
			b.Errorf("domain stat getting error: %v", err)
		}
	}
}
