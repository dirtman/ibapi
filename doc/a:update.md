# NAME

ibapi a update - update Infoblox A records

# USAGE

- ibapi a update &lt;options/args>

# DESCRIPTION

The update command is used to update Infoblox A records.

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

    Specify the view of the record to update.  Default: "default".

- -D &lt;true|false>, --Disable=&lt;true|false>:

    Update the record's disabled status to the specified value.
    Note this is not a boolean flag - the value "true" or "false"
    must be specified.

- -n &lt;new\_hostname>, --Name=&lt;new\_hostname>:

    Update the name of the specified record to "new\_hostname".

- -c &lt;comment>, --Comment=&lt;comment>:

    Update the record's comment.

- --TTL=&lt;ttl>:

    Update the record's TTL.

- -i &lt;new\_IP>, --ip=&lt;new\_ip>:

    Update the IP of the specified record to "new\_IP".

- -F &lt;fields>, --Fields=&lt;fields>:

    Specify fields and corresponding values to be updated.  For instance:
    "comment=RT100931",view=default,ttl=900".

- -f &lt;filename>, --filename=&lt;filename>:

    Specify a filename containing a list of records to update.
    Each line must contain a hostname and, depending on the specified options, an IP address.
    Blank lines and lines beginning with "#" are ignored, as is anything on a line
    following a "#".

- -C, --Check:

    Before updating a record's hostname or IP address, check if any "related" records
    already exist, and if so do not update the new record.
    Related records are those that share the same hostname and/or IP address, 
    such as an A record and Host record with the same hostname and/or IP address.
    Default: false.

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

- ibapi a update rb4.rice.edu 168.7.56.224 -i 168.7.56.225

    Update the "rb4.rice.edu/168.7.56.224" A record, changing the IP address 
    to "168.7.56.225".

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

ibapi(1),
host(1),
host:add(1),
ptr(1),
cname(1),
alias(1),
host:get(1),
host:delete(1),
host:update(1),
ptr:add(1),
ptr:delete(1),
ptr:get(1),
ptr:update(1),
cname:add(1),
cname:delete(1),
cname:get(1),
cname:update(1),
alias:add(1),
alias:delete(1),
alias:get(1),
alias:update(1),
url(1),
url:add(1),
url:delete(1),
url:get(1),
url:update(1),
a(1),
a:delete(1),
a:get(1),
a:add(1),
fixedaddress(1),
fixedaddress:add(1),
fixedaddress:delete(1),
fixedaddress:get(1),
fixedaddress:update(1),
grid(1),
mx:add(1),
mx:delete(1),
mx:get(1),
mx:update(1),
mx(1),
txt:add(1),
txt:delete(1),
txt:get(1),
txt:update(1),
txt(1),
authzone(1),
authzone:add(1),
authzone:delete(1),
authzone:get(1),
authzone:update(1),
aaaa(1),
aaaa:add(1),
aaaa:delete(1),
aaaa:get(1),
aaaa:update(1),
ibapi.conf(5)
