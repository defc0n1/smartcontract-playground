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

contract TemperatureMeasurementA {

    address owner;
    address temperatureWriter;
    string storageLocation;

    uint8 minTemperature;
    uint8 maxTemperature;
    
    uint32[] failedTimestampSeconds;
    uint16 maxFailureReports;
    uint32 measurements = 0;
    uint32 failures = 0;
    uint32 firstTimestamp = 0;
    uint32 lastTimestamp = 0;
    bytes32[] hashes;

    /* Constructor, set who is allowed to write and the temperature range */
    function TemperatureMeasurementA(address _temperatureWriter, 
            uint8 _minTemperature, uint8 _maxTemperature, 
            uint16 _maxFailureReports, string _storageLocation) {
        owner = msg.sender;
	    temperatureWriter = _temperatureWriter;
        minTemperature = _minTemperature;
        maxTemperature = _maxTemperature;
        if(_maxFailureReports < 1) {
            throw;
        }
        maxFailureReports = _maxFailureReports;
        storageLocation = _storageLocation;
    }

    /* Any remaining funds should be sent back to the sender 
       Both, sender and the reporter can call this function */
    function done() {
        if (msg.sender == owner || msg.sender == temperatureWriter) {
            suicide(msg.sender);
        }
    }
  
    function reportTemperature(uint8[] _temperatures, uint32[] _timestamps) public {
        /* Only temperature reporter can write temperature */
        if (msg.sender != temperatureWriter) {
            throw;
        }
            
        if(_temperatures.length != _timestamps.length) {
            throw;
        }
        
        /* calculate hash of input, store hash in array -> expensive operation */
        bytes memory b = new bytes(5*_temperatures.length);
        for (uint16 j = 0; j < _temperatures.length; j++) {
            b[(j*5)+0]=bytes1(_timestamps[j]);
            b[(j*5)+1]=bytes1(shr(_timestamps[j], 8));
            b[(j*5)+2]=bytes1(shr(_timestamps[j], 16));
            b[(j*5)+3]=bytes1(shr(_timestamps[j], 24));
            b[(j*5)+4]=bytes1(_temperatures[j]);
        }
        hashes.push(sha256(b));
        
        /* read state variable, writing directly to it is expensive */
        var _measurements = measurements;
        uint32 _failures = failures;
        var _firstTimestamp = firstTimestamp;
        var _lastTimestamp = lastTimestamp;
        uint8 _maxTemperature = maxTemperature;
        uint8 _minTemperature = minTemperature;
        uint16 _maxFailureReports = maxFailureReports;
        var _currentFailureReports = failedTimestampSeconds.length;
            
        /* set first timestamp of measurement */
        if(_firstTimestamp == 0 && _timestamps.length > 0) {
		    _firstTimestamp = _timestamps[0];
		}
		/* set last timestamp of measurement */
		    if(_timestamps.length > 0 && _timestamps[_timestamps.length-1] > _lastTimestamp) {
	        _lastTimestamp = _timestamps[_timestamps.length-1];
		}
        
        for (uint32 i = 0; i < _temperatures.length; i++) {
		       _measurements++;
		        
		       if(_temperatures[i] > _maxTemperature 
                    || _temperatures[i] < _minTemperature) {
                _failures++;
                if(_currentFailureReports < _maxFailureReports) {
                    _currentFailureReports++;
                    /* write state back, this is the expensive work */
                    failedTimestampSeconds.push(_timestamps[i]);
                }
                
            }
        }
        /* write state back, this is the expensive work */
        if(measurements != _measurements) {
            measurements = _measurements;
        }
        if(failures != _failures) {
            failures = _failures;
        }
        if(firstTimestamp != _firstTimestamp) {
            firstTimestamp = _firstTimestamp;
        }
        if(lastTimestamp != _lastTimestamp) {
            lastTimestamp = _lastTimestamp;
        }
        
    }

    /* Success is when no failed timestamp was reported and at least
       one measurement was carried out */
    function success() constant returns (bool) {
       return failedTimestampSeconds.length == 0 && measurements > 0;
    }
    
    /* Failed is when one failed timestamp was reported and at least
       one measurement was carried out */
    function failed() constant returns (bool) {
       return failedTimestampSeconds.length > 0 && measurements > 0;
    }
    
    /* The number of carried out measurements */
    function nrMeasurements() constant returns (uint32) {
       return measurements;
    }
    
    /* The number of reported failures */
    function nrFailures() constant returns (uint32) {
       return failures;
    }
    
    /* The timestamp of the frist temperature that was out of range */
    function failedTimestampSecondsAt(uint16 index) constant returns (uint32) {
       return failedTimestampSeconds[index];
    }
    
    /* The length of the failed timestamp array */
    function failedTimestampLength() constant returns (uint16) {
       return uint16(failedTimestampSeconds.length);
    }
    
    /* The temperature range to check */
    function temperatureMin() constant returns (uint8) {
       return minTemperature;
    }
    function temperatureMax() constant returns (uint8) {
       return maxTemperature;
    }

    /* The timestamp range */
    function timestampFirst() constant returns (uint32) {
       return firstTimestamp;
    }
    function timestampLast() constant returns (uint32) {
       return lastTimestamp;
    }

    /* shift right, currently not implemented: https://github.com/ethereum/solidity/issues/33 */
    function shr(uint32 input, byte bits) constant returns (uint32) {
        return input / (2 ** uint32(bits));
    }
}
