package gobag

import "regexp"

// TimeStampRegex is the universal Para regex for validating timestamps. This
// regex pattern validates a timestamp to an ISO8601 format that has an
// optional time zone.
var TimeStampISO8601Regex = regexp.MustCompile(`^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:Z|[+-][01]\d:?[0-5]\d)?$`)

// TimeStampRegex is the universal Para regex for validating timestamps. This
// regex pattern validates a timestamp to an RFC3339 format that requires the
// time zone to be present.
var TimeStampRFC3999Regex = regexp.MustCompile(`^(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)(T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d)(\.\d{1,4})?(Z|([+-][01]\d:?[0-5]\d))$`)

// BoolRegex ...
var BoolRegex = regexp.MustCompile(`^(true|false)$`)
