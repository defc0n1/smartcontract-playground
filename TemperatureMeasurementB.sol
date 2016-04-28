contract TemperatureMeasurementB {

    address owner;
    address temperatureWriter;

    uint8 minTemperature;
    uint8 maxTemperature;
    
    uint32 firstFailedTimestampSeconds = 0;
    uint32 measurements = 0;
    uint32 failures = 0;

    /* Constructor, set who is allowed to write and the temperature range */
    function TemperatureMeasurementB(address _temperatureWriter, 
            uint8 _minTemperature, uint8 _maxTemperature) {
        owner = msg.sender;
	    temperatureWriter = _temperatureWriter;
        minTemperature = _minTemperature;
        maxTemperature = _maxTemperature;
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
        if (msg.sender == temperatureWriter) {
            if(_temperatures.length != _timestamps.length) {
                throw;
            }
            /* read state variable, writing directly to it is expensive */
            var _measurements = measurements;
            var _failures = failures;
            var _maxTemperature = maxTemperature;
            var _minTemperature = minTemperature;
            var _firstFailedTimestampSeconds = firstFailedTimestampSeconds;
            for (uint32 i = 0; i < _temperatures.length; i++) {
		        _measurements++;
                if(_temperatures[i] > _maxTemperature 
                        || _temperatures[i] < _minTemperature) {
                    if(_firstFailedTimestampSeconds == 0) {
                        _firstFailedTimestampSeconds = _timestamps[i];
                        /* write state back, this is the expensive work */
                        firstFailedTimestampSeconds = _firstFailedTimestampSeconds;
                    }
                    _failures++;
                }
            }
            /* write state back, this is the expensive work */
            failures = _failures;
            measurements = _measurements;
        }
    }

    /* Success is when no failed timestamp was reported and at least
       one measurement was carried out */
    function success() constant returns (bool res) {
       return firstFailedTimestampSeconds == 0 && measurements > 0;
    }
    
    /* Failed is when one failed timestamp was reported and at least
       one measurement was carried out */
    function failed() constant returns (bool res) {
       return firstFailedTimestampSeconds > 0 && measurements > 0;
    }
    
    /* The number of carried out measurements */
    function nrMeasurements() constant returns (uint32) {
       return measurements;
    }
    
    /* The number of reported failures */
    function nrFailures() constant returns (uint32) {
       return failures;
    }
    
    /* The first timestamp of the frist temperature that was out of range */
    function failedTimestampSeconds() constant returns (uint32) {
       return firstFailedTimestampSeconds;
    }
    
    /* The temperature range to check */
    function temperatureRange() constant returns (uint8,uint8) {
       return (minTemperature, maxTemperature);
    }
}
