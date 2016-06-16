//install abigen as described here: https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
//go:generate abigen --sol TemperatureMeasurementA.sol --pkg main --out TemperatureMeasurementA.go
//go:generate abigen --sol TemperatureMeasurementB.sol --pkg main --out TemperatureMeasurementB.go

package main

import (
	"testing"
	"log"
)

func TestReportFailed1(t *testing.T) {
	bs := NewBlockSimulator()
	bs.TransferETH()
	temp := make([]int8, 30, 30)
	temp[0] = 31;
	timestamp := make([]uint32, 30, 30)
	for i := 0; i < 30; i++ {
		timestamp[i] = uint32(i + 1);
	}
	_, err := bs.contractA.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	bs.sim.Commit()

	retValBool1, _ := bs.contractA.Success(nil)
	if retValBool1 {
		t.Log("reported temps cannot to be a success!")
		t.Fail()
	}
	retValBool1, _ = bs.contractB.Success(nil)
	if retValBool1 {
		t.Log("reported temps cannot to be a success!")
		t.Fail()
	}

	retValBool2, _ := bs.contractA.Failed(nil)
	if !retValBool2 {
		t.Log("reported temps needs to fail")
		t.Fail()
	}
	retValBool2, _ = bs.contractB.Failed(nil)
	if !retValBool2 {
		t.Log("reported temps needs to fail")
		t.Fail()
	}

	retValNr1, _ := bs.contractA.FailedTemperaturesLength(nil)
	if retValNr1 != 1 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2, _ := bs.contractA.FailedTimestampSecondsAt(nil, 0)
	if retValNr2 != 1 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2_2, _ := bs.contractA.FailedTemperaturesAt(nil, 0)
	if retValNr2_2 != 31 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2, _ = bs.contractB.FailedTimestampSeconds(nil)
	if retValNr2 != 1 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}

	retValNr3, _ := bs.contractA.NrFailures(nil)
	if retValNr3 != 1 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr3, _ = bs.contractB.NrFailures(nil)
	if retValNr3 != 1 {
		t.Log("report temp has 1 failed temp")
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

func TestReportFailed2(t *testing.T) {
	bs := NewBlockSimulator()
	bs.TransferETH()
	temp := make([]int8, 30, 30)
	temp[0] = 31;
	temp[1] = -3;
	timestamp := make([]uint32, 30, 30)
	for i := 0; i < 30; i++ {
		timestamp[i] = uint32(i + 1);
	}
	_, err := bs.contractA.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	bs.sim.Commit()

	retValBool1, _ := bs.contractA.Success(nil)
	if retValBool1 {
		t.Log("reported temps cannot to be a success!")
		t.Fail()
	}
	retValBool1, _ = bs.contractB.Success(nil)
	if retValBool1 {
		t.Log("reported temps cannot to be a success!")
		t.Fail()
	}

	retValBool2, _ := bs.contractA.Failed(nil)
	if !retValBool2 {
		t.Log("reported temps needs to fail")
		t.Fail()
	}
	retValBool2, _ = bs.contractB.Failed(nil)
	if !retValBool2 {
		t.Log("reported temps needs to fail")
		t.Fail()
	}

	retValNr1, _ := bs.contractA.FailedTemperaturesLength(nil)
	if retValNr1 != 2 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2, _ := bs.contractA.FailedTimestampSecondsAt(nil, 0)
	if retValNr2 != 1 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2_2, _ := bs.contractA.FailedTemperaturesAt(nil, 0)
	if retValNr2_2 != 31 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2, _ = bs.contractB.FailedTimestampSeconds(nil)
	if retValNr2 != 1 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}

	retValNr3, _ := bs.contractA.NrFailures(nil)
	if retValNr3 != 2 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr3, _ = bs.contractB.NrFailures(nil)
	if retValNr3 != 2 {
		t.Log("report temp has 1 failed temp")
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

func TestReportFailed11(t *testing.T) {
	bs := NewBlockSimulator()
	bs.TransferETH()
	temp := make([]int8, 30, 30)
	temp[0] = -5;
	for i := 1; i < 11; i++ {
		temp[i] = -3;
	}
	timestamp := make([]uint32, 30, 30)
	for i := 0; i < 30; i++ {
		timestamp[i] = uint32(i + 1);
	}
	_, err := bs.contractA.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	bs.sim.Commit()

	retValBool1, _ := bs.contractA.Success(nil)
	if retValBool1 {
		t.Log("reported temps cannot to be a success!")
		t.Fail()
	}
	retValBool1, _ = bs.contractB.Success(nil)
	if retValBool1 {
		t.Log("reported temps cannot to be a success!")
		t.Fail()
	}

	retValBool2, _ := bs.contractA.Failed(nil)
	if !retValBool2 {
		t.Log("reported temps needs to fail")
		t.Fail()
	}
	retValBool2, _ = bs.contractB.Failed(nil)
	if !retValBool2 {
		t.Log("reported temps needs to fail")
		t.Fail()
	}

	retValNr1, _ := bs.contractA.FailedTemperaturesLength(nil)
	if retValNr1 != 10 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2, _ := bs.contractA.FailedTimestampSecondsAt(nil, 9)
	if retValNr2 != 10 {
		t.Log("report temp has 1 failed temp:", retValNr2)
		t.Fail()
	}
	retValNr2_2, _ := bs.contractA.FailedTemperaturesAt(nil, 0)
	if retValNr2_2 != -5 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr2, _ = bs.contractB.FailedTimestampSeconds(nil)
	if retValNr2 != 1 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}

	retValNr3, _ := bs.contractA.NrFailures(nil)
	if retValNr3 != 11 {
		t.Log("report temp has 1 failed temp")
		t.Fail()
	}
	retValNr3, _ = bs.contractB.NrFailures(nil)
	if retValNr3 != 11 {
		t.Log("report temp has 1 failed temp")
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



