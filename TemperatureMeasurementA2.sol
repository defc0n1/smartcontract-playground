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
pragma solidity ^0.4.4;

contract TemperatureMeasurementA2 {

    address owner;
    address temperatureWriter;
    string storageLocation;

    /* set at contract creation */
    bytes1 minTemperature;
    bytes1 maxTemperature;
    address signerAddress;

    /* state */
    uint32 measurements = 0;
    bool isFailed = false;
    bytes8 startTimestamp;
    bytes1 measurementInterval;
    bytes32 hash;
    // the signature is composed of values v, r, s
    uint8 v;
    bytes32 r;
    bytes32 s;

    /* Constructor, set who is allowed to write and the temperature range */
    function TemperatureMeasurementA2(address _temperatureWriter,
            bytes1 _minTemperature, bytes1 _maxTemperature,
            string _storageLocation, address _signerAddress) {

        owner = msg.sender;
        temperatureWriter = _temperatureWriter;
        minTemperature = _minTemperature;
        maxTemperature = _maxTemperature;
        storageLocation = _storageLocation;
        signerAddress = _signerAddress;
    }

    /* Any remaining funds should be sent back to the sender
       Both, sender and the reporter can call this function */
    function done() {
        if (msg.sender == owner || msg.sender == temperatureWriter) {
            suicide(msg.sender);
        }
    }

    // the signature is composed of values v, r, s
    function reportResult(uint32 _measurements, bytes8 _startTimestamp,
            bytes1 _measurementInterval, bytes32 _hash, bool _isFailed,
            uint8 _v, bytes32 _r, bytes32 _s) public {

        /* Only temperature reporter can write temperature */

        // only one temperature report allowed per shipment
        if (measurements != 0) {
            throw;
        }

        if (msg.sender != temperatureWriter) {
            throw;
        }

        measurements = _measurements;
        startTimestamp = _startTimestamp;
        measurementInterval = _measurementInterval;
        hash = _hash;
        isFailed = _isFailed;
        v = _v;
        r = _r;
        s = _s;
    }

    function generateReport(bytes1[] _temperatures) constant
            returns (bytes32 , bool) {

        var len = uint32 (_temperatures.length);

        bool _isFailed = false;

        // use hash of byte measurements concatenated with start timestamp and measurement interval for signature checking
        bytes memory b = new bytes(len + 9);
        for (uint16 i = 0; i < len; i++) {
            b[i] = _temperatures[i];

            if(_temperatures[i] > maxTemperature || _temperatures[i] < minTemperature) {
                _isFailed = true;
            }
        }

        
        for (uint8 j = 0; j < 8; j++) {
            b[len + j] = startTimestamp[j];
        }
        b[len + 8] = measurementInterval;

        return (sha256(b), _isFailed);
    }

    function verifyReport(bytes1[] _temperatures)
            constant returns (bool) {

        if(_temperatures.length != measurements) {
            return false;
        }

        var (_hash, _isFailed) = generateReport(_temperatures);

        if (_isFailed != isFailed) return false;
        if (_hash != hash) return false;

        address _keyToCompare = ecrecover(_hash, v, r, s);
        if (_keyToCompare != signerAddress) return false;

        return true;
    }

    /* Success is when no failed timestamp was reported and at least
       one measurement was carried out */
    function success() constant returns (bool) {
       return !isFailed && measurements > 0;
    }

    /* Failed is when one failed timestamp was reported and at least
       one measurement was carried out */
    function failed() constant returns (bool) {
       return isFailed && measurements > 0;
    }

    /* The number of carried out measurements */
    function nrMeasurements() constant returns (uint32) {
       return measurements;
    }

    /* The temperature range to check */
    function temperatureRange() constant returns (bytes1, bytes1) {
       return (minTemperature, maxTemperature);
    }

    /* The timestamp range */
    function getStartTimestamp() constant returns (bytes8) {
       return startTimestamp;
    }
    function getMeasurementInterval() constant returns (bytes1) {
       return measurementInterval;
    }

    function getHash() constant returns (bytes32) {
       return hash;
    }

    function getSignature() constant returns (uint8, bytes32, bytes32) {
        return (v, r, s);
    }
}

