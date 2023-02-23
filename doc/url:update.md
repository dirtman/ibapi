# NAME

ibapi url update - make an Infoblox PUT request

# USAGE

- ibapi url update <URL>

# DESCRIPTION

The update command is used to make an Infoblox PUT API request.

Note that the Infoblox PUT method requires an object reference.
The various ibapi record "get" commands (such as "ibapi address get") all
have a "--ref" option to get the object reference of an object.

# OPTIONS

Some options can be specified with either a short (i.e., -h) or long
(i.e., --help) form.  In the latter case, case is non-sensitive.

Most options have a corresponding configuration file setting
that is equal to the long option name without the "--" prefix.
Command line options always override configuration file settings.
See ibapi.conf(5) for more information.

Note: while the "default" value for an option is sometimes specified in the option's
description, do not rely on this value.  It is always best to use the --ShowConfig
option to verify the real value.

## OPTIONS - IBAPI Infoblox API Options

- --APIBaseURL=&lt;url>:

    API base URL.
    Default: "https://infoblox.rice.edu/wapi/v2.11"

- --APIAuthMethod=&lt;method>:

    WAPI authentication method for accessing the Infoblox API.  Currently,
    only "Basic" authentication (username and password) is supported.

- -u=&lt;tokenID>, --APIAuthTokenID=&lt;tokenID>:

    Name of an authentication token ID.  Unless the --password option is also
    specified, the token ID is taken as the name of a file stored in the "secrets"
    directory and containing a "username:password" entry to be used for Basic
    authentication.
    If the --password option is also specified, the tokenID is taken as the 
    username for basic authentication.
    See also --SecretsDir.

- --SecretsDir=&lt;pathname>:

    Specify a directory where optional "token files" are kept.
    These files can be used as a slightly safer alternative to
    storing authentication credentials in the standard configuration files.
    The name of a secret file corresponds to a "tokenID" configured via --APIAuthTokenID,
    and the file should contain a single "username:password" entry.
    If the specified pathname does not begin with a "/", the directory is searched
    for in the same directories as the configuration file.
    Obviously, these files should be safely guarded.

- --HTTPTimeout=&lt;seconds>:

    Timeout in seconds for the HTTPS WAPI connection.  Default: 10.

## OPTIONS - Common To All IBAPI Commands

- -h, --help:

    Help; show usage information and exit.

- --showConfig:

    Read in and show all configuration settings and exit.

- -q, --Quiet (--noQuiet):

    Be quieter than normal.

- --Quieter (--noQuieter):

    Quieter mode.  Suppress all messages except warning and error messages.

- -v, --Verbose (--noVerbose):

    Be louder than normal. Over-rides the "--Quiet"  and "-Quieter" options.

- --nopage:

    By default, usage information (see --help) is piped to the pager specified
    by the environment variable "PAGER", if this environment variable is set,
    or by the pager specified by --Pager.
    The --nopage option disables this paging.

- --Pager=&lt;pager>:

    Specify a pager command for paging the usage information (with --help).  By default,
    the environment variable PAGER is used.  If a full path is not specified, the command
    is searched for using the PATH environment variable.

# EXAMPLES

- ibapi url update "record:cname/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuZWR1LnJpY2Uuc2VjaSxkYngsMTAuMTAuMTAuMjAx:dbx.seci.rice.edu/default?canonical=somewhere.edu"

    Change the target, or canonical name, of the referenced CNAME record.

- ibapi cname get dbalias.seci.rice.edu
- ref=$(ibapi cname get -r dbalias.seci.rice.edu | awk '{print $2}')
- ibapi url update "${ref}?canonical=somewhere.edu"

    Get the reference for the dbalias.seci.rice.edu CNAME record, and then
    use the reference to update the record.

# FILES

- /usr/site/ibapi-0.0/etc/ibapi.conf
- /etc/opt/ibapi/ibapi.conf
- /etc/opt/ibapi-0.0//ibapi.conf
- ~/.ibapi/ibapi.conf
- ~/.ibapi-0.0/ibapi.conf

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
a(1),
a:delete(1),
a:get(1),
a:update(1),
a:add(1),
ibapi.conf(5)
