// Code generated by goa v3.5.5, DO NOT EDIT.
//
// identity HTTP client CLI support package
//
// Command:
// $ goa gen github.com/vinhphuctadang/authorizor/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	authorizorc "github.com/vinhphuctadang/authorizor/gen/http/authorizor/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `authorizor (register|ping)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` authorizor register --body '{
      "password": "Magnam velit et assumenda adipisci excepturi eos.",
      "username": "Cum doloribus molestias qui."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		authorizorFlags = flag.NewFlagSet("authorizor", flag.ContinueOnError)

		authorizorRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		authorizorRegisterBodyFlag = authorizorRegisterFlags.String("body", "REQUIRED", "")

		authorizorPingFlags = flag.NewFlagSet("ping", flag.ExitOnError)
	)
	authorizorFlags.Usage = authorizorUsage
	authorizorRegisterFlags.Usage = authorizorRegisterUsage
	authorizorPingFlags.Usage = authorizorPingUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "authorizor":
			svcf = authorizorFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "authorizor":
			switch epn {
			case "register":
				epf = authorizorRegisterFlags

			case "ping":
				epf = authorizorPingFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "authorizor":
			c := authorizorc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = authorizorc.BuildRegisterPayload(*authorizorRegisterBodyFlag)
			case "ping":
				endpoint = c.Ping()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// authorizorUsage displays the usage of the authorizor command and its
// subcommands.
func authorizorUsage() {
	fmt.Fprintf(os.Stderr, `Service for authenticating and authorizing
Usage:
    %[1]s [globalflags] authorizor COMMAND [flags]

COMMAND:
    register: Register implements register.
    ping: Ping implements ping.

Additional help:
    %[1]s authorizor COMMAND --help
`, os.Args[0])
}
func authorizorRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] authorizor register -body JSON

Register implements register.
    -body JSON: 

Example:
    %[1]s authorizor register --body '{
      "password": "Magnam velit et assumenda adipisci excepturi eos.",
      "username": "Cum doloribus molestias qui."
   }'
`, os.Args[0])
}

func authorizorPingUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] authorizor ping

Ping implements ping.

Example:
    %[1]s authorizor ping
`, os.Args[0])
}