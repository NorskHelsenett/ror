# RLog

## Why RLog

Rlog is a wrapper for the Uber Zap logging framework. The reason we want to wrap
zap and not use it directly is that it makes it easier to change logger in the
future if necessary. RLog was made when we wanted to switch from Logrus to Zap
and encountered this exact problem.

## What RLog offers

Rlog offers a global logger with both context (go context) and non context aware
logging functions. Context aware functions will add certain predefined fields
from a context and also correlate any trace information with the logs.

Rlog offers leveled and structured logging in the same manner as Zap does. The
levels supported are Info, Debug, Warn, Error and Fatal. Logs are normally
output as JSON but can be output in a more readable form when developing, its
important to note that in prod all logs shall be in the JSON format.

## Config

Rlog is configured using environment variables, from either an env file or the
environment its running in.

### LOG_LEVEL

sets the log level

### LOG_OUTPUT

set where the logs should be sent, defaults to stderr. Logs can be sent to a
file an url or stdout/stderr. Logs can be sent to multiple locations with a
comma separated string. E.g. LOG_OUTPUT="/home/user/foo/.ror/log,stderr".
This example logs to both a file and stderr

### LOG_OUTPUT_ERROR

Set where error level logs should be sent, works the same as LOG_OUTPUT

### LOG_DEVELOP

Set wheter or not logs should be on JSON format or a more human readable format.
False = JSON
True = more human readable

