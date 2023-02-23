# NAME

ibapi a add - create Infoblox A records

# USAGE

- ibapi a add &lt;options/args>

# DESCRIPTION

The add command is used to create Infoblox A records.
To create a single A record, a single hostname and IP address can be provided as
command line arguments, in either order.
Alternatively, a list of records to add can be specified in a file (see --filename).

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

## OPTIONS - General

- -V &lt;view>, --View=&lt;view>:

    Specify the view for the new record.  Default: "default".

- -D, --Disable:

    Disable the new record.  Default: false.

- -c &lt;comment>, --Comment=&lt;comment>:

    Specify the comment for the new record.
    Alternatively, you can specify this via the --fields option.
    Default: "ibapi:address:add".

- --TTL=&lt;ttl>:

    Specify the ttl for the new record.
    Alternatively, you can specify this via the --fields option.

- -F &lt;fields>, --Fields=&lt;fields>:

    Specify fields and corresponding values for the new record.  For intance:
    "comment=RT100931",view=default,ttl=900".

- -f &lt;filename>, --Filename=&lt;filename>:

    Specify a filename containing a list of A records to create.
    Each line should contain a hostname and an IP address, in either order, separated
    by one or more spaces.
    Blank lines and lines beginning with "#" are ignored, as is anything on a line
    following a "#".

- -C, --Check:

    Before creating a new record, check if any "related" records already exist, and if
    so do not create the new record.
    Related records are those that share the same name and/or content, 
    such as an Host record and A record with the same name and/or IP address.
    Default: false.

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

- ibapi a add rb4.rice.edu 168.7.56.224

    Create a new A record with hostname "rb4.rice.edu" and IP address "168.7.56.224".

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
url:update(1),
a(1),
a:delete(1),
a:get(1),
a:update(1),
ibapi.conf(5)
