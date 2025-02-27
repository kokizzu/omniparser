// Code generated - DO NOT EDIT.

package validation

const (
    JSONSchemaParserSettings =
`
{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "github.com/jf-tech/omniparser:parser_settings",
    "title": "omniparser schema: parser_settings",
    "type": "object",
    "properties": {
        "parser_settings": {
            "type": "object",
            "properties": {
                "version": { "type": "string" },
                "file_format_type": { "type": "string" },
                "encoding": {
                    "type": "string",
                    "enum": [ "utf-8", "iso-8859-1", "windows-1252" ]
                },
                "debug": { "type": "integer", "minimum": 0 }
            },
            "required": [ "version", "file_format_type" ],
            "additionalProperties": false
        }
    },
    "required": [ "parser_settings" ]
}

`
)
