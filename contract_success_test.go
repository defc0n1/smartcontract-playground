//install abigen as described here: https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
//go:generate abigen --sol TemperatureMeasurementA.sol --pkg main --out TemperatureMeasurementA.go
//go:generate abigen --sol TemperatureMeasurementB.sol --pkg main --out TemperatureMeasurementB.go

package main

import (
	"testing"
	"log"
)

func TestReportSuccess30(t *testing.T) {
	bs := NewBlockSimulator()
	bs.TransferETH()
	temp := make([]int8, 30, 30)
	temp[0] = -2;
	temp[1] = 30;
	timestamp := make([]uint32, 30, 30)
	_, err := bs.contractA.ReportTemperature(authTempWriter, temp, timestamp)
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp)
	if err == nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	bs.sim.Commit()

	retValBool1, _ := bs.contractA.Success(nil)
	if !retValBool1 {
		t.Log("reported temps need to be a success!")
		t.Fail()
	}
	retValBool1, _ = bs.contractB.Success(nil)
	if !retValBool1 {
		t.Log("reported temps need to be a success!")
		t.Fail()
	}

	retValBool2, _ := bs.contractA.Failed(nil)
	if retValBool2 {
		t.Log("reported temps cannot fail")
		t.Fail()
	}
	retValBool2, _ = bs.contractB.Failed(nil)
	if retValBool2 {
		t.Log("reported temps cannot fail")
		t.Fail()
	}

	retValNr1, _ := bs.contractA.FailedTemperaturesLength(nil)
	if retValNr1 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}
	retValNr2, _ := bs.contractB.FailedTimestampSeconds(nil)
	if retValNr2 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}

	retValNr3, _ := bs.contractA.NrFailures(nil)
	if retValNr3 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}
	retValNr3, _ = bs.contractB.NrFailures(nil)
	if retValNr3 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}

	retValNr4, _ := bs.contractA.NrMeasurements(nil)
	if retValNr4 != 30 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}
	retValNr4, _ = bs.contractB.NrMeasurements(nil)
	if retValNr4 != 30 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}

}

func TestReportSuccess60(t *testing.T) {
	bs := NewBlockSimulator()
	bs.TransferETH()
	temp := make([]int8, 30, 30)
	temp[0] = 30;
	temp[1] = -2;
	timestamp := make([]uint32, 30, 30)
	_, err := bs.contractA.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	bs.sim.Commit()

	_, err = bs.contractA.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	bs.sim.Commit()

	retValBool1, _ := bs.contractA.Success(nil)
	if !retValBool1 {
		t.Log("reported temps need to be a success!")
		t.Fail()
	}
	retValBool1, _ = bs.contractB.Success(nil)
	if !retValBool1 {
		t.Log("reported temps need to be a success!")
		t.Fail()
	}

	retValBool2, _ := bs.contractA.Failed(nil)
	if retValBool2 {
		t.Log("reported temps cannot fail")
		t.Fail()
	}
	retValBool2, _ = bs.contractB.Failed(nil)
	if retValBool2 {
		t.Log("reported temps cannot fail")
		t.Fail()
	}

	retValNr1, _ := bs.contractA.FailedTemperaturesLength(nil)
	if retValNr1 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}
	retValNr2, _ := bs.contractB.FailedTimestampSeconds(nil)
	if retValNr1 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}

	retValNr3, _ := bs.contractA.NrFailures(nil)
	if retValNr2 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}
	retValNr3, _ = bs.contractB.NrFailures(nil)
	if retValNr3 != 0 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}

	retValNr4, _ := bs.contractA.NrMeasurements(nil)
	if retValNr4 != 60 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}
	retValNr4, _ = bs.contractB.NrMeasurements(nil)
	if retValNr4 != 60 {
		t.Log("reported temps cannot have fail temps")
		t.Fail()
	}

}

