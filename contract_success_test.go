/*
 * Copyright 2016 Modum.io and the CSG Group at University of Zurich
 *
 * Licensed under the Apache License, Version 2.0 (the 'License'); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package smartcontract

import (
	"testing"
	//"log"
	"log"
)

func TestReportSuccess30(t *testing.T) {
	bs := NewBlockSimulator()
	bs.TransferETH()
	temp := make([]int8, 30, 30)
	temp[0] = -2;
	temp[1] = 30;
	timestamp := make([]uint32, 30, 30)

	failedTimestampSeconds, failedTemperatures, failures, measurements, hash, err := bs.contractA.GenerateReport(nil, temp, timestamp)
	//log.Printf("XX: %v %v %v %v %v %v\n", failedTimestampSeconds, failedTemperatures, failures, measurements, hash, err)
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}

	_, err = bs.contractA.ReportResult(authTempWriter, failedTimestampSeconds, failedTemperatures,
		failures, measurements, timestamp[0], timestamp[len(timestamp) - 1], hash)
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}

	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp)
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
	failedTimestampSeconds, failedTemperatures, failures, measurements, hash, err :=  bs.contractA.GenerateReport(nil, temp, timestamp)
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractA.ReportResult(authTempWriter, failedTimestampSeconds, failedTemperatures,
		failures, measurements, timestamp[0], timestamp[len(timestamp) - 1], hash)
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractB.ReportTemperature(authTempWriter, temp, timestamp);
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	bs.sim.Commit()

	failedTimestampSeconds, failedTemperatures, failures, measurements, hash, err =  bs.contractA.GenerateReport(nil, temp, timestamp)
	if err != nil {
		log.Printf("Failed to report temperatures: %v", err)
	}
	_, err = bs.contractA.ReportResult(authTempWriter, failedTimestampSeconds, failedTemperatures,
		failures, measurements, timestamp[0], timestamp[len(timestamp) - 1], hash)
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