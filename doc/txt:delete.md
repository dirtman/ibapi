# NAME

ibapi txt delete - delete Infoblox TXT records

# USAGE

- ibapi txt delete &lt;options/args>

# DESCRIPTION

The delete command is used to delete Infoblox TXT records.
To delete a single TXT record, a single hostname and optionally a TXT value may
be provided as command line arguments.
Alternatively, a list of records to delete can be specified in a file (see --filename).

If a TXT value is specified, the TXT record to delete must contain that TXT value, else
no TXT record will be deleted.  If no TXT value is specified and only one TXT record
is found for the specified name, that TXT record is deleted regardless of its TXT value.
If multiple TXT records are found for the same
request, the deletion process is aborted (no records are deleted) unless the --multiple
options is specified to allow mutliple record deletions per request.

# OPTIONS

Some options can be specified with either a short (i.e., -h) or long
(i.e., --help) form.  In the latter case, case is non-sensitive.

Boolean options (flags) do not require a value.  "-v" is equivalent to "-v=true".
To disable, set to "false" ("-v=false" or "--verbose=false").

Most options have a corresponding configuration file setting
that is equal to the long option name without the "--" prefix.
Command line options always override configuration file settings.
Use the --ShowConfig to view each option and its value.

## OPTIONS - General

- -V &lt;view>, --View=&lt;view>:

    Specify the view of the record to delete.  Default: "default".

- -f &lt;filename>, --filename=&lt;filename>:

    Specify a filename containing a list of TXT records to delete.
    Each line should contain a hostname and optionally a TXT value,
    separated by one or more spaces.
    Blank lines and lines beginning with "#" are ignored, as is anything on a line
    following a "#".

- -m, --multiple:

    If only a name is specified (no TXT value is specified), allow deletion of 
    multiple records if multiple records are found for the specified name.
    This option has no effect if both the name and data value are specified.

## OPTIONS - API Options

- --APIBaseURL=&lt;url>:

    API base URL.
    Default: "https://infoblox.rice.edu/wapi/v2.11"

- --HTTPTimeout=&lt;seconds>:

    Timeout in seconds for the HTTP connection.
    Default: 10.

- --APIAuthMethod=&lt;method>:

    WAPI authentication method for accessing the Infoblox API.  Currently,
    only "Basic" authentication (username and password) is supported.

- --username=&lt;username>:

    Specify the username used for basic auth.

- --password=&lt;password>:

    Specify the password used for basic auth.  If this option is specified
    and is non-empty, either the --username option can be used to specify
    the corresponding username, or the current user will be assumed for username.

- --PromptForPassword:

    Prompt the user for the password used for basic authentication.  This is done
    automatically unless one of these options is specified: --password,
    \--APIAuthToken, --APIAuthTokenID.

- --APIAuthToken=&lt;token>:

    As an alternative to specifing a username and password separately, an "authentication token" 
    containing both the username and password, separated by a ':', can be specified.

- --APIAuthTokenID=&lt;tokenID>:

    As an alternative to specifying a username/password or an APIAuthToken
    directly, the name of an authentication token ID which maps to an
    authentication token can be specified.  See also --SecretsDir.
    Default: "owlapi\_basic".

- --SecretsDir=&lt;pathname>:

    Specify a directory where optional "secret files" are kept.  These files can be
    used as a slightly safer alternative to storing authentication credentials in
    the standard configuration files.  The name of a secret file corresponds to a
    "tokenID" configured via --APIAuthTokenID, and the file contains an
    authentication token.  Obviously, these files should be safely guarded or
    avoided altogether.
    Default: "/etc/opt/ibapi/private".

- --GridReference=&lt;grid\_reference>:

    Specify the Infoblox grid reference ID. This can be used to save a fetch when
    the --restartServices option is specified.  While this option is only relevant
    to a few commands, it is allowed (ignored) by the other commands.

## OPTIONS - Common To All IBAPI Commands

- -h, --help:

    Help; show usage information and exit.

- --showConfig:

    Read in and show all configuration settings and exit.

- -q, --Quiet:

    Be quieter than normal.

- --Quieter:

    Quieter mode.  Suppress all messages except warning and error messages.

- -v, --Verbose:

    Be louder than normal. Over-rides the "--Quiet"  and "-Quieter" options.
    Note such extra details are printed to Stderr so that the normal output
    remains the same regardless of verbosity.

- --page:

    Page help/usage information via the command specified by the --Pager option or
    the environment variable "PAGER".  If neither of these is set, this option 
    is ignored.  Default: true.

- --Pager=&lt;pager>:

    Specify a pager command for paging the usage information (shown with --help).  By default,
    the environment variable PAGER is used.  If a full path is not specified, the command
    is searched for using the PATH environment variable.

# EXAMPLES

- ibapi txt delete t1.txt.rice.edu "v=spf1 a:mh.rice.edu a:a16.spf.rice.edu/16 -all"

    Delete the TXT record with hostname "t1.txt.rice.edu" and TXT value "v=spf1 a:mh.rice.edu a:a16.spf.rice.edu/16 -all".

- ibapi txt delete t1.txt.rice.edu

    Delete all TXT records named t1.txt.rice.edu.

# FILES

- /usr/site/ibapi-1.0/etc/ibapi.conf
- /etc/opt/ibapi/ibapi.conf
- /etc/opt/ibapi-1.0//ibapi.conf
- ~/.ibapi/ibapi.conf
- ~/.ibapi-1.0/ibapi.conf

    The IBAPI configuration files which can be used to
    set defaults for nearly all of the options described above.
    Any combination of these may be used.
    Each file found is read in turn, with settings in later files
    overriding those in previous files.  Note that command line
    options override all config file settings.

# SEE ALSO

a(1),
a:add(1),
a:delete(1),
a:get(1),
a:update(1),
alias(1),
alias:add(1),
alias:delete(1),
alias:get(1),
alias:update(1),
cname(1),
cname:add(1),
cname:delete(1),
cname:get(1),
cname:update(1),
fixedaddress(1),
fixedaddress:add(1),
fixedaddress:delete(1),
fixedaddress:get(1),
fixedaddress:update(1),
grid(1),
host(1),
host:add(1),
host:delete(1),
host:get(1),
host:update(1),
ibapi(1),
ptr(1),
ptr:add(1),
ptr:delete(1),
ptr:get(1),
ptr:update(1),
url(1),
url:add(1),
url:delete(1),
url:get(1),
url:update(1),
mx(1),
mx:add(1),
mx:delete(1),
mx:get(1),
mx:update(1),
txt(1),
txt:add(1),
txt:get(1),
txt:update(1),
ibapi.conf(5)
