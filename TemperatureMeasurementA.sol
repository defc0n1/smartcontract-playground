contract TemperatureMeasurementA {

    address owner;
    address temperatureWriter;

    uint8 minTemperature;
    uint8 maxTemperature;
    
    uint32[] failedTimestampSeconds;
    uint16 maxFailureReports;
    uint32 measurements = 0;
    uint32 failures = 0;
    uint32 firstTimestamp = 0;
    uint32 lastTimestamp = 0;

    /* Constructor, set who is allowed to write and the temperature range */
    function TemperatureMeasurementA(address _temperatureWriter, 
            uint8 _minTemperature, uint8 _maxTemperature, uint16 _maxFailureReports) {
        owner = msg.sender;
	    temperatureWriter = _temperatureWriter;
        minTemperature = _minTemperature;
        maxTemperature = _maxTemperature;
        if(_maxFailureReports < 1) {
            throw;
        }
        maxFailureReports = _maxFailureReports;
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
    function temperatureRange() constant returns (uint8,uint8) {
       return (minTemperature, maxTemperature);
    }

    /* The timestamp range */
    function timestampRange() constant returns (uint32,uint32) {
       return (firstTimestamp, lastTimestamp);
    }
}
