# NAME

ibapi alias update - update Infoblox Alias records

# USAGE

- ibapi alias update &lt;options/args>

# DESCRIPTION

The update command is used to update Infoblox Alias records.

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

    Specify the view of the record to update.  Default: "default".

- -D &lt;true|false>, --Disable=&lt;true|false>:

    Update the record's disabled status to the specified value.

- -n &lt;new\_hostname>, --Name=&lt;new\_hostname>:

    Update the name of the specified record to "new\_hostname".

- -c &lt;comment>, --Comment=&lt;comment>:

    Update the record's comment.

- --TTL=&lt;ttl>:

    Update the the record's TTL.

- -t &lt;new\_target>, --Target=&lt;new\_target>:

    Update the target of the specified alias to "new\_target".

- -T &lt;new\_target\_type>, --targetType=&lt;new\_target\_type>:

    Update the target\_type of the specified alias to "new\_target\_type".

- -F &lt;fields>, --Fields=&lt;fields>:

    Specify fields and corresponding values to be updated.  For intance:
    "comment=RT100931",view=default,ttl=900".

- -f &lt;filename>, --filename=&lt;filename>:

    Specify a filename containing a list of Alias records to update.
    Each line must contain a hostname and, depending on the specified options, a target.
    Blank lines and lines beginning with "#" are ignored, as is anything on a line
    following a "#".

- -C, --Check:

    Before updating a record's hostname or target, check if any "related" records
    already exist, and if so do not update the new record.
    Related records are those that share the same hostname and/or target, 
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

- ibapi alias update rb4.rice.edu somewhere.com -t nowhere.com

    Update the "rb4.rice.edu" Alias record, changing the target to "nowhere.com".

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
url(1),
url:add(1),
url:delete(1),
url:get(1),
url:update(1),
a(1),
a:delete(1),
a:get(1),
a:update(1),
a:add(1),
ibapi.conf(5)
